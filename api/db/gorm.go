package db

import (
    "log"
    "os"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var GormDB *gorm.DB

func InitGorm(autoMigrateModels ...interface{}) error {
    dsn := os.Getenv("DB_URL")
    if dsn == "" {
        log.Println("DB_URL not set; skipping GORM init")
        return nil
    }
    gdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return err
    }
    if len(autoMigrateModels) > 0 {
        if err := gdb.AutoMigrate(autoMigrateModels...); err != nil {
            return err
        }
    }
    GormDB = gdb
    log.Println("GORM initialized")
    return nil
}


