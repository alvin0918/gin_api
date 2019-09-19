package userModel

import (
	"fmt"
	ORM "github.com/alvin0918/gin_api/core/orm"
)

var (
	TableName  string = "luffy_user"
	PrimaryKey string = "id"
)

type User struct {
	ID          int    `json:"id"`
	Password    string `json:"password"`
	LastLogin   string `json:"last_login"`
	IsSuperuser string `json:"is_superuser"`
	Username    string `json:"username"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsStaff     string `json:"is_staff"`
	IsActive    string `json:"is_active"`
	DateJoined  string `json:"date_joined"`
	Mobile      string `json:"mobile"`
	Avatar      string `json:"avatar"`
}

func (user *User) CheckUserAndPassword() (data map[string]string, err error) {

	data = make(map[string]string)

	if data, err = ORM.DBConfig.TableName(TableName).
		Where(fmt.Sprintf("username = %s", user.Username), "and").
		Where(fmt.Sprintf("password = %s", user.Password), "and").
		IsPrintSql(true).
		Find(); err != nil {
		return
	}

	return
}
