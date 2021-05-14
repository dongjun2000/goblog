package user

import (
	"goblog/pkg/password"
	"gorm.io/gorm"
)

// GORM 的模型钩子，创建模型前调用
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = password.Hash(u.Password)
	return
}

// GORM 的模型钩子，更新模型前调用
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if !password.IsHashed(u.Password) {
		u.Password = password.Hash(u.Password)
	}
	return
}