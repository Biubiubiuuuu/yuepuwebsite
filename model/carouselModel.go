package model

import (
	"github.com/Biubiubiuuuu/yuepuwebsite/db/mysql"
)

// 轮播图
type Carousel struct {
	Model
	Url  string `json:"url"`  // 图片地址
	Link string `json:"link"` // 跳转链接
	Sort int64  `json:"sort"` // 排序 越大越靠前
}

// 添加轮播图
func (c *Carousel) AddCarousel() error {
	db := mysql.GetMysqlDB()
	return db.Create(&c).Error
}

// 查询所有轮播
func QueryCarousel(pageSize int, page int) (count int, carousels []Carousel) {
	db := mysql.GetMysqlDB()
	query := db.Table("carousel").Select("carousel.*")
	query.Count(&count)
	query.Limit(pageSize).Offset((page - 1) * pageSize).Order("sort desc").Find(&carousels)
	return
}

// 轮播详情
func (c *Carousel) QueryCarouselByID() error {
	db := mysql.GetMysqlDB()
	return db.First(&c, c.ID).Error
}

// 修改轮播
func (c *Carousel) EditCarousel(args map[string]interface{}) error {
	db := mysql.GetMysqlDB()
	return db.Model(&c).Updates(args).Error
}

// 删除轮播，返回受影响行数
func DelCarousel(ids []int64) int64 {
	db := mysql.GetMysqlDB()
	return db.Where("id in (?)", ids).Unscoped().Delete(&Carousel{}).RowsAffected
}
