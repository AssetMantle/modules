package helpers

type Migrations interface {
	Get() []Migration
}
