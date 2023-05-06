package tomato

import (
	"github.com/JuShangEnergy/framework/orm"
	"github.com/JuShangEnergy/framework/storage"
	"github.com/JuShangEnergy/framework/storage/postgres"
	"github.com/JuShangEnergy/framework/test"
	"testing"
)

func TestRun(t *testing.T) {
	Run()
}

func initPostgresEnv() {
	orm.InitOrm(getPostgresAdapter())
}

func getPostgresAdapter() storage.Adapter {
	return postgres.NewPostgresAdapter("tomato", test.OpenPostgreSQForTest())
}
