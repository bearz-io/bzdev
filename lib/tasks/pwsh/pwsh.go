package pwsh

import (
	"fmt"
	"strings"

	"github.com/bearz-io/bzdev/lib/os/exec"
)

func init() {
	exec.Register("pwsh", &exec.Executable{
		Name:     "pwsh",
		Variable: "POWERSHELL_PATH",
		Windows: []string{
			"${ProgramFiles}\\PowerShell\\7\\pwsh.exe",
			"${ProgramFiles}\\PowerShell\\6\\pwsh.exe",
		},
		Linux: []string{
			"/usr/bin/pwsh",
			"/opt/microsoft/powershell/7/pwsh",
			"/opt/microsoft/powershell/6/pwsh",
		},
	})
}

func WhichOrDefault() string {
	exe, _ := exec.Find("pwsh", nil)
	if exe == "" {
		return "pwsh"
	}

	return exe
}

func New(args ...string) *exec.Cmd {
	return exec.New(WhichOrDefault(), args...)
}

func File(file string) *exec.Cmd {
	args := []string{"-NoLogo", "-NonInteractive", "-NoProfile", "-ExecutionPolicy", "Bypass", "-File", file}
	return exec.New(WhichOrDefault(), args...)
}

func Script(script string) *exec.Cmd {
	exe := WhichOrDefault()
	args := []string{"-NoLogo", "-NonInteractive", "-NoProfile", "-ExecutionPolicy", "Bypass"}
	if !strings.Contains(script, "\n") {
		script = strings.TrimSpace(script)
		if strings.HasSuffix(script, ".ps1") {
			args = append(args, "-File", script)
			return exec.New(exe, args...)
		}
	}
	wrap := `
$ErrorActionPreference = "Stop"
%s

if ($LASTEXITCODE -ne 0) {
	exit $LASTEXITCODE
}`
	script = fmt.Sprintf(wrap, script)
	args = append(args, "-Command", script)
	return exec.New(exe, args...)
}

func Run(script string) (*exec.PsOutput, error) {
	return Script(script).Run()
}

func Output(script string) (*exec.PsOutput, error) {
	return Script(script).Output()
}
