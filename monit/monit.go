package monit

import (
	"regexp"

	"github.com/BooleanCat/igo/ios/iexec"
)

type Status int
type Statuses map[string]Status

const (
	StatusRunning Status = iota
	StatusNotMonitored
)

var statusMapping = Statuses{
	"running":       StatusRunning,
	"not monitored": StatusNotMonitored,
}

func getStatus(status string) Status {
	return statusMapping[status]
}

type Monit struct {
	MonitrcPath string

	exec iexec.Exec
}

func New() *Monit {
	return &Monit{
		exec: new(iexec.ExecWrap),
	}
}

func (monit *Monit) GetSummary() (Statuses, error) {
	rawSummary, err := monit.getRawSummary()
	if err != nil {
		return nil, err
	}

	processes := monit.getProcessesFromRawSummary(rawSummary)
	summaries := monit.newProcessMap(processes)

	return summaries, nil
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
	pattern := regexp.MustCompile(`(?m)^Process '([\w\-]+)'\s+([\w ]+)$`)
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
