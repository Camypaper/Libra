package libra

import "fmt"

/*
Src means source code.
*/
type Src struct {
	Name    string
	Compile string
	Exec    string
}

func (src Src) compiler() compiler {
	prog, _ := newProgram(src.Compile)
	return compiler{
		name:    src.Name,
		program: prog,
	}
}

func (src Src) String() string {
	return fmt.Sprintf("Name:%v, Compile:`%v`, Exec:`%v`]", src.Name, src.Compile, src.Exec)
}
