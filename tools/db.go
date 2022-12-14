package tools

import (
	"fmt"
	"gorm.io/gorm"
)

const (
	EquatorialRadius = 6378000 // 赤道半径，单位：米
)

// Paginate 分页封装
func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

// NotDeleted 条件为未删除的
func NotDeleted() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("deleted_at is null")
	}
}

// SortDistance 按照距离排序
func SortDistance(longitude, latitude float64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		distanceSelectSql := fmt.Sprintf("t_scenic.*, ACOS(COS(RADIANS(%v)) * COS(RADIANS(latitude)) * COS(RADIANS(longitude) - RADIANS(%v)) + SIN(RADIANS(%v)) * SIN(RADIANS(latitude))) * %d AS distance", latitude, longitude, latitude, EquatorialRadius)
		return db.Select(distanceSelectSql).Order("distance")
	}
}
