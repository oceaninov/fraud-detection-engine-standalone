package instances

import (
	"fmt"
	"gitlab.com/fds22/detection-sys/pkg/databaseConnector"
	"gitlab.com/fds22/detection-sys/pkg/environments"
	"gorm.io/gorm"
	"time"
)

func NewOrm(envs *environments.Envs) (*gorm.DB, error) {
	duration, err := time.ParseDuration(envs.DatabaseConMaxLifetime)
	if err != nil {
		return nil, err
	}
	stringConnection := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s",
		//"%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true",
		envs.DatabaseUser,
		envs.DatabasePass,
		envs.DatabaseHost,
		envs.DatabasePort,
		envs.DatabaseName,
		envs.DatabaseSchema,
	)
	opts := &databaseConnector.PostgresOption{
		ConnectionString:      stringConnection,
		MaxOpenConnection:     envs.DatabaseMaxOpenConn,
		MaxIdleConnection:     envs.DatabaseMaxIdleConn,
		MaxLifeTimeConnection: duration,
	}
	return databaseConnector.NewPostgres(opts)
}
