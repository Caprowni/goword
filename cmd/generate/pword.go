package generate

import (
	"crypto/rand"
	"errors"
	"github.com/spf13/cobra"
	"log"
	"math/big"
)

const (
	lowers        = "abcdefghijklmnopqrstuvwxyz"
	uppers        = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers       = "0123456789"
	special_chars = "~!@#$^&*()_+`-={}|[]\\:\"<>?,./"
	max           = 32
)

func NewCmd() *cobra.Command {

	pword := password{}

	cmd := &cobra.Command{
		Use:   "gen",
		Short: "Generate Password",
		Long:  "Generate Password",
		Run: func(cmd *cobra.Command, args []string) {
			pword.run()
		},
	}
	cmd.Flags().IntVar(
		&pword.passwords,
		"passwords",
		1,
		"Define how many passwords to produce",
	)
	cmd.Flags().IntVar(
		&pword.length,
		"length",
		8,
		"The length of the password",
	)

	cmd.Flags().BoolVar(
		&pword.specials,
		"specials",
		false,
		"Define if special chars should be used or not")

	cmd.Flags().BoolVar(
		&pword.upper,
		"upper",
		false,
		"Define if upper case chars should be used")

	cmd.Flags().IntVar(
		&pword.numbers,
		"numbers",
		0,
		"Define how many numbers should be in the password",
	)

	return cmd
}

type password struct {
	length    int
	numbers   int
	passwords int
	specials  bool
	upper     bool
}

func (p *password) run() (string, error) {
	pass := ""
	_, err := p.checkVariables()
	if err != nil {
		log.Printf("%v", err)
		return "", err
	}
	for i := 0; i < p.passwords; i++ {
		pass, err := p.generate()
		if err != nil {
			log.Printf("%v", err)
			return "", err
		}
		log.Printf("Password %v: %v", i+1, pass)
	}
	return pass, nil
}

func (p *password) generate() (string, error) {

	nums := p.numbers
	chars := p.constructCharSet()

	var result string

	for i := 0; i < p.length; i++ {
		c, _ := randomChar(chars)
		result += c
	}

	for i := 0; i < nums; i++ {
		d, _ := randomChar(numbers)

		result, _ = randomInsert(result, d)
	}
	return result, nil
}

func randomChar(chars string) (string, error) {
	char, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
	if err != nil {
		return "", err
	}
	return string(chars[char.Int64()]), nil
}

func randomInsert(result, val string) (string, error) {
	char, err := rand.Int(rand.Reader, big.NewInt(int64(len(result)+1)))
	if err != nil {
		return "", err
	}
	return result[0:char.Int64()] + val + result[char.Int64():], nil
}

func (p *password) checkVariables() (string, error) {

	if p.length > max {
		return "", errors.New("Password length cannot exceed max length")
	}

	if p.numbers > max {
		return "", errors.New("Numbers desired cannot exceed max length")
	}

	return "", nil
}

func (p *password) constructCharSet() string {
	chars := lowers

	if p.upper {
		chars += uppers
	}

	if p.specials {
		chars += special_chars
	}

	return chars
}
