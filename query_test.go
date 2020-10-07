package lightweight_db

import (
	"github.com/ssst0n3/lightweight_db/test/test_data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnector_QueryIdByGuid(t *testing.T) {
	id, err := Conn.QueryIdByGuid(test_data.TableNameChallenge, test_data.ColumnNameChallengeName, test_data.Challenge1.Name)
	assert.NoError(t, err)
	assert.Equal(t, int64(test_data.Challenge1.Id), id)
}
