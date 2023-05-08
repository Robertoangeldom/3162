// Filename: internal/models/equipments.go
package models

import (
	"database/sql"
	"fmt"
)

// The EquipmentType model will represent a single equipment type in our equipment_types table
type EquipmentType struct {
	EquipmentName       string
	EquipmentStatus     string
	EquipmentAvailability string
	TypeName            string
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
	SELECT equipments.equ_name, equipments.equ_status, equipments.equ_availability, equipment_types.type_name
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
		err := rows.Scan(&equipmentType.EquipmentName, &equipmentType.EquipmentStatus, &equipmentType.EquipmentAvailability, &equipmentType.TypeName)
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
