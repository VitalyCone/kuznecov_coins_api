-- Создание таблицы companies
CREATE TABLE companies (
    id SERIAL PRIMARY KEY,
    avatar BYTEA,  -- Хранение изображения в формате bytea
    name VARCHAR(100) NOT NULL UNIQUE,  -- Имя компании (не более 100 символов, уникально)
    description TEXT,  -- Описание компании
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_companies_id ON companies(id);
CREATE INDEX idx_companies_name ON companies(name);

-- Создание функции для обновления поля updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column_companies()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$
 LANGUAGE plpgsql;

-- Создание триггера, который будет вызываться перед обновлением записи в таблице companies
CREATE TRIGGER set_updated_at_companies
BEFORE UPDATE ON companies
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column_companies();