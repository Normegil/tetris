package main

func main() {
	g := newGame()
	err := g.run()
	if nil != err {
		panic(err)
	}
}
