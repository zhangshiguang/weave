package testing

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func AssertErrorInterface(t *testing.T, expected interface{}, actual error, desc string) {
	require.Error(t, actual, desc)
	require.Implements(t, expected, actual, desc)
}

func AssertErrorType(t *testing.T, expected interface{}, actual error, desc string) {
	require.Error(t, actual, desc)
	// require.IsType doesn't take the pointer element, and doesn't resolve
	// through interfaces, so we have to do this one ourselves.
	actualT, expectedT := reflect.TypeOf(actual), reflect.TypeOf(expected).Elem()
	if actualT != expectedT {
		require.FailNow(
			t,
			fmt.Sprintf("Expected %s but got %s", expectedT.String(), actualT.String()),
			desc,
		)
	}
}

// TrimTestArgs finds the first -- in os.Args and trim all args before
// that, returning true when a -- was in fact found.
func TrimTestArgs() bool {
	i, l := 0, len(os.Args)
	for ; i < l; i++ {
		if os.Args[i] == "--" {
			break
		}
	}
	if i == l {
		return false
	}
	os.Args = append(os.Args[:1], os.Args[i+1:l]...)
	return true
}
