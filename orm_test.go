package lightweight_db

import (
	"github.com/ssst0n3/lightweight_db/test/test_data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnector_OrmShowObjectByIdUsingReflect(t *testing.T) {
	Conn.InitTable(test_data.TableNameChallenge, test_data.Challenges)
	var challengeWithId test_data.ChallengeWithId
	assert.NoError(t, Conn.OrmShowObjectByIdUsingReflect(test_data.TableNameChallenge, 1, &challengeWithId))
	assert.Equal(t, test_data.Challenge1, challengeWithId)
}