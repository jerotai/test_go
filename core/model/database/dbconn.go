package database

import (
        "database/sql"

        _ "github.com/go-sql-driver/mysql"
)

var pool map[string]*sql.DB

func init() {
        pool = make(map[string]*sql.DB)
}

func GetConn(dbName string) *sql.DB {
        pool[dbName] = createConn(dbName)
        return pool[dbName]
}

func createConn(dbName string) *sql.DB {
        connStr := "Kdoctor_dev:kGalileiair_DEV@tcp(192.168.4.234:3306)/" + dbName
        db, err := sql.Open("mysql", connStr)
        if err != nil {
                //logger.Fatal("Connect Database Failed", zap.Error(err), zap.String("connStr", connStr))
        }

        db.SetMaxOpenConns(100)
        db.SetMaxIdleConns(100)
        db.SetConnMaxLifetime(3600)
        return db
}