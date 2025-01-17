package cmd

import (
    "database/sql"
    "fmt"
    "log"

    "github.com/golang-migrate/migrate/v4"
    "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
    _ "github.com/lib/pq"
    "test/config"
    "test/utils"
)

func main() {
    var cfg config.Config
    err := utils.LoadProperties("config.json", &cfg)
    if err != nil {
        log.Fatal("Error loading config:", err)
    }

    dbSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name, cfg.Database.SSLMode)

    db, err := sql.Open("postgres", dbSource)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        log.Fatal("Error pinging database:", err)
    }

    driver, err := postgres.WithInstance(db, &postgres.Config{})
    if err != nil {
        log.Fatal("Could not create database instance", err)
    }

    m, err := migrate.NewWithDatabaseInstance(
        "file://resources/migrations/postgres",
        "postgres", driver)
    if err != nil {
        log.Fatal("Could not create migrate instance", err)
    }
    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        log.Fatal("Could not run migrations", err)
    }

    fmt.Println("Database migrated successfully.")

    // Here is where you would add your application logic.
    // For now, we'll just print a message:
    fmt.Println("Application started.")
}