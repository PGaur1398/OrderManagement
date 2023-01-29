package repository

import (
	"OrderManagement/config"
	"context"
	"database/sql"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type gormClient struct {
	db *gorm.DB
}

func (gClient gormClient) Query(ctx context.Context, cmd Command, args ...interface{}) (*sql.Rows, error) {
	return gClient.db.Raw(cmd.GetQuery(), args...).Rows()

}

func (gClient gormClient) Exec(ctx context.Context, cmd Command, args ...interface{}) *gorm.DB {
	return gClient.db.Exec(cmd.GetQuery(), args...)
}

func (gClient gormClient) Insert(ctx context.Context, tableName string, insertData interface{}) *gorm.DB {
	return gClient.db.Table(tableName).Create(insertData)

}

func (gClient gormClient) Update(ctx context.Context, tableName string, updateData interface{}) *gorm.DB {
	return gClient.db.Table(tableName).Save(updateData)
}

func GetGORMDBConnection(dbcfg config.DatabaseConfig) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", dbcfg.GetDBUrl())
	if err != nil {
		return nil, err
	}
	return db, nil

}
func NewGORMClient(db *gorm.DB) Client {
	return gormClient{db}
}
