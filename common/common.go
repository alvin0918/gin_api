package common

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

func GetPostJsonData(c *gin.Context) (mapResult map[string]interface{}) {

	var (
		data []byte
		err  error
	)

	if data, err = ioutil.ReadAll(c.Request.Body); err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(data, &mapResult); err != nil {
		fmt.Println("JsonToMapDemo err: ", err)
	}

	return
}

/**
JSON 转 Map
*/
func JsonToMapDemo(jsonStr []byte) (mapResult map[string]string) {

	var (
		err error
	)

	if err = json.Unmarshal(jsonStr, &mapResult); err != nil {
		fmt.Println("JsonToMapDemo err: ", err)
	}

	return
}

