package error_

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	errors "github.com/m4gshm/expressions/error_"
	"github.com/m4gshm/expressions/error_/catch"
	"github.com/m4gshm/expressions/error_/try"
	"github.com/m4gshm/expressions/expr/get"
	"github.com/stretchr/testify/assert"
)

type SomeErr struct {
	Status int
}

// Error implements error
func (e SomeErr) Error() string {
	return fmt.Sprintf("error status %d", e.Status)
}

var _ error = SomeErr{}

func Test_As(t *testing.T) {
	expected := SomeErr{Status: 100}
	e, ok := errors.As[SomeErr](expected)
	assert.True(t, ok)
	assert.Equal(t, expected, e)
}

func Test_Run(t *testing.T) {
	catcher, userTempDir := catch.One(os.MkdirTemp(os.TempDir(), strconv.Itoa(rand.Int())))
	var file *os.File
	catcher.Run(func() { catcher, file = catch.One(os.Create(filepath.Join(userTempDir, "out.txt"))) })
	if file != nil {
		defer file.Close()
	}
	assert.NoError(t, catcher.Err)
}

func Test_Convert(t *testing.T) {
	catcher, userTempDir := catch.One(os.MkdirTemp(os.TempDir(), strconv.Itoa(rand.Int())))
	file := try.Convertt(catcher, filepath.Join(userTempDir, "out.txt"), os.Create)
	if file != nil {
		defer file.Close()
	}
	assert.NoError(t, catcher.Err)
}

func Test_NoCatch(t *testing.T) {
	file, err := func() (*os.File, error) {
		userTempDir, err := os.MkdirTemp(os.TempDir(), strconv.Itoa(rand.Int()))
		return get.If_(err == nil, func() (*os.File, error) { return os.Create(filepath.Join(userTempDir, "out.txt")) }).ElseErr(err)
	}()
	if file != nil {
		defer file.Close()
	}
	assert.NoError(t, err)
}
