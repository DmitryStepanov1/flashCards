package store_test

import (
	"testing"

	"github.com/DmitryStepanov1/flashCards/internal/app/model"
	"github.com/DmitryStepanov1/flashCards/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestCardRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("cards")

	c, err := s.Card().Create(&model.Card{
		ID:        123,
		Word:      "Greeting",
		Translate: "qweqwe",
	})

	assert.NoError(t, err)
	assert.NotNil(t, c)
}

func TestCardRepository_FindByWord(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("cards")

	word := "ttt"
	_, err := s.Card().FindByWord(word)

	assert.Error(t, err)
}
