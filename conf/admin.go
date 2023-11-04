package conf

import (
	"os"
	"strings"
)

func EnvAdmin() (res map[string]bool) {
	emailsStr := os.Getenv("ADMIN_EMAILS")
	emails := strings.Split(emailsStr, ",")

	res = map[string]bool{}

	for _, email := range emails {
		res[email] = true
	}

	return
}
