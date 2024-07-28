package handler

import (
	"net/http"
	"tracker/lib/model"

	"github.com/gofiber/fiber/v2"
)

func AddExpense(c *fiber.Ctx) error {
	var expenseReq model.ExpenseRequest
	if err := c.BodyParser(&expenseReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	user, ok := c.Locals("user").(*model.User)
	if !ok {
		return c.Status(http.StatusInternalServerError).SendString("User not found")
	}

	if err := model.ValidateAndCalculateAmounts(&expenseReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	expenseID, err := model.CreateExpenseWithParticipants(expenseReq.ExpenseName, expenseReq.TotalAmount, user.ID, expenseReq.Participants)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create expense with participants"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "expense created successfully", "expense_id": expenseID})
}

func TrackMe(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*model.User)
	if !ok {
		return c.Status(http.StatusInternalServerError).SendString("User not found")
	}
	expenses, err := model.GetMyExpenses(user.ID)
	if err != nil {
		return c.Status(http.StatusUnauthorized).SendString("Invalid User")
	}

	return c.JSON(expenses)
}

func TrackAll(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*model.User)
	if !ok {
		return c.Status(http.StatusInternalServerError).SendString("User not found")
	}

	expenses, err := model.GetAllUsersExpenses(user.ID)
	if err != nil {
		return c.Status(http.StatusUnauthorized).SendString("Invalid User")
	}

	return c.JSON(expenses)
}

func TrackAllAdmin(c *fiber.Ctx) error {

	expenses, err := model.GetAdminAllExpenses()
	if err != nil {
		return c.Status(http.StatusUnauthorized).SendString("Invalid User")
	}

	return c.JSON(expenses)
}
