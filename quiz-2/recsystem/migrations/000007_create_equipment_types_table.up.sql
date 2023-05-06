-- Filename: migrations/000015_create_equipment_type_table.up.sql
CREATE TABLE IF NOT EXISTS equipment_types (
  equipment_types_id bigserial PRIMARY KEY,
  type_name varchar(255) NOT NULL
);
