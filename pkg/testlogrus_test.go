package testlogrus_test

import (
	"context"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	testlogrus "github.com/kilianpaquier/testlogrus/pkg"
)

func TestCatchLogs(t *testing.T) {
	log := logrus.WithContext(context.Background())

	t.Run("error_panic", func(t *testing.T) {
		// Act & Assert
		assert.PanicsWithValue(t, testlogrus.ErrNoHook, func() { testlogrus.Logs() })
	})

	t.Run("success", func(t *testing.T) {
		// Arrange
		testlogrus.CatchLogs(t)

		// Act
		log.Warn("something")

		// Assert
		assert.Contains(t, testlogrus.Logs(), "something")
	})
}
