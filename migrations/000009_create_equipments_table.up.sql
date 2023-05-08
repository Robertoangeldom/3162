-- Filename: migrations/000007_create_equipments_table.up.sql
CREATE TABLE IF NOT EXISTS equipments (
  equipments_id bigserial PRIMARY KEY,
  equ_name varchar(255) NOT NULL,
  equ_image bytea,
  equipment_type_id int NOT NULL REFERENCES equipment_types (equipment_types_id),
  equ_status text NOT NULL,
  equ_availability text NOT NULL 
);
