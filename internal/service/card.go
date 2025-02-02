package service

import (
	"github.com/iremsha/oapicodegen-example/internal/entity"
	"github.com/iremsha/oapicodegen-example/internal/errors"
	"github.com/iremsha/oapicodegen-example/internal/model"
)

type CardRepository interface {
	Save(card *model.Card) error
	FindAll() ([]model.Card, error)
	FindByID(id int64) (*model.Card, error)
}

type CardService struct {
	repo CardRepository
}

func NewCardService(r CardRepository) *CardService {
	return &CardService{repo: r}
}

func (s *CardService) Create(card *entity.Card) error {
	model := model.Card{
		ID:         card.ID,
		Bank:       card.Bank,
		HolderName: card.HolderName,
	}
	if err := s.repo.Save(&model); err != nil {
		return err
	}
	card.ID = model.ID
	return nil
}

func (s *CardService) GetList() ([]entity.Card, error) {
	cardsModel, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var cards []entity.Card
	for _, card := range cardsModel {
		cards = append(cards, entity.Card{
			ID:         card.ID,
			Bank:       card.Bank,
			HolderName: card.HolderName,
		})
	}
	return cards, nil
}

func (s *CardService) GetByID(id int64) (entity.Card, error) {
	card, err := s.repo.FindByID(id)
	if err != nil {
		return entity.Card{}, err
	}

	if card == nil {
		return entity.Card{}, errors.ErrCardNotFound
	}

	return entity.Card{
		ID:         card.ID,
		Bank:       card.Bank,
		HolderName: card.HolderName,
	}, nil
}
