package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	playerStore PlayerStore
	in          *bufio.Scanner
}

func NewCLI(playerStore PlayerStore, in io.Reader) *CLI {
	return &CLI{playerStore, bufio.NewScanner(in)}
}

func (c *CLI) PlayPoker() {
	userInput := c.readLine()
	c.playerStore.RecordWin(extractWinner(userInput))
}

func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
