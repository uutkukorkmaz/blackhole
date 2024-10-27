package blackhole

import (
	"fmt"
	"testing"
)

type dummyGrammar struct {
	baseGrammar
}

// CompileCreateDatabase provides a dummy implementation for creating a database.
func (dg *dummyGrammar) CompileCreateDatabase(name string) (string, error) {
	return fmt.Sprintf("DUMMY CREATE DATABASE %s;", name), nil
}

// CompileDropDatabase provides a dummy implementation for dropping a database.
func (dg *dummyGrammar) CompileDropDatabase(name string) (string, error) {
	return fmt.Sprintf("DUMMY DROP DATABASE %s;", name), nil
}

// CompileCreateTable provides a dummy implementation for creating a table.
func (dg *dummyGrammar) CompileCreateTable(table string, callback func(Blueprint)) (string, error) {
	return fmt.Sprintf("DUMMY CREATE TABLE %s;", table), nil
}

// DefineColumn provides a dummy implementation for defining a column.
func (dg *dummyGrammar) CompileColumn(c *Column) (string, error) {
	return fmt.Sprintf("DUMMY COLUMN %s %s;", c.GetName(), c.GetDataType()), nil
}

// Unit tests for dummyGrammar
func TestDummyGrammar_CompileCreateDatabase(t *testing.T) {
	grammar := &dummyGrammar{}
	dbName := "test_db"
	expected := "DUMMY CREATE DATABASE test_db;"

	result, err := grammar.CompileCreateDatabase(dbName)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestDummyGrammar_CompileDropDatabase(t *testing.T) {
	grammar := &dummyGrammar{}
	dbName := "test_db"
	expected := "DUMMY DROP DATABASE test_db;"

	result, err := grammar.CompileDropDatabase(dbName)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestDummyGrammar_CompileCreateTable(t *testing.T) {
	grammar := &dummyGrammar{}
	tableName := "test_table"
	expected := "DUMMY CREATE TABLE test_table;"

	result, err := grammar.CompileCreateTable(tableName, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestDummyGrammar_DefineColumn(t *testing.T) {
	grammar := &dummyGrammar{}
	column := integerColumn("id")
	expected := "DUMMY COLUMN id integer;"

	result, err := grammar.CompileColumn(column)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}
