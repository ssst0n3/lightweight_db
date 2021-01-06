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
		return -1, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		awesome_error.CheckErr(err)
		return -1, err
	}
	return id, nil
}

func CheckErrorDuplicate(err error) bool {
	e := strings.ToLower(err.Error())
	return strings.Contains(e, "duplicate") || strings.Contains(e, "unique constraint")
}

func (c Connector) CreateObjectPreventDuplicate(tableName string, model interface{}) (exists bool, id int64, err error) {
	//	https://www.mysqltutorial.org/mysql-unique-constraint/
	// Error Code: 1062. Duplicate entry 'ABC Inc-4000 North 1st Street' for key 'uc_name_address'

	// https://www.sqlitetutorial.net/sqlite-unique-constraint/
	// https://www.sqlitetutorial.net/sqlite-unique-constraint/
	id, err = c.CreateObject(tableName, model)
	if err != nil {
		awesome_error.CheckWarning(err)
		if CheckErrorDuplicate(err) {
			exists = true
			err = nil
		}
	}
	return
}