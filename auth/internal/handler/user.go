package handler

import (
	"auth/lib/model"
	"auth/util"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func SignupHandler(c *fiber.Ctx) error {
    var user model.User
    if err := c.BodyParser(&user); err != nil {
        return c.Status(http.StatusBadRequest).SendString(err.Error())
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return c.Status(http.StatusInternalServerError).SendString(err.Error())
    }
    user.Password = string(hashedPassword)

    err = model.CreateUser(&user)
    if err != nil {
        return c.Status(http.StatusInternalServerError).SendString(err.Error())
    }

    return c.SendStatus(http.StatusCreated)
}

func LoginHandler(c *fiber.Ctx) error {
    var user model.User
    if err := c.BodyParser(&user); err != nil {
        return c.Status(http.StatusBadRequest).SendString(err.Error())
    }

    storedUser, err := model.GetUserByEmail(user.Email)
    if err != nil {
        return c.Status(http.StatusUnauthorized).SendString("Invalid email")
    }

    err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
    if err != nil {
        return c.Status(http.StatusUnauthorized).SendString("Invalid password")
    }
    token, err := util.GenerateJWT(storedUser.ID)
    if err != nil {
        return c.Status(http.StatusInternalServerError).SendString(err.Error())
    }

    return c.JSON(fiber.Map{
        "token": token,
    })
}

func CheckMeHandler(c *fiber.Ctx) error {
    user, ok := c.Locals("user").(*model.User)
    if !ok {
        return c.Status(http.StatusInternalServerError).SendString("User not found")
    }

    userResponse := map[string]interface{}{
        "user_id":    user.ID,
        "email":      user.Email,
        "name":       user.Name,
        "mobile":     user.Mobile,
        "created_at": user.CreatedAt,
    }

    return c.JSON(userResponse)
}


func ChangePasswordHandler(c *fiber.Ctx) error {
    var req struct {
        OldPassword string `json:"old_password"`
        NewPassword string `json:"new_password"`
    }
    if err := c.BodyParser(&req); err != nil {
        return c.Status(http.StatusBadRequest).SendString(err.Error())
    }

    user := c.Locals("user").(*model.User)
    if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword)) != nil {
        return c.Status(http.StatusUnauthorized).SendString("Invalid old password")
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
    if err != nil {
        return c.Status(http.StatusInternalServerError).SendString(err.Error())
    }

    err = model.UpdatePassword(user.ID, string(hashedPassword))
    if err != nil {
        return c.Status(http.StatusInternalServerError).SendString(err.Error())
    }

    return c.SendStatus(http.StatusOK)
}

func LogoutHandler(c *fiber.Ctx) error {
    return c.SendStatus(http.StatusOK)
}
