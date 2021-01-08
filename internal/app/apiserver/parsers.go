package apiserver

import (
	"fmt"

	"github.com/tealeg/xlsx/v3"
)

//Xlsx parser
func cellVisitor(c *xlsx.Cell) error {
	value, err := c.FormattedValue()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Cell value:", value)
	}
	return err
}

func rowVisitor(r *xlsx.Row) error {
	return r.ForEachCell(cellVisitor)
}

func rowStuff(data []byte) {

	wb, err := xlsx.OpenBinary(data)

	//wb, err := xlsx.OpenFile(file)
	if err != nil {
		panic(err)
	}

	for _, sh := range wb.Sheets {
		fmt.Println("Max row is", sh.MaxRow)
		sh.ForEachRow(rowVisitor)
	}
}
