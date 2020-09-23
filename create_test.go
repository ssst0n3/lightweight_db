package lightweight_db

import (
	"github.com/ssst0n3/lightweight_db/example/conn"
	"github.com/ssst0n3/lightweight_db/test/test_data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnector_CreateObject(t *testing.T) {
	conn.Conn.InitTable(test_data.TableNameChallenge, test_data.Challenges, conn.Conn.ResetAutoIncrementSqlite, nil)
	{
		_, err := conn.Conn.CreateObject(test_data.TableNameChallenge, test_data.Challenge1)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "UNIQUE")
	}

}

func TestConnector_CreateObjectPreventDuplicate(t *testing.T) {
	conn.Conn.InitTable(test_data.TableNameChallenge, test_data.Challenges, conn.Conn.ResetAutoIncrementSqlite, nil)
	exists, _, err := conn.Conn.CreateObjectPreventDuplicate(test_data.TableNameChallenge, test_data.Challenge1)
	assert.NoError(t, err)
	assert.Equal(t, true, exists)
}
