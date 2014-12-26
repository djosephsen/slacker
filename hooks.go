package main

func initHooks(b *bot) error{
	b.register(chores.Ping)
}
