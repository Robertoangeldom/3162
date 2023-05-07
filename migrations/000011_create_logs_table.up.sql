-- Filename: migrations/000011_create_logs_table.up.sql
CREATE TABLE IF NOT EXISTS logs (
  logs_id bigserial PRIMARY KEY,
  users_id bigserial REFERENCES "users" ("users_id"),
  purpose varchar(255) NOT NULL,
  duration int NOT NULL,
  sign_in_time TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);
