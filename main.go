package main

import (
	"flag"

	"math/rand"
	"time"

	"github.com/Sirupsen/logrus"
)

const (
	FONT_TUSJ   = "tusj.ttf"
	FONT_ANUDRG = "anudrg.ttf"
)

func init() {
	verbose := flag.Bool("v", false, "Verbose mode will display all debug informations")
	flag.Parse()

	if *verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}

	rand.Seed(time.Now().UnixNano())
}

func main() {
	g := newGame()
	err := g.run()
	if nil != err {
		panic(err)
	}
}
