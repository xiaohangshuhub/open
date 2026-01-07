-- Active: 1767504600613@@172.16.1.105@5432@open
-- Owner: open
-- Encoding: UTF8
CREATE DATABASE open;

-- 应用类型字典表
CREATE TABLE app_type (
    id SMALLINT NOT NULL PRIMARY KEY,
    name VARCHAR(20) NOT NULL,
    -- 唯一约束
    CONSTRAINT uk_app_type_name UNIQUE (name)
)

COMMENT ON TABLE app_type IS '应用类型表';

COMMENT ON COLUMN app_type.id IS '应用类型ID';

COMMENT ON COLUMN app_type.name IS '应用类型名称';

-- 插入应用类型数据
INSERT INTO
    app_type (id, name)
VALUES (1, 'web应用'),
    (2, '桌面应用'),
    (3, '移动应用'),
    (4, '小程序'),
    (5, '客户端应用'),
    (6, '物联网应用');

-- 应用表
CREATE TABLE app (
    id UUID NOT NULL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    app_type SMALLINT NOT NULL,
    key VARCHAR(255) NOT NULL,
    security VARCHAR(255) NOT NULL,
    desc TEXT,
    status SMALLINT NOT NULL,
    created_by UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_by UUID,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    -- 外键约束
    CONSTRAINT fk_app_type FOREIGN KEY (app_type) REFERENCES app_type (id),
    -- 唯一约束
    CONSTRAINT uk_app_name UNIQUE (name)
)

COMMENT ON TABLE app IS '应用表';

COMMENT ON COLUMN app.id IS '应用ID';

COMMENT ON COLUMN app.name IS '应用名称';

COMMENT ON COLUMN app.app_type IS '应用类型ID';

COMMENT ON COLUMN app.key IS '应用Key';

COMMENT ON COLUMN app.security IS '应用密钥';

COMMENT ON COLUMN app.desc IS '应用描述';

COMMENT ON COLUMN app.status IS '应用状态';

COMMENT ON COLUMN app.created_by IS '创建人ID';

COMMENT ON COLUMN app.created_at IS '创建时间';

COMMENT ON COLUMN app.updated_by IS '更新人ID';

COMMENT ON COLUMN app.updated_at IS '更新时间';

-- 主题表
CREATE TABLE topics (
    id UUID NOT NULL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    desc TEXT,
    status SMALLINT NOT NULL,
    app_id UUID NOT NULL,
    data_struct JSONB,
    created_by UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_by UUID,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
)

COMMENT ON TABLE topics IS '主题表';
COMMENT ON COLUMN topics.id IS '主题ID';
COMMENT ON COLUMN topics.name IS '主题名称';
COMMENT ON COLUMN topics.desc IS '主题描述';
COMMENT ON COLUMN topics.status IS '主题状态';
COMMENT ON COLUMN topics.app_id IS '应用ID';
COMMENT ON COLUMN topics.data_struct IS '数据结构定义';
COMMENT ON COLUMN topics.created_by IS '创建人ID';
COMMENT ON COLUMN topics.created_at IS '创建时间';
COMMENT ON COLUMN topics.updated_by IS '更新人ID';
COMMENT ON COLUMN topics.updated_at IS '更新时间';

-- 通道表
CREATE TABLE channels (
    id UUID NOT NULL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    direction SMALLINT NOT NULL,
    transport SMALLINT NOT NULL,
    auth_type SMALLINT NOT NULL,
    endpoint TEXT NOT NULL,
    accessKey TEXT,
    secretKey TEXT,
    app_id UUID NOT NULL,
    desc TEXT,
    status SMALLINT NOT NULL,
    created_by UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_by UUID,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
)

COMMENT ON TABLE channels IS '通道表';
COMMENT ON COLUMN channels.id IS '通道ID';
COMMENT ON COLUMN channels.name IS '通道名称';
COMMENT ON COLUMN channels.direction IS '通道方向';
COMMENT ON COLUMN channels.transport IS '传输方式';
COMMENT ON COLUMN channels.auth_type IS '认证类型';
COMMENT ON COLUMN channels.endpoint IS '终端地址';
COMMENT ON COLUMN channels.accessKey IS '访问密钥';
COMMENT ON COLUMN channels.secretKey IS '密钥';
COMMENT ON COLUMN channels.app_id IS '应用ID';
COMMENT ON COLUMN channels.desc IS '通道描述';
COMMENT ON COLUMN channels.status IS '通道状态';
COMMENT ON COLUMN channels.created_by IS '创建人ID';
COMMENT ON COLUMN channels.created_at IS '创建时间';
COMMENT ON COLUMN channels.updated_by IS '更新人ID';
COMMENT ON COLUMN channels.updated_at IS '更新时间';
