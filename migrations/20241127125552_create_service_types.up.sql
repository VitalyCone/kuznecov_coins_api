CREATE TABLE service_types(
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

CREATE INDEX idx_service_types_id ON service_types(id);