CREATE TABLE reviews(
    id SERIAL PRIMARY KEY,
    review_type_id INT NOT NULL,
    type_id INT NOT NULL,
    rating INT NOT NULL,
    creator_username VARCHAR(32) NOT NULL,  -- Assuming you have a user table
    header VARCHAR(255),
    text TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_review_type FOREIGN KEY (review_type_id) REFERENCES review_types(id) ON DELETE CASCADE,
    CONSTRAINT fk_creator_username FOREIGN KEY (creator_username) REFERENCES users(username) ON DELETE CASCADE
);

CREATE INDEX idx_reviews_id ON reviews(id);

CREATE OR REPLACE FUNCTION update_updated_at_column_reviews()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$
 LANGUAGE plpgsql;

-- Создание триггера, который будет вызываться перед обновлением записи в таблице reviews
CREATE TRIGGER set_updated_at_reviews
BEFORE UPDATE ON reviews
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column_reviews();