-- 用户表定义
CREATE TABLE fy_user
(
    -- 核心字段 (来自 DTO)
    user_uuid          UUID PRIMARY KEY,                                       -- 用户唯一标识符 (主键), 建议由应用层生成以保持一致性
    username           VARCHAR(255) UNIQUE NOT NULL,                           -- 用户名, 唯一且不能为空
    email              VARCHAR(255) UNIQUE NOT NULL,                           -- 电子邮箱, 唯一且不能为空
    phone              VARCHAR(30) UNIQUE,                                     -- 手机号码, 唯一 (可选)
    role               VARCHAR(50)         NOT NULL DEFAULT 'user',            -- 用户角色, 不能为空, 默认为 'user'
    permissions        JSONB,                                                  -- 用户权限, 使用 JSONB 类型以支持复杂结构和高效查询
    created_at         TIMESTAMPTZ         NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 记录创建时间, 不能为空, 默认为当前时间戳
    updated_at         TIMESTAMPTZ         NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 记录更新时间, 不能为空, 默认为当前时间戳

    -- 扩展字段：认证与安全
    password_hash      VARCHAR(255)        NOT NULL,                           -- 存储用户密码的哈希值, 不能为空
    email_verified_at  TIMESTAMPTZ,                                            -- 邮箱验证时间 (可选)
    phone_verified_at  TIMESTAMPTZ,                                            -- 手机验证时间 (可选)
    two_factor_enabled BOOLEAN             NOT NULL DEFAULT FALSE,             -- 是否启用两因素认证, 默认为 false
    two_factor_secret  VARCHAR(255),                                           -- 两因素认证密钥 (例如 TOTP secret), 在启用时设置

    -- 扩展字段：用户个人信息
    nickname           VARCHAR(100),                                           -- 昵称 (可选)
    avatar_url         TEXT,                                                   -- 用户头像链接 (可选)
    gender             VARCHAR(10),                                            -- 性别 (可选, 例如: 'male', 'female', 'other', 'unknown')
    birth_date         DATE,                                                   -- 出生日期 (可选)
    bio                TEXT,                                                   -- 个人简介 (可选)
    qq_email           VARCHAR(255) UNIQUE,                                    -- QQ邮箱 (可选, 唯一)

    -- 扩展字段：账户状态与活动
    status             VARCHAR(20)         NOT NULL DEFAULT 'active',          -- 用户账户状态 (例如: 'active', 'inactive', 'pending_verification', 'banned'), 默认为 'active'
    last_login_at      TIMESTAMPTZ,                                            -- 最后登录时间 (可选)
    last_login_ip      VARCHAR(45),                                            -- 最后登录IP地址 (支持IPv4和IPv6) (可选)
    registration_ip    VARCHAR(45),                                            -- 注册时IP地址 (可选)

    -- 扩展字段：软删除
    deleted_at         TIMESTAMPTZ                                             -- 软删除标记, 记录删除时间 (可选, 用于实现软删除功能)
);

-- 为常用查询字段创建索引
CREATE INDEX idx_fy_user_role ON fy_user (role);
CREATE INDEX idx_fy_user_status ON fy_user (status);
CREATE INDEX idx_fy_user_deleted_at ON fy_user (deleted_at); -- 配合软删除查询
CREATE INDEX idx_fy_user_qq_email ON fy_user (qq_email) WHERE qq_email IS NOT NULL;

-- 为 JSONB 类型的 permissions 字段创建 GIN 索引，以提高查询性能
CREATE INDEX idx_fy_user_permissions ON fy_user USING GIN (permissions);

-- 为每个字段添加注释
COMMENT ON TABLE fy_user IS '用户表';
COMMENT ON COLUMN fy_user.user_uuid IS '用户唯一标识符';
COMMENT ON COLUMN fy_user.username IS '用户名';
COMMENT ON COLUMN fy_user.email IS '电子邮箱';
COMMENT ON COLUMN fy_user.phone IS '手机号码';
COMMENT ON COLUMN fy_user.role IS '用户角色';
COMMENT ON COLUMN fy_user.permissions IS '用户权限';
COMMENT ON COLUMN fy_user.created_at IS '记录创建时间';
COMMENT ON COLUMN fy_user.updated_at IS '记录更新时间';
COMMENT ON COLUMN fy_user.password_hash IS '密码哈希值';
COMMENT ON COLUMN fy_user.email_verified_at IS '邮箱验证时间';
COMMENT ON COLUMN fy_user.phone_verified_at IS '手机验证时间';
COMMENT ON COLUMN fy_user.two_factor_enabled IS '是否启用两因素认证';
COMMENT ON COLUMN fy_user.two_factor_secret IS '两因素认证密钥';
COMMENT ON COLUMN fy_user.nickname IS '用户昵称';
COMMENT ON COLUMN fy_user.avatar_url IS '头像URL';
COMMENT ON COLUMN fy_user.gender IS '性别';
COMMENT ON COLUMN fy_user.birth_date IS '出生日期';
COMMENT ON COLUMN fy_user.bio IS '个人简介';
COMMENT ON COLUMN fy_user.qq_email IS 'QQ邮箱';
COMMENT ON COLUMN fy_user.status IS '用户账户状态';
COMMENT ON COLUMN fy_user.last_login_at IS '最后登录时间';
COMMENT ON COLUMN fy_user.last_login_ip IS '最后登录IP地址';
COMMENT ON COLUMN fy_user.registration_ip IS '注册IP地址';
COMMENT ON COLUMN fy_user.deleted_at IS '删除时间（软删除）';

-- 将触发器绑定到 fy_user 表的 UPDATE 操作上
CREATE TRIGGER trigger_fy_user_updated_at
    BEFORE UPDATE
    ON fy_user
    FOR EACH ROW
EXECUTE FUNCTION updated_at_column();