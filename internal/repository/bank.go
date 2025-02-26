package repository

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/iremsha/oapicodegen-example/internal/model"
	"gorm.io/gorm"
)

type BankRepository struct {
	db *gorm.DB
}

func NewBankRepository(db *gorm.DB) *BankRepository {
	return &BankRepository{db: db}
}

func (r *BankRepository) Create(bank *model.Bank) (*model.Bank, error) {
	if err := r.db.Create(bank).Error; err != nil {
		return nil, err
	}
	return bank, nil
}

func (r *BankRepository) Update(bank *model.Bank) (*model.Bank, error) {
	if err := r.db.Save(bank).Error; err != nil {
		return nil, err
	}
	return bank, nil
}

func (r *BankRepository) FindByID(id int) (*model.Bank, error) {
	var bank model.Bank
	err := r.db.First(&bank, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &bank, nil
}

func (r *BankRepository) FindAll() (*[]model.Bank, error) {
	var banks []model.Bank
	query := sq.Select("*").From("banks")
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	err = r.db.Raw(sql, args...).Scan(&banks).Error
	return &banks, err
}
