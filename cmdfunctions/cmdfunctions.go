package cmdfunctions

import (
	"fmt"
	"fs/util"
	"os"
	"os/exec"
)

func Response(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
}

func RunCommand(command string) *exec.Cmd {
	cmd := exec.Command("bash", "-c", command)
	err := cmd.Run()
	util.Error(err)
	return cmd
}

func Mkdir(dir string, root bool) {
	var sudo string = ""
	if root {
		sudo = "sudo "
	}

	command := fmt.Sprintf("%vmkdir -p %v", sudo, dir)
	cmd := RunCommand(command)

	Response(cmd)
}

func Touch(pathToFile string, root bool) {
	var sudo string = ""
	if root {
		sudo = "sudo "
	}

	command := fmt.Sprintf("%vtouch %v", sudo, pathToFile)
	cmd := RunCommand(command)

	Response(cmd)
}

func Echo(input string, output string, typeChange string, root bool) {
	var sudo string = ""
	if root {
		sudo = "sudo "
	}

	command := fmt.Sprintf("%vecho -e '%v' %v %v %v", sudo, input, typeChange, sudo ,output)
	cmd := RunCommand(command)

	Response(cmd)
}