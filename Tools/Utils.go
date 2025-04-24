package Tools

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func WelcomePrint() []byte {
	file, err := os.ReadFile("Chat/welcome.txt")

	if err != nil {
		fmt.Println(err.Error())
	}

	return file
}

func NameVerification(Name string) ([]byte, error) {
	Name = strings.Trim(Name, " ")
	Name = strings.Join(strings.Fields(Name), " ")
	for _, char := range Name {
		if !(unicode.IsDigit(char) || unicode.IsLetter(char)) || unicode.IsControl(char) {
			return []byte(Name), fmt.Errorf("only alphanumeric Characters Allowed")
		}
	}
	if Name == "" {
		return []byte(Name), fmt.Errorf("invalid Name")
	}
	for _, name := range Clients {
		if name == Name {
			return []byte(Name), fmt.Errorf("this Name Already exist")
		}
	}
	return []byte(Name), nil
}
