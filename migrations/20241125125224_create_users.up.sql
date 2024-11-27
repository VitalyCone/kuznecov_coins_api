-- create_users.up.sql

CREATE TABLE users (
    id SERIAL PRIMARY KEY,                     -- Уникальный идентификатор пользователя
    username VARCHAR(32) NOT NULL UNIQUE,     -- Имя пользователя (не более 32 символов, уникально)
    password_hash VARCHAR(255) NOT NULL,       -- Хеш пароля (достаточно длинный для хранения хеша bcrypt)
    first_name VARCHAR(50),                    -- Имя пользователя (не более 50 символов)
    second_name VARCHAR(50),                   -- Фамилия пользователя (не более 50 символов)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Дата и время создания записи
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Дата и время последнего обновления
);

-- Индексы для повышения производительности (опционально)
CREATE INDEX idx_username ON users(username);

-- Создание функции для обновления поля updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Создание триггера, который будет вызываться перед обновлением записи
CREATE TRIGGER set_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();