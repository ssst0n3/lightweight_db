package lightweight_db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnector_KVGetValueByKey(t *testing.T) {
	assert.NoError(t, Conn.CreateTableConfig())
	config := Config{
		Key:   "test",
		Value: "aaa",
	}
	_, err := Conn.CreateObject(TableNameConfig, config)
	assert.NoError(t, err)
	configResult, err := Conn.KVGetValueByKey(TableNameConfig, ColumnNameConfigValue, ColumnNameConfigKey, config.Key)
	assert.NoError(t, err)
	assert.Equal(t, config, configResult)
}

func TestConnector_ShouldInitialize(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		Conn.DeleteAllObjects(TableNameConfig)
		shouldInitialize, err := Conn.ShouldInitialize()
		assert.NoError(t, err)
		assert.Equal(t, true, shouldInitialize)
	})
	t.Run("is_initialized", func(t *testing.T) {
		Conn.DeleteAllObjects(TableNameConfig)
		_, err := Conn.CreateObject(TableNameConfig, Config{
			Key:   "is_initialized",
			Value: "true",
		})
		assert.NoError(t, err)
		shouldInitialize, err := Conn.ShouldInitialize()
		assert.NoError(t, err)
		assert.Equal(t, false, shouldInitialize)
	})
	t.Run("not_initialized", func(t *testing.T) {
		Conn.DeleteAllObjects(TableNameConfig)
		_, err := Conn.CreateObject(TableNameConfig, Config{
			Key:   "is_initialized",
			Value: "false",
		})
		assert.NoError(t, err)
		shouldInitialize, err := Conn.ShouldInitialize()
		assert.NoError(t, err)
		assert.Equal(t, true, shouldInitialize)
	})
}