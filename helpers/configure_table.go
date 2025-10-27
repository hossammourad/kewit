package helpers

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func ConfigureTable(t table.Writer) {
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleRounded)
	t.Style().Options.SeparateRows = true
	t.Style().Options.DoNotColorBordersAndSeparators = true
	t.Style().Color.RowAlternate = text.Colors{text.FgYellow}
	t.Style().Color.Header = text.Colors{text.Bold}
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Colors: text.Colors{text.Bold}},
	})
}
