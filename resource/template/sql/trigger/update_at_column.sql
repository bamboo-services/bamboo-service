-- 创建一个触发器函数，用于在行更新时自动更新 updated_at 字段
CREATE OR REPLACE FUNCTION updated_at_column()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP; -- 将 updated_at 设置为当前时间戳
    RETURN NEW; -- 返回修改后的行数据
END;
$$ LANGUAGE plpgsql;