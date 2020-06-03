package lightweight_db

import (
	"github.com/ssst0n3/lightweight_db/test/test_data"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

// TODO
func ExampleRetColsValues() {

}

func TestRetColsValues(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		colsRet, valuesRet := RetColsValues(struct {
			Name string `json:"name"`
		}{"test"})
		assert.Equal(t, []string{"name"}, colsRet)
		assert.Equal(t, []interface{}{"test"}, valuesRet)
	})
	t.Run("time.Time", func(t *testing.T) {
		now := time.Now()
		colsRet, valuesRet := RetColsValues(struct {
			T time.Time `json:"t"`
		}{now})
		assert.Equal(t, []string{"t"}, colsRet)
		assert.Equal(t, []interface{}{now}, valuesRet)
	})
	t.Run("nested struct", func(t *testing.T) {
		type Nested struct {
			Nested bool `json:"nested"`
		}
		colsRet, valuesRet := RetColsValues(struct {
			Id uint `json:"id"`
			Nested
		}{
			Id: 1,
			Nested: Nested{
				Nested: true,
			},
		})
		assert.Equal(t, []string{"id", "nested"}, colsRet)
		assert.Equal(t, []interface{}{uint(1), true}, valuesRet)
	})
}

func TestBindModelFromMap(t *testing.T) {
	var challengeWithId test_data.ChallengeWithId
	assert.NoError(t, BindModelFromMap(&challengeWithId, test_data.Challenge1FromDbSimulate))
	assert.Equal(t, test_data.Challenge1, challengeWithId)
}

func TestRetModelFromMap(t *testing.T) {
	type args struct {
		model  interface{}
		object map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "struct with bool",
			args: args{
				model:  test_data.ChallengeWithId{},
				object: test_data.Challenge1FromDbSimulate,
			},
			want:    test_data.Challenge1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RetModelFromMap(tt.args.model, tt.args.object)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReflectModelFromMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReflectModelFromMap() got = %v, want %v", got, tt.want)
			}
		})
	}
}
