-- 角色表定义
CREATE TABLE fy_role
(
    role_uuid        UUID PRIMARY KEY,
    role_name        VARCHAR(64)  NOT NULL,
    role_nickname    VARCHAR(64)  NOT NULL,
    role_description VARCHAR(255) NOT NULL DEFAULT '开发者好像很懒，没有写这个角色的描述',
    role_permission  JSONB        NOT NULL DEFAULT '[]',
    role_status      BOOLEAN      NOT NULL DEFAULT TRUE,
    created_at       TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- 角色权限表定义
COMMENT ON TABLE fy_role IS '角色表';
COMMENT ON COLUMN fy_role.role_name IS '角色名称';
COMMENT ON COLUMN fy_role.role_nickname IS '角色昵称';
COMMENT ON COLUMN fy_role.role_description IS '角色描述';
COMMENT ON COLUMN fy_role.role_permission IS '角色权限';
COMMENT ON COLUMN fy_role.role_status IS '角色状态(开启和关闭)';
COMMENT ON COLUMN fy_role.created_at IS '创建时间';
COMMENT ON COLUMN fy_role.updated_at IS '更新时间';

-- 角色表索引定义
CREATE INDEX fy_role_role_name_idx ON fy_role (role_name);
CREATE INDEX fy_role_role_nickname_idx ON fy_role (role_nickname);
CREATE INDEX fy_role_role_status_idx ON fy_role (role_status);

-- 预先构建部分角色
INSERT INTO fy_role (role_uuid, role_name, role_nickname, role_description, role_permission, role_status)
VALUES
    (uuid_in('60adc6c5-ae95-4051-b408-31a376654cf9'), 'super_admin',  '超级管理员', '超级管理员，拥有最高权限', jsonb_build_array(), TRUE),
    (uuid_in('fec12068-0c50-4198-8ece-968599ea4155'), 'admin',  '管理员', '管理员，可以管理系统内部的内容', jsonb_build_array(), TRUE),
    (uuid_in('058db8f6-664b-4d2d-99dc-1bdc3404da95'), 'duser',  '普通用户', '普通用户，可以进行部分操作', jsonb_build_array(), TRUE);