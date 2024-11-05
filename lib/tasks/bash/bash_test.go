package bash_test

import (
	"strings"
	"testing"

	"github.com/bearz-io/bzdev/lib/tasks/bash"
	"github.com/stretchr/testify/assert"
)

func TestBashScript(t *testing.T) {
	println(bash.WhichOrDefault())
	r, e := bash.Script("echo 'Hello World'").Run()
	if e != nil {
		t.Errorf("Error running bash script: %v", e)
	}

	assert.Equal(t, 0, r.Code)
}

func TestBashRun(t *testing.T) {
	r, e := bash.Run("echo 'Hello World'")
	if e != nil {
		t.Errorf("Error running bash script: %v", e)
	}

	assert.Equal(t, 0, r.Code)
}

func TestBashOutput(t *testing.T) {
	r, e := bash.Script("echo 'Hello World'").Output()
	if e != nil {
		t.Errorf("Error running bash script: %v", e)
	}

	assert.Equal(t, 0, r.Code)
	assert.Equal(t, "Hello World", strings.TrimSpace(r.Text()))
}
