/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.153.128
 Source Server Type    : MySQL
 Source Server Version : 50649
 Source Host           : 192.168.153.128:33061
 Source Schema         : miaosha

 Target Server Type    : MySQL
 Target Server Version : 50649
 File Encoding         : 65001

 Date: 06/09/2020 01:01:58
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order`  (
  `id` int(10) UNSIGNED NOT NULL COMMENT '用户ID',
  `order_id` int(10) UNSIGNED NOT NULL COMMENT '订单ID',
  `goods_id` int(10) UNSIGNED NOT NULL COMMENT '商品ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Compact;

SET FOREIGN_KEY_CHECKS = 1;
