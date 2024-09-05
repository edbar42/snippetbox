package mysql

import (
	"database/sql"

	"github.com/edbar42/snippetbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

// Returns a specific snippet given an id
func (m *SnippetModel) Get(id int) (*models.Snipet, error) {
	return nil, nil
}

// Returns the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*models.Snipet, error) {
	return nil, nil
}
