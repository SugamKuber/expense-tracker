package handler

import (
	"auth/lib/model"
	"auth/util"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func setupApp() *fiber.App {
	app := fiber.New()
	// Set up routes
	app.Post("/signup", SignupHandler)
	app.Post("/login", LoginHandler)
	app.Get("/checkme", CheckMeHandler)
	app.Post("/changepassword", ChangePasswordHandler)
	app.Post("/logout", LogoutHandler)
	return app
}

func TestSignupHandler(t *testing.T) {
	app := setupApp()

	user := model.User{
		Email:    "test@example.com",
		Password: "password123",
		Name:     "Test User",
	}

	body, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error making request: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}
}

func TestLoginHandler(t *testing.T) {
	app := setupApp()

	// Mock the model functions
	model.GetUserByEmail = func(email string) (*model.User, error) {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		return &model.User{
			Email:    email,
			Password: string(hashedPassword),
			ID:       1,
		}, nil
	}
	util.GenerateJWT = func(userID float64) (string, error) {
		return "fake-jwt-token", nil
	}

	user := model.User{
		Email:    "test@example.com",
		Password: "password123",
	}

	body, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error making request: %v", err)
	}

	var result map[string]string
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		t.Fatalf("Error decoding response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	if result["token"] != "fake-jwt-token" {
		t.Errorf("Expected token 'fake-jwt-token', got %s", result["token"])
	}
}

func TestCheckMeHandler(t *testing.T) {
	app := setupApp()

	// Create a mock user and set it in the context
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("user", &model.User{
			ID:        1,
			Email:     "test@example.com",
			Name:      "Test User",
			Mobile:    "1234567890",
			CreatedAt: "2024-07-29T12:34:56Z",
		})
		return c.Next()
	})

	req := httptest.NewRequest(http.MethodGet, "/checkme", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error making request: %v", err)
	}

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		t.Fatalf("Error decoding response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	if result["email"] != "test@example.com" {
		t.Errorf("Expected email 'test@example.com', got %v", result["email"])
	}
}

func TestChangePasswordHandler(t *testing.T) {
	app := setupApp()

	// Mock the model functions
	model.UpdatePassword = func(userID float64, newPassword string) error {
		return nil
	}

	user := &model.User{
		Password: "$2a$10$KIXzF6nUOL1vvj00D/rteO78P5OOWv/RnY9jMFDzFGKX6v0mLRiCy", // hashed "oldpassword"
	}

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("user", user)
		return c.Next()
	})

	reqBody := struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}{
		OldPassword: "oldpassword",
		NewPassword: "newpassword123",
	}
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/changepassword", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error making request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestLogoutHandler(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest(http.MethodPost, "/logout", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error making request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}
