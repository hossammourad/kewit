package helpers

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

var ConfigureTable = func(t table.Writer) {
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleRounded)
	t.Style().Options.SeparateRows = true
}
