package store

import (
	"database/sql"
	"net/http"

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

// Add card
func (s *Store) AddCard(w http.ResponseWriter, req *http.Request) {

	card := req.URL.Query().Get("card")
	translate := req.URL.Query().Get("translate")

	sqlStatement := `
		INSERT INTO cards (card, trasnlate)
		VALUES ($1, $2)`
	_, err := s.db.Exec(sqlStatement, card, translate)
	if err != nil {
		panic(err)
	}

}

// List cards
func (s *Store) ListCards(w http.ResponseWriter, req *http.Request) {

	card := req.URL.Query().Get("card")
	translate := req.URL.Query().Get("translate")

	sqlStatement := `
		INSERT INTO cards (card, trasnlate)
		VALUES ($1, $2)`
	_, err := s.db.Exec(sqlStatement, card, translate)
	if err != nil {
		panic(err)
	}

}
