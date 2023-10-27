package drift

import (
	"io"

	"github.com/jedib0t/go-pretty/v6/table"

	"github.com/OvictorVieira/promeheux.api/pkg/dynamic"
	"github.com/OvictorVieira/promeheux.api/pkg/style"
)

func (s *Strategy) Print(f io.Writer, pretty bool, withColor ...bool) {
	var tableStyle *table.Style
	if pretty {
		tableStyle = style.NewDefaultTableStyle()
	}
	dynamic.PrintConfig(s, f, tableStyle, len(withColor) > 0 && withColor[0], dynamic.DefaultWhiteList()...)
}
