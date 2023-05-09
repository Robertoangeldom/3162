// Filename: internal/models/equipments.go
package models

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

// The EquipmentType model will represent a single equipment type in our equipment_types table
type EquipmentType struct {
	ID                    string
	EquipmentName         string
	EquipmentStatus       string
	EquipmentAvailability string
	TypeName              string
}

// The EquipmentModel type will encapsulate the
// DB connection pool that will be initialized
// in the main() function
type EquipmentModel struct {
	DB *sql.DB
}

// The Display() function retrieves all equipment types from the database
func (m *EquipmentModel) Display() ([]EquipmentType, error) {
	query := `
	SELECT equipments.equipments_id, equipments.equ_name, equipments.equ_status, equipments.equ_availability, equipment_types.type_name
	FROM equipments
	INNER JOIN equipment_types ON equipments.equipment_type_id = equipment_types_id; 
	`
	rows, err := m.DB.Query(query)
	if err != nil {
		fmt.Println("Error querying database:", err)
		return nil, err
	}
	defer rows.Close()

	var equipmentTypes []EquipmentType

	// Iterate over the rows and create a slice of structs
	for rows.Next() {
		var equipmentType EquipmentType
		err := rows.Scan(&equipmentType.ID, &equipmentType.EquipmentName, &equipmentType.EquipmentStatus, &equipmentType.EquipmentAvailability, &equipmentType.TypeName)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		equipmentTypes = append(equipmentTypes, equipmentType)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
		return nil, err
	}

	return equipmentTypes, nil
}

// The Delete() function deletes an equipment type from the database
func (m *EquipmentModel) Delete(id2 string) error {
	var int_id int
	int_id, err := strconv.Atoi(id2)
	if err != nil {
		fmt.Println("Error in converting string to int:", err)
		return err
	}
	query := `
	DELETE FROM equipments
	WHERE equ_name = $1
	`
	result, err := m.DB.Exec(query, int_id)
	if err != nil {
		fmt.Println("Error deleting from database:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error getting affected rows:", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no equipment type with name %s found", id2)
	}

	return nil
}

func (m *EquipmentModel) Update(value, value2, id string) error {

	query := `
		UPDATE equipments
		SET equ_status = $1, equ_availability = $2
		WHERE equ_name = $3;
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := m.DB.ExecContext(ctx, query, value, value2, id)
	if err != nil {
		fmt.Println("Error updating equipment:", err)
		return err
	}

	return nil
}
