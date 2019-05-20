package input

import (
	"os"
	"os/user"
)

//Create the interactive shell to execute the commands.
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

	_, err = proc.Wait()
	if err != nil {
		panic(err)
	}

}
