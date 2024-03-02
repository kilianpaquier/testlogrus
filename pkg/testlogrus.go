package testlogrus

import (
	"errors"
	"strings"
	"testing"

	"github.com/sirupsen/logrus/hooks/test"
)

var hook *test.Hook

// ErrNoHook is the panic error in case Logs is called without a previous call to CatchLogs.
var ErrNoHook = errors.New("testlogrus hook wasn't initialized with CatchLogs")

// CatchLogs creates a hook over logrus logs.
// Easy to use with Logs function to catch written logs in tests.
//
// Useful to confirm logging behavior over application.
//
// Note that at the end of t.Run a cleanup is called to reset the hook.
// As such CatchLogs must either be called at the beginning of every t.Run needing a log verification
// or at the beginning of every test function (TestXXX) needing logs verification.
func CatchLogs(t testing.TB) {
	// clean hook to ensure every test needs to call CatchLogs
	t.Cleanup(func() { hook = nil })

	// create global hook
	hook = test.NewGlobal()
}

// Logs resets the logrus test logs hooks and returns all catched logs from when CatchLogs was executed.
func Logs() string {
	if hook == nil {
		panic(ErrNoHook)
	}

	entries := hook.AllEntries()

	results := make([]string, len(entries))
	for i, entry := range entries {
		v, _ := entry.String()
		results[i] = v
	}

	hook.Reset()
	return strings.Join(results, "\n")
}
