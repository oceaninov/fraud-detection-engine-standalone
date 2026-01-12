package databaseConnector

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type MySqlOption struct {
	ConnectionString                     string
	MaxLifeTimeConnection                time.Duration
	MaxIdleTimeConnection                time.Duration
	SlowThreshold                        time.Duration
	MaxIdleConnection, MaxOpenConnection int
	SkipDefaultTransaction               bool
	PrepareStmt                          bool
}

func NewMySql(option *MySqlOption) (*gorm.DB, error) {
	var (
		opts = &gorm.Config{
			QueryFields:            true,
			SkipDefaultTransaction: option.SkipDefaultTransaction,
			PrepareStmt:            option.PrepareStmt,
			Logger:                 logger.Default.LogMode(logger.Info),
		}
	)
	db, err := gorm.Open(mysql.Open(option.ConnectionString), opts)
	if err != nil {
		return nil, err
	}
	sql, err := db.DB()
	if err != nil {
		return nil, err
	}
	sql.SetConnMaxIdleTime(option.MaxIdleTimeConnection)
	sql.SetConnMaxLifetime(option.MaxLifeTimeConnection)
	sql.SetMaxOpenConns(option.MaxOpenConnection)
	sql.SetMaxIdleConns(option.MaxIdleConnection)
	return db, nil
}

type PostgresOption struct {
	ConnectionString                     string
	MaxLifeTimeConnection                time.Duration
	MaxIdleTimeConnection                time.Duration
	SlowThreshold                        time.Duration
	MaxIdleConnection, MaxOpenConnection int
	SkipDefaultTransaction               bool
	PrepareStmt                          bool
}

func NewPostgres(option *PostgresOption) (*gorm.DB, error) {
	var (
		opts = &gorm.Config{
			SkipDefaultTransaction: option.SkipDefaultTransaction,
			PrepareStmt:            option.PrepareStmt,
			Logger:                 logger.Default.LogMode(logger.Info),
		}
	)
	db, err := gorm.Open(postgres.Open(option.ConnectionString), opts)
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetConnMaxIdleTime(option.MaxIdleTimeConnection)
	sqlDB.SetConnMaxLifetime(option.MaxLifeTimeConnection)
	sqlDB.SetMaxOpenConns(option.MaxOpenConnection)
	sqlDB.SetMaxIdleConns(option.MaxIdleConnection)

	return db, nil
}
