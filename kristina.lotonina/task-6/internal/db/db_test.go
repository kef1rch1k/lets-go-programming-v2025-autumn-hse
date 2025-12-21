package db_test

import (
	"errors"
	"testing"

	"github.com/kef1rch1k/task-6/internal/db"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetNames(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock error: %v", err)
	}
	defer mockDB.Close()

	service := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Ivan").
		AddRow("Petr")

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(rows)

	names, err := service.GetNames()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(names) != 2 {
		t.Fatalf("expected 2 names, got %d", len(names))
	}
}

func TestGetNames_QueryError(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()

	service := db.New(mockDB)

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnError(errors.New("query error"))

	_, err := service.GetNames()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestGetUniqueNames(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()

	service := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Ivan").
		AddRow("Petr")

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(rows)

	names, err := service.GetUniqueNames()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(names) != 2 {
		t.Fatalf("expected 2 names, got %d", len(names))
	}
}

func TestGetUniqueNames_QueryError(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()

	service := db.New(mockDB)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnError(errors.New("query error"))

	_, err := service.GetUniqueNames()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestGetNames_ScanError(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()

	service := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow(nil)

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(rows)

	_, err := service.GetNames()
	if err == nil {
		t.Fatal("expected scan error, got nil")
	}
}

func TestGetNames_RowsError(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()

	service := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Ivan").
		RowError(0, errors.New("rows error"))

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(rows)

	_, err := service.GetNames()
	if err == nil {
		t.Fatal("expected rows error, got nil")
	}
}

func TestGetUniqueNames_ScanError(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()

	service := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow(nil)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(rows)

	_, err := service.GetUniqueNames()
	if err == nil {
		t.Fatal("expected scan error, got nil")
	}
}

func TestGetUniqueNames_RowsError(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()

	service := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Ivan").
		RowError(0, errors.New("rows error"))

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(rows)

	_, err := service.GetUniqueNames()
	if err == nil {
		t.Fatal("expected rows error, got nil")
	}
}

