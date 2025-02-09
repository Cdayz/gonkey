package fixtures

import (
	"database/sql"
	"strings"

	_ "github.com/lib/pq"

	"github.com/lamoda/gonkey/fixtures/mysql"
	"github.com/lamoda/gonkey/fixtures/postgres"
)

type DbType int

const (
	Postgres DbType = iota
	Mysql
)

const (
	PostgresParam = "postgres"
	MysqlParam    = "mysql"
)

type Config struct {
	DB       *sql.DB
	DbType   DbType
	Location string
	Debug    bool
}

type Loader interface {
	Load(names []string) error
}

func NewLoader(cfg *Config) Loader {

	var loader Loader

	location := strings.TrimRight(cfg.Location, "/")

	switch cfg.DbType {
	case Postgres:
		loader = postgres.New(
			cfg.DB,
			location,
			cfg.Debug,
		)
	case Mysql:
		loader = mysql.New(
			cfg.DB,
			location,
			cfg.Debug,
		)
	default:
		panic("unknown db type")
	}

	return loader
}

func FetchDbType(dbType string) DbType {
	switch dbType {
	case PostgresParam:
		return Postgres
	case MysqlParam:
		return Mysql
	default:
		panic("unknown db type param")
	}
}
