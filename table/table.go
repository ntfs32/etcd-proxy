package table

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

func Example()  {
	data := [][]string{
		[]string{"A", "The Good", "500"},
		[]string{"B", "The Very very Bad Man", "288"},
		[]string{"C", "The Ugly", "120"},
		[]string{"D", "", "800"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"租户名称", "环境", "总数"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}

/**
输出合并纵向单元格的表
 */
func Example2()  {
	data := [][]string{
		[]string{"1/1/2014", "Domain name", "1234", "$10.98"},
		[]string{"1/1/2014", "January Hosting", "2345", "$54.95"},
		[]string{"1/4/2014", "February Hosting", "3456", "$51.00"},
		[]string{"1/4/2014", "February Extra Bandwidth", "4567", "$30.00"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Date", "Description", "CV2", "Amount"})
	table.SetFooter([]string{"", "", "Total", "$146.93"})
	table.SetAutoMergeCells(true)
	table.SetAutoWrapText(true)
	table.SetRowLine(true)
	table.AppendBulk(data)
	table.Render()
}