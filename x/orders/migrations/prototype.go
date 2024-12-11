package migrations

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/orders/migrations/v2"
)

func Prototype() helpers.Migrations {
	return base.NewMigrations(v2.Migration)
}
