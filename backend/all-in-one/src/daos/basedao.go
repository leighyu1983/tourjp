package daos

import (
	"entities"
	"database/sql"
	// go get github.com/go-sql-driver/mysql
	_"github.com/go-sql-driver/mysql"
	"utils"
)

/*
CREATE TRIGGER trigger_created_on BEFORE UPDATE 
ON shops FOR EACH ROW 
BEGIN
SET created_on = NOW();
END
*/

var (
	db *sql.DB
	err error
    config  entities.ConfigEntity
)

func initDb() {
	if config.MysqlUrl == "" {
		config, err = util.GetConfig();
	}
	
	if err != nil {
		panic(err.Error())
	}

	// "root:123456@tcp(192.168.121.186:3306)/tour_jp?charset=utf8"
	db, err = sql.Open("mysql", config.MysqlUrl)
	db.SetMaxOpenConns(config.MysqlMaxConnection)
    db.SetMaxIdleConns(10)
    db.Ping()
}


//插入
//func insert(db *sql.DB, sqlstr string, args ...interface{}) (int64, error) {
func InsertUpdate(sqlstr string, args ...interface{}) {
	
	initDb()
	
	stmtIns, err := db.Prepare(sqlstr)
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(args...)
	if err != nil {
		panic(err.Error())
	}
}


func FetchRows(sqlstr string, args ...interface{}) (*[]map[string]string, error) {
	initDb()
	
	if err != nil {
		panic(err.Error())
	}

	stmtOut, err := db.Prepare(sqlstr)
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()
 
	rows, err := stmtOut.Query(args...)
	if err != nil {
		panic(err.Error())
	}
 
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}
 
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
 
	ret := make([]map[string]string, 0)
	for i := range values {
		scanArgs[i] = &values[i]
	}
 
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		var value string
		vmap := make(map[string]string, len(scanArgs))
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			vmap[columns[i]] = value
		}
		ret = append(ret, vmap)
	}
	return &ret, nil
}