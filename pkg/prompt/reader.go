package prompt

import (
	"bufio"
	"os"
)

func (pr *prompt) Read(p []byte) (n int, err error) {
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')
	p = []byte(value)

	return len(value), nil
}

func (pr *prompt) ReadString(p []byte) (n int, err error) {
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')
	p = []byte(value)

	return len(value), nil
}
