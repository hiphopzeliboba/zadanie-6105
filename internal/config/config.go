package config

import (
	"github.com/joho/godotenv"
)

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}

//type PostgreConn interface {
//	Connect() (*sql.DB, error)
//}
//type pDSN struct {
//	POSTGRES_USERNAME string
//	POSTGRES_PASSWORD string
//	POSTGRES_DATABASE string
//	POSTGRES_HOST     string
//	POSTGRES_PORT     string
//}
//
//func PGDSN()
