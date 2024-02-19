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
	testlogrus.CatchLogs()

	t.Run("success", func(t *testing.T) {
		// Act
		log.Warn("something")

		// Assert
		assert.Contains(t, testlogrus.Logs(), "something")
	})
}
