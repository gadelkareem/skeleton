package commands

import (
    "database/sql"
    "fmt"

    "backend/kernel"
    h "github.com/gadelkareem/go-helpers"
    "github.com/rubenv/sql-migrate"
)

var directions = []string{"up", "down"}

type migrator struct {
    db         *sql.DB
    migrations migrate.MigrationSource
}

func NewMigrator(db *sql.DB) kernel.Command {
    migrations := &migrate.FileMigrationSource{
        Dir: "migrations/sql",
    }
    return &migrator{db: db, migrations: migrations}
}

func (c *migrator) Run(args []string) {
    dir := migrate.Up
    if len(args) > 0 && args[0] == "down" {
        dir = migrate.Down
    }

    c.exec(dir)
}

func (c *migrator) exec(dir migrate.MigrationDirection) {
    fmt.Printf("Migration going %s...\n", directions[dir])
    n, err := migrate.Exec(c.db, "postgres", c.migrations, dir)
    h.PanicOnError(err)

    ms, err := c.migrations.FindMigrations()
    h.PanicOnError(err)

    for i := n - 1; i >= 0; i-- {
        fmt.Printf("Applied %s migration id: %s\n", directions[dir], ms[i].Id)
    }

    fmt.Printf("Applied %d %s migrations!\n", n, directions[dir])
}

func (c *migrator) Help() {
    fmt.Printf(`
Usage: skeleton migrate [options] ...
    
    Controls database migrations
    
Available options:
    up        Migrates the database to the most recent version available
    down      Undo a database migration
`)
}
