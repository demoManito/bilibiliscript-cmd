package main

import (
	"github.com/spf13/cobra"

	_ "github.com/demoManito/bilibiliscript-cmd/cmds"
	"github.com/demoManito/bilibiliscript-cmd/core"
)

var rootCmd = &cobra.Command{
	Use:     "bs",
	Short:   "阿B同事吧盖楼脚本",
	Long:    "阿B同事吧盖楼脚本源码参见：https://github.com/demoManito/bilibiliscript (欢迎各位加星)",
	Version: "v1.0.0",
}

func main() {
	if err := core.Execute(rootCmd); err != nil {
		panic(err)
	}
}
