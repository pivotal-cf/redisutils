package monit

import "github.com/BooleanCat/igo/ios/iexec"

const (
	StatusRunning = iota
	StatusNotMonitored
)

type Monit struct {
	MonitrcPath string

	exec iexec.Exec
}

func New() *Monit {
	return &Monit{
		exec: new(iexec.ExecFake),
	}
}

func (monit *Monit) GetSummary() (map[string]int, error) {
	cmd := monit.getMonitCommand("summary")

	_, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return map[string]int{
		"process-watcher":   StatusRunning,
		"process-destroyer": StatusRunning,
		"cf-redis-broker":   StatusRunning,
		"broker-nginx":      StatusRunning,
		"route_registrar":   StatusRunning,
	}, nil
}

func (monit *Monit) getMonitCommand(args ...string) iexec.Cmd {
	var allArgs []string

	if monit.MonitrcPath != "" {
		allArgs = []string{"-c", monit.MonitrcPath}
	}

	allArgs = append(allArgs, args...)
	return monit.exec.Command("monit", allArgs...)
}
