package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iremsha/oapicodegen-example/internal/entity"
	"github.com/iremsha/oapicodegen-example/internal/errors"
)

type CardService interface {
	Create(card *entity.Card) error
	GetList() ([]entity.Card, error)
	GetByID(id int64) (entity.Card, error)
}

type CardHandler struct {
	service CardService
}

func NewCardHandler(s CardService) *CardHandler {
	return &CardHandler{service: s}
}

func (h *CardHandler) Create(c *fiber.Ctx) error {
	var card entity.Card
	if err := c.BodyParser(&card); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	if err := h.service.Create(&card); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot create card"})
	}
	return c.Status(fiber.StatusCreated).JSON(card)
}

func (h *CardHandler) GetList(c *fiber.Ctx) error {
	cards, err := h.service.GetList()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot get cards"})
	}
	return c.JSON(cards)
}

func (h *CardHandler) GetByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseInt(idParam, 0, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	card, err := h.service.GetByID(id)
	if err != nil {
		if err == errors.ErrCardNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "card not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot retrieve card"})
	}

	return c.JSON(card)
}
