package lightweight_db

import (
	"database/sql"
	"reflect"
	"testing"
)

func TestConnector_Close(t *testing.T) {
	type fields struct {
		DriverName string
		Dsn        string
		db         *sql.DB
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Connector{
				DriverName: tt.fields.DriverName,
				Dsn:        tt.fields.Dsn,
				DB:         tt.fields.db,
			}
			if err := c.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConnector_CreateObject(t *testing.T) {
	type fields struct {
		DriverName string
		Dsn        string
		db         *sql.DB
	}
	type args struct {
		tableName string
		model     interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Connector{
				DriverName: tt.fields.DriverName,
				Dsn:        tt.fields.Dsn,
				DB:         tt.fields.db,
			}
			got, err := c.CreateObject(tt.args.tableName, tt.args.model)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateObject() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnector_DeleteObjectById(t *testing.T) {
	type fields struct {
		DriverName string
		Dsn        string
		db         *sql.DB
	}
	type args struct {
		tableName string
		id        uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Connector{
				DriverName: tt.fields.DriverName,
				Dsn:        tt.fields.Dsn,
				DB:         tt.fields.db,
			}
			if err := c.DeleteObjectById(tt.args.tableName, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteObjectById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConnector_Exec(t *testing.T) {
	type fields struct {
		DriverName string
		Dsn        string
		db         *sql.DB
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    sql.Result
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Connector{
				DriverName: tt.fields.DriverName,
				Dsn:        tt.fields.Dsn,
				DB:         tt.fields.db,
			}
			got, err := c.Exec(tt.args.query, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Exec() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnector_IsObjectIdExists(t *testing.T) {
	type fields struct {
		DriverName string
		Dsn        string
		db         *sql.DB
	}
	type args struct {
		tableName string
		id        uint64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Connector{
				DriverName: tt.fields.DriverName,
				Dsn:        tt.fields.Dsn,
				DB:         tt.fields.db,
			}
			if got := c.IsObjectIdExists(tt.args.tableName, tt.args.id); got != tt.want {
				t.Errorf("IsObjectIdExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnector_IsResourceExists(t *testing.T) {
	type fields struct {
		DriverName string
		Dsn        string
		db         *sql.DB
	}
	type args struct {
		tableName string
		colName   string
		content   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Connector{
				DriverName: tt.fields.DriverName,
				Dsn:        tt.fields.Dsn,
				DB:         tt.fields.db,
			}
			got, err := c.IsResourceExists(tt.args.tableName, tt.args.colName, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsResourceExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsResourceExists() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnector_IsResourceNameExists(t *testing.T) {
	type fields struct {
		DriverName string
		Dsn        string
		db         *sql.DB
	}
	type args struct {
		tableName   string
		guidColName string
		guidValue   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Connector{
				DriverName: tt.fields.DriverName,
				Dsn:        tt.fields.Dsn,
				DB:         tt.fields.db,
			}
			got, err := c.IsResourceNameExists(tt.args.tableName, tt.args.guidColName, tt.args.guidValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsResourceNameExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsResourceNameExists() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnector_IsResourceNameExistsExceptSelf(t *testing.T) {
	type fields struct {
		DriverName string
		Dsn        string
		db         *sql.DB
	}
	type args struct {
		tableName   string
		guidColName string
		guidValue   string
		id          uint64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Connector{
				DriverName: tt.fields.DriverName,
				Dsn:        tt.fields.Dsn,
				DB:         tt.fields.db,
			}
			if got := c.IsResourceNameExistsExceptSelf(tt.args.tableName, tt.args.guidColName, tt.args.guidValue, tt.args.id); got != tt.want {
				t.Errorf("IsResourceNameExistsExceptSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnector_ListAllPropertiesByTableName(t *testing.T) {
	type fields struct {
		DriverName string
		Dsn        string
		db         *sql.DB
	}
	type args struct {
		tableName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Connector{
				DriverName: tt.fields.DriverName,
				Dsn:        tt.fields.Dsn,
				DB:         tt.fields.db,
			}
			got, err := c.ListAllPropertiesByTableName(tt.args.tableName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListAllPropertiesByTableName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAllPropertiesByTableName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnector_ListObjects(t *testing.T) {
	type fields struct {
		DriverName string
		Dsn        string
		db         *sql.DB
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Connector{
				DriverName: tt.fields.DriverName,
				Dsn:        tt.fields.Dsn,
				DB:         tt.fields.db,
			}
			got, err := c.ListObjects(tt.args.query, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListObjects() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListObjects() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnector_Query(t *testing.T) {
	type fields struct {
		DriverName string
		Dsn        string
		db         *sql.DB
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *sql.Rows
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Connector{
				DriverName: tt.fields.DriverName,
				Dsn:        tt.fields.Dsn,
				DB:         tt.fields.db,
			}
			got, err := c.Query(tt.args.query, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Query() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnector_ShowObjectById(t *testing.T) {
	type fields struct {
		DriverName string
		Dsn        string
		db         *sql.DB
	}
	type args struct {
		tableName string
		id        uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Connector{
				DriverName: tt.fields.DriverName,
				Dsn:        tt.fields.Dsn,
				DB:         tt.fields.db,
			}
			got, err := c.ShowObjectById(tt.args.tableName, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShowObjectById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShowObjectById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnector_ShowObjectOnePropertyById(t *testing.T) {
	type fields struct {
		DriverName string
		Dsn        string
		db         *sql.DB
	}
	type args struct {
		tableName  string
		columnName string
		id         uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Connector{
				DriverName: tt.fields.DriverName,
				Dsn:        tt.fields.Dsn,
				DB:         tt.fields.db,
			}
			got, err := c.ShowObjectOnePropertyById(tt.args.tableName, tt.args.columnName, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShowObjectOnePropertyById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShowObjectOnePropertyById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnector_UpdateObject(t *testing.T) {
	type fields struct {
		DriverName string
		Dsn        string
		db         *sql.DB
	}
	type args struct {
		id        uint64
		tableName string
		model     interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Connector{
				DriverName: tt.fields.DriverName,
				Dsn:        tt.fields.Dsn,
				DB:         tt.fields.db,
			}
			if err := c.UpdateObject(tt.args.id, tt.args.tableName, tt.args.model); (err != nil) != tt.wantErr {
				t.Errorf("UpdateObject() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConnector_UpdateObjectSingleColumnById(t *testing.T) {
	type fields struct {
		DriverName string
		Dsn        string
		db         *sql.DB
	}
	type args struct {
		id         uint64
		tableName  string
		columnName string
		value      interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Connector{
				DriverName: tt.fields.DriverName,
				Dsn:        tt.fields.Dsn,
				DB:         tt.fields.db,
			}
			if err := c.UpdateObjectSingleColumnById(tt.args.id, tt.args.tableName, tt.args.columnName, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("UpdateObjectSingleColumnById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFetchOneRow(t *testing.T) {
	type args struct {
		rows *sql.Rows
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FetchOneRow(tt.args.rows)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchOneRow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchOneRow() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFetchRows(t *testing.T) {
	type args struct {
		rows *sql.Rows
	}
	tests := []struct {
		name    string
		args    args
		want    []map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FetchRows(tt.args.rows)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchRows() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchRows() got = %v, want %v", got, tt.want)
			}
		})
	}
}