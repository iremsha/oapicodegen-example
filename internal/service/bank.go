package service

import (
	"github.com/iremsha/oapicodegen-example/internal/entity"
	"github.com/iremsha/oapicodegen-example/internal/errors"
	"github.com/iremsha/oapicodegen-example/internal/model"
)

type BankRepository interface {
	Save(bank *model.Bank) error
	FindAll() ([]model.Bank, error)
	FindByID(id int64) (*model.Bank, error)
}

type BankService struct {
	repo BankRepository
}

func NewBankService(r BankRepository) *BankService {
	return &BankService{repo: r}
}

func (s *BankService) Create(bank *entity.Bank) error {
	model := model.Bank{
		Name: bank.Name,
	}
	if err := s.repo.Save(&model); err != nil {
		return err
	}
	bank.ID = model.ID
	return nil
}

func (s *BankService) GetList() ([]entity.Bank, error) {
	banksModel, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var banks []entity.Bank
	for _, card := range banksModel {
		banks = append(banks, entity.Bank{
			ID:   card.ID,
			Name: card.Name,
		})
	}
	return banks, nil
}

func (s *BankService) GetByID(id int64) (entity.Bank, error) {
	card, err := s.repo.FindByID(id)
	if err != nil {
		return entity.Bank{}, err
	}

	if card == nil {
		return entity.Bank{}, errors.ErrBankNotFound
	}

	return entity.Bank{
		ID:   card.ID,
		Name: card.Name,
	}, nil
}
