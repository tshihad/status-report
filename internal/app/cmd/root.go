package cmd

import (
	"fmt"

	"github.com/status-report/internal/app/services/mailapi"
)

func Execute() {
	sender := mailapi.NewMail("shihad@qburst.com")
	err := sender.SendHTMLEmail([]string{"tshihad9@gmail.com"}, nil, "test", "This is a test message")
	fmt.Println(err)
	// mailapi.Quickstart()
}
