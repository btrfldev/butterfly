package main

import "github.com/gofiber/fiber/v2"

func (s *Server) UploadUI(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}