package bash

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"unicode"

	"github.com/bearz-io/bzdev/lib/os/exec"
)

var wslEnabled = false

func init() {
	exec.Register("bash", &exec.Executable{
		Name:     "bash",
		Variable: "BASH_PATH",
		Windows: []string{
			"${ProgramFiles}\\Git\\bin\\bash.exe",
			"${ProgramFiles(x86)}\\Git\\bin\\bash.exe",
			"${ChocolateyInstall}\\msys2\\usr\\bin\\bash.exe",
			"${SystemRoom}\\System32\\bash.exe",
		},
		Linux: []string{
			"/bin/bash",
			"/usr/bin/bash",
		},
	})

	if runtime.GOOS == "windows" {
		fi, err := os.Stat("C:\\Windows\\System32\\bash.exe")
		wslEnabled = err == nil && !fi.IsDir()
	}
}

func WhichOrDefault() string {
	exe, _ := exec.Find("bash", nil)
	if exe == "" {
		return "bash"
	}

	return exe
}

func New(args ...string) *exec.Cmd {
	return exec.New(WhichOrDefault(), args...)
}

func File(file string) *exec.Cmd {
	args := []string{"-noprofile", "--norc", "-e", "-o", "pipefail"}
	exe := WhichOrDefault()
	if wslEnabled {
		if strings.HasSuffix("System32\\bash.exe", exe) {
			f, err := filepath.Abs(file)
			if err == nil {
				file = f
			}

			file = "/mnt/" + string(unicode.ToLower(rune(file[0]))) + file[2:]
			file = filepath.ToSlash(file)
		}
	}

	args = append(args, file)
	return exec.New(WhichOrDefault(), file)
}

func Script(script string) *exec.Cmd {
	if !strings.Contains(script, "\n") {
		script = strings.TrimSpace(script)

		if strings.HasSuffix(script, ".sh") {
			return File(script)
		}
	}

	args := []string{"-noprofile", "--norc", "-e", "-o", "pipefail", "-c", script}
	println("args", strings.Join(args, " "))
	return exec.New(WhichOrDefault(), args...)
}

func Run(script string) (*exec.PsOutput, error) {
	return Script(script).Run()
}

func Output(script string) (*exec.PsOutput, error) {
	return Script(script).Output()
}
