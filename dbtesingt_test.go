package dbtesting

import (
	"github.com/wind-coco/dbtesting/config"
	"os"
	"testing"
)

func TestRunDBInDocker(t *testing.T) {

}

func TestMain(m *testing.M) {
	os.Exit(RunDBInDocker(m, &config.Config{
		Image:         "mysql",
		User:          "root",
		Password:      "123456",
		DBName:        "test_db",
		DB:            config.Mysql,
		ContainerPort: "3306/tcp",
	}))
}
