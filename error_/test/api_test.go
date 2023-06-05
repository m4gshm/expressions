package error_

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/m4gshm/expressions/error_"
	"github.com/m4gshm/expressions/expr/get"
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
	e, ok := error_.As[SomeErr](expected)
	assert.True(t, ok)
	assert.Equal(t, expected, e)
}

func Test_Convert(t *testing.T) {
	file, err := error_.Convertt(error_.Convert(error_.Catch(os.MkdirTemp(os.TempDir(), strconv.Itoa(rand.Int()))),
		func(userTempDir string) string { return filepath.Join(userTempDir, "out.txt") }), os.Create,
	).Get()

	if file != nil {
		defer file.Close()
	}
	assert.NoError(t, err)
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
