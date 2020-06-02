package lightweight_db

import (
	"github.com/ssst0n3/lightweight_db/test/test_data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnector_OrmShowObjectByIdUsingReflectBind(t *testing.T) {
	Conn.InitTable(test_data.TableNameChallenge, test_data.Challenges)
	var challengeWithId test_data.ChallengeWithId
	assert.NoError(t, Conn.OrmShowObjectByIdUsingReflectBind(test_data.TableNameChallenge, 1, &challengeWithId))
	assert.Equal(t, test_data.Challenge1, challengeWithId)
}

func TestConnector_OrmShowObjectOnePropertyByIdUsingJsonBind(t *testing.T) {
	Conn.InitTable(test_data.TableNameChallenge, test_data.Challenges)
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
