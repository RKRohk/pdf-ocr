package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/ocr", func(c *fiber.Ctx) error {
		multipartForm, err := c.MultipartForm()
		if err != nil {
			return err
		}

		files := multipartForm.File["file"]

		if len(files) == 0 {

			return fiber.NewError(fiber.ErrBadRequest.Code, "no files uploaded")
		} else {

			//TODO: do something with the files
			_ = files[0]

		}

		return nil
	})

	app.Listen(":3000")
}
