
DROP TRIGGER IF EXISTS set_addresses_updated_at ON addresses;
DROP TRIGGER IF EXISTS set_cities_updated_at ON cities;
DROP TRIGGER IF EXISTS set_states_updated_at ON states;
DROP TRIGGER IF EXISTS set_countries_updated_at ON countries;
DROP TRIGGER IF EXISTS set_users_updated_at ON users;

-- Drop tables in reverse order of creation (to handle foreign key dependencies)
DROP TABLE IF EXISTS addresses;
DROP TABLE IF EXISTS cities;
DROP TABLE IF EXISTS states;
DROP TABLE IF EXISTS countries;
DROP TABLE IF EXISTS users;

-- Drop the shared timestamp function
DROP FUNCTION IF EXISTS update_timestamp();