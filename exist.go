package lightweight_db

import (
	"fmt"
	"github.com/ssst0n3/awesome_libs/awesome_error"
)

func (c Connector) IsResourceExistsById(tableName string, id int64) (bool, error) {
	var result int
	query := fmt.Sprintf("SELECT COUNT(id) FROM %s WHERE id=?", tableName)
	if err := c.DB.QueryRow(query, id).Scan(&result); err != nil {
		awesome_error.CheckErr(err)
		return false, err
	}
	if result > 0 {
		return true, nil
	}
	return false, nil
}

func (c Connector) IsResourceExistsByGuid(tableName string, guidColName, guidValue interface{}) (bool, error) {
	var result int
	query := fmt.Sprintf("SELECT COUNT(id) FROM %s WHERE %s=?", tableName, guidColName)
	if err := c.QueryRow(query, &result, guidValue); err != nil {
		awesome_error.CheckErr(err)
		return false, err
	}

	if result > 0 {
		return true, nil
	}
	return false, nil
}

func (c Connector) IsResourceExistsExceptSelfByGuid(tableName string, guidColName string, guidValue interface{}, id int64) (bool, error) {
	var result int
	query := fmt.Sprintf("SELECT COUNT(id) FROM %s WHERE %s=? AND id != ?", tableName, guidColName)
	if err := c.DB.QueryRow(query, guidValue, id).Scan(&result); err != nil {
		awesome_error.CheckErr(err)
		return false, err
	}
	Logger.Debugf("in function IsResourceNameExists, count: %#v", result)
	if result > 0 {
		return true, nil
	}
	return false, nil
}
