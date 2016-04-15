package main

import (
	"fmt"
	"os"

	"github.com/vrischmann/envconfig"
)

var opts = envconfig.Options{
	AllOptional: true,
}

func main() {
	plugin := &Plugin{}

	if err := envconfig.InitWithOptions(&plugin, opts); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if true {
		fmt.Println(plugin.Config.Submodules)
		return
	}

	if err := plugin.Exec(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
