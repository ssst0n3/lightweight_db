package lightweight_db

import (
	"fmt"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"strings"
)

/*
!!!reflect attention, may cause panic!!!
model can be struct, can also be pointer(reference)
*/
func (c Connector) CreateObject(tableName string, model interface{}) (int64, error) {
	cols, args := RetColsValues(model)
	query := fmt.Sprintf("INSERT INTO %s (`%s`) VALUES (%s)", tableName, strings.Join(cols, "`,`"), strings.Repeat("?,", len(cols))[:2*len(cols)-1])
	res, err := c.Exec(query, args...)
	if err != nil {
		awesome_error.CheckErr(err)
		return -1, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		awesome_error.CheckErr(err)
		return -1, err
	}
	return id, nil
}
