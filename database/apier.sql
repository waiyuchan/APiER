/*
 Navicat Premium Data Transfer

 Source Server         : MySQL
 Source Server Type    : MySQL
 Source Server Version : 80031 (8.0.31)
 Source Host           : localhost:3306
 Source Schema         : apier

 Target Server Type    : MySQL
 Target Server Version : 80031 (8.0.31)
 File Encoding         : 65001

 Date: 06/04/2024 00:14:18
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin`  (
  `admin_id` int NOT NULL AUTO_INCREMENT COMMENT '管理员ID',
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '电子邮箱',
  `role` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`admin_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '管理员' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for admin_action
-- ----------------------------
DROP TABLE IF EXISTS `admin_action`;
CREATE TABLE `admin_action`  (
  `action_id` int NOT NULL AUTO_INCREMENT COMMENT '操作记录ID',
  `admin_id` int NOT NULL COMMENT '管理员ID',
  `action_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '操作类型',
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '操作描述',
  `created_at` datetime NOT NULL COMMENT '操作时间',
  PRIMARY KEY (`action_id`) USING BTREE,
  INDEX `admin_id`(`admin_id` ASC) USING BTREE,
  CONSTRAINT `admin_action_ibfk_1` FOREIGN KEY (`admin_id`) REFERENCES `admin` (`admin_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '管理员操作记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for admin_permission
-- ----------------------------
DROP TABLE IF EXISTS `admin_permission`;
CREATE TABLE `admin_permission`  (
  `permission_id` int NOT NULL AUTO_INCREMENT COMMENT '权限ID',
  `permission_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '权限名称',
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '权限描述',
  PRIMARY KEY (`permission_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '管理员权限' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for admin_role_permission
-- ----------------------------
DROP TABLE IF EXISTS `admin_role_permission`;
CREATE TABLE `admin_role_permission`  (
  `role_permissions_id` int NOT NULL COMMENT '管理员角色权限关系ID',
  `role` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色名称',
  `permission_id` int NOT NULL COMMENT '权限ID',
  PRIMARY KEY (`role_permissions_id`) USING BTREE,
  INDEX `permission_id`(`permission_id` ASC) USING BTREE,
  CONSTRAINT `admin_role_permission_ibfk_1` FOREIGN KEY (`permission_id`) REFERENCES `admin_permission` (`permission_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '管理员角色权限关联关系' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for api_approval
-- ----------------------------
DROP TABLE IF EXISTS `api_approval`;
CREATE TABLE `api_approval`  (
  `approval_id` int NOT NULL AUTO_INCREMENT COMMENT '审批记录ID',
  `api_id` int NOT NULL COMMENT '接口ID',
  `developer_id` int NOT NULL COMMENT '开发者ID',
  `status` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '审批状态，三种状态：0-待审批，1-通过，2-拒绝',
  `reason` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '驳回理由',
  `created_at` datetime NOT NULL COMMENT '提交审批时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '审批更新时间',
  PRIMARY KEY (`approval_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '接口审批' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for api_call_stats
-- ----------------------------
DROP TABLE IF EXISTS `api_call_stats`;
CREATE TABLE `api_call_stats`  (
  `stat_id` int NOT NULL AUTO_INCREMENT COMMENT '统计ID',
  `api_id` int NOT NULL COMMENT '接口ID',
  `call_count` int NOT NULL COMMENT '调用次数',
  `success_count` int NOT NULL COMMENT '成功次数',
  `fail_count` int NOT NULL COMMENT '失败次数',
  `last_called` datetime NOT NULL COMMENT '最后调用时间',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`stat_id`) USING BTREE,
  INDEX `api_id`(`api_id` ASC) USING BTREE,
  CONSTRAINT `api_call_stats_ibfk_1` FOREIGN KEY (`api_id`) REFERENCES `api_info` (`api_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '接口调用统计表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for api_info
-- ----------------------------
DROP TABLE IF EXISTS `api_info`;
CREATE TABLE `api_info`  (
  `api_id` int NOT NULL AUTO_INCREMENT COMMENT '接口ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '接口名称',
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '接口描述',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '接口URL',
  `method` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求方法',
  `request_params` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '请求参数（JSON格式）',
  `request_headers` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '请求头（JSON格式）',
  `response_structure` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '响应结构体（JSON格式）',
  `is_active` tinyint(1) NOT NULL COMMENT '是否激活',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`api_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '接口信息' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for api_permission
-- ----------------------------
DROP TABLE IF EXISTS `api_permission`;
CREATE TABLE `api_permission`  (
  `permission_id` int NOT NULL AUTO_INCREMENT COMMENT '权限ID',
  `api_id` int NOT NULL COMMENT '接口ID',
  `developer_id` int NOT NULL COMMENT '开发者ID',
  `status` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '审批状态',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`permission_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '接口权限' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for api_quota
-- ----------------------------
DROP TABLE IF EXISTS `api_quota`;
CREATE TABLE `api_quota`  (
  `quota_id` int NOT NULL AUTO_INCREMENT COMMENT '额度记录ID',
  `developer_id` int NOT NULL COMMENT '开发者ID',
  `quota` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'API调用额度',
  PRIMARY KEY (`quota_id`) USING BTREE,
  INDEX `api_quota_developer_fgk`(`developer_id` ASC) USING BTREE,
  CONSTRAINT `api_quota_developer_fgk` FOREIGN KEY (`developer_id`) REFERENCES `developer` (`developer_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '接口调用额度' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for api_tag
-- ----------------------------
DROP TABLE IF EXISTS `api_tag`;
CREATE TABLE `api_tag`  (
  `tag_id` int NOT NULL AUTO_INCREMENT COMMENT '标签ID',
  `tag_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标签名称',
  PRIMARY KEY (`tag_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '接口标签' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for api_tag_relation
-- ----------------------------
DROP TABLE IF EXISTS `api_tag_relation`;
CREATE TABLE `api_tag_relation`  (
  `api_id` int NOT NULL COMMENT '接口ID',
  `tag_id` int NOT NULL COMMENT '标签ID',
  PRIMARY KEY (`api_id`, `tag_id`) USING BTREE,
  INDEX `tag_id`(`tag_id` ASC) USING BTREE,
  CONSTRAINT `api_tag_relation_ibfk_1` FOREIGN KEY (`api_id`) REFERENCES `api_info` (`api_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `api_tag_relation_ibfk_2` FOREIGN KEY (`tag_id`) REFERENCES `api_tag` (`tag_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '接口与标签关联关系' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for api_test_record
-- ----------------------------
DROP TABLE IF EXISTS `api_test_record`;
CREATE TABLE `api_test_record`  (
  `test_id` int NOT NULL AUTO_INCREMENT COMMENT '测试记录ID',
  `api_id` int NULL DEFAULT NULL COMMENT '接口ID',
  `created_at` datetime NOT NULL COMMENT '测试时间',
  `user_id` int NOT NULL COMMENT '用户ID（开发者或管理员ID）',
  `user_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户类型（开发者、管理员）',
  `request_method` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求方法',
  `request_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '请求URL',
  `request_headers` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '请求头（JSON格式）',
  `request_body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '请求体（JSON格式）',
  `response_body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '响应体（JSON格式）',
  `status_code` int NULL DEFAULT NULL COMMENT '响应状态码',
  `test_result` int NULL DEFAULT NULL COMMENT '测试结果',
  PRIMARY KEY (`test_id`) USING BTREE,
  INDEX `fk_api_test_records_api_id`(`api_id` ASC) USING BTREE,
  CONSTRAINT `fk_api_test_records_api_id` FOREIGN KEY (`api_id`) REFERENCES `api_info` (`api_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '接口测试记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for api_usage_record
-- ----------------------------
DROP TABLE IF EXISTS `api_usage_record`;
CREATE TABLE `api_usage_record`  (
  `record_id` int NOT NULL AUTO_INCREMENT COMMENT '使用记录ID',
  `developer_id` int NOT NULL COMMENT '开发者ID',
  `api_id` int NOT NULL COMMENT '接口ID',
  `result` int NOT NULL COMMENT '调用结果',
  `last_used` datetime NOT NULL COMMENT '最后使用时间',
  `created_at` datetime NOT NULL COMMENT '记录创建时间',
  PRIMARY KEY (`record_id`) USING BTREE,
  INDEX `api_id`(`api_id` ASC) USING BTREE,
  INDEX `developer_id`(`developer_id` ASC) USING BTREE,
  CONSTRAINT `api_usage_record_ibfk_1` FOREIGN KEY (`api_id`) REFERENCES `api_info` (`api_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `api_usage_record_ibfk_2` FOREIGN KEY (`developer_id`) REFERENCES `developer` (`developer_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '接口调用记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for developer
-- ----------------------------
DROP TABLE IF EXISTS `developer`;
CREATE TABLE `developer`  (
  `developer_id` int NOT NULL AUTO_INCREMENT COMMENT '开发者ID',
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '电子邮箱',
  `api_token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'API令牌',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`developer_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '开发者信息' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for developer_setting
-- ----------------------------
DROP TABLE IF EXISTS `developer_setting`;
CREATE TABLE `developer_setting`  (
  `setting_id` int NOT NULL AUTO_INCREMENT COMMENT '设置ID',
  `developer_id` int NOT NULL COMMENT '开发者ID',
  `email_notifications` tinyint(1) NOT NULL COMMENT '邮件通知开关',
  `theme` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '主题设置',
  `created_at` datetime NOT NULL COMMENT '记录创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`setting_id`) USING BTREE,
  INDEX `developer_id`(`developer_id` ASC) USING BTREE,
  CONSTRAINT `developer_setting_ibfk_1` FOREIGN KEY (`developer_id`) REFERENCES `developer` (`developer_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '个人设置' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for domain_prefix_mapping
-- ----------------------------
DROP TABLE IF EXISTS `domain_prefix_mapping`;
CREATE TABLE `domain_prefix_mapping`  (
  `mapping_id` int NOT NULL AUTO_INCREMENT COMMENT '映射ID',
  `original_domain` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '原始域名',
  `target_domain` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '目标域名',
  `original_prefix` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '原始前缀',
  `target_prefix` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '目标前缀',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`mapping_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '域名前缀映射' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for internal_message
-- ----------------------------
DROP TABLE IF EXISTS `internal_message`;
CREATE TABLE `internal_message`  (
  `message_id` int NOT NULL AUTO_INCREMENT COMMENT '消息ID',
  `subject` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '消息主题',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '消息内容',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`message_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '站内信' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for message_recipient
-- ----------------------------
DROP TABLE IF EXISTS `message_recipient`;
CREATE TABLE `message_recipient`  (
  `recipient_id` int NOT NULL AUTO_INCREMENT COMMENT '接收者ID',
  `message_id` int NOT NULL COMMENT '消息ID',
  `user_id` int NOT NULL COMMENT '用户ID',
  `user_type` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户类型，类型：develop / admin',
  `is_read` tinyint(1) NOT NULL COMMENT '是否已读',
  `read_at` datetime NULL DEFAULT NULL COMMENT '阅读时间',
  PRIMARY KEY (`recipient_id`) USING BTREE,
  INDEX `message_id`(`message_id` ASC) USING BTREE,
  CONSTRAINT `message_recipient_ibfk_1` FOREIGN KEY (`message_id`) REFERENCES `internal_message` (`message_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '消息接收' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for super_admin
-- ----------------------------
DROP TABLE IF EXISTS `super_admin`;
CREATE TABLE `super_admin`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '超级管理员' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
