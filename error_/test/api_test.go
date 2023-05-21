package error_

import (
	"fmt"
	"testing"

	"github.com/m4gshm/expressions/error_"
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
	e, ok := error_.As[SomeErr](expected)
	assert.True(t, ok)
	assert.Equal(t, expected, e)
}
