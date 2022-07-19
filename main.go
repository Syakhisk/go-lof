package main

import (
	"fmt"
	"os"

	"github.com/bitfield/script"
	"github.com/syakhisk/go-lof/program"
)

type MyProgs map[string]*program.Program

func main() {
	myProgs := make(MyProgs)
	initProgs(myProgs)

	args, _ := script.Args().Slice()

	if len(args) == 0 || len(args) > 1 {
		ExitError("Only one arguments are allowed")
		return
	}

	progKey := args[0]
	prog, ok := myProgs[progKey]
	if !ok {
		ExitError("key `" + progKey + "` doesn't exists.")
		return
	}

	if isRunning(prog) {
		fmt.Println("running")
		prog.Focus()
	} else {
		fmt.Println("not running")
		prog.Launch()
	}
}

func initProgs(myProgs MyProgs) {
	var p *program.Program
	p = program.NewProgram()
	p.ClassName = "Code"
	p.LaunchCmd = "code"

	myProgs["vscode"] = p

	p = program.NewProgram()
	p.Name = "DevTools"
	p.LaunchCmd = program.WithKeybind("ctrl+shift+j", "Google-chrome", "classname")

	myProgs["devtools"] = p
}

func isRunning(prog *program.Program) bool {
	cmd := `xdotool search --desktop "$(xdotool get_desktop)"`

	if prog.Name != "" {
		cmd = cmd + " --name " + prog.Name
	} else if prog.ClassName != "" {
		cmd = cmd + " --classname " + prog.ClassName
	} else if prog.Class != "" {
		cmd = cmd + " --class " + prog.Class
	} else {
		return false
	}

	count, err := script.Exec(cmd).CountLines()
	if err != nil || count < 1 {
		return false
	}

	return true
}

func ExitError(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
