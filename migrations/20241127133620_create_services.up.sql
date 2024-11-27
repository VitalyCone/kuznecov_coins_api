-- Создание таблицы services
-- migrations/20231010_create_services_table.sql

CREATE TABLE services (
    id SERIAL PRIMARY KEY,
    company_id INT NOT NULL,
    service_type_id INT NOT NULL,
    text TEXT NOT NULL,
    price FLOAT8 NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE,
    FOREIGN KEY (service_type_id) REFERENCES service_types(id) ON DELETE CASCADE
);

CREATE INDEX idx_services_id ON services(id);
CREATE INDEX idx_services_company_id ON services(company_id);
CREATE INDEX idx_services_service_type_id ON services(service_type_id);

-- Создание функции для обновления поля updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column_services()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$
 LANGUAGE plpgsql;

-- Создание триггера, который будет вызываться перед обновлением записи в таблице services
CREATE TRIGGER set_updated_at_services
BEFORE UPDATE ON services
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column_services();