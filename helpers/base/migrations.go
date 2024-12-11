package base

import (
	"github.com/AssetMantle/modules/helpers"
)

type migrations struct {
	migrationList []helpers.Migration
}

var _ helpers.Migrations = (*migrations)(nil)

func (migrations migrations) Get() []helpers.Migration {
	return migrations.migrationList
}

func NewMigrations(migrationList ...helpers.Migration) helpers.Migrations {
	return migrations{
		migrationList: migrationList,
	}
}
