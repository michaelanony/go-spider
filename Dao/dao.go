package Dao

import (
	"github.com/garyburd/redigo/redis"
	"github.com/jmoiron/sqlx"
)

type DBDao struct {
	MysqlPool *sqlx.DB
	RedisPool *redis.Pool
}
