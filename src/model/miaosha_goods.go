package model

import "time"

type MiaoshaGoods struct {
	Id               uint
	GoodsName        string
	GoodsStock       int
	MiaoshaStartDate time.Time
	MiaoshaEndDate   time.Time
}
