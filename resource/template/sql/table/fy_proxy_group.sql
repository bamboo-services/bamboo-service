-- 代理组表
CREATE TABLE fy_proxy_group
(
    group_uuid  UUID PRIMARY KEY,
    user_uuid   UUID         NOT NULL,
    name        VARCHAR(64)  NOT NULL,
    file_name   VARCHAR(64)  NOT NULL DEFAULT 'proxy',
    description VARCHAR(255) NOT NULL DEFAULT '开发者好像很懒，没有写这个代理组的描述',
    proxy       JSONB        NOT NULL DEFAULT '[]',
    partition   JSONB        NOT NULL DEFAULT '[]',
    rule        JSONB        NOT NULL DEFAULT '[]',
    is_enabled  BOOLEAN      NOT NULL DEFAULT FALSE,
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- 评论
COMMENT ON TABLE fy_proxy_group IS '代理组';
COMMENT ON COLUMN fy_proxy_group.group_uuid IS '代理组UUID';
COMMENT ON COLUMN fy_proxy_group.user_uuid IS '用户UUID';
COMMENT ON COLUMN fy_proxy_group.name IS '代理组名称';
COMMENT ON COLUMN fy_proxy_group.file_name IS '代理组文件名';
COMMENT ON COLUMN fy_proxy_group.proxy IS '代理组代理';
COMMENT ON COLUMN fy_proxy_group.description IS '代理组描述';
COMMENT ON COLUMN fy_proxy_group.partition IS '代理组分区';
COMMENT ON COLUMN fy_proxy_group.rule IS '代理组规则';
COMMENT ON COLUMN fy_proxy_group.is_enabled IS '代理组是否启用';
COMMENT ON COLUMN fy_proxy_group.created_at IS '代理组创建时间';
COMMENT ON COLUMN fy_proxy_group.updated_at IS '代理组更新时间';

-- 代理组索引
CREATE INDEX fy_proxy_group_user_uuid_idx ON fy_proxy_group (user_uuid);
CREATE INDEX fy_proxy_group_name_idx ON fy_proxy_group (name);
CREATE INDEX fy_proxy_group_is_enabled_idx ON fy_proxy_group (is_enabled);

-- 触发器
CREATE TRIGGER trigger_fy_proxy_group_updated_at
    BEFORE UPDATE
    ON fy_proxy_group
    FOR EACH ROW
EXECUTE FUNCTION updated_at_column();

-- 外键约束
ALTER TABLE fy_proxy_group
    ADD CONSTRAINT fy_proxy_group_user_uuid_fkey FOREIGN KEY (user_uuid) REFERENCES fy_user (user_uuid)
        ON DELETE CASCADE ON UPDATE CASCADE;