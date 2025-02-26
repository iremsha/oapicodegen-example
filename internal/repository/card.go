package repository

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/iremsha/oapicodegen-example/internal/model"
	"gorm.io/gorm"
)

type CardRepository struct {
	db *gorm.DB
}

func NewCardRepository(db *gorm.DB) *CardRepository {
	return &CardRepository{db: db}
}

func (r *CardRepository) Create(card *model.Card) (*model.Card, error) {
	if err := r.db.Create(card).Error; err != nil {
		return nil, err
	}
	return card, nil
}

func (r *CardRepository) Update(card *model.Card) (*model.Card, error) {
	if err := r.db.Save(card).Error; err != nil {
		return nil, err
	}
	return card, nil
}

func (s *CardRepository) FindAll() (*[]model.Card, error) {
	var cards []model.Card
	query := sq.Select("*").From("cards")
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	err = s.db.Raw(sql, args...).Scan(&cards).Error
	return &cards, err
}

func (r *CardRepository) FindByID(id int) (*model.Card, error) {
	var card model.Card
	err := r.db.First(&card, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &card, nil
}
