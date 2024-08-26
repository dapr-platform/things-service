package entity

type DbColumnInfo struct {
	TableSchema  string      `json:"table_schema"`
	TableName    string      `json:"table_name"`
	Position     int         `json:"position"`
	ColumnName   string      `json:"column_name"`
	DataType     string      `json:"data_type"`
	MaxLength    int         `json:"max_length"`
	IsNullable   string      `json:"is_nullable"`
	IsGenerated  string      `json:"is_generated"`
	IsUpdatable  string      `json:"is_updatable"`
	DefaultValue interface{} `json:"default_value"`
}
