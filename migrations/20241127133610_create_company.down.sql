DROP INDEX IF EXISTS idx_companies_id;
DROP INDEX IF EXISTS idx_companies_name;
DROP TRIGGER IF EXISTS set_updated_at_companies ON companies;
DROP FUNCTION IF EXISTS update_updated_at_column_companies();

-- Удаление таблицы companies
DROP TABLE IF EXISTS companies;