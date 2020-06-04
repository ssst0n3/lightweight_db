package lightweight_db

import (
	"github.com/sirupsen/logrus"
	awesomeError "github.com/ssst0n3/awesome_libs/error"
	"github.com/ssst0n3/lightweight_db/test/test_data"
	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
	"testing"
)

var Conn Connector

func init() {
	Conn = Connector{
		DriverName: "sqlite",
		Dsn:        "test/test_data/base.sqlite",
	}
	Conn.Init()
	Logger.SetLevel(logrus.DebugLevel)
}

func (c Connector) InitTable(tableName string, r []test_data.ResourceWrapper) {
	Logger.Info("remove all in table categories")
	c.DeleteAllObjects(tableName)

	c.ResetAutoIncrementSqlite(tableName)

	Logger.Info("import test data")
	if len(r) > 0 {
		for _, wrapper := range r {
			resource := wrapper.Resource
			_, err := c.CreateObject(tableName, resource)
			if err != nil {
				awesomeError.CheckErr(err)
				Logger.Fatal(err)
			}
		}
	}
}

func TestConnector_IsResourceExistsByGuid(t *testing.T) {
	t.Run("not exist", func(t *testing.T) {
		Conn.DeleteAllObjects(test_data.TableNameChallenge)
		exists, err := Conn.IsResourceExistsByGuid(test_data.TableNameChallenge, test_data.ColumnNameChallengeName, test_data.Challenge1.Name)
		assert.NoError(t, err)
		assert.Equal(t, false, exists)
	})

	t.Run("exist", func(t *testing.T) {
		Conn.InitTable(test_data.TableNameChallenge, test_data.Challenges)
		exists, err := Conn.IsResourceExistsByGuid(test_data.TableNameChallenge, test_data.ColumnNameChallengeName, test_data.Challenge1.Name)
		assert.NoError(t, err)
		assert.Equal(t, true, exists)
	})
}

func TestConnector_UpdateObject(t *testing.T) {
	Conn.InitTable(test_data.TableNameChallenge, test_data.Challenges)
	t.Run("simple struct", func(t *testing.T) {
		err := Conn.UpdateObject(int64(test_data.Challenge1.Id), test_data.TableNameChallenge, test_data.Challenge1Update)
		assert.NoError(t, err)
	})
	t.Run("nested struct", func(t *testing.T) {
		// TODO: use another model
	})
}
