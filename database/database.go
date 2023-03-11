package database

import (
	"fmt"

	"github.com/damoncoo/yi-api-core/logger"
	"github.com/go-redis/redis"
	"xorm.io/xorm"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type Storage struct {
	RedisEngine *RedisDBEngine
	ORMEngine   *xorm.Engine
}

type Connect interface {
	Driver() string
	ConnectString() string
	RedisHost() string
	RedisPassword() string
	RedisPort() int
	RedisDB() int
}

func (Storage) Init(c Connect) *Storage {

	Engine, err := xorm.NewEngine(c.Driver(), c.ConnectString())
	if err != nil {
		panic(err)
	}

	Engine.SetMaxIdleConns(300)
	Engine.SetMaxOpenConns(500)
	Engine.ShowSQL(logger.Debug == true)

	rdis := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.RedisHost(), c.RedisPort()),
		Password: c.RedisPassword(), // no password set
		DB:       c.RedisDB(),
	})

	reply, err := rdis.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)

	storage := &Storage{
		ORMEngine: Engine,
		RedisEngine: &RedisDBEngine{
			DBConn: rdis,
		},
	}
	return storage
}
