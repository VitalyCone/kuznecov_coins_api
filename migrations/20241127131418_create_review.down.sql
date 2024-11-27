DROP INDEX IF EXISTS idx_reviews_id;
DROP TRIGGER IF EXISTS set_updated_at_reviews ON reviews;
DROP FUNCTION IF EXISTS update_updated_at_column_reviews();
DROP TABLE IF EXISTS reviews;