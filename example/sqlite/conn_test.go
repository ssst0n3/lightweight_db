package sqlite

import (
	"github.com/ssst0n3/lightweight_db"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetDsnFromEnvNormal(t *testing.T) {
	assert.NoError(t, os.Setenv(lightweight_db.EnvDbDsn, "../test/test_data/base.sqlite"))
	lightweight_db.Logger.Info(lightweight_db.GetDsnFromEnvNormal())
}