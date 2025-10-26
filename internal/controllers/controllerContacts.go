package controllers

import (
	"database/sql"
	"fmt"
	"main/internal/models"

	_ "github.com/mattn/go-sqlite3"
)

func CreateContactsTable(db *sql.DB) error {
	createTableSQL := `CREATE TABLE IF NOT EXISTS contacts (
		"id" INTEGER PRIMARY KEY ,
		"first_name" TEXT,
		"last_name" TEXT,
		"email" TEXT,
		"age" INTEGER
	  );`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf("error creating contacts table: %w", err)
	}
	return nil
}

func InsertContact(db *sql.DB, firstName, lastName, email string, age uint64) error {
	insertSQL := `INSERT INTO contacts (first_name, last_name, email, age) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(insertSQL, firstName, lastName, email, age)
	if err != nil {
		return fmt.Errorf("error inserting contact: %w", err)
	}
	return nil
}

func GetContactByID(db *sql.DB, id uint64) (models.Contact, error) {
	contact := models.Contact{}
	searchSQL := `SELECT id, first_name, last_name, email, age FROM contacts WHERE id = ?`
	err := db.QueryRow(searchSQL, id).Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Email, &contact.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Contact{ID: 0}, fmt.Errorf("no contact found with the given ID")
		}
		return models.Contact{ID: 0}, fmt.Errorf("error retrieving contact: %w", err)
	}
	return contact, nil
}

func UpdateContact(db *sql.DB, id uint64, firstName, lastName, email string, age uint64) error {
	updateSQL := `UPDATE contacts SET first_name = ?, last_name = ?, email = ?, age = ? WHERE id = ?`
	result, err := db.Exec(updateSQL, firstName, lastName, email, age, id)
	if err != nil {
		return fmt.Errorf("error updating contact: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no contact found with the given ID")
	}
	return nil
}

func DeleteContact(db *sql.DB, id uint64) error {
	deleteSQL := `DELETE FROM contacts WHERE id = ?`
	result, err := db.Exec(deleteSQL, id)
	if err != nil {
		return fmt.Errorf("error deleting contact: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no contact found with the given ID")
	}
	return nil
}

func GetAllContacts(db *sql.DB) ([]models.Contact, error) {
	searchSQL := `SELECT id, first_name, last_name, email, age FROM contacts`
	rows, err := db.Query(searchSQL)
	if err != nil {
		return nil, fmt.Errorf("error querying contacts: %w", err)
	}
	defer rows.Close()

	var contacts []models.Contact
	for rows.Next() {
		c := models.Contact{}
		if err := rows.Scan(&c.ID, &c.FirstName, &c.LastName, &c.Email, &c.Age); err != nil {
			return nil, fmt.Errorf("error scanning contact: %w", err)
		}
		contacts = append(contacts, c)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over contacts: %w", err)
	}
	return contacts, nil
}
func SearchedContacts(db *sql.DB, searchTerm string) ([]models.Contact, error) {
	contacts := []models.Contact{}
	searchSQL := `SELECT id, first_name, last_name, email, age FROM contacts WHERE CAST(id AS TEXT) LIKE ? OR UPPER(first_name) LIKE ? OR UPPER(last_name) LIKE ? OR UPPER(email) LIKE ? OR CAST(age AS TEXT) LIKE ?`
	rows, err := db.Query(searchSQL, "%"+searchTerm+"%", "%"+searchTerm+"%", "%"+searchTerm+"%", "%"+searchTerm+"%", "%"+searchTerm+"%")
	if err != nil {
		return nil, fmt.Errorf("error querying contacts: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		contact := models.Contact{}
		if err := rows.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Email, &contact.Age); err != nil {
			return nil, fmt.Errorf("error scanning contact: %w", err)
		}
		contacts = append(contacts, contact)
	}

	return contacts, nil
}
