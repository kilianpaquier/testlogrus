/*
The testlogrus package exposes two functions to use during testing.

When calling `Logs` function, the logs are reset,
as such if multiple verification are to be done,
you should first affect `Logs` result to a variable.

	func TestSome(t *testing.T) {
		t.Run("error_test", func(t *testing.T) {
			// Arrange
			// mandatory call to setup logrus logging hook
			// logging hook is set to nil at the end of t.Run
			testlogrus.CatchLogs(t)

			// Act
			// call some functions that log with logrus

			// Assert
			logs := testlogrus.Logs()
			// assert some things around function expected logs
		})
	}
*/
package testlogrus
