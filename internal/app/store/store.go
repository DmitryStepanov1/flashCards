package store

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/DmitryStepanov1/flashCards/internal/app/model"
	_ "github.com/lib/pq"
)

type Store struct {
	config         *Config
	db             *sql.DB
	cardRepository *CardRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {

	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

// Card
func (s *Store) Card() *CardRepository {
	if s.cardRepository != nil {
		return s.cardRepository
	}

	s.cardRepository = &CardRepository{
		store: s,
	}

	return s.cardRepository
}

// List cards
func (s *Store) ListCards() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var cards []model.Card

		rows, err := s.db.Query("SELECT * FROM fc_test.cards")

		if err != nil {
			panic(err)
		}

		defer rows.Close()

		for rows.Next() {
			var card model.Card
			if err := rows.Scan(&card.ID, &card.Word, &card.Translate); err != nil {
				fmt.Println(err)
			}
			cards = append(cards, card)
		}
		fmt.Println(cards)
	}

}
