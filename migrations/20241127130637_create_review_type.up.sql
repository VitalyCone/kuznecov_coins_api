CREATE TABLE review_types(
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

CREATE INDEX idx_review_types_id ON review_types(id);