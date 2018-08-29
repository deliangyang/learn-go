package main

import (
	"github.com/tealeg/xlsx"
	"fmt"
	"log"
)

func main() {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell

	style := xlsx.NewStyle()

	fill := *xlsx.NewFill("solid", "00FF0000", "FF000000")
	font := *xlsx.NewFont(20, "Verdana")
	border := *xlsx.NewBorder("thin", "thin", "thin", "thin")
	style.Fill = fill
	style.Font = font
	style.Border = border

	style.ApplyFill = true
	style.ApplyFont = true
	style.ApplyBorder = true

	file = xlsx.NewFile()
	sheet, er := file.AddSheet("SheetName")
	if er != nil {
		log.Fatal(er)
	}
	row = sheet.AddRow()

	cell = row.AddCell()
	cell.Value = "100000"
	cell.SetStyle(style)

	cell = row.AddCell()
	cell.Value = "test"

	err := file.Save("D:\\test.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
