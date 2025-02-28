package main

import (
	"database/sql"
	"log"
	"net/smtp"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func connDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/Login_DB?parseTime=true")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createTable() {
	db, err := connDB()
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	defer db.Close()

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS user_Login (
		id INT AUTO_INCREMENT PRIMARY KEY,
		email VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Error creating table:", err)
	} else {
		log.Println("Table 'user_Login' is ready!")
	}

}

func sendEmail(to string) {
	from := "vishal.9dot@gmail.com"
	password := "cwdd nisn kwcm bdat" 
	to = strings.TrimSpace(to)
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	subject := "Subject: It just a Noraml Email!\n"
	body := "\n It is just a normal email. \n"

	msg := []byte(subject + body)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
	if err != nil {
		log.Println("Failed to send email:", err)
	} else {
		log.Println("Email sent successfully to", to)
	}
}



func main() {
	app := fiber.New()

	app.Use(cors.New())

	createTable()
	app.Post("/api/login", func(c *fiber.Ctx) error {
		user := new(User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
		}

		db, err := connDB()
		if err != nil {
			log.Println("Database connection error:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database connection error"})
		}
		defer db.Close()

		var existingUser string
		err = db.QueryRow("SELECT email FROM user_Login WHERE email = ?", user.Email).Scan(&existingUser)
		if err != nil {
			if err == sql.ErrNoRows {
				_, err = db.Exec("INSERT INTO user_Login (email, password) VALUES (?, ?)", user.Email, user.Password)
				if err != nil {
					log.Printf("Error inserting data into database: %v", err)
					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error inserting data into database"})
				}
				sendEmail(user.Email)
				return c.JSON(fiber.Map{"message": "User login successful! Welcome email sent."})
			}
		}
		return c.JSON(fiber.Map{"message": "User login successful!"})
	})
	
	log.Println("Server running on port 8000")
	log.Fatal(app.Listen(":8000"))
}
