-- Filename: migrations/000017_create_roles_table.up.sql
CREATE TABLE IF NOT EXISTS roles (
  roles_id bigserial PRIMARY KEY,
  roles int NOT NULL
);
