package xmysql

import (
	"database/sql"
	migrate "github.com/rubenv/sql-migrate"
)

const mysqlPlatform = "mysql"

type MysqlDatabaseMigrator struct {
	dbClient        *sql.DB
	migrationSet    *migrate.MigrationSet
	migrationSource migrate.MigrationSource
	platform        string
}

func NewMysqlDatabaseMigrator(dbClient *sql.DB, migrationsLocation string, tableName string) *MysqlDatabaseMigrator {
	source := &migrate.FileMigrationSource{
		Dir: migrationsLocation,
	}

	migrationSet := &migrate.MigrationSet{TableName: tableName}

	return &MysqlDatabaseMigrator{dbClient: dbClient, migrationSet: migrationSet, migrationSource: source, platform: mysqlPlatform}
}

func (m *MysqlDatabaseMigrator) Up() (int, error) {
	return m.migrationSet.Exec(m.dbClient, m.platform, m.migrationSource, migrate.Up)
}

func (m *MysqlDatabaseMigrator) Down() (int, error) {
	return m.migrationSet.Exec(m.dbClient, m.platform, m.migrationSource, migrate.Down)
}
