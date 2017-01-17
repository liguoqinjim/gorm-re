package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
)

type Config struct {
	DBHost string
	DBUser string
	DBPwd  string
	DBName string
}

var config Config

func LoadConfig(data []byte) {
	err := json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}
	if config.DBName == "DBName" {
		log.Fatal("请修改配置文件中的数据库配置")
	}
}

type Column struct {
	TableCataLog           string
	TableSchema            string
	TableName              string
	ColumnName             string
	OrdinalPosition        string
	ColumnDefault          string
	IsNullable             string
	DataType               string
	CharacterMaximumLength string
	CharacterOctetLength   string
	NumericPrecision       string
	NumericScale           string
	CharacterSetName       string
	CollationName          string
	ColumnType             string
	ColumnKey              string
	Extra                  string
	Privileges             string
	ColumnComment          string
}

func GetColumns() []*Column {
	connectInfo := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", config.DBUser, config.DBPwd, config.DBHost, "information_schema")
	db, err := sql.Open("mysql", connectInfo)
	if err != nil {
		log.Fatal(err)
	}

	//ping
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//select
	sql := "select * from `COLUMNS` where TABLE_SCHEMA = ? order by TABLE_NAME,ORDINAL_POSITION"
	rows, err := db.Query(sql, config.DBName)
	if err != nil {
		log.Fatal(err)
	}

	columns := make([]*Column, 0)
	for rows.Next() {
		column := new(Column)
		rows.Scan(column.TableCataLog, column.TableSchema, column.TableName, column.ColumnName, column.OrdinalPosition,
			column.ColumnDefault, column.IsNullable, column.DataType, column.CharacterMaximumLength, column.CharacterOctetLength,
			column.NumericPrecision, column.NumericScale, column.CharacterSetName, column.CollationName, column.ColumnType, column.ColumnKey,
			column.Extra, column.Privileges, column.ColumnComment)
		columns = append(columns, column)
	}

	fmt.Println("length =", len(columns))

	return nil
}

func main() {
	//读取config.json
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	//读取配置文件
	LoadConfig(data)

	//连接数据库，查询columns表
	GetColumns()
}
