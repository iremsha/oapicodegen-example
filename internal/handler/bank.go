package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/iremsha/oapicodegen-example/internal/entity"
	"github.com/iremsha/oapicodegen-example/internal/errors"
)

type BankService interface {
	Create(Bank *entity.Bank) error
	GetList() ([]entity.Bank, error)
	GetByID(id int64) (entity.Bank, error)
}

type BankHandler struct {
	service BankService
}

func NewBankHandler(s BankService) *BankHandler {
	return &BankHandler{service: s}
}

func (h *BankHandler) Create(c *fiber.Ctx) error {
	var bank entity.Bank
	if err := c.BodyParser(&bank); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	if err := h.service.Create(&bank); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot create bank"})
	}
	return c.Status(fiber.StatusCreated).JSON(bank)
}

func (h *BankHandler) GetList(c *fiber.Ctx) error {
	banks, err := h.service.GetList()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot get banks"})
	}
	return c.JSON(banks)
}

func (h *BankHandler) GetByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseInt(idParam, 0, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	bank, err := h.service.GetByID(id)
	if err != nil {
		if err == errors.ErrBankNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "bank not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot retrieve bank"})
	}

	return c.JSON(bank)
}
