package os_test

import (
	"runtime"
	"testing"

	"github.com/bearz-io/bzdev/lib/os"
	"github.com/stretchr/testify/assert"
)

func TestOs(t *testing.T) {
	assert.Equal(t, os.PLATFORM, runtime.GOOS)
}
