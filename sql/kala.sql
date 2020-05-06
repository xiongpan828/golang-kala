/*
 Navicat Premium Data Transfer

 Source Server         : aliyun.go_orm
 Source Server Type    : MySQL
 Source Server Version : 50725
 Source Host           : qocat.com:3306
 Source Schema         : go_orm

 Target Server Type    : MySQL
 Target Server Version : 50725
 File Encoding         : 65001

 Date: 06/05/2020 17:23:28
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for kala
-- ----------------------------
DROP TABLE IF EXISTS `kala`;
CREATE TABLE `kala`  (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `status` bigint(2) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 39 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
