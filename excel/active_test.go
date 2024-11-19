package test

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"testing"
)

func TestExcel1(t *testing.T) {
	f := excelize.NewFile()

	index := f.NewSheet("Sheet2")
	f.SetCellValue("Sheet2", "A2", "Hello, world.")
	f.SetActiveSheet(index)

}
