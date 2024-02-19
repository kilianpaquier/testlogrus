# testlogrus <!-- omit in toc -->

- [How to use ?](#how-to-use-)
- [Features](#features)

## How to use ?

```sh
go get -u github.com/kilianpaquier/testlogrus@latest
```

## Features

The testlogrus package exposes two functions to use during testing. When calling `Logs` function, the logs are reset, as such if multiple verification are to be done, you should first affect `Logs` result to a variable.

```go
func TestSome(t *testing.T) {
  t.Run("error_test", func(t *testing.T) {
    // Arrange
    testlogrus.CatchLogs()

    // Act
    // call some functions that log with logrus

    // Assert
    logs := testlogrus.Logs()
    // assert some things around function expected logs
  })
}
```
