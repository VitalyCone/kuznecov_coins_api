DROP INDEX IF EXISTS idx_services_id;
DROP INDEX IF EXISTS idx_services_company_id;
DROP INDEX IF EXISTS idx_services_service_type_id;
DROP TRIGGER IF EXISTS set_updated_at_services ON services;
DROP FUNCTION IF EXISTS update_updated_at_column_services();
DROP TABLE IF EXISTS services;