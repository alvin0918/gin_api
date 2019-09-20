package HomeModel

import (
	ORM "github.com/alvin0918/gin_api/core/orm"
	"log"
)

var (
	navTableName string = "luffy_nav"
)

type Nav struct {
	Id          int    `json:"id"`
	Orders      int    `json:"orders"`
	IsShow      int8   `json:"is_show"`
	IsDelete    int8   `json:"is_delete"`
	Name        string `json:"name"`
	Link        string `json:"link"`
	Opt         int8   `json:"opt"`
	CreatedTime string `json:"created_time"`
	UpdatedTime string `json:"updated_time"`
}

/**
获取导航
*/
func GetNav() (res map[int]map[string]string, code string, err error) {

	var (
		result map[string]string
	)

	if res, err = ORM.DBConfig.TableName(navTableName).
		Where("opt = 0", "").
		Select(); err != nil {
		log.Fatal(err)
	}

	if result, err = ORM.DBConfig.TableName(navTableName).
		Field("count(*) as num").
		Where("opt = 0", "").
		Find(); err != nil {
		log.Fatal(err)
	}

	code = result["num"]

	return
}

/**
获取友情链接
*/
func GetFooter() (res map[int]map[string]string, code string, err error) {

	var (
		result map[string]string
	)

	if res, err = ORM.DBConfig.TableName(navTableName).
		Where("opt = 1", "").
		Select(); err != nil {
		log.Fatal(err)
	}

	if result, err = ORM.DBConfig.TableName(navTableName).
		Field("count(*) as num").
		Where("opt = 1", "").
		Find(); err != nil {
		log.Fatal(err)
	}

	code = result["num"]

	return
}
