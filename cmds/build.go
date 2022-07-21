package cmds

import (
	"github.com/demoManito/bilibiliscript-cmd/core"

	"github.com/demoManito/bilibiliscript/building"
	"github.com/spf13/cobra"
)

var buildcmd core.ICommand = func() *BuildCmd {
	buildcmd := &BuildCmd{
		build: building.New(""),
	}
	buildcmd.build.Conf.URL = ""
	buildcmd.build.Conf.MaxLimit = 2
	buildcmd.build.Conf.TickerDuration = 2000
	return buildcmd
}()

func init() {
	core.Register(buildcmd)
}

type BuildCmd struct {
	*core.ParentCommand

	build *building.Building
}

func (bc *BuildCmd) Use() string {
	return "run"
}

func (bc *BuildCmd) Long() string {
	return "该功能源码参见：https://github.com/demoManito/bilibiliscript/building (欢迎各位加星)"
}

func (bc *BuildCmd) Run() core.CmdFunc {
	return func(cmd *cobra.Command, args []string) {
		if bc.build.Conf.Timing.EndTime != "" || bc.build.Conf.Timing.StartTime != "" {
			bc.build.Conf.Timing.Enable = true
		}
		bc.build.Run()
	}
}

func (bc *BuildCmd) AppendFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&bc.build.Conf.ArticleBusinessID, "article_id", "i", "", "article business id")
	cmd.PersistentFlags().StringVarP(&bc.build.Conf.XCSRF, "xcsrf", "x", "", "x-csrf in request header")
	cmd.PersistentFlags().StringVarP(&bc.build.Conf.Cookie, "cookie", "c", "", "cookie in request header")
	cmd.PersistentFlags().StringVarP(&bc.build.Conf.Timing.StartTime, "starttime", "s", "", "bilding start time")
	cmd.PersistentFlags().StringVarP(&bc.build.Conf.Timing.EndTime, "endtime", "e", "", "bilding start time")
}
