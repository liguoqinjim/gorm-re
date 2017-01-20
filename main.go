package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Config struct {
	DBHost string
	DBUser string
	DBPwd  string
	DBName string

	ModelFileName string
	PackageName   string

	Mysql8 string
}

var myConfig Config

func LoadConfig(data []byte) {
	err := json.Unmarshal(data, &myConfig)
	if err != nil {
		log.Fatal(err)
	}
	if myConfig.DBName == "DBName" {
		log.Fatal("请修改配置文件中的数据库配置")
	}
}

type Column struct {
	TableCataLog           sql.NullString
	TableSchema            sql.NullString
	TableName              sql.NullString
	ColumnName             sql.NullString
	OrdinalPosition        sql.NullString
	ColumnDefault          sql.NullString
	IsNullable             sql.NullString
	DataType               sql.NullString
	CharacterMaximumLength sql.NullString
	CharacterOctetLength   sql.NullString
	NumericPrecision       sql.NullString
	NumericScale           sql.NullString
	DatetimePrecision      sql.NullString
	CharacterSetName       sql.NullString
	CollationName          sql.NullString
	ColumnType             sql.NullString
	ColumnKey              sql.NullString
	Extra                  sql.NullString
	Privileges             sql.NullString
	ColumnComment          sql.NullString
	GenerationExpression   sql.NullString
}

func GetColumns() []*Column {
	connectInfo := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", myConfig.DBUser, myConfig.DBPwd, myConfig.DBHost, "information_schema")
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
	querySql := "select * from `COLUMNS` where TABLE_SCHEMA = ? order by TABLE_NAME,ORDINAL_POSITION"
	rows, err := db.Query(querySql, myConfig.DBName)
	if err != nil {
		log.Fatal(err)
	}

	//columnNames, err := rows.Columns()
	//if err != nil {
	//	log.Fatal(err)
	//}

	columns := make([]*Column, 0)

	for rows.Next() {
		column := new(Column)
		if myConfig.Mysql8 == "false" {
			err = rows.Scan(&column.TableCataLog, &column.TableSchema, &column.TableName, &column.ColumnName, &column.OrdinalPosition,
				&column.ColumnDefault, &column.IsNullable, &column.DataType, &column.CharacterMaximumLength, &column.CharacterOctetLength,
				&column.NumericPrecision, &column.NumericScale, &column.CharacterSetName, &column.CollationName, &column.ColumnType, &column.ColumnKey,
				&column.Extra, &column.Privileges, &column.ColumnComment)
		} else {
			err = rows.Scan(&column.TableCataLog, &column.TableSchema, &column.TableName, &column.ColumnName, &column.OrdinalPosition,
				&column.ColumnDefault, &column.IsNullable, &column.DataType, &column.CharacterMaximumLength, &column.CharacterOctetLength,
				&column.NumericPrecision, &column.NumericScale, &column.DatetimePrecision, &column.CharacterSetName, &column.CollationName, &column.ColumnType, &column.ColumnKey,
				&column.Extra, &column.Privileges, &column.ColumnComment, &column.GenerationExpression)
		}
		if err != nil {
			log.Fatal(err)
		}

		columns = append(columns, column)
	}

	return columns
}

func GenerateStructs(modelFile *os.File, columns []*Column) { //逆向工程所有的表
	tableName := ""

	//写入package名字
	packageName := fmt.Sprintf("package %s\n", myConfig.PackageName)
	modelFile.WriteString(packageName)

	//写每个Struct
	var tableColumns []*Column
	for _, v := range columns {
		if v.TableName.String != tableName { //新的一个表
			if tableName != "" {
				GenerateStruct(modelFile, tableColumns) //生成Struct
			}
			tableName = v.TableName.String
			tableColumns = make([]*Column, 0)
		}
		tableColumns = append(tableColumns, v)
	}
}
func GenerateStruct(modelFile *os.File, columns []*Column) string { //逆向工程一个表
	structName := GetStructName(columns[0].TableName.String)

	structContent := fmt.Sprintf("type %s struct{\n", structName)

	for _, v := range columns {
		fieldContent := GetField(v)
		structContent += fmt.Sprintf("%s\n", fieldContent)
	}

	structContent += "}\n"

	modelFile.WriteString(structContent)

	//写Struct对应的表命 (gorm中的TableName())
	structTableName := GetStructTableName(structName, columns[0].TableName.String)
	modelFile.WriteString(structTableName + "\n")

	return structContent
}

func GetStructTableName(structName, tableName string) string {
	return fmt.Sprintf("func (%s) TableName() string {\n return \"%s\"\n }\n", structName, tableName)
}

func GetStructName(tableName string) string { //表名转换到类名
	tableName = strings.Replace(tableName, "t_", "", 1)
	names := strings.Split(tableName, "_")

	structName := ""
	for _, v := range names {
		structName = fmt.Sprintf("%s%s", structName, strings.Title(v))
	}

	return structName
}

func GetField(column *Column) string {
	fieldContent := ""

	fieldName := GetFieldName(column.ColumnName.String)
	fieldContent = fmt.Sprintf("%s %s", fieldContent, fieldName)

	fieldType := GetFieldType(column)
	fieldContent = fmt.Sprintf("%s %s", fieldContent, fieldType)

	fieldTag := GetFieldTag(column)
	fieldContent = fmt.Sprintf("%s %s", fieldContent, fieldTag)

	return fieldContent
}
func GetFieldName(columnName string) string {
	names := strings.Split(columnName, "_")
	fieldName := ""
	for _, v := range names {
		fieldName = fmt.Sprintf("%s%s", fieldName, strings.Title(v))
	}

	return fieldName
}
func GetFieldType(column *Column) string {
	switch column.DataType.String {
	case "varchar":
		return "string"
	case "char":
		return "string"
	default:
		return column.DataType.String
	}
}
func GetFieldTag(column *Column) string {
	tag := "`gorm:\""
	tags := []string{}

	//添加列名
	tags = append(tags, "column:"+column.ColumnName.String)

	//添加type
	tags = append(tags, "type:"+column.ColumnType.String)

	//添加主键
	switch column.ColumnKey.String {
	case "PRI":
		tags = append(tags, "primary_key")
	}

	//添加自增
	switch column.Extra.String {
	case "auto_increment":
		tags = append(tags, "AUTO_INCREMENT")
	}

	//添加varchar,char长度
	if column.DataType.String == "varchar" || column.DataType.String == "char" {
		tags = append(tags, "size:"+column.CharacterMaximumLength.String)
	}

	//添加not null
	if column.IsNullable.String == "NO" {
		tags = append(tags, "not null")
	}

	//添加default
	if column.ColumnDefault.Valid {
		tags = append(tags, "default:"+column.ColumnDefault.String)
	}

	for n, v := range tags {
		tag += v
		if n != len(tags)-1 {
			tag += ";"
		}
	}

	tag += "\"`"
	return tag
}

func GetFileName() string { //得到要生成的go文件的文件名
	if myConfig.ModelFileName == "" {
		return "model.go"
	} else {
		if !strings.Contains(myConfig.ModelFileName, ".go") {
			return fmt.Sprintf("%s.go", myConfig.ModelFileName)
		}
	}
	return myConfig.ModelFileName
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
	columns := GetColumns()

	if len(columns) > 0 {
		//新建文件
		f, err := os.Create(GetFileName())
		if err != nil {
			log.Fatal(err)
		}

		//开始生成struct
		GenerateStructs(f, columns)

		err = f.Sync()
		if err != nil {
			log.Fatal(err)
		}

		f.Close()
	}

	//go fmt
	goFmtCmd := exec.Command("go", "fmt", GetFileName())
	_, err = goFmtCmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("生成成功")
}
