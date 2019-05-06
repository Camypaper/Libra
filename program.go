package libra

import (
	"os/exec"

	shellwords "github.com/mattn/go-shellwords"
)

/*
Runnable shows executable program.
*/
type Runnable interface {
	Run() Status
	Kill()
}

type program struct {
	cmd *exec.Cmd
}

func (prog program) Run() Status {
	cmd := prog.cmd
	if cmd != nil {
		err := cmd.Run()
		if err != nil {
			return Status{Code: RE, Msg: "Runtime Error"}
		}
		return Status{Code: OK, Msg: "Success"}
	}
	return Status{Code: NG, Msg: "Command not found"}
}
func (prog program) Kill() {
	cmd := prog.cmd
	if cmd != nil {
		proc := cmd.Process
		if proc != nil {
			proc.Kill()
		}
	}
}

func newProgram(cmd string) (program, error) {
	p := shellwords.NewParser()
	p.ParseEnv = true
	args, err := p.Parse(cmd)
	if err != nil {
		return program{}, err
	}
	switch len(args) {
	case 0:
		return program{}, nil
	case 1:
		return program{cmd: exec.Command(args[0])}, nil
	default:
		return program{cmd: exec.Command(args[0], args[1:]...)}, nil
	}
}
