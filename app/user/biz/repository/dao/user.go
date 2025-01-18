package dao

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              int64  `gorm:"primaryKey,autoIncrement"`
	Name            string `gorm:"unique;not null;type:varchar(64)"`
	Password        string `gorm:"not null;type:varchar(128)"`
	Avatar          string `gorm:"type:varchar(128)"`
	BackgroundImage string `gorm:"type:varchar(128)"`
	Signature       string `gorm:"type:varchar(128)"`
	Ctime           int64
	Utime           int64
}

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

func (dao *UserDao) Create(ctx context.Context, user User) (int64, error) {
	now := time.Now().Unix()
	user.Ctime = now
	user.Utime = now
	err := dao.db.WithContext(ctx).Model(&User{}).Create(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return -1, err
		}
		return 0, err
	}

	return user.ID, nil
}

func (dao *UserDao) FindUserByName(ctx context.Context, name string) (User, error) {
	var user User
	err := dao.db.WithContext(ctx).Model(&User{}).Where("name = ?", name).First(&user).Error
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (dao *UserDao) FindUserByID(ctx context.Context, uid int64) (User, error) {
	var user User
	err := dao.db.WithContext(ctx).Model(&User{}).Where("id = ?", uid).First(&user).Error
	if err != nil {
		return User{}, err
	}

	return user, nil
}
