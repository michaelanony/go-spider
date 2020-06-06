package initSpider

import (
	"github.com/BurntSushi/toml"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

var (
	MysqlPool *sqlx.DB
	RedisPool *redis.Pool
)

type Config struct {
	MysqlDB string
	Redis   string
}

func Init() (err error) {
	var config Config
	if _, err = toml.DecodeFile("./conf/config.toml", &config); err != nil {
		log.Fatal(err)
		return
	}
	if err = createMysqlPool(config.MysqlDB); err != nil {
		log.Fatal(err)
		return
	}
	if err = createRedisPool(config.MysqlDB); err != nil {
		log.Fatal(err)
		return
	}
	return
}

func createMysqlPool(mysqlDns string) (err error) {
	MysqlPool, err = sqlx.Open("mysql", mysqlDns)
	if err != nil {
		return
	}
	if err = MysqlPool.Ping(); err != nil {
		return
	}
	MysqlPool.SetMaxIdleConns(100)
	MysqlPool.SetMaxOpenConns(16)
	return
}

func createRedisPool(RedisDns string) (err error) {
	RedisPool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (conn redis.Conn, err error) {
			return redis.Dial("tcp", RedisDns)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return
}
