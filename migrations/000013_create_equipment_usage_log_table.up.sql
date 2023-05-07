-- Filename: migrations/000013_create_equipment_usage_log_table.up.sql
CREATE TABLE IF NOT EXISTS equipment_usage_log (
  equipment_usage_log_id bigserial PRIMARY KEY,
  equipments_id bigserial NOT NULL REFERENCES "equipments" ("equipments_id"),
  logs_id bigserial NOT NULL REFERENCES "logs" ("logs_id") ,
  time_borrowed TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
  returned_status boolean NOT NULL
);
