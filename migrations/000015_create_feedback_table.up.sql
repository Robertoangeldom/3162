-- Filename: migrations/000009_create_feedback_table.up.sql
CREATE TABLE IF NOT EXISTS feedback (
  feedback_id bigserial PRIMARY KEY,
  users_id bigserial NOT NULL REFERENCES "users" ("users_id"),
  star_count int NOT NULL,
  feedback_message varchar(255) NOT NULL,
  created_at timestamp NOT NULL
);
