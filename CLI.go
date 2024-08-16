package poker

import "io"

type CLI struct {
	playerStore PlayerStore
	in          io.Reader
}

func (c *CLI) PlayPoker() {
	c.playerStore.RecordWin("Chris")
}
