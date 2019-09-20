package ORM

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/alvin0918/gin_api/core/config"
	_ "github.com/go-sql-driver/mysql"
	"html"
	"strings"
)

type MysqlConfig struct {
	where     string
	orderBy   string
	groupBy   string
	join      string
	field     string
	tableName string
	alias     string
	isSql     bool
	query     string
}

var DBConfig *MysqlConfig = &MysqlConfig{ isSql:false,}

/**
连接数据库
*/
func conn() (db *sql.DB, err error) {

	// 连接MySQL
	if db, err = sql.Open("mysql", config.G_config.ApiDatabaseConf); err != nil {
		panic(errors.New("数据库连接失败！原因是：" + err.Error()))
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	err = db.Ping()

	return
}

/**
设置WHERE条件，支持链式重复调用
*/
func (MysqlDBConfig *MysqlConfig) Where(str string, mode string) *MysqlConfig {
	if MysqlDBConfig.where == "" {
		MysqlDBConfig.where = " WHERE " + str
	} else {
		MysqlDBConfig.where += " " + mode + " " + str
	}

	return MysqlDBConfig
}

/**
设置查询字段，支持链式，支持重复调用
*/
func (MysqlDBConfig *MysqlConfig) Field(str string) *MysqlConfig {
	if MysqlDBConfig.field == "" {
		MysqlDBConfig.field = str
	} else {
		MysqlDBConfig.field = ", " + str
	}

	return MysqlDBConfig
}

/**
设置排序条件，支持链式，不支持重复调用
*/
func (MysqlDBConfig *MysqlConfig) OrderBy(str string, mode string) *MysqlConfig {

	MysqlDBConfig.orderBy = str + " " + mode

	return MysqlDBConfig
}

/**
设置数据表
*/
func (MysqlDBConfig *MysqlConfig) TableName(str string) *MysqlConfig {

	MysqlDBConfig.tableName = str

	return MysqlDBConfig
}

/**
设置数据表
*/
func (MysqlDBConfig *MysqlConfig) Alias(str string) *MysqlConfig {

	MysqlDBConfig.alias = str

	return MysqlDBConfig
}

/**
设置分组排序，支持链式调用，不支持重复使用
*/
func (MysqlDBConfig *MysqlConfig) GroupBy(str string, mode string) *MysqlConfig {

	MysqlDBConfig.orderBy = str + " " + mode

	return MysqlDBConfig
}

/**
连表操作，支持链式，支持重复调用
*/
func (MysqlDBConfig *MysqlConfig) Join(str string, mode string) *MysqlConfig {

	switch strings.ToLower(mode) {
	case "inner":
		MysqlDBConfig.join = " INNER JOIN " + str
	case "right":
		MysqlDBConfig.join = " RIGHT JOIN " + str
	case "left":
		MysqlDBConfig.join = " LEFT JOIN " + str
	default:
		panic(errors.New("Abnormal parameter！"))
	}

	return MysqlDBConfig
}

/**
是否打印SQL
 */
func (MysqlDBConfig *MysqlConfig) IsPrintSql(mode bool) *MysqlConfig {

	MysqlDBConfig.isSql = mode

	return MysqlDBConfig
}

/**
查询数据， 返回结果map
*/
func (MysqlDBConfig *MysqlConfig) Select() (result map[int]map[string]string, err error) {

	var (
		query string
		rows  *sql.Rows
		str   string
		cols  []string
		val   [][]byte
		scans []interface{}
		i     int
		row   map[string]string
		db    *sql.DB
	)

	defer func() {
		err = db.Close()
	}()

	// 每次调用首相初始化一个连接
	if db, err = conn(); err != nil {
		panic(errors.New(err.Error()))
	}

	str = "SELECT"

	// 获取SQL语句
	query = MysqlDBConfig.analysisSql(str)

	if rows, err = db.Query(query); err != nil {
		panic(errors.New("查询失败!" + err.Error()))
	}

	// 查出字段
	if cols, err = rows.Columns(); err != nil {
		panic(errors.New("查询失败!" + err.Error()))
	}

	// 查出每一列的值
	val = make([][]byte, len(cols))

	// rows.Scan()的参数， 因为每次查询出来的列是不定长的，用len(cols)定住每次查询的长度
	scans = make([]interface{}, len(cols))

	// 让每一行数据填充到val中
	for i := range val {
		scans[i] = &val[i]
	}

	// 得到最后的map
	result = make(map[int]map[string]string)

	i = 0

	// 循环游标，向下推移
	for rows.Next() {
		if err = rows.Scan(scans...); err != nil {
			panic(errors.New(err.Error()))
		}

		// 获取每一行的数据
		row = make(map[string]string)

		for k, v := range val {
			key := cols[k]
			row[key] = string(v)
		}

		// 装入结果集中
		result[i] = row

		i++
	}

	DBConfig = &MysqlConfig{isSql:false,}

	return
}

/**
指查询一条
 */
func (MysqlDBConfig *MysqlConfig) Find() (result map[string]string, err error) {

	var (
		rows map[int]map[string]string
	)

	rows, err = MysqlDBConfig.Select()

	if len(rows[0]) > 0 {
		result = rows[0]
	}

	return
}

/**
插入数据 isRows true 返回影响的行数 FALSE 返回最后一行的主键ID
*/
func (MysqlDBConfig *MysqlConfig) Insert(data map[string]string, isRows bool) (rows int64, err error) {

	var (
		query  string
		stmt   *sql.Stmt
		result sql.Result
		str    string
		db     *sql.DB
	)

	defer func() {
		err = db.Close()
	}()

	// 每次调用首相初始化一个连接
	if db, err = conn(); err != nil {
		panic(errors.New(err.Error()))
	}

	str = "INSERT"

	// 获取SQL语句
	query = MysqlDBConfig.analysisSqls(data, str)

	if stmt, err = db.Prepare(query); err != nil {
		panic(errors.New(err.Error()))
	}

	if result, err = stmt.Exec(); err != nil {
		panic(errors.New(err.Error()))
	}

	if isRows {
		if rows, err = result.RowsAffected(); err != nil {
			panic(errors.New(err.Error()))
		}
	} else {
		if rows, err = result.LastInsertId(); err != nil {
			panic(errors.New(err.Error()))
		}
	}

	return

}

/**
修改数据 isRows true 返回影响的行数 FALSE 返回最后一行的主键ID
*/
func (MysqlDBConfig *MysqlConfig) Update(data map[string]string, isRows bool) (rows int64, err error) {

	var (
		query  string
		stmt   *sql.Stmt
		result sql.Result
		str    string
		db     *sql.DB
	)

	defer func() {
		err = db.Close()
	}()

	str = "Upload"

	// 获取SQL语句
	query = MysqlDBConfig.analysisSqls(data, str)

	// 每次调用首相初始化一个连接
	if db, err = conn(); err != nil {
		panic(errors.New(err.Error()))
	}

	if stmt, err = db.Prepare(query); err != nil {
		panic(errors.New(err.Error()))
	}

	if result, err = stmt.Exec(); err != nil {
		panic(errors.New(err.Error()))
	}

	if isRows {
		if rows, err = result.RowsAffected(); err != nil {
			panic(errors.New(err.Error()))
		}
	} else {
		if rows, err = result.LastInsertId(); err != nil {
			panic(errors.New(err.Error()))
		}
	}

	return

}

/**
删除数据 isRows true 返回影响的行数 FALSE 返回最后一行的主键ID
*/
func (MysqlDBConfig *MysqlConfig) Delete(isRows bool) (rows int64, err error) {

	var (
		query  string
		stmt   *sql.Stmt
		result sql.Result
		str    string
		db     *sql.DB
	)

	defer func() {
		err = db.Close()
	}()

	// 每次调用首相初始化一个连接
	if db, err = conn(); err != nil {
		panic(errors.New(err.Error()))
	}

	str = "DELETE"

	// 获取SQL语句
	query = MysqlDBConfig.analysisSql(str)

	if stmt, err = db.Prepare(query); err != nil {
		panic(errors.New(err.Error()))
	}

	if result, err = stmt.Exec(); err != nil {
		panic(errors.New(err.Error()))
	}

	if isRows {
		if rows, err = result.RowsAffected(); err != nil {
			panic(errors.New(err.Error()))
		}
	} else {
		if rows, err = result.LastInsertId(); err != nil {
			panic(errors.New(err.Error()))
		}
	}

	return

}

/**
根据查询模式，获取SQL
*/
func (MysqlDBConfig *MysqlConfig) analysisSql(mode string) (str string) {

	str = strings.ToUpper(mode)

	switch str {
	case "UPDATE":
		if MysqlDBConfig.tableName == "" {
			panic(errors.New("不能没有表名呀兄弟！"))
		}

		str += " " + MysqlDBConfig.tableName + " SET "

		if MysqlDBConfig.field == "" {
			panic(errors.New("需要修改字段及数据"))
		}

		str += " " + MysqlDBConfig.field

		if MysqlDBConfig.where != "" {
			str += " " + MysqlDBConfig.where
		}

	case "DELETE":

		if MysqlDBConfig.tableName == "" {
			panic(errors.New("不能没有表名呀兄弟！"))
		}

		str += " FROM " + MysqlDBConfig.tableName

		if MysqlDBConfig.where != "" {
			str += " " + MysqlDBConfig.where
		} else {
			panic(errors.New("这个操作太危险啦！真那么想不开的话设置成 1 = 1吧！"))
		}

	case "SELECT":

		// 格式化查询字段
		if MysqlDBConfig.field != "" {
			str += " " + MysqlDBConfig.field
		} else {
			str += " * "
		}

		// 设置表名
		if MysqlDBConfig.tableName != "" {
			str += " FROM " + MysqlDBConfig.tableName
		} else {
			panic(errors.New("Can't Find TableName！"))
		}

		// 设置表别名
		if MysqlDBConfig.alias != "" {
			str += " AS " + MysqlDBConfig.alias
		}

		// 格式化查询条件
		if MysqlDBConfig.where != "" {
			str += " " + MysqlDBConfig.where
		}

		// 格式化分组
		if MysqlDBConfig.groupBy != "" {
			str += " " + MysqlDBConfig.groupBy
		}

		// 格式化排序
		if MysqlDBConfig.orderBy != "" {
			str += " " + MysqlDBConfig.orderBy
		}
	default:
		// 执行原生SQL
		return MysqlDBConfig.query
	}

	// SQL语句格式化，简要避免SQL注入
	str = html.EscapeString(str)

	// 是否打印SQL
	if MysqlDBConfig.isSql {
		fmt.Println(str)
	}

	return
}

/**
根据查询模式，获取SQL
*/
func (MysqlDBConfig *MysqlConfig) analysisSqls(data map[string]string, mode string) (str string) {

	str = strings.ToUpper(mode)

	switch str {
	case "INSERT":

		if MysqlDBConfig.tableName == "" {
			panic(errors.New("不能没有表名呀兄弟！"))
		}

		str += " INTO " + MysqlDBConfig.tableName

		var key string = "("
		var value string = "("

		for k, v := range data{
			if key == "(" {
				key += k + ","
			}

			if value == "(" {
				value += v
			}
		}

		key += ")"
		value += ")"

		str += " " + key + " VALUES " + value

		if MysqlDBConfig.where != "" {
			str += " " + MysqlDBConfig.where
		}

	case "UPDATE":
		if MysqlDBConfig.tableName == "" {
			panic(errors.New("不能没有表名呀兄弟！"))
		}

		str += " " + MysqlDBConfig.tableName + " SET "

		for k, v := range data{
			str += " " + k + " = " + v + ","
		}

		if MysqlDBConfig.where != "" {
			str += " " + MysqlDBConfig.where
		}

	default:
		// 执行原生SQL
		return MysqlDBConfig.query
	}

	// SQL语句格式化，简要避免SQL注入
	str = html.EscapeString(str)

	// 是否打印SQL
	if MysqlDBConfig.isSql {
		fmt.Println(str)
	}

	return
}
