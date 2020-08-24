package lightweight_db

import (
	"github.com/sirupsen/logrus"
	"github.com/ssst0n3/lightweight_db/test/test_data"
	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
	"testing"
)

const (
	driverName = "sqlite"
)

var Conn Connector

func init() {
	Conn = Connector{
		DriverName: driverName,
		Dsn:        "test/test_data/base.sqlite",
	}
	Conn.Init()
	Logger.SetLevel(logrus.DebugLevel)
}

//func (c Connector) InitTable(tableName string, r []test_data.ResourceWrapper) {
//	Logger.Info("remove all in table categories")
//	c.DeleteAllObjects(tableName)
//
//	c.ResetAutoIncrementSqlite(tableName)
//
//	Logger.Info("import test data")
//	if len(r) > 0 {
//		for _, wrapper := range r {
//			resource := wrapper.Resource
//			_, err := c.CreateObject(tableName, resource)
//			if err != nil {
//				awesome_error.CheckErr(err)
//				Logger.Fatal(err)
//			}
//		}
//	}
//}

func TestConnector_UpdateObject(t *testing.T) {
	Conn.InitTable(test_data.TableNameChallenge, test_data.Challenges, Conn.ResetAutoIncrementSqlite, nil)
	t.Run("simple struct", func(t *testing.T) {
		err := Conn.UpdateObject(int64(test_data.Challenge1.Id), test_data.TableNameChallenge, test_data.Challenge1Update)
		assert.NoError(t, err)
	})
	t.Run("nested struct", func(t *testing.T) {
		// TODO: use another model
	})
}

func TestConnector_CountTable(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		Conn.InitTable(test_data.TableNameChallenge, test_data.Challenges, Conn.ResetAutoIncrementSqlite, nil)
		count, err := Conn.CountTable(test_data.TableNameChallenge)
		assert.NoError(t, err)
		assert.Equal(t, uint(1), count)
	})
	t.Run("empty", func(t *testing.T) {
		Conn.DeleteAllObjects(test_data.TableNameChallenge)
		count, err := Conn.CountTable(test_data.TableNameChallenge)
		assert.NoError(t, err)
		assert.Equal(t, uint(0), count)
	})
	t.Run("table not exists", func(t *testing.T) {
		_, err := Conn.CountTable("not_exists")
		assert.Error(t, err)
	})
}
