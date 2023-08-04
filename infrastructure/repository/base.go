package repository

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"sync"
)

var db *PGRepository
var pgSingleton sync.Once

var redisDB *RedisDB
var redisSingleton sync.Once

type RedisDB struct {
	Client *redis.Client
}

type PGRepository struct {
	DB *gorm.DB
}

func RedisInit() *RedisDB {
	redisSingleton.Do(func() {
		redisDB = &RedisDB{
			Client: createRedisConnection(),
		}
	})
	return redisDB
}

func NewGormDatabase() *PGRepository {
	pgSingleton.Do(func() {
		println("hi")
		client, _ := GormInit()
		db = &PGRepository{DB: client}
	})

	return db
}
