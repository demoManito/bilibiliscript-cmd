package cmds

import (
	"github.com/demoManito/bilibiliscript/floor"
	"github.com/spf13/cobra"

	"github.com/demoManito/bilibiliscript-cmd/core"
)

var floorcmd core.ICommand = func() *FloorCmd {
	floorcmd := &FloorCmd{
		floor: floor.New(""),
	}
	floorcmd.floor.Conf.URL = ""
	return floorcmd
}()

func init() {
	core.Register(floorcmd)
}

type FloorCmd struct {
	*core.ParentCommand

	floor *floor.Floor
}

func (fc *FloorCmd) Use() string {
	return "floor"
}

func (fc *FloorCmd) Short() string {
	return "查询指定目标楼层信息"
}

func (fc *FloorCmd) Long() string {
	return "该功能源码参见：https://github.com/demoManito/bilibiliscript/floor (欢迎各位加星)"
}

func (fc *FloorCmd) Run() core.CmdFunc {
	return func(cmd *cobra.Command, args []string) {
		fc.floor.SetPageNum(fc.floor.Conf.FloorNum)
		fc.floor.Report()
	}
}

func (fc *FloorCmd) AppendFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&fc.floor.Conf.ArticleBusinessID, "article_id", "i", "", "article business id")
	cmd.PersistentFlags().StringVarP(&fc.floor.Conf.XCSRF, "xcsrf", "x", "", "x-csrf in request header")
	cmd.PersistentFlags().StringVarP(&fc.floor.Conf.Cookie, "cookie", "c", "", "cookie in request header")
	cmd.PersistentFlags().Int64VarP(&fc.floor.Conf.FloorNum, "floor_num", "n", 0, "target floor num")
}
