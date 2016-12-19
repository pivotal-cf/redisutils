package monit

import (
	"errors"
	"regexp"
	"time"

	"github.com/BooleanCat/igo/ios/iexec"
)

//ErrTimeout indicates that some monit action took too long
var ErrTimeout = errors.New("timed out waiting for monit operation")

//Status is an enumeration of monit statuses
type Status int

//Statuses is a mapping of monit statuses to Status
type Statuses map[string]Status

const (
	_ Status = iota

	//StatusRunning indicates `Process 'foo' running`
	StatusRunning

	//StatusNotMonitored indicates `Process 'foo' not monitored`
	StatusNotMonitored

	//StatusNotMonitoredStartPending indicates `Process 'foo' not monitored - start pending`
	StatusNotMonitoredStartPending

	//StatusInitializing indicates `Process 'foo' initializing`
	StatusInitializing

	//StatusDoesNotExist indicates `Process 'foo' Does not exist`
	StatusDoesNotExist

	//StatusNotMonitoredStopPending indicates `Process 'foo' not monitored - stop pending`
	StatusNotMonitoredStopPending

	//StatusRunningRestartPending indicates `Process 'foo' running - restart pending`
	StatusRunningRestartPending
)

var statusMapping = Statuses{
	"running":                       StatusRunning,
	"not monitored":                 StatusNotMonitored,
	"not monitored - start pending": StatusNotMonitoredStartPending,
	"initializing":                  StatusInitializing,
	"Does not exist":                StatusDoesNotExist,
	"not monitored - stop pending":  StatusNotMonitoredStopPending,
	"running - restart pending":     StatusRunningRestartPending,
}

func getStatus(status string) Status {
	return statusMapping[status]
}

//Monit is a controller for the monit CLI
type Monit struct {
	MonitrcPath string

	interval time.Duration
	timeout  time.Duration
	exec     iexec.Exec
}

//New is the correct way to initialise a new Monit
func New() *Monit {
	return &Monit{
		interval: time.Millisecond * 100,
		timeout:  time.Second * 15,
		exec:     new(iexec.ExecWrap),
	}
}

//GetSummary is synonymous with `monit summary`
func (monit *Monit) GetSummary() (Statuses, error) {
	rawSummary, err := monit.getRawSummary()
	if err != nil {
		return nil, err
	}

	processes := monit.getProcessesFromRawSummary(rawSummary)
	return monit.newProcessMap(processes), nil
}

//GetStatus a job specific Status from GetSummary
func (monit *Monit) GetStatus(job string) (Status, error) {
	summary, err := monit.GetSummary()
	if err != nil {
		return 0, err
	}
	return summary[job], nil
}

//Start is synonymous with `monit start {job}`
func (monit *Monit) Start(job string) error {
	cmd := monit.getMonitCommand("start", job)
	return cmd.Run()
}

//Stop is synonymous with `monit stop {job}`
func (monit *Monit) Stop(job string) error {
	cmd := monit.getMonitCommand("stop", job)
	return cmd.Run()
}

//StartAndWait runs Start(job) and waits for GetStatus(job) to report StatusRunning
func (monit *Monit) StartAndWait(job string) error {
	err := monit.Start(job)
	if err != nil {
		return err
	}

	return monit.waitFor(job, StatusRunning)
}

//StopAndWait runs Stop(job) and waits for GetStatus(job) to report StatusNotMonitored
func (monit *Monit) StopAndWait(job string) error {
	err := monit.Stop(job)
	if err != nil {
		return err
	}

	return monit.waitFor(job, StatusNotMonitored)
}

func (monit *Monit) waitFor(job string, status Status) error {
	for elapsed := time.Duration(0); elapsed < monit.timeout; elapsed = elapsed + monit.interval {
		currentStatus, err := monit.GetStatus(job)
		if err != nil {
			return err
		}

		if status == currentStatus {
			return nil
		}

		time.Sleep(monit.interval)
	}

	return ErrTimeout
}

func (monit *Monit) getRawSummary() (string, error) {
	cmd := monit.getMonitCommand("summary")

	rawSummary, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(rawSummary), nil
}

func (monit *Monit) getProcessesFromRawSummary(summary string) [][]string {
	pattern := regexp.MustCompile(`(?m)^Process '([\w\-]+)'\s+([\w \-]+)$`)
	return pattern.FindAllStringSubmatch(summary, -1)
}

func (monit *Monit) newProcessMap(processes [][]string) Statuses {
	processMap := make(Statuses)
	for _, process := range processes {
		processMap[process[1]] = getStatus(process[2])
	}

	return processMap
}

func (monit *Monit) getMonitCommand(args ...string) iexec.Cmd {
	var allArgs []string

	if monit.MonitrcPath != "" {
		allArgs = []string{"-c", monit.MonitrcPath}
	}

	allArgs = append(allArgs, args...)
	return monit.exec.Command("monit", allArgs...)
}
