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

 Date: 06/09/2020 01:01:49
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for miaosha_goods
-- ----------------------------
DROP TABLE IF EXISTS `miaosha_goods`;
CREATE TABLE `miaosha_goods`  (
  `id` int(10) UNSIGNED NOT NULL COMMENT '商品ID',
  `goods_name` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL COMMENT '商品名称',
  `goods_stock` int(11) NOT NULL COMMENT '商品库存',
  `miaosha_start_date` datetime(0) NULL DEFAULT NULL COMMENT '秒杀开始时间',
  `miaosha_end_date` datetime(0) NULL DEFAULT NULL COMMENT '秒杀结束时间\r\n',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Compact;

SET FOREIGN_KEY_CHECKS = 1;
