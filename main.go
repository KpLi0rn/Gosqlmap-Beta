package main

import (
	"github.com/KpLi0rn/Gosqlmap/lib/controller"
	"github.com/KpLi0rn/Gosqlmap/lib/core"
	"github.com/KpLi0rn/Gosqlmap/lib/parse"
)
func main()  {
	parse.ParserInput()
	core.Init()

	controller.Start()
	// checkConnection()
}
