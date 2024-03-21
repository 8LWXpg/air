package runner

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func (e *Engine) killCmd(cmd *exec.Cmd) (pid int, err error) {
	pid = cmd.Process.Pid
	// https://stackoverflow.com/a/44551450
	kill := exec.Command("TASKKILL", "/T", "/F", "/PID", strconv.Itoa(pid))
	return pid, kill.Run()
}

func (e *Engine) startCmd(cmd string) (c *exec.Cmd, err error) {
	if !strings.Contains(cmd, ".exe") {
		e.runnerLog("CMD will not recognize non .exe file for execution, path: %s", cmd)
	}
	c = exec.Command("cmd", "/c", cmd)

	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	err = c.Start()
	if err != nil {
		return nil, err
	}
	return c, nil
}
