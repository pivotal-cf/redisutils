package monit

import "github.com/BooleanCat/igo/ios/iexec"

type Monit struct {
	exec iexec.Exec
}

func New() *Monit {
	return &Monit{
		exec: new(iexec.ExecFake),
	}
}

func (monit *Monit) GetSummary() ([]byte, error) {
	cmd := monit.exec.Command("monit", "summary")
	summary, _ := cmd.CombinedOutput()
	return summary, nil
}
