package cmd

import (
	"fmt"

	"github.com/status-report/internal/app/services/mailapi"
	"github.com/status-report/internal/app/utils/log"
)

// Execute task starts from here
func Execute() {
	sender := mailapi.NewMail("shihad@qburst.com")
	err := sender.SendHTMLEmail([]string{"tshihad9@gmail.com"}, nil, "test", "This is a test message")
	fmt.Println(err)
	// mailapi.Quickstart()
}

// Serving at given port
func serve() {
	if err := mustPrepareApp(); err != nil {
		log.Panic(err.Error())
	}
}
