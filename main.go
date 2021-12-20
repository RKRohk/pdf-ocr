package main

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
			file := files[0]

			//check if filename is not pdf 
			if !strings.HasSuffix(file.Filename, ".pdf") {
				return fiber.NewError(fiber.ErrBadRequest.Code, "file is not pdf")
			}

			c.SaveFile(file, os.TempDir()+"/"+file.Filename)

			id := uuid.New()
			performOCR(os.TempDir()+"/"+file.Filename, file.Filename, id)

			c.Redirect("/ocr/" + id.String() + ".pdf")

		}

		return nil
	})

	app.Static("/ocr", "./output")

	app.Listen(":3000")
}
