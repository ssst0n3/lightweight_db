package lightweight_db

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/ssst0n3/lightweight_db/test/test_data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnector_QueryRow(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var result int
		assert.NoError(t, Conn.QueryRow("SELECT 1", &result))
	})
	t.Run("empty", func(t *testing.T) {
		var result int
		assert.NoError(t, Conn.QueryRow("SELECT COUNT(*) WHERE 1=0", &result))
	})
}

func TestConnector_OrmMapTableByIdRet(t *testing.T) {
	result, err := Conn.OrmListTableUsingReflectRet(test_data.TableNameChallenge, test_data.Challenge{})
	assert.NoError(t, err)
	spew.Dump(result)
}

func TestConnector_OrmListTableUsingReflectRet(t *testing.T) {
	_, err := Conn.OrmListTableUsingReflectRet(test_data.TableNameChallenge, test_data.ChallengeWithId{})
	assert.NoError(t, err)
}

func TestConnector_OrmShowObjectByIdUsingReflectBind(t *testing.T) {
	Conn.InitTable(test_data.TableNameChallenge, test_data.Challenges, Conn.ResetAutoIncrementSqlite, nil)
	var challengeWithId test_data.ChallengeWithId
	assert.NoError(t, Conn.OrmShowObjectByIdUsingReflectBind(test_data.TableNameChallenge, 1, &challengeWithId))
	assert.Equal(t, test_data.Challenge1, challengeWithId)
}

func TestConnector_OrmShowObjectOnePropertyByIdUsingJsonBind(t *testing.T) {
	Conn.InitTable(test_data.TableNameChallenge, test_data.Challenges, Conn.ResetAutoIncrementSqlite, nil)
	var name string
	assert.NoError(t, Conn.OrmShowObjectOnePropertyByIdUsingJsonBind(
		test_data.TableNameChallenge,
		test_data.ColumnNameChallengeName,
		1,
		&name,
	),
	)
	assert.Equal(t, test_data.Challenge1.Name, name)
}

// TODO
func TestConnector_OrmQueryRowBind(t *testing.T) {
	type Category struct {
		Name        string `json:"name"`
		Parent      uint64 `json:"parent"`
		Description string `json:"description"`
	}

	type CategoryWithId struct {
		Id uint `json:"id"`
		Category
	}

	type CategoriesRecursive struct {
		CategoryWithId
		Sub []CategoriesRecursive `json:"sub"`
	}
	//var categoriesRecursiveList []CategoriesRecursive

	//Conn.OrmQueryRowsBind(&categoriesRecursiveList, query)
}

func TestConnector_OrmListTableByColumnBind(t *testing.T) {
	Conn.InitTable(test_data.TableNameChallenge, test_data.Challenges, Conn.ResetAutoIncrementSqlite, nil)
	var challenges []test_data.ChallengeWithId
	err := Conn.OrmListTableByColumnBind(test_data.TableNameChallenge, test_data.ColumnNameChallengeName, "name", &challenges)
	assert.NoError(t, err)
	assert.Equal(t, int64(test_data.Challenges[0].Id), challenges[0].Id)
}