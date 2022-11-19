package boot

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	g "main/app/global"
	"time"
)

func MysqlDBSetup() {
	config := g.Config.DataBase.Mysql

	db, err := gorm.Open(mysql.Open(config.GetDsn()))
	if err != nil {
		g.Logger.Fatalf("initialize mysql db failed, err: %v", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxIdleTime(10 * time.Second)
	sqlDB.SetConnMaxLifetime(100 * time.Second)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	err = sqlDB.Ping()
	if err != nil {
		g.Logger.Fatalf("connect to mysql db failed, err: %v", err)
	}

	g.Logger.Infof("initialize mysql db successfully")
	g.MysqlDB = db
}

func MongoDBSetup() {
	clientOptions := options.Client().ApplyURI(g.Config.DataBase.Mongo.GetAddr())
	clientOptions.SetAuth(options.Credential{
		Username: g.Config.DataBase.Mongo.Username,
		Password: g.Config.DataBase.Mongo.Password,
	})

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		g.Logger.Fatalf("initialize mongodb failed, err: %v", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		g.Logger.Fatalf("initialize mongodb failed, err: %v", err)
	}

	g.MongoDB = client

	g.Logger.Infof("initiate mongodb successfully")
}

func RedisSetup() {
	config := g.Config.DataBase.Redis

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Addr, config.Port),
		Username: "",
		Password: config.Password,
		DB:       config.Db,
		PoolSize: 10000,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		g.Logger.Fatalf("connect to redis instance failed, err: %v", err)
	}

	g.Rdb = rdb

	g.Logger.Info("initialize redis client successfully")
}
