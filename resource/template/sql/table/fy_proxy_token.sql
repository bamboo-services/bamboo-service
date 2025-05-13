-- 代理授权令牌表
CREATE TABLE fy_proxy_token
(
    proxy_token_uuid UUID PRIMARY KEY,
    user_uuid        UUID         NOT NULL,
    name             VARCHAR(64)  NOT NULL,
    description      VARCHAR(255) NOT NULL DEFAULT '开发者好像很懒，没有写这个代理令牌的描述',
    created_at       TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expired_at        TIMESTAMPTZ  NOT NULL
);

-- 评论
COMMENT ON TABLE fy_proxy_token IS '代理令牌表';
COMMENT ON COLUMN fy_proxy_token.proxy_token_uuid IS '代理令牌UUID';
COMMENT ON COLUMN fy_proxy_token.user_uuid IS '用户UUID';
COMMENT ON COLUMN fy_proxy_token.name IS '代理令牌名称';
COMMENT ON COLUMN fy_proxy_token.description IS '代理令牌描述';
COMMENT ON COLUMN fy_proxy_token.created_at IS '代理令牌创建时间';
COMMENT ON COLUMN fy_proxy_token.expired_at IS '代理令牌过期时间';

-- 索引
CREATE INDEX fy_proxy_token_proxy_token_uuid_idx ON fy_proxy_token (proxy_token_uuid);
CREATE INDEX fy_proxy_token_name_idx ON fy_proxy_token (name);
CREATE INDEX fy_proxy_token_user_uuid_idx ON fy_proxy_token (user_uuid);

-- 外键约束
ALTER TABLE fy_proxy_token
    ADD CONSTRAINT fy_proxy_token_user_uuid_fkey FOREIGN KEY (user_uuid) REFERENCES fy_user (user_uuid)
        ON UPDATE CASCADE ON DELETE CASCADE;