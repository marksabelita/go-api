package user_response

import "github.com/gofiber/fiber/v2"

type ErrorResponse struct {
    Message string     `json:"message"`
}

type UserResponse struct {
    Status  int        `json:"status"`
    Message string     `json:"message"`
    Data    *fiber.Map `json:"data"`
}

type GetResponse struct {
    
}