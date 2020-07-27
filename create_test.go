package lightweight_db

import (
	"github.com/ssst0n3/lightweight_db/test/test_data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnector_CreateObject(t *testing.T) {
	Conn.InitTable(test_data.TableNameChallenge, test_data.Challenges, Conn.ResetAutoIncrementSqlite)
	{
		_, err := Conn.CreateObject(test_data.TableNameChallenge, test_data.Challenge1)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "UNIQUE")
	}

}
