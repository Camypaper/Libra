package libra

import (
	"bytes"
	"fmt"
)

type compiler struct {
	name string
	program
}

func (c compiler) Name() string {
	return c.name
}
func (c compiler) Run() Status {
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	c.program.cmd.Stdout = stdout
	c.program.cmd.Stderr = stderr
	res := c.program.Run()
	if res.Code != OK {
		return Status{
			Code: CE,
			Msg:  fmt.Sprintf("stdout: %v\nstderr:%v", stdout.String(), stderr.String()),
		}
	}
	return res
}
