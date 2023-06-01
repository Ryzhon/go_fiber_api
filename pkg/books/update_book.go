package books

import (
	"github.com/Ryzhon/go-fiber-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type UpdateBookRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (h handler) UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	body := UpdateBookRequestBody{}
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}
	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	h.DB.Save(&book)

	return c.Status(fiber.StatusOK).JSON(&book)
}
