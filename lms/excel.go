package lms

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"nik-cli/gpee/model"
	"strconv"
	"strings"
)

func MakeExcelFile(date string, data []model.HistoryDataRow) error {
	f := excelize.NewFile()
	f.SetCellValue("Sheet1", "A1", "Дата:")
	f.SetCellValue("Sheet1", "B1", date)
	f.SetCellValue("Sheet1", "C1", "W-код:")
	f.SetCellValue("Sheet1", "D1", "62W7675366327771")

	f.SetCellValue("Sheet1", "A2", "Година")
	f.SetCellValue("Sheet1", "B2", "00-14")
	f.SetCellValue("Sheet1", "C2", "15-29")
	f.SetCellValue("Sheet1", "D2", "30-44")
	f.SetCellValue("Sheet1", "E2", "45-59")

	for i, r := range data {
		idx := strconv.Itoa(i + 3)
		_ = f.SetCellValue("Sheet1", "A"+idx, r.Hour)
		e := r.E
		if strings.HasPrefix(e, "-") || e == "0.000" {
			e = "0"
		}

		_ = f.SetCellValue("Sheet1", "B"+idx, e)
		_ = f.SetCellValue("Sheet1", "C"+idx, e)
		_ = f.SetCellValue("Sheet1", "D"+idx, e)
		_ = f.SetCellValue("Sheet1", "E"+idx, e)
	}

	if err := f.SaveAs("1.xlsx"); err != nil {
		fmt.Println(err)
	}
	return nil
}
