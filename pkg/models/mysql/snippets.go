package mysql

import (
	"database/sql"

	"github.com/edbar42/snippetbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

// Inserts a snippet into the database
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	statement := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(statement, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Returns a specific snippet given an id
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	statement := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(statement, id)

	s := &models.Snippet{}
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return s, nil
}

// Returns the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	statement := `SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	snippets := []*models.Snippet{}

	for rows.Next() {
		s := &models.Snippet{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}
	// Don`t assume that a successful iteration was completed over the whole result set
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
