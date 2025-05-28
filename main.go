package main

import (
	"github.com/gofiber/fiber/v2"
)

// Struct Mahasiswa
type Mahasiswa struct {
	Nama    string `json:"nama"`
	NPM     string `json:"npm"`
	Umur    int    `json:"umur"`
	Jurusan string `json:"jurusan"`
}

// Fungsi handler untuk mengembalikan data Mahasiswa dalam JSON
func GetMahasiswa(c *fiber.Ctx) error {
	data := Mahasiswa{
		Nama:    "Muhammad Nizar Akmal",
		NPM:     "714234567",
		Umur:    18,
		Jurusan: "Teknik Informatika",
	}
	return c.JSON(data)
}

func main() {
	app := fiber.New()

	// Endpoint GET /mahasiswa
	app.Get("/mahasiswa", GetMahasiswa)

	app.Listen(":3000")
}
