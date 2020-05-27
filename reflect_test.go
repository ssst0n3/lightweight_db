package lightweight_db

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestFieldByJsonTag(t *testing.T) {
	type args struct {
		v       reflect.Value
		jsonTag string
	}
	tests := []struct {
		name  string
		args  args
		want  reflect.Value
		want1 bool
	}{
		{
			name: "struct",
			args: args{
				v: Reflect(struct {
					Name string `json:"name"`
				}{
					Name: "john",
				}),
				jsonTag: "name",
			},
			want:  Reflect("john"),
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FieldByJsonTag(tt.args.v, tt.args.jsonTag)
			assert.Equal(t, tt.want.Interface(), got.Interface())
			if got1 != tt.want1 {
				t.Errorf("FieldByJsonTag() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestIsPointer(t *testing.T) {
	type args struct {
		model interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPointer(tt.args.model); got != tt.want {
				t.Errorf("IsPointer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReflect(t *testing.T) {
	type args struct {
		model interface{}
	}
	tests := []struct {
		name string
		args args
		want reflect.Value
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reflect(tt.args.model); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reflect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReflectByModel(t *testing.T) {
	type args struct {
		model interface{}
	}
	tests := []struct {
		name string
		args args
		want reflect.Value
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReflectByModel(tt.args.model)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReflectByModel() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReflectByPtr(t *testing.T) {
	type args struct {
		modelPtr interface{}
	}
	tests := []struct {
		name string
		args args
		want reflect.Value
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReflectByPtr(tt.args.modelPtr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReflectByPtr() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReflectModelFromMap(t *testing.T) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReflectModelFromMap(tt.args.model, tt.args.object)
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

func TestReflectModelPtrFromMap(t *testing.T) {
	type args struct {
		modelPtr interface{}
		object   map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ReflectModelPtrFromMap(tt.args.modelPtr, tt.args.object); (err != nil) != tt.wantErr {
				t.Errorf("ReflectModelPtrFromMap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReflectRetColsValues(t *testing.T) {
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
			cols, args := ReflectRetColsValues(tt.args.model)
			assert.Equal(t, []string{"id", "name"}, cols)
			assert.Equal(t, []interface{}{1, "test_name"}, args)
		})
	}
}
