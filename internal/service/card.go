package service

import (
	"github.com/iremsha/oapicodegen-example/internal/entity"
	"github.com/iremsha/oapicodegen-example/internal/model"
)

type CardRepository interface {
	Create(bank *model.Card) (*model.Card, error)
	Update(bank *model.Card) (*model.Card, error)
	FindAll() (*[]model.Card, error)
	FindByID(id int) (*model.Card, error)
}

type CardService struct {
	repo CardRepository
}

func NewCardService(r CardRepository) *CardService {
	return &CardService{repo: r}
}

func (s *CardService) Create(req *entity.CardRequest) (*entity.CardResponse, error) {
	model := model.Card{
		Name: req.Name,
		Type: req.Type,
		Cvv:  req.Cvv,
	}

	result, err := s.repo.Create(&model)
	if err != nil {
		return nil, err
	}

	res := &entity.CardResponse{
		ID:   result.ID,
		Name: result.Name,
		Type: result.Type,
		Cvv:  result.Cvv,
	}

	return res, nil
}

func (s *CardService) Update(id int, req *entity.CardRequest) (*entity.CardResponse, error) {
	model := model.Card{
		ID:   id,
		Name: req.Name,
		Type: req.Type,
		Cvv:  req.Cvv,
	}

	result, err := s.repo.Update(&model)
	if err != nil {
		return nil, err
	}

	res := &entity.CardResponse{
		ID:   result.ID,
		Name: result.Name,
		Type: result.Type,
		Cvv:  result.Cvv,
	}

	return res, nil
}

func (s *CardService) GetByID(id int) (*entity.CardResponse, error) {
	panic("unimplemented")
}

func (s *CardService) GetList() (*[]entity.CardResponse, error) {
	panic("unimplemented")
}
