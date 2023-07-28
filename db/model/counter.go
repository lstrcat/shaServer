package model

import (
	"time"
)

// CounterModel 计数器模型
type CounterModel struct {
	Id        int32     `gorm:"column:id" json:"id"`
	Count     int32     `gorm:"column:count" json:"count"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}

type GoodsModel struct {
	Id   int32  `gorm:"column:id" json:"id"`
	Json string `gorm:"column:json" json:"json"`
}
