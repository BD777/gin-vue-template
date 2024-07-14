package models

import (
	"fmt"
	"gin-vue-template/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Model struct {
	ID         int64 `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt  int64 `gorm:"not null" json:"createdAt"`
	ModifiedAt int64 `gorm:"not null" json:"modifiedAt"`
}

type GormDAO struct {
	db *gorm.DB
}

func NewDAO(cfg *config.Config) *GormDAO {
	var db *gorm.DB
	var (
		err                          error
		dbName, user, password, host string
	)

	dbName = cfg.Mysql.Name
	user = cfg.Mysql.User
	password = cfg.Mysql.Password
	host = cfg.Mysql.Host

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		logrus.Fatalf("gorm open error: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		logrus.Fatalf("get db error: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	initTables(db)

	return &GormDAO{db: db}
}

func initTables(db *gorm.DB) {
	db.AutoMigrate(&Admin{})
}

type Pagination struct {
	Page     int64  `form:"page" json:"page"`
	PageSize int64  `form:"pageSize" json:"pageSize"`
	Total    int64  `form:"total" json:"total"`
	Cursor   *int64 `form:"cursor" json:"cursor"` // after cursor
}

type ListRequest struct {
	QueryCount bool        `form:"queryCount" json:"queryCount"`
	Pagination *Pagination `form:"pagination" json:"pagination"`
}

type ListResponse[T any] struct {
	List       []T         `json:"list"`
	Pagination *Pagination `json:"pagination"`
}
