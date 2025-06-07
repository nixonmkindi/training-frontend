package database

import (
	"context"
	"training-frontend/package/config"
	"training-frontend/package/log"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
)

var once sync.Once

// type DriverPg struct {
// 	conn string
// }

// variavel Global
var instance *pgxpool.Pool
var err error

func Connect() (*pgxpool.Pool, error) {
	//thread safe singletone initialised
	//connectionString := "postgresql://root@192.168.51.227:26257/teleradiology?sslcert=/home/kangwe/cockroachdb/certs%2Fclient.root.crt&sslkey=/home/kangwe/cockroachdb/certs%2Fclient.root.key&sslmode=verify-full&sslrootcert=/home/kangwe/cockroachdb/certs%2Fca.crt"
	//connectionString := "postgresql://root@192.168.51.227:26257/defaultdb?sslcert=/home/kangwe/cockroachdb/certs%2Fclient.root.crt&sslkey=/home/kangwe/cockroachdb/certs%2Fclient.root.key&sslmode=verify-full&sslrootcert=/home/kangwe/cockroachdb/certs%2Fca.crt"
	//connectionString := "postgresql://root@localhost:26257/defaultdb?sslmode=disable&pool_max_conns=100"
	once.Do(func() {
		connectionString := config.GetDatabaseConnection()
		instance, err = pgxpool.Connect(context.Background(), connectionString)
		if err != nil {
			log.Errorf("unable to create a database instance")
		}

	})
	if err != nil {
		log.Errorf("unable to connect to database: %v\n", err)
		return nil, err
	}
	return instance, err
}

func Close() {
	instance.Close()
}
