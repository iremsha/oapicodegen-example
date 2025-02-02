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

func TestCardHandler_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCardService := NewMockCardService(ctrl)
	handler := NewCardHandler(mockCardService)

	app := fiber.New()
	app.Get("/cards/:id", handler.GetByID)

	testID := int64(1)
	expectedCard := entity.Card{ID: testID, Bank: "Test Bank", HolderName: "Test HolderName"}
	mockCardService.EXPECT().GetByID(testID).Return(expectedCard, nil)

	req := httptest.NewRequest(http.MethodGet, "/cards/1", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var actualCard entity.Card
	err := json.NewDecoder(resp.Body).Decode(&actualCard)
	assert.Nil(t, err)
	assert.Equal(t, expectedCard.ID, actualCard.ID)
	assert.Equal(t, expectedCard.Bank, actualCard.Bank)
	assert.Equal(t, expectedCard.HolderName, actualCard.HolderName)
}

func TestCardHandler_GetByID_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCardService := NewMockCardService(ctrl)
	handler := NewCardHandler(mockCardService)

	app := fiber.New()
	app.Get("/cards/:id", handler.GetByID)

	testID := int64(2)
	mockCardService.EXPECT().GetByID(testID).Return(entity.Card{}, errors.ErrCardNotFound)

	req := httptest.NewRequest(http.MethodGet, "/cards/2", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestCardHandler_GetByID_InvalidID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCardService := NewMockCardService(ctrl)
	handler := NewCardHandler(mockCardService)

	app := fiber.New()
	app.Get("/cards/:id", handler.GetByID)

	req := httptest.NewRequest(http.MethodGet, "/cards/invalid", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestCardHandler_GetList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCardService := NewMockCardService(ctrl)
	handler := NewCardHandler(mockCardService)

	app := fiber.New()
	app.Get("/cards", handler.GetList)

	expectedCards := []entity.Card{
		{ID: 1, Bank: "Bank A", HolderName: "HolderName A"},
		{ID: 2, Bank: "Bank B", HolderName: "Test HolderName B"},
	}
	mockCardService.EXPECT().GetList().Return(expectedCards, nil)

	req := httptest.NewRequest(http.MethodGet, "/cards", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var actualCards []entity.Card
	err := json.NewDecoder(resp.Body).Decode(&actualCards)
	assert.Nil(t, err)
	assert.Equal(t, expectedCards, actualCards)
}

func TestCardHandler_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCardService := NewMockCardService(ctrl)
	handler := NewCardHandler(mockCardService)

	app := fiber.New()
	app.Post("/cards", handler.Create)

	newCard := entity.Card{Bank: "New Bank", HolderName: "New HolderName"}
	mockCardService.EXPECT().Create(&newCard).Return(nil)

	body, _ := json.Marshal(newCard)
	req := httptest.NewRequest(http.MethodPost, "/cards", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var actualCard entity.Card
	err := json.NewDecoder(resp.Body).Decode(&actualCard)
	assert.Nil(t, err)
	assert.Equal(t, newCard.Bank, actualCard.Bank)
	assert.Equal(t, newCard.HolderName, actualCard.HolderName)
}
