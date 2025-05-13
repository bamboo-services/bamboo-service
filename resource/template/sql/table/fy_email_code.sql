CREATE TABLE fy_email_code
(
    code_uuid  UUID PRIMARY KEY,
    email      VARCHAR     NOT NULL,
    code       VARCHAR(10)  NOT NULL,
    purpose    VARCHAR     NOT NULL,
    expired_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- 邮箱验证码表评论
COMMENT ON TABLE fy_email_code IS '邮箱验证码表';
COMMENT ON COLUMN fy_email_code.code_uuid IS '验证码主键';
COMMENT ON COLUMN fy_email_code.email IS '邮箱地址';
COMMENT ON COLUMN fy_email_code.code IS '验证码';
COMMENT ON COLUMN fy_email_code.purpose IS '验证码用途：register-注册,reset-重置密码,bind-绑定';
COMMENT ON COLUMN fy_email_code.expired_at IS '过期时间';
COMMENT ON COLUMN fy_email_code.created_at IS '创建时间';

-- 邮箱验证码表索引
CREATE INDEX fy_email_code_email_index ON fy_email_code (email);
CREATE INDEX fy_email_code_code_index ON fy_email_code (code);
CREATE INDEX fy_email_code_expired_at_index ON fy_email_code (expired_at);