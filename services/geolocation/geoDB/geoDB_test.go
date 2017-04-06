package geoDB

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestMain(t *testing.T) {
	t.Run("unit", func(t *testing.T) {
		suite.Run(t, new(GeoDBUnitSuite))
	})

	t.Run("functional", func(t *testing.T) {
		suite.Run(t, new(GeoDBFunctionalSuite))
	})
}
