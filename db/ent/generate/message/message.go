// Code generated by ent, DO NOT EDIT.

package message

import (
	"time"
)

const (
	// Label holds the string label denoting the message type in the database.
	Label = "message"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldLong holds the string denoting the long field in the database.
	FieldLong = "long"
	// FieldPriority holds the string denoting the priority field in the database.
	FieldPriority = "priority"
	// FieldSequenceID holds the string denoting the sequenceid field in the database.
	FieldSequenceID = "sequence_id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the message in the database.
	Table = "messages"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "messages"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_messages"
)

// Columns holds all SQL columns for message fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldTitle,
	FieldContent,
	FieldLong,
	FieldPriority,
	FieldSequenceID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "messages"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_messages",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// ContentValidator is a validator for the "content" field. It is called by the builders before save.
	ContentValidator func(string) error
	// PriorityValidator is a validator for the "priority" field. It is called by the builders before save.
	PriorityValidator func(string) error
)