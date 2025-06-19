-- Drop foreign key constraints first
ALTER TABLE users DROP CONSTRAINT IF EXISTS fk_created_by;
ALTER TABLE users DROP CONSTRAINT IF EXISTS fk_updated_by;

-- Drop the columns
ALTER TABLE users DROP COLUMN IF EXISTS created_by;
ALTER TABLE users DROP COLUMN IF EXISTS updated_by;
