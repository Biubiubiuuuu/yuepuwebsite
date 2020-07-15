package model

import (
	"github.com/Biubiubiuuuu/yuepuwebsite/db/mysql"
)

// 基础数据
//  AreaType 面积分类
type AreaType struct {
	Model
	MinArea float64 `json:"min_area"` // 最小面积（单位：平方米）
	MaxArea float64 `json:"max_area"` // 最大面积（单位：平方米）
}

// 添加面积分类
func (a *AreaType) AddAreaType() error {
	db := mysql.GetMysqlDB()
	return db.Create(&a).Error
}

// 修改面积分类
func (a *AreaType) EditAreaType(args map[string]interface{}) error {
	db := mysql.GetMysqlDB()
	return db.Model(&a).Updates(args).Error
}

// 查询面积所在分组
func (a *AreaType) QueryAreaTypeByArea(area float64) error {
	db := mysql.GetMysqlDB()
	return db.Where("min_area < ? AND max_area >= ?", area, area).First(&a).Error
}

// 查询面积分类
func (a *AreaType) QueryAreaType() (AreaTypes []AreaType) {
	db := mysql.GetMysqlDB()
	db.Find(&AreaTypes)
	return
}

// 删除面积分类(可批量)
func (a *AreaType) DeleteAreaType(ids []int64) error {
	db := mysql.GetMysqlDB()
	tx := db.Begin()
	if err := tx.Unscoped().Delete("id in (?)", ids).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
