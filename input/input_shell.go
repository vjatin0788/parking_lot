package input

import (
	"fmt"
	"os"
	"os/user"
)

func CreateIntractiveShell() {

	me, err := user.Current()
	if err != nil {
		panic(err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	os.Setenv("PL_ENV", "1")

	pa := os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Dir:   cwd,
	}

	proc, err := os.StartProcess("/usr/bin/login", []string{"login", "-fpl", me.Username}, &pa)
	if err != nil {
		panic(err)
	}

	state, err := proc.Wait()
	if err != nil {
		panic(err)
	}

	fmt.Printf("<< Exited shell: %s\n", state.String())
}
