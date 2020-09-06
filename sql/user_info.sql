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

 Date: 06/09/2020 15:51:00
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user_info
-- ----------------------------
DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info`  (
  `id` int(10) UNSIGNED NOT NULL COMMENT '用户ID',
  `phone` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL COMMENT '手机号码',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Compact;

SET FOREIGN_KEY_CHECKS = 1;
