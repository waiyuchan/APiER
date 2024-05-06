package model

import (
	"apier/internal/global/errors"
	"apier/internal/global/variable"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type BaseModel struct {
	*gorm.DB  `gorm:"-" json:"-"`
	Id        int64  `gorm:"primaryKey" json:"id"`
	CreatedAt string `gorm:"-" json:"created_at,omitempty"` //日期时间字段统一设置为字符串即可
	UpdatedAt string `gorm:"-" json:"updated_at,omitempty"`
}

func UseDbConn(sqlType string) *gorm.DB {
	fmt.Println("use database connection")
	var db *gorm.DB
	sqlType = strings.Trim(sqlType, " ")
	if sqlType == "" {
		//sqlType = variable.ConfigGormYaml.GetString("Gorm.UseDbType")
		sqlType = "mysql"
	}
	if variable.GormDbMysql == nil {
		variable.ZapLog.Fatal(fmt.Sprintf(errors.ErrorsGormNotInitGlobalPointer, sqlType, sqlType))
	}
	db = variable.GormDbMysql

	//switch strings.ToLower(sqlType) {
	//case "mysql":
	//	if variable.GormDbMysql == nil {
	//		variable.ZapLog.Fatal(fmt.Sprintf(errors.ErrorsGormNotInitGlobalPointer, sqlType, sqlType))
	//	}
	//	db = variable.GormDbMysql
	//case "sqlserver":
	//
	//case "postgres", "postgre", "postgresql":
	//
	//default:
	//	variable.ZapLog.Error(errors.ErrorsDbDriverNotExists + sqlType)
	//}
	return db
}
