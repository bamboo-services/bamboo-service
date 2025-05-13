-- 代理订阅地址表
CREATE TABLE fy_proxy_subscription
(
    subscription_uuid UUID PRIMARY KEY,
    proxy_group_uuid  UUID         NOT NULL,
    user_uuid         UUID         NOT NULL,
    name              VARCHAR(64)  NOT NULL,
    merchant          VARCHAR(64)  NOT NULL,
    description       VARCHAR(255) NOT NULL DEFAULT '开发者好像很懒，没有写这个代理订阅地址的描述',
    url               VARCHAR      NOT NULL,
    content           TEXT,
    created_at        TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at        TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    subscribe_at      TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- 评论
COMMENT ON TABLE fy_proxy_subscription IS '代理订阅地址表';
COMMENT ON COLUMN fy_proxy_subscription.subscription_uuid IS '代理订阅地址唯一标识符';
COMMENT ON COLUMN fy_proxy_subscription.proxy_group_uuid IS '代理组唯一标识符';
COMMENT ON COLUMN fy_proxy_subscription.user_uuid IS '用户唯一标识符';
COMMENT ON COLUMN fy_proxy_subscription.name IS '代理订阅地址名称';
COMMENT ON COLUMN fy_proxy_subscription.merchant IS '代理订阅地址商户';
COMMENT ON COLUMN fy_proxy_subscription.description IS '代理订阅地址描述';
COMMENT ON COLUMN fy_proxy_subscription.url IS '代理订阅地址URL';
COMMENT ON COLUMN fy_proxy_subscription.content IS '代理订阅地址所订阅获取的内容';
COMMENT ON COLUMN fy_proxy_subscription.created_at IS '代理订阅地址创建时间';
COMMENT ON COLUMN fy_proxy_subscription.updated_at IS '代理订阅地址更新时间';
COMMENT ON COLUMN fy_proxy_subscription.subscribe_at IS '代理订阅地址订阅时间';

-- 触发器
CREATE TRIGGER trigger_fy_proxy_subscription_updated_at
    BEFORE UPDATE
    ON fy_proxy_subscription
    FOR EACH ROW
EXECUTE FUNCTION updated_at_column();

-- 外键约束
ALTER TABLE fy_proxy_subscription
    ADD CONSTRAINT fy_proxy_subscription_user_uuid_fkey FOREIGN KEY (user_uuid) REFERENCES fy_user (user_uuid)
        ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE fy_proxy_subscription
    ADD CONSTRAINT fy_proxy_subscription_proxy_group_uuid_fkey FOREIGN KEY (proxy_group_uuid) REFERENCES fy_proxy_group (group_uuid)
        ON DELETE CASCADE ON UPDATE CASCADE;
