package pkg

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

// ExecuteSQL is a generic function that can be used to execute any SQL query.
func ExecuteSQL(db *sql.DB, query string, args ...interface{}) (sql.Result, error) {
	return db.Exec(query, args...)
}

// QuerySQL is a generic function that can be used to execute any SELECT SQL query.
// It returns a slice of map[string]interface{}, with each map representing a row in the result set.
// The keys of the map are the column names, and the values are the column values for that row.
func QuerySQL(db *sql.DB, query string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]map[string]interface{}, 0)
	for rows.Next() {
		columns, err := rows.Columns()
		if err != nil {
			return nil, err
		}

		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, err
		}

		row := make(map[string]interface{})
		for i, column := range columns {
			row[column] = values[i]
		}
		result = append(result, row)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// MapToStruct converts a map[string]interface{} to a struct.
// The map keys must match the struct field names.
func MapToStruct(m map[string]interface{}, s interface{}) error {
	// Get the value of the interface as a reflect.Value
	v := reflect.ValueOf(s)
	// Check that the argument is a pointer to a struct
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("argument must be a pointer to a struct, got %T", s)
	}
	// Get the element that the pointer points to
	v = v.Elem()
	// Get the type of the element
	t := v.Type()
	// Iterate over the fields in the struct
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		// Check if the field has a "db" tag
		name := field.Tag.Get("db")
		// If no "db" tag is present, use the field name as the map key
		if name == "" {
			name = field.Name
		}
		// Check if the map has a value for the key
		if value, ok := m[name]; ok {
			// Set the field value to the value in the map
			v.Field(i).Set(reflect.ValueOf(value))
		}
	}
	return nil
}

// CreateMySqlTable creates a new table in the database based on the provided struct.
func CreateMySqlTable(db *sql.DB, tableName string, s interface{}) error {
	// Get the type and value of the struct
	val := reflect.ValueOf(s)
	typ := reflect.TypeOf(s)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return fmt.Errorf("CreateTable: s must be a struct or a pointer to a struct")
	}

	// Build the column definitions
	columnDefs := make([]string, 0)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		name := field.Tag.Get("db")
		if name == "" {
			name = field.Name
		}

		// Get the column type from the field type
		var colType string
		switch field.Type.Kind() {
		case reflect.Bool:
			colType = "BOOLEAN"
		case reflect.Int:
			colType = "INT"
		case reflect.Int8:
			colType = "TINYINT"
		case reflect.Int16:
			colType = "SMALLINT"
		case reflect.Int32:
			colType = "INT"
		case reflect.Int64:
			colType = "BIGINT"
		case reflect.Uint:
			colType = "INT UNSIGNED"
		case reflect.Uint8:
			colType = "TINYINT UNSIGNED"
		case reflect.Uint16:
			colType = "SMALLINT UNSIGNED"
		case reflect.Uint32:
			colType = "INT UNSIGNED"
		case reflect.Uint64:
			colType = "BIGINT UNSIGNED"
		case reflect.Float32:
			colType = "FLOAT"
		case reflect.Float64:
			colType = "DOUBLE"
		case reflect.String:
			colType = "VARCHAR(255)"
		default:
			colType = "VARCHAR(255)"
		}

		// Add any additional column options specified in the field tag
		options := field.Tag.Get("options")
		if options != "" {
			colType += " " + options
		}

		columnDefs = append(columnDefs, fmt.Sprintf("%s %s", name, colType))
	}

	// Build the CREATE TABLE statement
	query := fmt.Sprintf("CREATE TABLE %s (%s)", tableName, strings.Join(columnDefs, ", "))
	// Execute the query
	_, err := db.Exec(query)
	return err
}
