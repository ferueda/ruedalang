// Package object contains the core-definitions for objects.
package object

import (
	"fmt"
)

// Pre-defined constant Type
const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ    = "NULL"
)

type ObjectType string

// Object is the interface that all object-types must implmenet.
type Object interface {
	// Type returns the type of this object.
	Type() ObjectType
	// Inspect returns a string-representation of the given object.
	Inspect() string
}

// Integer wraps int64 and implements the Object interface.
type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

// Boolean wraps bool and implements the Object interface.
type Boolean struct {
	Value bool
}

func (i *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (i *Boolean) Inspect() string  { return fmt.Sprintf("%t", i.Value) }

// Null wraps an empty value and implements the Object interface.
type Null struct{}

func (i *Null) Type() ObjectType { return NULL_OBJ }
func (i *Null) Inspect() string  { return "null" }
