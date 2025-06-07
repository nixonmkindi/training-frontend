package excel

import (
	"errors"
	"training-frontend/package/log"
	"training-frontend/package/util"

	"github.com/xuri/excelize/v2"
)

const (
	compareLoansTemplate     = "template_compare_loan.xlsx"
	teachingLoadTemplate     = "teaching_load_report_template.xlsx"
	failedStudentsTemplate   = "failed_students.xlsx"
	teachingLoadStartingRow  = 11
	failedStudentStartingRow = 8
	cmstartingRow            = 4
	loanCompare              = 3
	tamisemiStartingRow      = 3
)

type Tamisemi struct {
	FirstName      string `json:"first_name"`
	MiddleName     string `json:"middle_name"`
	LastName       string `json:"last_name"`
	BirthDate      string `json:"birth_date"`
	Gender         string `json:"gender"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone_number"`
	FivIndexNumber string `json:"fiv_index_num"`
	ProgramID      string `json:"program_id"`
}

//func GenerateTeachingLoadExcel(data []*models.TeachingLoad) (string, error) {
//
//	path := config.TemplatePath()
//	xlsx, err := excelize.OpenFile(path + teachingLoadTemplate)
//	if util.CheckError(err) {
//		return "", err
//	}
//	xlsx.WriteToBuffer()
//
//	//check integrity
//	sheetName, err := checkModuleResultIntegrity(xlsx)
//	if !util.CheckError(err) {
//		log.Infoln("teaching load template is good")
//	}
//
//	//insert header info
//	departmentCell := "A3"
//	academicyrCell := "F6"
//	semesterCell := "F7"
//	//academicYear := "C7"
//	xlsx.SetCellValue(sheetName, departmentCell, "Department of "+data[0].Staff.Department)
//	xlsx.SetCellValue(sheetName, academicyrCell, data[0].AcademicYear)
//	xlsx.SetCellValue(sheetName, semesterCell, data[0].DataSemester)
//	xv := int32(1)
//	for r, row := range data {
//
//		//serial number
//		cell := fmt.Sprintf("A%d", r+teachingLoadStartingRow)
//		cell2 := fmt.Sprintf("A%d", r+teachingLoadStartingRow)
//
//		if row.Staff.FullName != "" {
//
//			xlsx.SetCellValue(sheetName, cell, xv)
//			xv++
//		}
//		if row.Staff.FullName == "" {
//			//	fmt.Printf("in merging cells\n\n\n\n")
//			cell = fmt.Sprintf("A%d", r+teachingLoadStartingRow)
//			cell2 = fmt.Sprintf("A%d", (r+teachingLoadStartingRow)-1)
//			//cell2 = fmt.Sprintf("B%d", r+teachingLoadStartingRow)
//			errMerge := xlsx.MergeCell(sheetName, cell2, cell)
//			if errMerge != nil {
//				log.Errorf("error occurred :%v", errMerge)
//				return "", errMerge
//			}
//		}
//
//		//Staff full Name
//		if row.Staff.FullName != "" {
//			//	fmt.Printf("in UN merging cells\n")
//			cell = fmt.Sprintf("B%d", r+teachingLoadStartingRow)
//			xlsx.SetCellValue(sheetName, cell, row.Staff.FullName)
//		}
//		if row.Staff.FullName == "" {
//			//	fmt.Printf("in merging cells\n\n\n\n")
//			cell = fmt.Sprintf("B%d", r+teachingLoadStartingRow)
//			cell2 = fmt.Sprintf("B%d", (r+teachingLoadStartingRow)-1)
//			//cell2 = fmt.Sprintf("B%d", r+teachingLoadStartingRow)
//			errMerge := xlsx.MergeCell(sheetName, cell2, cell)
//			if errMerge != nil {
//				log.Errorf("error occurred :%v", errMerge)
//				return "", errMerge
//			}
//		}
//
//		//staff Gender
//		if row.Staff.Gender != "" {
//			//fmt.Printf("in UN merging cells\n")
//			cell = fmt.Sprintf("C%d", r+teachingLoadStartingRow)
//			xlsx.SetCellValue(sheetName, cell, row.Staff.Gender)
//		}
//		if row.Staff.Gender == "" {
//			//	fmt.Printf("in merging cells\n\n\n\n")
//			cell = fmt.Sprintf("C%d", r+teachingLoadStartingRow)
//			cell2 = fmt.Sprintf("C%d", (r+teachingLoadStartingRow)-1)
//			//cell2 = fmt.Sprintf("B%d", r+teachingLoadStartingRow)
//			errMerge := xlsx.MergeCell(sheetName, cell2, cell)
//			if errMerge != nil {
//				log.Errorf("error occurred :%v", errMerge)
//				return "", errMerge
//			}
//		}
//		//employment type
//
//		if row.Staff.EmploymentType != "" {
//
//			//fmt.Printf("in UN merging cells\n")
//			cell = fmt.Sprintf("D%d", r+teachingLoadStartingRow)
//			xlsx.SetCellValue(sheetName, cell, row.Staff.EmploymentType)
//		}
//		if row.Staff.EmploymentType == "" {
//			//	fmt.Printf("in merging cells\n\n\n\n")
//			cell = fmt.Sprintf("D%d", r+teachingLoadStartingRow)
//			cell2 = fmt.Sprintf("D%d", (r+teachingLoadStartingRow)-1)
//			//cell2 = fmt.Sprintf("B%d", r+teachingLoadStartingRow)
//			errMerge := xlsx.MergeCell(sheetName, cell2, cell)
//			if errMerge != nil {
//				log.Errorf("error occurred :%v", errMerge)
//				return "", errMerge
//			}
//		}
//		//module code
//		cell = fmt.Sprintf("E%d", r+teachingLoadStartingRow)
//		xlsx.SetCellValue(sheetName, cell, row.Module.ModuleCode)
//		//module name
//		cell = fmt.Sprintf("F%d", r+teachingLoadStartingRow)
//		xlsx.SetCellValue(sheetName, cell, row.Module.ModuleName)
//		//class name
//		cell = fmt.Sprintf("G%d", r+teachingLoadStartingRow)
//		xlsx.SetCellValue(sheetName, cell, row.ClassName)
//		//Credit
//		cell = fmt.Sprintf("H%d", r+teachingLoadStartingRow)
//		xlsx.SetCellValue(sheetName, cell, row.Module.Credit)
//
//		//lecture hours
//		cell = fmt.Sprintf("I%d", r+teachingLoadStartingRow)
//		xlsx.SetCellValue(sheetName, cell, row.Module.LectureHours)
//
//		//practical hours
//		cell = fmt.Sprintf("J%d", r+teachingLoadStartingRow)
//		xlsx.SetCellValue(sheetName, cell, row.Module.PracticalHours)
//
//		//independent studies hours
//		cell = fmt.Sprintf("K%d", r+teachingLoadStartingRow)
//
//		xlsx.SetCellValue(sheetName, cell, row.Module.TutorialHours)
//
//		//tutorials hours
//		//cell = fmt.Sprintf("L%d", r+teachingLoadStartingRow)
//		//xlsx.SetCellValue(sheetName, cell, row.Module.IndependentStudies)
//
//		//tutorials hours
//		cell = fmt.Sprintf("L%d", r+teachingLoadStartingRow)
//		xlsx.SetCellValue(sheetName, cell, row.Module.TotalHours)
//
//		if row.Staff.EmploymentType != "" {
//			//totalHours = row.Module.TotalHours
//
//			cell = fmt.Sprintf("M%d", r+teachingLoadStartingRow)
//			xlsx.SetCellValue(sheetName, cell, row.Module.TotalHours)
//
//			//fmt.Printf("in UN merging cells\n")
//		}
//		if row.Staff.FullName == "" {
//			row.Module.TotalHours = row.Module.TotalHours + data[r-1].Module.TotalHours
//			//	fmt.Printf("in merging cells\n\n\n\n")
//			cell = fmt.Sprintf("M%d", r+teachingLoadStartingRow)
//			cell2 = fmt.Sprintf("M%d", (r+teachingLoadStartingRow)-1)
//			//cell2 = fmt.Sprintf("B%d", r+teachingLoadStartingRow)
//			errMerge := xlsx.MergeCell(sheetName, cell2, cell)
//			if errMerge != nil {
//				log.Errorf("error occurred :%v", errMerge)
//				return "", errMerge
//			}
//			xlsx.SetCellValue(sheetName, cell, row.Module.TotalHours)
//		}
//
//	}
//	filePath := fmt.Sprintf("%s\\teaching_load.xlsx", path)
//	xlsx.SaveAs(filePath)
//	return filePath, nil
//}
////
//func FailedStudents(data []*examModel.Student) (string, error) {
//
//	path := config.TemplatePath()
//	xlsx, err := excelize.OpenFile(path + failedStudentsTemplate)
//	if util.CheckError(err) {
//		return "", err
//	}
//	xlsx.WriteToBuffer()
//
//	//check integrity
//	sheetName, err := checkModuleResultIntegrity(xlsx)
//	if !util.CheckError(err) {
//		log.Infoln("teaching load template is good")
//	}
//
//	for r, row := range data {
//		cell := fmt.Sprintf("A%d", r+failedStudentStartingRow)
//		fivIndex := strings.Replace(row.CSEENumber, "/", ".", -1)
//		//serial number
//		cell = fmt.Sprintf("A%d", r+failedStudentStartingRow)
//		xlsx.SetCellValue(sheetName, cell, r+1)
//
//		//fivIndex number
//		cell = fmt.Sprintf("B%d", r+failedStudentStartingRow)
//		xlsx.SetCellValue(sheetName, cell, fivIndex)
//
//		//amount
//		cell = fmt.Sprintf("C%d", r+failedStudentStartingRow)
//		xlsx.SetCellValue(sheetName, cell, row.Amount)
//		//remarks
//		cell = fmt.Sprintf("D%d", r+failedStudentStartingRow)
//		xlsx.SetCellValue(sheetName, cell, row.Description)
//	}
//	filePath := fmt.Sprintf("%s\\studentFailedUpload.xlsx", path)
//	xlsx.SaveAs(filePath)
//	return filePath, nil
//}
//
//func CMDDetails(filename string) (string, error) {
//	path := config.TemplatePath()
//	xlsx, err := excelize.OpenFile(path + compareLoansTemplate)
//	if util.CheckError(err) {
//		return "", err
//	}
//	xlsx.WriteToBuffer()
//
//	//check integrity
//	sheetS, err := getSheetMultipleName(xlsx)
//	if util.CheckError(err) {
//		log.Infoln("error reading from file")
//	}
//	sheetName := sheetS[0]
//	//insert header info
//	style, _ := xlsx.NewStyle(`{"border":[{"type":"left","color":"000000","style":2},
//			{"type":"top","color":"000000","style":2},{"type":"bottom","color":"000000","style":2},
//			{"type":"right","color":"000000","style":2}]}`)
//
//	xlsxe, err := excelize.OpenFile(filename)
//	if util.CheckError(err) {
//		log.Errorf("error importing result:%v\n", err)
//		return "nil", err
//	}
//
//	//get sheet name
//	sheetNam, err := getSheetMultipleName(xlsxe)
//	if util.CheckError(err) {
//		log.Errorf("error getting sheet name:%v\n", err)
//		return "nil", err
//	}
//
//	rw, rw1 := int32(-1), int32(0)
//
//	for {
//
//		rw++
//		fivIndex := "B"
//		cellFivIndex := fmt.Sprintf("%s%d", fivIndex, loanCompare+rw)
//
//		fivIndexNumber, _ := xlsxe.GetCellValue(sheetNam[0], cellFivIndex)
//		if fivIndexNumber == "" {
//			break //terminate if the cell value is empty
//		}
//		//colns := make(map[string]string)
//		columns := []string{"C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N"}
//		columnDataFromDB := []string{"O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
//		var rowValues, somaValues [12]string
//		//
//
//		//get column values from xlsx
//		for i, c := range columns {
//			cellName := fmt.Sprintf("%s%d", c, loanCompare+rw)
//			//cellNamename := fmt.Sprintf("%s%d", name, rw)
//			value, _ := xlsxe.GetCellValue(sheetNam[0], cellName)
//			if value == "" || value == "0" {
//				continue
//			} else {
//				rowValues[i] = value
//			}
//		}
//		billService := bills.NewService()
//		// check in db if fiv index has loan
//		cseeNumber := strings.Replace(fivIndexNumber, ".", "/", -1)
//
//		byCSEENumber, err := billService.ListBillsByCSEENumber(cseeNumber)
//		if err != nil {
//			pp.Println(err)
//			//append all data to xlsx file
//			for i, data := range rowValues {
//				if data != "" && data != "0" {
//					//fiv index
//					celliv := fmt.Sprintf("B%d", loanCompare+rw1)
//					xlsx.SetCellStyle(sheetName, celliv, celliv, style)
//					xlsx.SetCellValue(sheetName, celliv, fivIndexNumber)
//					cLetter := columns[i]
//					cell := fmt.Sprintf("%s%d", cLetter, loanCompare+rw1)
//					xlsx.SetCellStyle(sheetName, cell, cell, style)
//					xlsx.SetCellValue(sheetName, cell, data)
//
//				}
//			}
//			rw1++
//			continue
//
//		}
//		//if no loan at all,  append all data to xlsx file
//		if len(byCSEENumber) == 0 {
//			//write all column values for this csee number
//			for i, data := range rowValues {
//				if data != "" && data != "0" {
//					//fiv index
//					celliv := fmt.Sprintf("B%d", loanCompare+rw1)
//					xlsx.SetCellStyle(sheetName, celliv, celliv, style)
//					xlsx.SetCellValue(sheetName, celliv, fivIndexNumber)
//
//					cLetter := columns[i]
//					cell := fmt.Sprintf("%s%d", cLetter, loanCompare+rw1)
//					xlsx.SetCellStyle(sheetName, cell, cell, style)
//					xlsx.SetCellValue(sheetName, cell, data)
//
//				}
//			}
//			rw1++
//			continue
//		}
//		// if exists
//		for _, data := range byCSEENumber {
//			if data.BillAmount == "" || data.BillAmount == "0" {
//				continue
//			}
//			if data.ItemID == 1 {
//				//0 index
//				somaValues[0] = data.BillAmount
//			} else if data.ItemID == 14 {
//				//i
//				somaValues[1] = data.BillAmount
//			} else if data.ItemID == 15 {
//				somaValues[2] = data.BillAmount
//
//			} else if data.ItemID == 20 {
//				somaValues[3] = data.BillAmount
//
//			} else if data.ItemID == 24 {
//				somaValues[4] = data.BillAmount
//
//			} else if data.ItemID == 7 {
//				somaValues[5] = data.BillAmount
//
//			} else if data.ItemID == 3 {
//				somaValues[6] = data.BillAmount
//
//			} else if data.ItemID == 6 {
//				somaValues[7] = data.BillAmount
//
//			} else if data.ItemID == 4 {
//				somaValues[8] = data.BillAmount
//
//			} else if data.ItemID == 5 {
//				somaValues[9] = data.BillAmount
//
//			} else if data.ItemID == 30 {
//				somaValues[10] = data.BillAmount
//
//			} else if data.ItemID == 25 {
//				somaValues[11] = data.BillAmount
//
//			}
//		}
//		//compare items amount
//		if rowValues == somaValues {
//			continue
//		}
//
//		var diff bool
//		for i, data := range rowValues {
//			if data != "" && data != "0" {
//				diff = true
//				//fiv index
//				celliv := fmt.Sprintf("B%d", loanCompare+rw1)
//				xlsx.SetCellStyle(sheetName, celliv, celliv, style)
//				xlsx.SetCellValue(sheetName, celliv, fivIndexNumber)
//				cLetter := columns[i]
//				cell := fmt.Sprintf("%s%d", cLetter, loanCompare+rw1)
//				xlsx.SetCellStyle(sheetName, cell, cell, style)
//				xlsx.SetCellValue(sheetName, cell, data)
//
//			}
//		}
//
//		for i, data := range somaValues {
//			if data != "" && data != "0" {
//				diff = true
//				//write data to xlsx
//				//fiv index
//				celliv := fmt.Sprintf("B%d", loanCompare+rw1)
//				xlsx.SetCellStyle(sheetName, celliv, celliv, style)
//				xlsx.SetCellValue(sheetName, celliv, fivIndexNumber)
//				cLetter := columnDataFromDB[i]
//				cell := fmt.Sprintf("%s%d", cLetter, loanCompare+rw1)
//				xlsx.SetCellStyle(sheetName, cell, cell, style)
//				xlsx.SetCellValue(sheetName, cell, data)
//			}
//		}
//		if diff {
//			rw1++
//		}
//	}
//
//	filePath := fmt.Sprintf("%s\\studentFailedUpload.xlsx", path)
//	xlsx.SaveAs(filePath)
//	return filePath, nil
//}
//
//func HESLCompareDetails(filename string, students []*registration_entity.Student) (string, error) {
//	path := config.TemplatePath()
//
//	xlsxe, err := excelize.OpenFile(filename)
//	if util.CheckError(err) {
//		log.Errorf("error importing result:%v\n", err)
//		return "nil", err
//	}
//
//	//get sheet name
//	sheetNam, err := getSheetMultipleName(xlsxe)
//	if util.CheckError(err) {
//		log.Errorf("error getting sheet name:%v\n", err)
//		return "nil", err
//	}
//
//	data1 := make(map[int32]string)
//	data2 := make(map[string]string)
//	rw, rw1 := int32(0), int32(0)
//
//	//	fmt.Printf("I am here...\n")
//
//	for {
//		//cellName := fmt.Sprintf("%s%d", "B", cmstartingRow+rw)
//		//cellValue, err := xlsxe.GetCellValue(sheetNam[0], cellName)
//		//if util.CheckError(err) {
//		//	log.Errorf("error getting module code:%v\n", err)
//		//	return "", err
//		//}
//		//if cellValue == "" {
//		//	break //terminate if the cell value is empty
//		//}
//		regNo := "C"
//		cellNameregN := fmt.Sprintf("%s%d", regNo, cmstartingRow+rw)
//
//		//cellNamename := fmt.Sprintf("%s%d", name, rw)
//		regNoVal, _ := xlsxe.GetCellValue(sheetNam[0], cellNameregN)
//		if regNoVal == "" {
//			break //terminate if the cell value is empty
//		}
//		//nameVal, _ := xlsx.GetCellValue(sheetName[0], cellNamename)
//		data1[cmstartingRow+rw] = regNoVal
//		rw++
//
//	}
//
//	for {
//		name := "B"
//		regNo := "C"
//		cellNameregNo := fmt.Sprintf("%s%d", regNo, cmstartingRow+rw1)
//		cellNamename := fmt.Sprintf("%s%d", name, cmstartingRow+rw1)
//
//		regNoVal, _ := xlsxe.GetCellValue(sheetNam[1], cellNameregNo)
//		nameVal, _ := xlsxe.GetCellValue(sheetNam[1], cellNamename)
//		if regNoVal == "" {
//			break //terminate if the cell value is empty
//		}
//		data2[regNoVal] = nameVal
//		rw1++
//	}
//
//	for reg, valueName := range data2 {
//		for key, val := range data1 {
//			if reg == val {
//				//write to the excel
//				//CELL Value
//				l := "L"
//				m := "M"
//				cell4RegNo := fmt.Sprintf("%s%d", l, key)
//				cell4name := fmt.Sprintf("%s%d", m, key)
//
//				xlsxe.SetCellValue(sheetNam[0], cell4RegNo, key)
//
//				//cell = fmt.Sprintf("D%d", r+failedStudentStartingRow)
//				xlsxe.SetCellValue(sheetNam[0], cell4RegNo, reg)
//
//				//fivIndex number
//				xlsxe.SetCellValue(sheetNam[0], cell4name, valueName)
//				break
//			}
//
//		}
//
//	}
//
//	filePath := fmt.Sprintf("%s\\studentFailedUpload.xlsx", path)
//	xlsxe.SaveAs(filePath)
//	return filePath, nil
//}
//
//func GeneralFileComparison(filename string) (string, error) {
//	path := config.TemplatePath()
//	//xlsx, err := excelize.OpenFile(path + compareLoansTemplate)
//	//if util.CheckError(err) {
//	//	return "", err
//	//}
//	//xlsx.WriteToBuffer()
//	//check integrity
//	//sheetS, err := getSheetMultipleName(xlsx)
//	//if util.CheckError(err) {
//	//	log.Infoln("error reading from file")
//	//}
//	//sheetName := sheetS[0]
//	//insert header info
//
//	xlsxe, err := excelize.OpenFile(filename)
//	if util.CheckError(err) {
//		log.Errorf("error importing result:%v\n", err)
//		return "nil", err
//	}
//	_, _ = xlsxe.NewStyle(`{"border":[{"type":"left","color":"000000","style":2},
//			{"type":"top","color":"000000","style":2},{"type":"bottom","color":"000000","style":2},
//			{"type":"right","color":"000000","style":2}]}`)
//
//	fillStyle, err := xlsxe.NewStyle(
//		&excelize.Style{
//			Alignment: &excelize.Alignment{
//				Vertical: "center",
//				WrapText: true,
//			},
//			Border: []excelize.Border{
//				{Type: "left", Color: "000000", Style: 1},
//				{Type: "right", Color: "000000", Style: 1},
//				{Type: "top", Color: "000000", Style: 1},
//				{Type: "bottom", Color: "000000", Style: 1},
//			},
//			Fill: excelize.Fill{
//				Type:    "pattern",
//				Pattern: 1,
//				Color:   []string{"#208637"},
//			},
//			//Font: &excelize.Font{
//			//	Bold: true,
//			//},
//		},
//	)
//
//	fillStyle2, err := xlsxe.NewStyle(
//		&excelize.Style{
//			Alignment: &excelize.Alignment{
//				Vertical: "center",
//				WrapText: true,
//			},
//			Border: []excelize.Border{
//				{Type: "left", Color: "000000", Style: 1},
//				{Type: "right", Color: "000000", Style: 1},
//				{Type: "top", Color: "000000", Style: 1},
//				{Type: "bottom", Color: "000000", Style: 1},
//			},
//			Fill: excelize.Fill{
//				Type:    "pattern",
//				Pattern: 1,
//				Color:   []string{"#FF5A69"},
//			},
//			//Font: &excelize.Font{
//			//	Bold: true,
//			//},
//		},
//	)
//	//get sheet name
//	sheetNam, err := getSheetMultipleName(xlsxe)
//	if util.CheckError(err) {
//		log.Errorf("error getting sheet name:%v\n", err)
//		return "nil", err
//	}
//
//	data1 := make(map[string]string)
//	data2 := make(map[string]string)
//	rw, rw1 := int32(-1), int32(-1)
//
//	for {
//		rw1++
//		//name := "B"
//		regNo := "B"
//		amount := "C"
//		cellNameregNo := fmt.Sprintf("%s%d", regNo, loanCompare+rw1)
//		cellAmount := fmt.Sprintf("%s%d", amount, loanCompare+rw1)
//		//	cellNamename := fmt.Sprintf("%s%d", name, cmstartingRow+rw1)
//
//		regNoVal, _ := xlsxe.GetCellValue(sheetNam[0], cellNameregNo)
//		nameVal, _ := xlsxe.GetCellValue(sheetNam[0], cellAmount)
//		if regNoVal == "" {
//			break //terminate if the cell value is empty
//		}
//		data2[regNoVal] = nameVal
//	}
//	var count int32
//
//	for {
//		rw++
//		fivIndex := "B"
//		amount := "C"
//		cellFivIndex := fmt.Sprintf("%s%d", fivIndex, loanCompare+rw)
//
//		cellAmount := fmt.Sprintf("%s%d", amount, loanCompare+rw)
//		fivIndexNumber, _ := xlsxe.GetCellValue(sheetNam[1], cellFivIndex)
//		if fivIndexNumber == "" {
//			break //terminate if the cell value is empty
//		}
//		nameVal, _ := xlsxe.GetCellValue(sheetNam[1], cellAmount)
//		data1[fivIndexNumber] = nameVal
//		//does he/she has additional??
//		v, ok := data2[fivIndexNumber]
//		if ok {
//			if v == nameVal {
//
//				a, _ := strconv.ParseFloat(v, 64)
//				b, _ := strconv.ParseFloat(nameVal, 64)
//				if a+b > 1350000 {
//					pp.Println(v, ">>", nameVal)
//					xlsxe.SetCellStyle(sheetNam[1], cellFivIndex, cellFivIndex, fillStyle)
//				} else {
//					xlsxe.SetCellValue(sheetNam[1], cellAmount, fmt.Sprintf("%.0f", a+b))
//				}
//
//			} else {
//				a, _ := strconv.ParseFloat(v, 64)
//				b, _ := strconv.ParseFloat(nameVal, 64)
//				if a+b > 1350000 {
//					pp.Println(v, ">>", nameVal)
//					xlsxe.SetCellStyle(sheetNam[1], cellFivIndex, cellFivIndex, fillStyle2)
//				} else {
//					xlsxe.SetCellValue(sheetNam[1], cellAmount, fmt.Sprintf("%.0f", a+b))
//				}
//			}
//
//			//delete
//		}
//
//	}
//	for k, v := range data2 {
//		_, ok := data1[k]
//		if ok {
//			continue
//		} else {
//			pp.Println(k, ">>", v)
//			pp.Println("fiv:", k, ">>", v)
//			rw++
//			fivIndex := "B"
//			amount := "C"
//			cellFivIndex := fmt.Sprintf("%s%d", fivIndex, loanCompare+rw)
//			cellAmount := fmt.Sprintf("%s%d", amount, loanCompare+rw)
//			xlsxe.SetCellValue(sheetNam[1], cellFivIndex, k)
//			xlsxe.SetCellValue(sheetNam[1], cellAmount, v)
//		}
//
//	}
//	pp.Println(count)
//	filePath := fmt.Sprintf("%s\\studentFailedUpload.xlsx", path)
//	xlsxe.SaveAs(filePath)
//	return filePath, nil
//}
//
//func GeneralFileComparison2(filename string) (string, error) {
//	path := config.TemplatePath()
//	//xlsx, err := excelize.OpenFile(path + compareLoansTemplate)
//	//if util.CheckError(err) {
//	//	return "", err
//	//}
//	//xlsx.WriteToBuffer()
//	//check integrity
//	//sheetS, err := getSheetMultipleName(xlsx)
//	//if util.CheckError(err) {
//	//	log.Infoln("error reading from file")
//	//}
//	//sheetName := sheetS[0]
//	//insert header info
//
//	xlsxe, err := excelize.OpenFile(filename)
//	if util.CheckError(err) {
//		log.Errorf("error importing result:%v\n", err)
//		return "nil", err
//	}
//	_, _ = xlsxe.NewStyle(`{"border":[{"type":"left","color":"000000","style":2},
//			{"type":"top","color":"000000","style":2},{"type":"bottom","color":"000000","style":2},
//			{"type":"right","color":"000000","style":2}]}`)
//
//	_, _ = xlsxe.NewStyle(
//		&excelize.Style{
//			Alignment: &excelize.Alignment{
//				Vertical: "center",
//				WrapText: true,
//			},
//			Border: []excelize.Border{
//				{Type: "left", Color: "000000", Style: 1},
//				{Type: "right", Color: "000000", Style: 1},
//				{Type: "top", Color: "000000", Style: 1},
//				{Type: "bottom", Color: "000000", Style: 1},
//			},
//			Fill: excelize.Fill{
//				Type:    "pattern",
//				Pattern: 1,
//				Color:   []string{"#208637"},
//			},
//			//Font: &excelize.Font{
//			//	Bold: true,
//			//},
//		},
//	)
//
//	_, _ = xlsxe.NewStyle(
//		&excelize.Style{
//			Alignment: &excelize.Alignment{
//				Vertical: "center",
//				WrapText: true,
//			},
//			Border: []excelize.Border{
//				{Type: "left", Color: "000000", Style: 1},
//				{Type: "right", Color: "000000", Style: 1},
//				{Type: "top", Color: "000000", Style: 1},
//				{Type: "bottom", Color: "000000", Style: 1},
//			},
//			Fill: excelize.Fill{
//				Type:    "pattern",
//				Pattern: 1,
//				Color:   []string{"#FF5A69"},
//			},
//			//Font: &excelize.Font{
//			//	Bold: true,
//			//},
//		},
//	)
//	//get sheet name
//	sheetNam, err := getSheetMultipleName(xlsxe)
//	if util.CheckError(err) {
//		log.Errorf("error getting sheet name:%v\n", err)
//		return "nil", err
//	}
//
//	data := make(map[string]string)
//	data2 := make(map[string]string)
//	rw, rw2 := int32(-1), int32(-1)
//
//	for {
//		rw++
//		//name := "B"
//		regNo := "B"
//		amount := "C"
//		cellNameregNo := fmt.Sprintf("%s%d", regNo, loanCompare+rw)
//		cellAmount := fmt.Sprintf("%s%d", amount, loanCompare+rw)
//		//	cellNamename := fmt.Sprintf("%s%d", name, cmstartingRow+rw1)
//
//		regNoVal, _ := xlsxe.GetCellValue(sheetNam[0], cellNameregNo)
//		nameVal, _ := xlsxe.GetCellValue(sheetNam[0], cellAmount)
//		if regNoVal == "" {
//			break //terminate if the cell value is empty
//		}
//
//		data[regNoVal] = nameVal
//	}
//
//	billService := bills.NewService()
//	loansBills, err := billService.ListALLLoansBills()
//	if err != nil {
//		log.Errorln("error occurred:", err)
//		return "", err
//	}
//	for _, dt := range loansBills {
//		fivIndex := strings.Replace(dt.CSEENumber, "/", ".", -1)
//
//		data2[fivIndex] = dt.PaidAmount
//		vc, ss := data[fivIndex]
//		if ss {
//			if dt.PaidAmount != vc {
//				//update loan amount
//				ac, err := strconv.ParseFloat(vc, 64)
//				if err != nil {
//					log.Errorln(err)
//					break
//				}
//				if vc != "" {
//					if float32(ac) > 1350000 {
//
//						pp.Println(float32(ac))
//					}
//					_, err = billService.UpdateOldPayment(dt.BillID, 1, float32(ac))
//					if err != nil {
//						log.Errorln(err)
//						break
//					}
//				} else {
//					pp.Println("loan to delete:%v", dt.PaidAmount, dt.BillID)
//					err = billService.HardDeleteBills(dt.BillID)
//					if err != nil {
//						log.Errorln(err)
//						break
//					}
//				}
//
//			} else if dt.PaidAmount == vc {
//				continue
//			}
//		} else {
//			//delete loans
//			//pp.Println("loan to delete:%v", dt.PaidAmount, dt.BillID)
//			err = billService.HardDeleteBills(dt.BillID)
//			if err != nil {
//				log.Errorln(err)
//				break
//			}
//
//		}
//	}
//	for k, val := range data {
//		_, s := data2[k]
//		if s {
//			continue
//		} else {
//			if val != "" {
//				rw2++
//				fivIndex := "B"
//				amount := "C"
//				cellFivIndex := fmt.Sprintf("%s%d", fivIndex, loanCompare+rw2)
//				cellAmount := fmt.Sprintf("%s%d", amount, loanCompare+rw2)
//				xlsxe.SetCellValue(sheetNam[1], cellFivIndex, k)
//				xlsxe.SetCellValue(sheetNam[1], cellAmount, val)
//			}
//
//		}
//	}
//	filePath := fmt.Sprintf("%s\\studentFailedUpload.xlsx", path)
//	xlsxe.SaveAs(filePath)
//	return filePath, nil
//}

func UploadTamisemiStudents(filename string) ([]*Tamisemi, error) {
	var Students []*Tamisemi
	xlsxe, err := excelize.OpenFile(filename)
	if util.IsError(err) {
		log.Errorf("error importing result:%v\n", err)
		return nil, err
	}

	_, _ = xlsxe.NewStyle(
		&excelize.Style{
			Alignment: &excelize.Alignment{
				Vertical: "center",
				WrapText: true,
			},
			Border: []excelize.Border{
				{Type: "left", Color: "000000", Style: 1},
				{Type: "right", Color: "000000", Style: 1},
				{Type: "top", Color: "000000", Style: 1},
				{Type: "bottom", Color: "000000", Style: 1},
			},
			Fill: excelize.Fill{
				Type:    "pattern",
				Pattern: 1,
				Color:   []string{"#208637"},
			},
			//Font: &excelize.Font{
			//	Bold: true,
			//},
		},
	)

	_, _ = xlsxe.NewStyle(
		&excelize.Style{
			Alignment: &excelize.Alignment{
				Vertical: "center",
				WrapText: true,
			},
			Border: []excelize.Border{
				{Type: "left", Color: "000000", Style: 1},
				{Type: "right", Color: "000000", Style: 1},
				{Type: "top", Color: "000000", Style: 1},
				{Type: "bottom", Color: "000000", Style: 1},
			},
			Fill: excelize.Fill{
				Type:    "pattern",
				Pattern: 1,
				Color:   []string{"#FF5A69"},
			},
			//Font: &excelize.Font{
			//	Bold: true,
			//},
		},
	)
	//get sheet name
	sheetNam, err := getSheetMultipleName(xlsxe)
	if util.IsError(err) {
		log.Errorf("error getting sheet name:%v\n", err)
		return nil, err
	}

	//rw := int32(-1)
	//
	//for {
	//	rw++
	//	Fname := "B"
	//	cell := fmt.Sprintf("%s%d", Fname, tamisemiStartingRow+rw)
	//	firstName, _ := xlsxe.GetCellValue(sheetNam[0], cell)
	//	if firstName == "" {
	//		pp.Println("empty value:", firstName)
	//		break //terminate if the cell value is empty
	//	}
	//	var v Tamisemi
	//	columns := []string{"B", "C", "D", "E", "F", "G", "H", "I"}
	//	for i, col := range columns {
	//		cellLetter := fmt.Sprintf("%s%d", col, tamisemiStartingRow+rw)
	//		cellValue, _ := xlsxe.GetCellValue(sheetNam[0], cellLetter)
	//		if i == 0 {
	//			v.FirstName = cellValue
	//		} else if i == 1 {
	//			v.MiddleName = cellValue
	//		} else if i == 2 {
	//			v.LastName = cellValue
	//		} else if i == 3 {
	//			v.BirthDate = cellValue
	//		} else if i == 4 {
	//			v.Gender = cellValue
	//		} else if i == 5 {
	//			v.Email = cellValue
	//		} else if i == 6 {
	//			v.PhoneNumber = cellValue
	//		} else if i == 7 {
	//			v.FivIndexNumber = cellValue
	//		}
	//
	//	}
	//	Students = append(Students, &v)
	//
	//}
	rows, _ := xlsxe.GetRows(sheetNam[0])
	for i, r := range rows {
		if i > 1 {
			Students = append(Students, &Tamisemi{
				FirstName:      r[1],
				MiddleName:     r[2],
				LastName:       r[3],
				BirthDate:      r[4],
				Gender:         r[5],
				Email:          r[6],
				PhoneNumber:    r[7],
				FivIndexNumber: r[8],
				ProgramID:      r[9],
			})
		}
	}

	return Students, nil
	//billService := bills.NewService()
	//loansBills, err := billService.ListALLLoansBills()
	//if err != nil {
	//	log.Errorln("error occurred:", err)
	//	return "", err
	//}
	//for _, dt := range loansBills {
	//	fivIndex := strings.Replace(dt.CSEENumber, "/", ".", -1)
	//
	//	data2[fivIndex] = dt.PaidAmount
	//	vc, ss := data[fivIndex]
	//	if ss {
	//		if dt.PaidAmount != vc {
	//			//update loan amount
	//			ac, err := strconv.ParseFloat(vc, 64)
	//			if err != nil {
	//				log.Errorln(err)
	//				break
	//			}
	//			if vc != "" {
	//				if float32(ac) > 1350000 {
	//
	//					pp.Println(float32(ac))
	//				}
	//				_, err = billService.UpdateOldPayment(dt.BillID, 1, float32(ac))
	//				if err != nil {
	//					log.Errorln(err)
	//					break
	//				}
	//			} else {
	//				pp.Println("loan to delete:%v", dt.PaidAmount, dt.BillID)
	//				err = billService.HardDeleteBills(dt.BillID)
	//				if err != nil {
	//					log.Errorln(err)
	//					break
	//				}
	//			}
	//
	//		} else if dt.PaidAmount == vc {
	//			continue
	//		}
	//	} else {
	//		//delete loans
	//		//pp.Println("loan to delete:%v", dt.PaidAmount, dt.BillID)
	//		err = billService.HardDeleteBills(dt.BillID)
	//		if err != nil {
	//			log.Errorln(err)
	//			break
	//		}
	//
	//	}
	//}
	//for k, val := range data {
	//	_, s := data2[k]
	//	if s {
	//		continue
	//	} else {
	//		if val != "" {
	//			rw2++
	//			fivIndex := "B"
	//			amount := "C"
	//			cellFivIndex := fmt.Sprintf("%s%d", fivIndex, loanCompare+rw2)
	//			cellAmount := fmt.Sprintf("%s%d", amount, loanCompare+rw2)
	//			xlsxe.SetCellValue(sheetNam[1], cellFivIndex, k)
	//			xlsxe.SetCellValue(sheetNam[1], cellAmount, val)
	//		}
	//
	//	}
	//}
	//filePath := fmt.Sprintf("%s\\studentFailedUpload.xlsx", path)
	//xlsxe.SaveAs(filePath)
	//return filePath, nil
}

func getSheetMultipleName(ex *excelize.File) ([]string, error) {
	sheets := ex.GetSheetList()
	if len(sheets) < 1 {
		return nil, errors.New("the excel file must have at lest one sheet")
	} else {
		return sheets, nil
	}
}
