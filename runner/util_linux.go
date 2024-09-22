package runner

import (
	"io"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/creack/pty"
)

func (e *Engine) killCmd(cmd *exec.Cmd) (pid int, err error) {
	pid = cmd.Process.Pid

	if e.config.Build.SendInterrupt {
		// Sending a signal to make it clear to the process that it is time to turn off
		if err = syscall.Kill(-pid, syscall.SIGINT); err != nil {
			return
		}
		time.Sleep(e.config.killDelay())
	}

	// https://stackoverflow.com/questions/22470193/why-wont-go-kill-a-child-process-correctly
	err = syscall.Kill(-pid, syscall.SIGKILL)

	// Wait releases any resources associated with the Process.
	_, _ = cmd.Process.Wait()
	return
}

func (e *Engine) startCmd(cmd string) (c *exec.Cmd, err error) {
	c = exec.Command("/bin/sh", "-c", cmd)
	f, err := pty.Start(c)
	if err != nil {
		return nil, err
	}

	go func() {
		_, _ = io.Copy(os.Stdin, f)
	}()
	go func() {
		_, _ = io.Copy(os.Stdout, f)
	}()
	go func() {
		_, _ = io.Copy(os.Stderr, f)
	}()
	return c, nil
}
