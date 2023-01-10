package ghworktree

import (
	"os"
	"os/exec"

	"github.com/black-desk/lib/go/errwrap"
	"github.com/cli/safeexec"
)

func RunCommand(command []string) error {
	exe, err := safeexec.LookPath(command[0])
	if err != nil {
		return errwrap.Trace(err)
	}

	cmd := exec.Command(exe, command[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	log.Debugf("run command: %v", cmd)

	return cmd.Run()
}
