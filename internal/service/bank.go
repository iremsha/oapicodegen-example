package service

import (
	"github.com/iremsha/oapicodegen-example/internal/entity"
	"github.com/iremsha/oapicodegen-example/internal/model"
)

type BankRepository interface {
	Create(bank *model.Bank) (*model.Bank, error)
	Update(bank *model.Bank) (*model.Bank, error)
	FindByID(id int) (*model.Bank, error)
	FindAll() (*[]model.Bank, error)
}

type BankService struct {
	repo BankRepository
}

func NewBankService(r BankRepository) *BankService {
	return &BankService{repo: r}
}

func (s *BankService) Create(req *entity.BankRequest) (*entity.BankResponse, error) {
	model := model.Bank{
		Name:    req.Name,
		Address: req.Address,
	}

	result, err := s.repo.Create(&model)
	if err != nil {
		return nil, err
	}

	res := &entity.BankResponse{
		ID:      result.ID,
		Name:    result.Name,
		Address: result.Address,
	}

	return res, nil
}

func (s *BankService) Update(id int, req *entity.BankRequest) (*entity.BankResponse, error) {
	model := model.Bank{
		ID:      id,
		Name:    req.Name,
		Address: req.Address,
	}

	result, err := s.repo.Update(&model)
	if err != nil {
		return nil, err
	}

	res := &entity.BankResponse{
		ID:      result.ID,
		Name:    result.Name,
		Address: result.Address,
	}

	return res, nil
}

func (s *BankService) GetByID(id int) (*entity.BankResponse, error) {
	result, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	res := &entity.BankResponse{
		ID:      result.ID,
		Name:    result.Name,
		Address: result.Address,
	}

	return res, nil
}

func (s *BankService) GetList() (*[]entity.BankResponse, error) {
	panic("unimplemented")
}
