package sqldb

type DatabaseMigrator interface {
	Up() (int, error)
	Down() (int, error)
}
