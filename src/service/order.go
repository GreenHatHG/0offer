package service

import (
	"0offer/global"
	"0offer/model"
	"errors"
)

func Miaosha(goodsId uint, userId uint) error {
	goods := model.MiaoshaGoods{Id: goodsId}
	if err := global.DB.First(&model.UserInfo{Id: userId}).Error; err != nil {
		return err
	}

	if err := global.DB.First(&model.Order{UserId: userId, GoodsId: goodsId}).Error; err == nil {
		return errors.New("该用户已秒杀成功")
	}

	err := global.DB.First(&goods).Error
	if err != nil {
		return err
	}
	if goods.GoodsStock <= 0 {
		return errors.New("商品库存不足")
	}

	err = global.DB.Model(&goods).Update("goods_stock", goods.GoodsStock-1).Error
	if err != nil {
		return err
	}

	err = global.DB.Save(&model.Order{GoodsId: goodsId, UserId: userId}).Error
	if err != nil {
		return err
	}

	return nil
}
