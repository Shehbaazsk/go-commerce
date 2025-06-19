-- These columns are nullable by default
ALTER TABLE users
ADD COLUMN created_by INTEGER,
ADD COLUMN updated_by INTEGER;

-- These constraints allow nulls and set them to NULL on delete
ALTER TABLE users
ADD CONSTRAINT fk_created_by FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE SET NULL,
ADD CONSTRAINT fk_updated_by FOREIGN KEY (updated_by) REFERENCES users(id) ON DELETE SET NULL;
