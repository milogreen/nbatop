package nbatop
import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func schedule() *ui.List{
	scheduleData = []string{"one", "two", "three"}
	scheduleList = widgets.NewList()
	scheduleList.Rows = scheduleData
	scheduleList.setRect(0, 0, 25, 5)
	return scheduleList
}
