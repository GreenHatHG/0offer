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

 Date: 06/09/2020 15:50:53
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '订单ID',
  `user_id` int(10) UNSIGNED NOT NULL COMMENT '用户ID',
  `goods_id` int(10) UNSIGNED NOT NULL COMMENT '商品ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Compact;

SET FOREIGN_KEY_CHECKS = 1;
