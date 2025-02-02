package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/iremsha/oapicodegen-example/internal/entity"
	"github.com/iremsha/oapicodegen-example/internal/errors"
	"github.com/stretchr/testify/assert"
)

func TestBankHandler_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBankService := NewMockBankService(ctrl)
	handler := NewBankHandler(mockBankService)

	app := fiber.New()
	app.Get("/banks/:id", handler.GetByID)

	testID := int64(1)
	expectedBank := entity.Bank{ID: testID, Name: "Test Bank"}
	mockBankService.EXPECT().GetByID(testID).Return(expectedBank, nil)

	req := httptest.NewRequest(http.MethodGet, "/banks/1", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var actualBank entity.Bank
	err := json.NewDecoder(resp.Body).Decode(&actualBank)
	assert.Nil(t, err)
	assert.Equal(t, expectedBank.ID, actualBank.ID)
	assert.Equal(t, expectedBank.Name, actualBank.Name)
}

func TestBankHandler_GetByID_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBankService := NewMockBankService(ctrl)
	handler := NewBankHandler(mockBankService)

	app := fiber.New()
	app.Get("/banks/:id", handler.GetByID)

	testID := int64(2)
	mockBankService.EXPECT().GetByID(testID).Return(entity.Bank{}, errors.ErrBankNotFound)

	req := httptest.NewRequest(http.MethodGet, "/banks/2", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestBankHandler_GetByID_InvalidID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBankService := NewMockBankService(ctrl)
	handler := NewBankHandler(mockBankService)

	app := fiber.New()
	app.Get("/banks/:id", handler.GetByID)

	req := httptest.NewRequest(http.MethodGet, "/banks/invalid", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestBankHandler_GetList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBankService := NewMockBankService(ctrl)
	handler := NewBankHandler(mockBankService)

	app := fiber.New()
	app.Get("/banks", handler.GetList)

	expectedBanks := []entity.Bank{
		{ID: 1, Name: "Bank A"},
		{ID: 2, Name: "Bank B"},
	}
	mockBankService.EXPECT().GetList().Return(expectedBanks, nil)

	req := httptest.NewRequest(http.MethodGet, "/banks", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var actualBanks []entity.Bank
	err := json.NewDecoder(resp.Body).Decode(&actualBanks)
	assert.Nil(t, err)
	assert.Equal(t, expectedBanks, actualBanks)
}

func TestBankHandler_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBankService := NewMockBankService(ctrl)
	handler := NewBankHandler(mockBankService)

	app := fiber.New()
	app.Post("/banks", handler.Create)

	newBank := entity.Bank{Name: "New Bank"}
	mockBankService.EXPECT().Create(&newBank).Return(nil)

	body, _ := json.Marshal(newBank)
	req := httptest.NewRequest(http.MethodPost, "/banks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var actualBank entity.Bank
	err := json.NewDecoder(resp.Body).Decode(&actualBank)
	assert.Nil(t, err)
	assert.Equal(t, newBank.Name, actualBank.Name)
}
