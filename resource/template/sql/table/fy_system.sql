-- 系统表定义
CREATE TABLE fy_system
(
    system_uuid UUID PRIMARY KEY,                               -- 系统唯一标识符
    key         VARCHAR     NOT NULL,                           -- 键
    value       VARCHAR,                                        -- 值
    version     BIGINT      NOT NULL,                           -- 版本
    created_at  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP  -- 更新时间
);

-- 索引定义
CREATE INDEX fy_system_key_idx ON fy_system (key);

-- 评论表定义
COMMENT ON TABLE fy_system IS '系统表';
COMMENT ON COLUMN fy_system.system_uuid IS '系统唯一标识符';
COMMENT ON COLUMN fy_system.key IS '键';
COMMENT ON COLUMN fy_system.value IS '值';
COMMENT ON COLUMN fy_system.version IS '版本';
COMMENT ON COLUMN fy_system.created_at IS '创建时间';
COMMENT ON COLUMN fy_system.updated_at IS '更新时间';

-- 触发器绑定
CREATE TRIGGER trigger_fy_system_updated_at
    BEFORE UPDATE
    ON fy_system
    FOR EACH ROW
EXECUTE FUNCTION updated_at_column();

-- 添加基础数据
INSERT INTO fy_system (system_uuid, key, value, version)
VALUES
    (uuid_in('e721deac-ed22-4f17-a02d-6c324813dc30'), 'author_name', '筱锋xiao_lfeng', 1),
    (uuid_in('69d32777-db0a-4ea5-861b-e1e31f01eeb4'), 'author_chinese_name', '筱锋', 1),
    (uuid_in('fda6079e-2765-450f-b3f9-ec0a3905d5f5'), 'author_english_name', 'xiao_lfeng', 1),
    (uuid_in('1eac6c01-0340-4ea1-b41a-11fc4ac50fc4'), 'author_email', 'gm@x-lf.cn', 1),
    (uuid_in('a2342da0-02ad-494d-835a-4738a70ed39c'), 'author_qq', '1144939537', 1),
    (uuid_in('32287459-9716-42a5-ab1b-ad2c764e0e73'), 'author_link', 'https://www.x-lf.com', 1),
    (uuid_in('a70ab77e-a0f9-49f3-8ad6-09e6794930ae'), 'system_version', '1.0.0-SNAPSHOT', 1),
    (uuid_in('407e7a1e-7cbf-4f44-a65e-ed40d0d25446'), 'system_name', '筱服务', 1),
    (uuid_in('6d6288ce-38a7-4b1d-b65f-c176465710d8'), 'system_description', '提供高可用、高性能的RESTful API接口服务',1),
    (uuid_in('34f898df-1a0d-460f-a749-8c2dc8af915b'), 'system_able_register', 'false', 1),
    (uuid_in('ee86b909-0570-4c3c-b8f0-98d4d0c6fbd9'), 'system_super_admin_uuid', null, 1);