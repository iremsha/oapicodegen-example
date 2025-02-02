package service

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/iremsha/oapicodegen-example/internal/entity"
	"github.com/iremsha/oapicodegen-example/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestBankService_FindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBankRepo := NewMockBankRepository(ctrl)
	service := NewBankService(mockBankRepo)

	t.Run("full", func(t *testing.T) {
		mockedBanks := []model.Bank{
			{ID: 1, Name: "Bank A"},
			{ID: 2, Name: "Bank B"},
		}
		expectedBanks := []entity.Bank{
			{ID: 1, Name: "Bank A"},
			{ID: 2, Name: "Bank B"},
		}

		mockBankRepo.EXPECT().FindAll().Return(mockedBanks, nil)

		banks, err := service.GetList()

		assert.Nil(t, err)
		assert.Equal(t, expectedBanks, banks)
	})

	t.Run("empty", func(t *testing.T) {
		mockedBanks := []model.Bank{}
		expectedBanks := []entity.Bank(nil)

		mockBankRepo.EXPECT().FindAll().Return(mockedBanks, nil)

		banks, err := service.GetList()

		assert.Nil(t, err)
		assert.Equal(t, expectedBanks, banks)
	})
}

func TestBankService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBankRepo := NewMockBankRepository(ctrl)
	service := NewBankService(mockBankRepo)

	t.Run("success", func(t *testing.T) {
		newEntityBank := &entity.Bank{Name: "New Bank"}
		sendedModelBank := &model.Bank{Name: "New Bank"}

		mockBankRepo.EXPECT().Save(sendedModelBank).Return(nil)

		err := service.Create(newEntityBank)

		assert.Nil(t, err)
	})

	t.Run("repository error", func(t *testing.T) {
		newEntityBankError := &entity.Bank{Name: "Error Bank"}
		sendedModelBank := &model.Bank{Name: "Error Bank"}

		expectedError := errors.New("repository save error")
		mockBankRepo.EXPECT().Save(sendedModelBank).Return(expectedError)

		err := service.Create(newEntityBankError)

		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)
	})
}

func TestBankService_FindByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBankRepo := NewMockBankRepository(ctrl)
	service := NewBankService(mockBankRepo)

	t.Run("success - bank found", func(t *testing.T) {
		testID := int64(1)
		expectedBank := entity.Bank{ID: testID, Name: "Test Bank"}
		mockedBank := &model.Bank{ID: testID, Name: "Test Bank"}

		mockBankRepo.EXPECT().FindByID(testID).Return(mockedBank, nil)

		bank, err := service.GetByID(testID)

		assert.Nil(t, err)
		assert.Equal(t, expectedBank, bank)
	})

	t.Run("failure - bank not found", func(t *testing.T) {
		testID := int64(2)
		mockBankRepo.EXPECT().FindByID(testID).Return(nil, nil)

		_, err := service.GetByID(testID)

		assert.NotNil(t, err)
		assert.Equal(t, errors.New("bank not found"), err)
	})

	t.Run("failure - repository error", func(t *testing.T) {
		testID := int64(3)
		expectedError := errors.New("database error")
		mockBankRepo.EXPECT().FindByID(testID).Return(nil, expectedError)

		_, err := service.GetByID(testID)

		assert.NotNil(t, err)
		assert.Equal(t, expectedError, err)
	})
}
