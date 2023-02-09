package database

import (
	"database/sql"
	"time"

	_ "gitlab.id.vin/gami/go-agent/v3/integrations/nrmysql"
	"gitlab.id.vin/gami/ps2-gami-common/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

// DBAdapterV2 interface represent adapter connect to DB
type DBAdapterV2 interface {
	Open(connectionString string, config gorm.Config) error
	OpenSalve(connectionString string) error
	Begin() DBAdapterV2
	RollbackUselessCommitted()
	Commit()
	Close()
	Gormer() *gorm.DB
	DB() (*sql.DB, error)
}

type adapterV2 struct {
	gormer      *gorm.DB
	isCommitted bool
}

// NewDB returns a new instance of DB.
func NewDatabase() DBAdapterV2 {
	return &adapterV2{}
}

// Open opens a DB connection.
func (db *adapterV2) Open(connectionString string, config gorm.Config) error {
	nrdb, err := sql.Open("nrmysql", connectionString)

	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: nrdb}), &config)
	if err != nil {
		return err
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(configs.AppConfig.DB.MaxIdleConnection)
	sqlDB.SetMaxOpenConns(configs.AppConfig.DB.MaxOpenConnection)
	sqlDB.SetConnMaxLifetime(time.Hour * time.Duration(configs.AppConfig.DB.MaxLifeTime))

	db.gormer = gormDB
	return nil
}

// Begin starts a DB transaction.
func (db *adapterV2) Begin() DBAdapterV2 {
	tx := db.gormer.Begin()
	return &adapterV2{
		gormer:      tx,
		isCommitted: false,
	}
}

// RollbackUselessCommitted rollbacks useless DB transaction committed.
func (db *adapterV2) RollbackUselessCommitted() {
	if !db.isCommitted {
		db.gormer.Rollback()
	}
}

// Commit commits a DB transaction.
func (db *adapterV2) Commit() {
	if !db.isCommitted {
		db.gormer.Commit()
		db.isCommitted = true
	}
}

// Open opens a DB connection.
func (db *adapterV2) OpenSalve(connectionString string) error {

	err := db.gormer.Use(dbresolver.Register(dbresolver.Config{
		Replicas: []gorm.Dialector{mysql.Open(connectionString)},
	}).
		SetConnMaxLifetime(time.Hour * time.Duration(configs.AppConfig.DB.MaxLifeTime)).
		SetMaxIdleConns(configs.AppConfig.DB.MaxIdleConnection).
		SetMaxOpenConns(configs.AppConfig.DB.MaxOpenConnection))

	return err
}

// Close closes DB connection.
func (db *adapterV2) Close() {
	sqlDB, err := db.gormer.DB()
	if err != nil {
		return
	}

	_ = sqlDB.Close()
}

// Gormer returns an instance of gorm.DB.
func (db *adapterV2) Gormer() *gorm.DB {
	return db.gormer
}

// DB returns an instance of sql.DB.
func (db *adapterV2) DB() (*sql.DB, error) {
	return db.gormer.DB()
}
