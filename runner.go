package libra

import (
	"fmt"
	"time"
)

/*
Runner execute runnable
*/
type Runner struct {
	TL float64
}

/*
Exec !
*/
func (runner Runner) Exec(target Runnable) Status {
	ch := make(chan Status, 1)
	go func() { ch <- target.Run() }()
	select {
	case res := <-ch:
		return res
	case <-time.After(time.Duration(runner.TL) * time.Second):
		target.Kill()
		return Status{Code: TLE, Msg: fmt.Sprintf("%.1f seconds over", runner.TL)}
	}
}
