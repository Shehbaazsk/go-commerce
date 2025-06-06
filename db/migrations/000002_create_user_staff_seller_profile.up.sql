-- Customer Profile
CREATE TABLE IF NOT EXISTS customer_profiles (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    contact_preference JSONB DEFAULT '{}'::JSONB
);

-- Departments

CREATE TABLE departments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    manager_id INTEGER NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE
);

-- Positions
CREATE TABLE positions (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) UNIQUE NOT NULL,
    description TEXT
);

-- Staff Profile
CREATE TABLE IF NOT EXISTS staff_profiles (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    employee_id VARCHAR(100) NOT NULL UNIQUE,
    department_id INTEGER NOT NULL REFERENCES departments(id) ON DELETE SET NULL,
    position_id INTEGER NOT NULL REFERENCES positions(id) ON DELETE SET NULL,
    joining_date DATE
);

-- Seller Profile
CREATE TABLE IF NOT EXISTS seller_profiles (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    store_name VARCHAR(255) NOT NULL,
    gst_number VARCHAR(100),
    average_rating DECIMAL(3, 2) DEFAULT 0.00
);
