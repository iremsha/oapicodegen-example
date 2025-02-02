package service

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/iremsha/oapicodegen-example/internal/entity"
	"github.com/iremsha/oapicodegen-example/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestCardService_FindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCardRepo := NewMockCardRepository(ctrl)
	service := NewCardService(mockCardRepo)

	t.Run("full", func(t *testing.T) {
		mockedCards := []model.Card{
			{ID: 1, Bank: "Bank A", HolderName: "HolderName A"},
			{ID: 1, Bank: "Bank B", HolderName: "HolderName B"},
		}
		expectedCards := []entity.Card{
			{ID: 1, Bank: "Bank A", HolderName: "HolderName A"},
			{ID: 1, Bank: "Bank B", HolderName: "HolderName B"},
		}

		mockCardRepo.EXPECT().FindAll().Return(mockedCards, nil)

		cards, err := service.GetList()

		assert.Nil(t, err)
		assert.Equal(t, expectedCards, cards)
	})

	t.Run("empty", func(t *testing.T) {
		mockedCards := []model.Card{}
		expectedCards := []entity.Card(nil)

		mockCardRepo.EXPECT().FindAll().Return(mockedCards, nil)

		cards, err := service.GetList()

		assert.Nil(t, err)
		assert.Equal(t, expectedCards, cards)
	})
}

func TestCardService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCardRepo := NewMockCardRepository(ctrl)
	service := NewCardService(mockCardRepo)

	t.Run("success", func(t *testing.T) {
		newEntityCard := &entity.Card{Bank: "Bank A", HolderName: "HolderName A"}
		sendedModelCard := &model.Card{Bank: "Bank A", HolderName: "HolderName A"}

		mockCardRepo.EXPECT().Save(sendedModelCard).Return(nil)

		err := service.Create(newEntityCard)

		assert.Nil(t, err)
	})

	t.Run("repository error", func(t *testing.T) {
		newEntityCardError := &entity.Card{Bank: "Error Card", HolderName: "Error HolderName"}
		sendedModelCard := &model.Card{Bank: "Error Card", HolderName: "Error HolderName"}

		expectedError := errors.New("repository save error")
		mockCardRepo.EXPECT().Save(sendedModelCard).Return(expectedError)

		err := service.Create(newEntityCardError)

		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func TestCardService_FindByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCardRepo := NewMockCardRepository(ctrl)
	service := NewCardService(mockCardRepo)

	t.Run("success - card found", func(t *testing.T) {
		testID := int64(1)
		expectedCard := entity.Card{ID: testID, Bank: "Test Bank", HolderName: "Test HolderName"}
		mockedCard := &model.Card{ID: testID, Bank: "Test Bank", HolderName: "Test HolderName"}

		mockCardRepo.EXPECT().FindByID(testID).Return(mockedCard, nil)

		card, err := service.GetByID(testID)

		assert.Nil(t, err)
		assert.Equal(t, expectedCard, card)
	})

	t.Run("failure - card not found", func(t *testing.T) {
		testID := int64(2)
		mockCardRepo.EXPECT().FindByID(testID).Return(nil, nil)

		_, err := service.GetByID(testID)

		assert.NotNil(t, err)
		assert.Equal(t, errors.New("card not found"), err)
	})

	t.Run("failure - repository error", func(t *testing.T) {
		testID := int64(3)
		expectedError := errors.New("database error")
		mockCardRepo.EXPECT().FindByID(testID).Return(nil, expectedError)

		_, err := service.GetByID(testID)

		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)
	})
}
