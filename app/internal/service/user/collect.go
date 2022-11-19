package user

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	g "main/app/global"
	"main/app/internal/model"
)

type sCollect struct{}

var insCollect = sCollect{}

func (s *sCollect) CheckCollectionIsExist(ctx context.Context, collectType int32, userId int64, id interface{}) error {
	whereSql := ""
	switch collectType {
	case 1:
		whereSql = fmt.Sprintf("user_id = ? AND restaurant_id = ?")
	case 2:
		whereSql = fmt.Sprintf("user_id = ? AND recipe_id = ?")
	default:

	}

	userCollection := &model.UserCollection{}
	err := g.MysqlDB.WithContext(ctx).
		Table("user_collection").
		Where(whereSql, userId, id).
		First(userCollection).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			g.Logger.Errorf("query [user_collection] record failed, err: %v", err)
			return fmt.Errorf("internal err")
		}
	} else {
		return fmt.Errorf("duplicate collect")
	}

	return nil
}

func (s *sCollect) CheckCollectionIdIsExist(ctx context.Context, id, userId int64) error {
	userCollection := &model.UserCollection{}
	err := g.MysqlDB.WithContext(ctx).
		Table("user_collection").
		Select("id,user_id").
		Where("id= ? AND user_id = ?", id, userId).
		First(userCollection).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("collection not found")
		}
		g.Logger.Errorf("query [user_collection] record failed, err: %v", err)
		return fmt.Errorf("internal err")
	}

	return nil
}

func (s *sCollect) CreateCollection(ctx context.Context, userCollection *model.UserCollection) {
	g.MysqlDB.WithContext(ctx).
		Table("user_collection").
		Create(userCollection)
}

func (s *sCollect) DeleteCollection(ctx context.Context, id int64) error {
	err := g.MysqlDB.WithContext(ctx).
		Table("user_collection").
		Delete(&model.UserCollection{}, id).Error
	if err != nil {
		g.Logger.Errorf("delete [user_collection] record failed, err: %v", err)
		return fmt.Errorf("internal err")
	}

	return nil
}

func (s *sCollect) GetUserCollectionCount(ctx context.Context, userId int64, collectType int32) (int64, error) {
	var cnt int64
	err := g.MysqlDB.WithContext(ctx).
		Table("user_collection").
		Where("user_id = ? AND collect_type = ?", userId, collectType).
		Count(&cnt).Error
	if err != nil {
		g.Logger.Errorf("query [user_collection] record failed ,err: %v", err)
		return -1, fmt.Errorf("internal err")
	}

	return cnt, nil
}

func (s *sCollect) GetUserCollectionsWithLimit(ctx context.Context, userId int64, collectType int32, limit, page int) ([]*model.UserCollection, error) {
	var userCollections []*model.UserCollection
	err := g.MysqlDB.WithContext(ctx).
		Table("user_collection").
		Limit(limit).Offset(limit*(page-1)).
		Where("user_id = ? AND collect_type = ?", userId, collectType).
		Find(&userCollections).Error
	if err != nil {
		g.Logger.Errorf("query [user_collection] failed, err: %v", err)
		return nil, fmt.Errorf("internal err")
	}

	return userCollections, nil
}
