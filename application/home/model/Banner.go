package HomeModel

import (
	ORM "github.com/alvin0918/gin_api/core/orm"
	"log"
)

var (
	bannerTableName string = "luffy_banner"
)

type Banner struct {
	Id          int    `json:"id"`
	Orders      int    `json:"orders"`
	IsShow      int8   `json:"is_show"`
	IsDelete    int8   `json:"is_delete"`
	Name        string `json:"name"`
	Link        string `json:"link"`
	Note        string `json:"note"`
	Image       string `json:"image"`
	CreatedTime string `json:"created_time"`
	UpdatedTime string `json:"updated_time"`
}

/**
获取导航
*/
func GetBanner() (res map[int]map[string]string, code string, err error) {

	var (
		result map[string]string
	)

	if res, err = ORM.DBConfig.TableName(bannerTableName).Select(); err != nil {
		log.Fatal(err)
	}

	if result, err = ORM.DBConfig.TableName(bannerTableName).Field("count(*) as num").Find(); err != nil {
		log.Fatal(err)
	}

	code = result["num"]

	return
}
