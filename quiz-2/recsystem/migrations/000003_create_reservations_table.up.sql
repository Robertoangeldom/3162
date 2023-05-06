-- Filename: migrations/000003_create_reservations_table.up.sql

CREATE TABLE IF NOT EXISTS reservations (
  reservations_id bigserial PRIMARY KEY,
  users_id bigserial REFERENCES "users" ("users_id"),
  reservation_date date NOT NULL,
  reservation_time time NOT NULL,
  duration int NOT NULL,
  people_count int NOT NULL,
  notes varchar(255),
  approval boolean,
  created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);
