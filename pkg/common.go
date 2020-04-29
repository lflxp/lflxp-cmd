package pkg

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
)

func ChangeYourCmdEnvironment(cmd *exec.Cmd) error {
	env := os.Environ()
	cmdEnv := []string{}

	for _, e := range env {
		i := strings.Index(e, "=")
		if i > 0 && (e[:i] == "ENV_NAME") {
			// do yourself
		} else {
			cmdEnv = append(cmdEnv, e)
		}
	}
	cmd.Env = cmdEnv

	return nil
}

func ExecCommandString(cmd string) (string, error) {
	pipeline := exec.Command("/bin/bash", "-c", cmd)
	ChangeYourCmdEnvironment(pipeline)
	var out bytes.Buffer
	var stderr bytes.Buffer
	pipeline.Stdout = &out
	pipeline.Stderr = &stderr
	err := pipeline.Run()
	if err != nil {
		return stderr.String(), err
	}
	return out.String(), nil
}

// Home returns the home directory for the executing user.
//
// This uses an OS-specific method for discovering the home directory.
// An error is returned if a home directory cannot be detected.
func Home() (string, error) {
	user, err := user.Current()
	if nil == err {
		return user.HomeDir, nil
	}

	// cross compile support

	if "windows" == runtime.GOOS {
		return homeWindows()
	}

	// Unix-like system, so just assume Unix
	return homeUnix()
}

func homeUnix() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}

func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
}

func IsPathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
