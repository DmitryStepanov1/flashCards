package store

import "github.com/DmitryStepanov1/flashCards/internal/app/model"

type CardRepository struct {
	store *Store
}

func (r *CardRepository) Create(c *model.Card) (*model.Card, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO cards (word, translate) VALUES ($1, $2) RETURNING id",
		c.Word,
		c.Translate,
	).Scan(&c.ID); err != nil {
		return nil, err
	}

	return c, nil
}
