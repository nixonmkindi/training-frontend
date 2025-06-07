package report

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"training-frontend/package/config"
	"training-frontend/package/log"
	"training-frontend/package/util"

	"github.com/signintech/gopdf"
)

//Generate report
// [title] report title
// [data] contains two dimension data, both in string, including serial number, the first row is the table title
// [columnWidth] array of width of each column, number of table columns should be equal to column dimensions of the data
// [isLandscape] defines the page setup, true for landscape and false for potrait

func GeneralReport(mainTitle, title string, data [][]string, columnWidth []float64, fileName string, fontSize int, isLandcape bool) string {
	begin := time.Now()

	d, err := json.Marshal(data)
	if util.IsError(err) {
		log.Errorf("error getting json data %v", err)
		return ""
	}

	qrs, doi, err := util.GetQRString(d)
	if util.IsError(err) {
		log.Errorf("error getting qr string and doi %v", err)
		return ""
	}
	//variable initalisation
	//fmt.Printf("Is landscape: %v\n", isLandcape)
	currentPageNumber = 1
	totalPageNumber = 0
	pdf := initPDF(isLandcape)
	pdf.SetMargins(leftMargin, topMargin, rightMargin, bottomMargin)

	//add fonts
	addFonts(pdf)
	pdf.AddPage()
	pdf.SetX(rightMargin)
	pdf.SetY(topMargin)
	totalRows = float64(len(data)) //initalise the number of rows

	//display header
	title = strings.ToUpper(title)
	mainHeader(pdf, mainTitle, title, qrs, doi)

	xp := leftMargin
	xy := headingMargin

	columnWidth = normaliseTableWidth(columnWidth, availablePageWidth)

	tableHeading := data[0]
	data = data[1:]
	addTable(pdf, xp, xy, 20, 5, columnWidth, tableHeading, data, fontSize)

	reportDir, err := config.ReportDir()
	if util.IsError(err) {
		log.Errorf("error getting report directory %v", err)
		return ""
	}
	timeFileName := fmt.Sprintf("-%d", time.Now().Unix())
	path := reportDir + fileName + timeFileName + ".pdf"
	pdf.WritePdf(path)
	pdf.Close()

	end := time.Now()
	fmt.Printf("PDF Report generated in %v\n", end.Sub(begin))
	return path
}

func TimetablePDF(mainTitle, title string, data [][]string, columnWidth []float64, fileName string, fontSize int, timetable []map[string]interface{}, isLandcape bool, whichTimetable int) string {
	begin := time.Now()
	var folder string
	var filename string
	d, err := json.Marshal(data)
	if util.IsError(err) {
		log.Errorf("error getting json data %v", err)
		return ""
	}

	qrs, doi, err := util.GetQRString(d)
	if util.IsError(err) {
		log.Errorf("error getting qr string and doi %v", err)
		return ""
	}
	//variable initalisation
	//fmt.Printf("Is landscape: %v\n", isLandcape)
	currentPageNumber = 1
	totalPageNumber = 0
	pdf := initPDF(isLandcape)
	pdf.SetMargins(leftMargin, topMargin, rightMargin, bottomMargin)

	//add fonts
	addFonts(pdf)
	pdf.AddPage()
	pdf.SetX(rightMargin)
	pdf.SetY(topMargin)
	totalRows = float64(len(data)) //initalise the number of rows

	//display header
	title = strings.ToUpper(title)
	mainHeader(pdf, mainTitle, title, qrs, doi)

	xp := leftMargin
	xy := headingMargin

	columnWidth = normaliseTableWidth(columnWidth, availablePageWidth)

	tableHeading := data[0]
	data = data[1:]
	addTable(pdf, xp, xy, 28, 5, columnWidth, tableHeading, data, fontSize)

	//for the days(horizontal line monday[131.378], tuesday[267.378], wednesday[403.378], thursday[539.378], friday[675.378])
	//for the hours(vertical line not yet calculated 150,178,205,233,261,289,317,345.2,373.2,401,429,457,485.2,514)

	conversionDays := map[string]float64{
		"Monday":    131.378,
		"Tuesday":   267.378,
		"Wednesday": 403.378,
		"Thursday":  539.378,
		"Friday":    675.378,
	}
	conversionTimes := map[string]float64{
		"08:00": 150,
		"09:00": 178,
		"10:00": 206,
		"11:00": 234,
		"12:00": 262,
		"13:00": 290,
		"14:00": 318,
		"15:00": 346,
		"16:00": 374,
		"17:00": 402,
		"18:00": 430,
		"19:00": 458,
		"20:00": 486,
		"21:00": 514,
	}
	var lineWidth, lineHeight float64
	lineWidth = 156.378
	// var block string

	for i := 0; i < len(timetable); i++ {
		students := fmt.Sprintf("%v", timetable[i]["StudentSets"])

		tags := fmt.Sprintf("%v", timetable[i]["ActivityTags"])
		var tag string
		if tags != "" {
			tag = strings.Split(tags, "")[0]
		} else {
			tag = ""
		}
		subjects := fmt.Sprintf("%v", timetable[i]["Subject"])
		// subject := strings.Split(strings.Split(subjects, "(")[1], ")")[0]
		subjectAndTag := subjects + "(" + tag + ")"

		day := fmt.Sprintf("%v", timetable[i]["Day"])
		room := fmt.Sprintf("%v", timetable[i]["Room"])

		teacher := fmt.Sprintf("%v", timetable[i]["Teachers"])
		teacherName := strings.Split(teacher, " ")
		teacherFirstName := teacherName[0]
		var t, teacherSecondName string
		if len(teacherName) > 2 {
			teacherSecondName = teacherName[2]
			t = strings.Split(teacherSecondName, "")[0]

		} else {
			teacherSecondName = ""
			t = ""

		}
		TeacherName := teacherFirstName + "," + t

		dayToXY := conversionDays[day]
		time := fmt.Sprintf("%v", timetable[i]["StartHour"])
		timeToXY := conversionTimes[time]

		endTime := fmt.Sprintf("%v", timetable[i]["EndHour"])
		a := strings.Split(time, ":")
		b := strings.Split(endTime, ":")

		startTimeInt, _ := strconv.Atoi(a[0])
		endTimeInt, _ := strconv.Atoi(b[0])
		duration := endTimeInt - startTimeInt
		if duration == 1 {
			lineHeight = 28.0
		}
		if duration == 2 {
			lineHeight = 56.0
		}
		if duration == 3 {
			lineHeight = 84.0
		}
		if duration == 4 {
			lineHeight = 112.0
		}
		var arry []string
		//room timetable
		if whichTimetable == 1 {
			arry = []string{subjectAndTag, TeacherName, students}
			folder = "rooms/"
			filename = fmt.Sprintf("%v", timetable[i]["Room"])

		}
		//students timetable
		if whichTimetable == 2 {
			arry = []string{subjectAndTag, TeacherName, room}
			folder = "students/"
			filename = fmt.Sprintf("%v", timetable[i]["StudentSets"])
		}
		//teachers timetable
		if whichTimetable == 3 {
			arry = []string{subjectAndTag, students, room}
			folder = "teachers/"
			filename = fmt.Sprintf("%v", timetable[i]["Teachers"])

		}
		addCustomMultiLineBlock2(pdf, dayToXY, timeToXY, lineWidth, lineHeight, arry, true)

	}

	reportDir, err := config.ReportDir()
	if util.IsError(err) {
		log.Errorf("error getting report directory %v", err)
		return ""
	}
	timeFileName := fmt.Sprintf("-%d", time.Now().Unix())
	path := reportDir + folder + filename + timeFileName + ".pdf"
	pdf.WritePdf(path)
	pdf.Close()

	end := time.Now()
	fmt.Printf("PDF Report generated in %v\n", end.Sub(begin))
	return path
}

//GeneralReportAppendPages returns pdf that can be appended other pages
// [title] report title
// [data] contains two dimension data, both in string, including serial number, the first row is the table title
// [columnWidth] array of width of each column, number of table columns should be equal to column dimensions of the data
// [appendPages] indicate the number of pages to be added at the end
// [isLandscape] defines the page setup, true for landscape and false for potrait

func GeneralReportAppendPages(mainTitle, title string, data [][]string, columnWidth []float64, fontSize int, appendPages int, isLandcape bool) *gopdf.GoPdf {

	d, err := json.Marshal(data)
	if util.IsError(err) {
		log.Errorf("error getting json data %v", err)
	}

	qrs, doi, err := util.GetQRString(d)
	if util.IsError(err) {
		log.Errorf("error getting qr string and doi %v", err)
	}
	//variable initalisation
	//fmt.Printf("Is landscape: %v\n", isLandcape)
	currentPageNumber = 1
	totalPageNumber = appendPages
	pdf := initPDF(isLandcape)
	pdf.SetMargins(leftMargin, topMargin, rightMargin, bottomMargin)

	//add fonts
	addFonts(pdf)
	pdf.AddPage()
	pdf.SetX(rightMargin)
	pdf.SetY(topMargin)
	totalRows = float64(len(data)) //initalise the number of rows

	//display header
	//title = strings.ToUpper(title)
	mainHeader(pdf, mainTitle, title, qrs, doi)

	xp := leftMargin
	xy := headingMargin

	//normalise table width
	totalWidth := 0.0
	for _, w := range columnWidth {
		totalWidth += w
	}
	for c, w := range columnWidth {
		columnWidth[c] = w / totalWidth * availablePageWidth
	}

	tableHeading := data[0]
	data = data[1:]
	addTable(pdf, xp, xy, 20, 5, columnWidth, tableHeading, data, fontSize)

	return pdf
}

func CombineTimetablePDF(whichTimetable string) string {
	// Set the working directory to the folder
	timetablesPath := ".storage/reports/" + whichTimetable

	pdfFiles, err := filepath.Glob(timetablesPath + "/*.pdf")
	if err != nil {
		fmt.Println("Error:", err)
	}
	if len(pdfFiles) == 0 {
		fmt.Println("Error: No PDF files found")
	}
	fileName := "ALL_COMBINED_" + whichTimetable + ".pdf"
	args := append(pdfFiles, "cat", "output", fileName)
	cmd := exec.Command("pdftk", args...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Details:", string(output))
	}
	fmt.Println("Output:", string(output))
	//deleting the content of the folder for  latter use
	d, err := os.Open(timetablesPath)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		fmt.Println("Error:", err)
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(timetablesPath, name))
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
	fmt.Println("Deleted all contents of the folder:", timetablesPath)
	return fileName
}

func AdmissionReportPage(mainTitle, title string, data [][]string, columnWidth []float64, fontSize int, appendPages int, isLandcape bool) *gopdf.GoPdf {

	//variable initalisation
	//fmt.Printf("Is landscape: %v\n", isLandcape)
	currentPageNumber = 1
	totalPageNumber = appendPages
	pdf := initPDF(isLandcape)
	pdf.SetMargins(leftMargin, topMargin, rightMargin, bottomMargin)

	//add fonts
	addFonts(pdf)
	pdf.AddPage()
	pdf.SetX(rightMargin)
	pdf.SetY(topMargin)
	totalRows = float64(len(data)) //initalise the number of rows

	//display header
	//title = strings.ToUpper(title)
	admissionHeader(pdf, mainTitle, title)

	xp := leftMargin
	xy := headingMargin

	//normalise table width
	totalWidth := 0.0
	for _, w := range columnWidth {
		totalWidth += w
	}
	for c, w := range columnWidth {
		columnWidth[c] = w / totalWidth * availablePageWidth
	}

	tableHeading := data[0]
	data = data[1:]

	addTableleftAlign(pdf, xp, xy, 20, 2, columnWidth, tableHeading, data, fontSize)

	return pdf
}
