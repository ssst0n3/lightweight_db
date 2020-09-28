package mysql

import (
	"github.com/ssst0n3/lightweight_db"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetDsnFromEnvNormal(t *testing.T) {
	assert.NoError(t, os.Setenv(lightweight_db.EnvDbName, "db"))
	assert.NoError(t, os.Setenv(lightweight_db.EnvDbHost, "host"))
	assert.NoError(t, os.Setenv(lightweight_db.EnvDbPort, "3306"))
	assert.NoError(t, os.Setenv(lightweight_db.EnvDbUser, "user"))
	assert.NoError(t, os.Setenv(lightweight_db.EnvDbPasswordFile, "/etc/hostname"))
	lightweight_db.Logger.Info(lightweight_db.GetDsnFromEnvNormal())
}
