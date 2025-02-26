package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iremsha/oapicodegen-example/internal/entity"
	cardgen "github.com/iremsha/oapicodegen-example/internal/gen/card"
)

type CardService interface {
	Create(card *entity.CardRequest) (*entity.CardResponse, error)
	Update(id int, card *entity.CardRequest) (*entity.CardResponse, error)
	GetList() (*[]entity.CardResponse, error)
	GetByID(id int) (*entity.CardResponse, error)
}

type CardHandler struct {
	service CardService
}

func NewCardHandler(s CardService) *CardHandler {
	return &CardHandler{service: s}
}

func (h *CardHandler) CreateApiV1Card(c *fiber.Ctx) error {
	req := &cardgen.CardRequest{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	cardReq := entity.CardRequest{Name: req.Name, Type: req.Type, Cvv: req.Cvv}

	cardRes, err := h.service.Create(&cardReq)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot create card"})
	}

	res := &cardgen.CardResponse{Id: cardRes.ID, Name: cardRes.Name, Type: cardRes.Type, Cvv: &cardRes.Cvv}

	return c.Status(fiber.StatusCreated).JSON(res)
}

func (h *CardHandler) UpdateApiV1Card(c *fiber.Ctx, cardId int) error {
	req := &cardgen.CardRequest{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	cardReq := entity.CardRequest{Name: req.Name, Type: req.Type, Cvv: req.Cvv}

	cardRes, err := h.service.Update(cardId, &cardReq)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot update card"})
	}

	res := &cardgen.CardResponse{Id: cardRes.ID, Name: cardRes.Name, Type: cardRes.Type, Cvv: &cardRes.Cvv}

	return c.Status(fiber.StatusCreated).JSON(res)
}

func (h *CardHandler) GetApiV1Card(c *fiber.Ctx, cardId int) error {
	panic("unimplemented")
}

func (h *CardHandler) GetApiV1Cards(c *fiber.Ctx) error {
	panic("unimplemented")
}
