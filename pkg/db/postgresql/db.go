package postgresql

import (
	_ "fmt"

	"github.com/go-pg/pg/v10"
	_ "github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBname   string
	//SSLMode  string
}

func NewPostgresDBClient(cfg Config) (*pg.DB, error) {
	//dataSource := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
	//	cfg.Host, cfg.Port, cfg.Username, cfg.DBname, cfg.Password, cfg.SSLMode)
	//conn, err := sqlx.Connect("postgres", dataSource)
	conn := pg.Connect(&pg.Options{
		User:     cfg.Username,
		Addr:     cfg.Host,
		Password: cfg.Password,
		Database: cfg.DBname,
	})
	/*if err != nil {
		return nil, err
	}*/

	/*err = conn.Ping()
	if err != nil {
		return nil, err
	}*/

	return conn, nil
}
