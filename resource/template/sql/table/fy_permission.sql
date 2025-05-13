-- 权限表定义
CREATE TABLE fy_permission
(
    permission_key         VARCHAR PRIMARY KEY,
    permission_name        VARCHAR     NOT NULL,
    permission_description VARCHAR     NOT NULL DEFAULT '开发者好像很懒，没有写这个权限的描述',
    permission_status      BOOLEAN     NOT NULL DEFAULT TRUE,
    created_at             TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at             TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- 权限表评论
COMMENT ON TABLE fy_permission IS '权限表';
COMMENT ON COLUMN fy_permission.permission_key IS '权限标识';
COMMENT ON COLUMN fy_permission.permission_name IS '权限名称';
COMMENT ON COLUMN fy_permission.permission_description IS '权限描述';
COMMENT ON COLUMN fy_permission.permission_status IS '权限状态';
COMMENT ON COLUMN fy_permission.created_at IS '创建时间';
COMMENT ON COLUMN fy_permission.updated_at IS '更新时间';

-- 将触发器绑定到 fy_permission 表的 UPDATE 操作上
CREATE TRIGGER trigger_fy_permission_updated_at
    BEFORE UPDATE
    ON fy_permission
    FOR EACH ROW
EXECUTE FUNCTION updated_at_column();

-- 权限表索引
CREATE INDEX fy_permission_key_index ON fy_permission (permission_key);
CREATE INDEX fy_permission_name_index ON fy_permission (permission_name);
CREATE INDEX fy_permission_status_index ON fy_permission (permission_status);