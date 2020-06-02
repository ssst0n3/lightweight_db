package lightweight_db

import (
	"github.com/ssst0n3/lightweight_db/test/test_data"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestRetColsValues(t *testing.T) {
	type args struct {
		model     interface{}
		colsPtr   *[]string
		valuesPtr *[]interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				model: struct {
					Id   int    `json:"id"`
					Name string `json:"name"`
				}{
					Id:   1,
					Name: "test_name",
				},
				colsPtr:   nil,
				valuesPtr: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cols, args := RetColsValues(tt.args.model)
			assert.Equal(t, []string{"id", "name"}, cols)
			assert.Equal(t, []interface{}{1, "test_name"}, args)
		})
	}
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
