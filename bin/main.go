package main

import (
	"github.com/demoManito/bilibiliscript/building"
	"github.com/spf13/cobra"
)

var build *building.Building

func init() {
	build = building.New("")
	build.Conf.URL = ""
	build.Conf.MaxLimit = 2
	build.Conf.TickerDuration = 2000
}

func main() {
	rootCmd := &cobra.Command{
		Use:     "bs",
		Short:   "阿B同事吧盖楼脚本",
		Long:    "阿B同事吧盖楼脚本源码参见：https://github.com/demoManito/bilibiliscript (欢迎各位加星)",
		Version: "v1.0.0",
		Run: func(cmd *cobra.Command, args []string) {
			if build.Conf.Timing.EndTime != "" || build.Conf.Timing.StartTime != "" {
				build.Conf.Timing.Enable = true
			}
			build.Run()
		},
	}
	rootCmd.PersistentFlags().StringVarP(&build.Conf.ArticleBusinessID, "article_id", "i", "", "article business id")
	rootCmd.PersistentFlags().StringVarP(&build.Conf.XCSRF, "xcsrf", "x", "", "x-csrf in request header")
	rootCmd.PersistentFlags().StringVarP(&build.Conf.Cookie, "cookie", "c", "", "cookie in request header")
	rootCmd.PersistentFlags().StringVarP(&build.Conf.Timing.StartTime, "starttime", "s", "", "bilding start time")
	rootCmd.PersistentFlags().StringVarP(&build.Conf.Timing.EndTime, "endtime", "e", "", "bilding start time")
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
