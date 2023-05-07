-- Filename: migrations/000003_create_reservations_table.up.sql

CREATE TABLE IF NOT EXISTS reservations (
  reservations_id bigserial PRIMARY KEY,
  users_id bigserial REFERENCES "users" ("users_id"),
  reservation_date text NOT NULL,
  reservation_time text NOT NULL,
  duration text NOT NULL,
  people_count int NOT NULL,
  notes varchar(255),
  approval boolean DEFAULT true,
  created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);
