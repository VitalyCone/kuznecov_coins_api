-- create_users.down.sql

-- Удаление триггера, если он существует
DROP TRIGGER IF EXISTS set_updated_at ON users;

-- Удаление функции, если она существует
DROP FUNCTION IF EXISTS update_updated_at_column();

-- Удаление таблицы users
DROP TABLE IF EXISTS users;