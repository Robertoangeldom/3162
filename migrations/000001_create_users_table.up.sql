-- Filename: migrations/000001_create_users_table.up.sql
CREATE TABLE IF NOT EXISTS users(
  users_id bigserial PRIMARY KEY,
  email VARCHAR(255) UNIQUE NOT NULL,
  first_name text NOT NULL,
  last_name text NOT NULL,
  age int NOT NULL,
  user_address VARCHAR(255) NOT NULL,
  phone_number text NOT NULL,
  roles_id int DEFAULT 2, 
  user_password VARCHAR(255) UNIQUE NOT NULL,
  activated boolean DEFAULT true,
  created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);
