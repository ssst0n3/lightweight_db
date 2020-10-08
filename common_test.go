package lightweight_db

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
	"github.com/ssst0n3/awesome_libs"
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

func TestFetchRows(t *testing.T) {
	query := "SELECT 1"
	rows, err := Conn.Query(query)
	assert.NoError(t, err)
	objects, err := FetchRows(rows)
	spew.Dump(objects)
	expect := []awesome_libs.Dict{
		{"1": int64(1)},
	}
	assert.Equal(t, expect, objects)
}

func TestFetchOneRow(t *testing.T) {
	query := "SELECT 1"
	rows, err := Conn.Query(query)
	assert.NoError(t, err)
	objects, err := FetchOneRow(rows)
	spew.Dump(objects)
	expect := awesome_libs.Dict{
		"1": int64(1),
	}
	assert.Equal(t, expect, objects)
}

func TestConnector_ListObjects(t *testing.T) {
	query := "SELECT 1"
	objects, err := Conn.ListObjects(query)
	assert.NoError(t, err)
	expect := []awesome_libs.Dict{
		{"1": int64(1)},
	}
	assert.Equal(t, expect, objects)
}

func TestConnector_MapObjectById(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		query := "SELECT 1 as id"
		objects, err := Conn.MapObjectById(query)
		assert.NoError(t, err)
		expect := map[int64]awesome_libs.Dict{
			int64(1): {"id": int64(1)},
		}
		assert.Equal(t, expect, objects)
	})
	t.Run("empty", func(t *testing.T) {
		query := "SELECT 1 as id limit 0"
		objects, err := Conn.MapObjectById(query)
		assert.NoError(t, err)
		expect := map[int64]awesome_libs.Dict{}
		assert.Equal(t, expect, objects)
	})
}

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
