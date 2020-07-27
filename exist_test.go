package lightweight_db

import (
	"github.com/ssst0n3/lightweight_db/test/test_data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnector_IsResourceExistsByGuid(t *testing.T) {
	t.Run("not exist", func(t *testing.T) {
		Conn.DeleteAllObjects(test_data.TableNameChallenge)
		exists, err := Conn.IsResourceExistsByGuid(test_data.TableNameChallenge, test_data.ColumnNameChallengeName, test_data.Challenge1.Name)
		assert.NoError(t, err)
		assert.Equal(t, false, exists)
	})

	t.Run("exist", func(t *testing.T) {
		Conn.InitTable(test_data.TableNameChallenge, test_data.Challenges, Conn.ResetAutoIncrementSqlite)
		exists, err := Conn.IsResourceExistsByGuid(test_data.TableNameChallenge, test_data.ColumnNameChallengeName, test_data.Challenge1.Name)
		assert.NoError(t, err)
		assert.Equal(t, true, exists)
	})
}
