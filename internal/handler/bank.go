package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iremsha/oapicodegen-example/internal/entity"
	bankgen "github.com/iremsha/oapicodegen-example/internal/gen/bank"
)

type BankService interface {
	Create(Bank *entity.BankRequest) (*entity.BankResponse, error)
	Update(id int, Bank *entity.BankRequest) (*entity.BankResponse, error)
	GetByID(id int) (*entity.BankResponse, error)
	GetList() (*[]entity.BankResponse, error)
}

type BankHandler struct {
	service BankService
}

func NewBankHandler(s BankService) *BankHandler {
	return &BankHandler{service: s}
}

func (h *BankHandler) CreateApiV1Bank(c *fiber.Ctx) error {
	req := &bankgen.BankRequest{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	bankReq := entity.BankRequest{Name: req.Name, Address: req.Address}

	bankRes, err := h.service.Create(&bankReq)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot create bank"})
	}

	res := &bankgen.BankResponse{Id: bankRes.ID, Name: bankRes.Name, Address: bankRes.Address}

	return c.Status(fiber.StatusCreated).JSON(res)
}

func (h *BankHandler) UpdateApiV1Bank(c *fiber.Ctx, bankId int) error {
	req := &bankgen.BankRequest{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	bankReq := entity.BankRequest{Name: req.Name, Address: req.Address}

	bankRes, err := h.service.Update(bankId, &bankReq)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot update bank"})
	}

	res := &bankgen.BankResponse{Id: bankRes.ID, Name: bankRes.Name, Address: bankRes.Address}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *BankHandler) GetApiV1Bank(c *fiber.Ctx, bankId int) error {
	bankRes, err := h.service.GetByID(bankId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "bank not found"})
	}

	res := &bankgen.BankResponse{Id: bankRes.ID, Name: bankRes.Name, Address: bankRes.Address}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *BankHandler) GetApiV1BankCards(c *fiber.Ctx, bankId int) error {
	panic("unimplemented")
}

// GetApiV1Banks implements bankgen.ServerInterface.
func (h *BankHandler) GetApiV1Banks(c *fiber.Ctx) error {
	panic("unimplemented")
}