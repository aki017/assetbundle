package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"text/tabwriter"
)

// Table is TerminalTable
type Table struct {
	Title string
	rows  [][]string
}

// AddRow is add row
func (t *Table) AddRow(cs ...string) {
	t.rows = append(t.rows, cs)
}

func (t Table) String() string {
	var buffer bytes.Buffer
	bw := bufio.NewWriter(&buffer)
	writer := tabwriter.NewWriter(bw, 4, 8, 0, '\t', 0)

	buffer.WriteString(fmt.Sprintf("# %s\n", t.Title))
	for _, r := range t.rows {
		writer.Write([]byte(strings.Join(r, "\t")))
		writer.Write([]byte("\n"))
	}
	writer.Flush()
	bw.Flush()
	return buffer.String()
}
