package dao

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Relation struct {
	ID       int64 `gorm:"primaryKey,autoIncrement"`
	UserId   int64 `gorm:"index_id"`
	ToUserId int64 `gorm:"index_id"`
	Ctime    int64
	Utime    int64 `gorm:"index"`
}

type RelationDao struct {
	db *gorm.DB
}

func NewRelationDao(db *gorm.DB) *RelationDao {
	return &RelationDao{db: db}
}

func (d *RelationDao) GetFollowExists(ctx context.Context, uid, toUid int64) (bool, error) {
	var r Relation
	err := d.db.WithContext(ctx).Model(&Relation{}).Where("user_id = ? AND to_user_id = ?", uid, toUid).Find(&r).Error
	if err != nil {
		return false, err
	}
	if r.ID == 0 {
		return false, nil
	}

	return true, nil
}

func (d *RelationDao) AddFollow(ctx context.Context, relation Relation) error {
	now := time.Now().Unix()
	relation.Ctime = now
	relation.Utime = now

	return d.db.WithContext(ctx).Create(&relation).Error
}

func (d *RelationDao) DelFollow(ctx context.Context, relation Relation) error {
	return d.db.WithContext(ctx).Where("user_id = ? AND to_user_id = ?", relation.UserId, relation.ToUserId).Delete(&Relation{}).Error
}

func (d *RelationDao) GetFollowList(ctx context.Context, uid int64) ([]Relation, error) {
	var relations []Relation
	err := d.db.WithContext(ctx).Model(&Relation{}).Where("user_id = ?", uid).Find(&relations).Error
	if err != nil {
		return nil, err
	}

	return relations, nil
}

func (d *RelationDao) GetFollowerList(ctx context.Context, uid int64) ([]Relation, error) {
	var relations []Relation
	err := d.db.WithContext(ctx).Model(&Relation{}).Where("to_user_id = ?", uid).Find(&relations).Error
	if err != nil {
		return nil, err
	}

	return relations, nil
}

func (d *RelationDao) GetFriendList(ctx context.Context, uid int64) ([]int64, error) {
	var follows, followers []Relation
	err := d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.WithContext(ctx).Model(&Relation{}).Where("user_id = ?", uid).Find(&follows).Error
		if err != nil {
			return err
		}

		err = tx.WithContext(ctx).Model(&Relation{}).Where("to_user_id = ?", uid).Find(&followers).Error
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	var result []int64
	for _, f := range follows {
		for _, f2 := range followers {
			if f.ToUserId == f2.UserId {
				result = append(result, f.ToUserId)
			}
		}
	}

	return result, nil
}

func (d *RelationDao) GetIsFollow(ctx context.Context, uid, toUid int64) (bool, error) {
	var r Relation
	err := d.db.WithContext(ctx).Model(&Relation{}).Where("user_id = ? AND to_user_id = ?", uid, toUid).First(&r).Error
	if err != nil {
		return false, err
	}

	if r.ID == 0 {
		return false, nil
	}

	return true, nil
}

func (d *RelationDao) GetFollowCount(ctx context.Context, toUid int64) (int64, int64, error) {
	var followCount, followerCount int64
	err := d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.WithContext(ctx).Model(&Relation{}).Where("user_id = ?", toUid).Count(&followCount).Error
		if err != nil {
			return err
		}

		err = tx.WithContext(ctx).Model(&Relation{}).Where("to_user_id = ?", toUid).Count(&followerCount).Error
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return -1, -1, err
	}

	return followCount, followerCount, nil
}
