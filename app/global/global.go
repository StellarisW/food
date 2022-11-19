package g

import (
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"main/app/internal/model/config"
)

var (
	Config  *config.Config
	Logger  *zap.SugaredLogger
	MysqlDB *gorm.DB
	MongoDB *mongo.Client
	Rdb     *redis.Client
)
