package program

import (
	"fmt"
	"os"

	"github.com/bitfield/script"
)

type Program struct {
	Name      string
	ClassName string
	Class     string

	LaunchCmd string
	FocusCmd  string
}

func NewProgram() *Program {
	p := new(Program)
	return p
}

func WithKeybind(keybind string, target string, targetType string) string {
	cmd := `xdotool search --desktop "$(xdotool get_desktop)"`

	if targetType == "class" {
		cmd = cmd + " --class " + target
	} else if targetType == "classname" {
		cmd = cmd + " --classname " + target
	} else {
		cmd = cmd + " --name " + target
	}

	cmd = cmd + " windowactivate --sync %1 key " + keybind

	return cmd
}

func (p *Program) Launch() {
	script.Exec(p.LaunchCmd).Stdout()
}

func (p *Program) Focus() {
	if p.FocusCmd != "" {
		script.Exec(p.FocusCmd).Stdout()
		return
	}

	cmd := ""

	if p.Name != "" {
		cmd = fmt.Sprintf(
			`xdotool search --desktop "$(xdotool get_desktop)" --name "%s" windowactivate`, p.Name,
		)
	} else if p.ClassName != "" {
		cmd = fmt.Sprintf(
			`xdotool search --desktop "$(xdotool get_desktop)" --classname "%s" windowactivate`, p.ClassName,
		)
	} else if p.Class != "" {
		cmd = fmt.Sprintf(
			`xdotool search --desktop "$(xdotool get_desktop)" --class "%s" windowactivate`, p.Class,
		)
	} else {
		fmt.Println("at least focus command, classname, or name should be provided")
		os.Exit(1)
		return
	}

	script.Exec(cmd).Stdout()
}
