package controller

import (
	"github.com/gofiber/fiber/v2"
	"app/model"
	"app/utils"
	"fmt"
)

// Register
//
// @Tags Auth
// @Param payload body model.AuthRequest true "AuthRequest payload"
// @Accept json
// @Produce json
// @Success 200
// @Router /app/auth/register [post]
func Register(c *fiber.Ctx) error {
	var req model.AuthRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
	    })
	}
	user := model.User{
		Email: req.Email,
	    PasswordHash: utils.GeneratePassword(req.Password),
	}
	fmt.Println(user)

	//res := database.DB.Create(&user)
	// if res.Error != nil {
	//  return c.Status(400).JSON(fiber.Map{
	//   "message": res.Error.Error(),
	//  })
	// }
	return c.Status(201).JSON(fiber.Map{
		"message": "user created",
	})
}

// Login
//
// @Tags Auth
// @Param payload body model.AuthRequest true "AuthRequest payload"
// @Accept json
// @Produce json
// @Success 200
// @Router /app/auth/login [post]
func Login(c *fiber.Ctx) error {
	var req model.AuthRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
	    })
	}
	// var user model.User
	// res := database.DB.Where("email = ?", req.Email).First(&user)
	// if res.Error != nil {
	// 	return c.Status(400).JSON(fiber.Map{
	// 		"message": "user not found",
	//     })
	// }
	pw := "$2a$10$iQJTyS9J/mBU4p6EeXhjpuWvSBm9mzV0MaR3nG/fdINbGGodhjRqK"
	if !utils.ComparePassword(pw, req.Password) {
		return c.Status(400).JSON(fiber.Map{
			"message": "incorrect password",
	    })
	}
   
	token, err := utils.GenerateToken(1, req.Email)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
	    })
	}
	c.Set("Authorization", token)
	return c.JSON(fiber.Map{
		"token": token,
	})
}