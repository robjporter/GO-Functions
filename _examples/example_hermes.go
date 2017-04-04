package main

import (
	"io/ioutil"

	"github.com/robjporter/go-functions/hermes"
)

func main() {
	// Configure hermes by setting a theme and your product info
	h := hermes.Hermes{
		// Optional Theme
		// Theme: new(Default)
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name: "Hermes",
			Link: "https://example-hermes.com/",
			// Optional product logo
			Logo: "https:upload.wikimedia.org/wikipedia/commons/thumb/6/64/Cisco_logo.svg/320px-Cisco_logo.svg.png",
			// Custom copyright notice
			Copyright: "Copyright © 2017 Dharma Initiative. All rights reserved.",
		},
	}

	email := hermes.Email{
		Body: hermes.Body{
			Name: "Jon Snow",
			Intros: []string{
				"Welcome to Hermes! We're very excited to have you on board.",
			},
			Greeting:  "Dear",
			Signature: "Sincerely",

			Actions: []hermes.Action{
				{
					Instructions: "To get started with Hermes, please click here:",
					Button: hermes.Button{
						Color: "#22BC66", // Optional action button color
						Text:  "Confirm your account",
						Link:  "https://hermes-example.com/confirm?token=d9729feb74992cc3482b350163a1a010",
					},
				},
			},
			Table: hermes.Table{
				Data: [][]hermes.Entry{
					// List of rows
					{
						// Key is the column name, Value is the cell value
						// First object defines what columns will be displayed
						{Key: "Item", Value: "Golang"},
						{Key: "Description", Value: "Open source programming language that makes it easy to build simple, reliable, and efficient software"},
						{Key: "Price", Value: "$10.99"},
					},
					{
						{Key: "Item", Value: "Hermes"},
						{Key: "Description", Value: "Programmatically create beautiful e-mails using Golang."},
						{Key: "Price", Value: "$1.99"},
					},
				},
				Columns: hermes.Columns{
					// Custom style for each rows
					CustomWidth: map[string]string{
						"Item":  "20%",
						"Price": "15%",
					},
					CustomAlignement: map[string]string{
						"Price": "right",
					},
				},
			},
			Outros: []string{
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
		},
	}

	// Generate an HTML email with the provided contents (for modern clients)
	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}

	// Generate the plaintext version of the e-mail (for clients that do not support xHTML)
	//emailText, err := h.GeneratePlainText(email)
	//if err != nil {
	//	panic(err) // Tip: Handle error with something else than a panic ;)
	//}

	// Optionally, preview the generated HTML e-mail by writing it to a local file
	err = ioutil.WriteFile("preview.html", []byte(emailBody), 0644)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}
}
