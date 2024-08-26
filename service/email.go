package service

import (
	"fmt"
	"net/smtp"
	"97HW/config"
	"97HW/models"
)

func SendNotifications() error {
	users, err := getUsersFromDB()
	if err != nil {
		return err
	}

	for _, user := range users {
		if err := sendEmail(user.Email, "Updated Weather Data", "The weather data has been updated."); err != nil {
			return err
		}
	}

	return nil
}

func sendEmail(to, subject, body string) error {
	auth := smtp.PlainAuth("", config.SMTPUser, config.SMTPPassword, config.SMTPServer)
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	return smtp.SendMail(fmt.Sprintf("%s:%s", config.SMTPServer, config.SMTPPort), auth, config.SMTPUser, []string{to}, msg)
}

func getUsersFromDB() ([]models.User, error) {
	db := getDB()
	rows, err := db.Query("SELECT id, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
