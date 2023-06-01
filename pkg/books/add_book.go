package books

import (
	"github.com/Ryzhon/go-fiber-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type AddBookRequestsBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (h handler) AddBook(c *fiber.Ctx) error {
	body := AddBookRequestsBody{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())

	}
	var book models.Book

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	if result := h.DB.Create(&book); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())

	}
	return c.Status(fiber.StatusCreated).JSON(&book)
}
