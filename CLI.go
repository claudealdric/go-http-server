package poker

type CLI struct {
	playerStore PlayerStore
}

func (c *CLI) PlayPoker() {
	c.playerStore.RecordWin("")
}
