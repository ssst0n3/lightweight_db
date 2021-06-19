package lightweight_db

import "github.com/ssst0n3/awesome_libs/awesome_error"

func CreateTable(query string, c Connector) (err error) {
	statement, err := c.DB.Prepare(query)
	if err != nil {
		awesome_error.CheckFatal(err)
		return
	}
	_, err = statement.Exec()
	awesome_error.CheckErr(err)
	return
}
