function getReportType() {
  let reportTypeID = document.getElementById("reportTypeSelect").value;
  let classSelectBox = document.getElementById("classSelectBox");
  let programSelectBox = document.getElementById("programSelectBox");
  let departmentSelectBox = document.getElementById("departmentSelectBox");
  let campusSelectBox = document.getElementById("campusSelectBox");
  let ntaLevelSelectBox = document.getElementById("ntaLevelSelectBox");
  let bestStudentTypeSelectBox = document.getElementById(
    "bestStudentTypeSelectBox"
  );
  let yosSelectBox = document.getElementById("yosSelectBox");
  let genderSelectBox = document.getElementById("genderSelectBox");
  let awardClassificationSelectBox = document.getElementById(
    "awardClassificationSelectBox"
  );
  let gpaMeasurementSelectBox = document.getElementById(
    "gpaMeasurementSelectBox"
  );
  let gpaInputBox = document.getElementById("gpaInputBox");

  let reportClassID = document.getElementById("reportClassID");
  let reportProgramID = document.getElementById("reportProgramID");
  let reportDepartmentID = document.getElementById("reportDepartmentID");
  let reportCampusID = document.getElementById("reportCampusID");
  let gpaValue = document.getElementById("gpaValue");

  classSelectBox.style.display = "none";
  programSelectBox.style.display = "none";
  departmentSelectBox.style.display = "none";
  campusSelectBox.style.display = "none";

  reportClassID.removeAttribute("required");
  reportProgramID.removeAttribute("required");
  reportDepartmentID.removeAttribute("required");
  reportCampusID.removeAttribute("required");
  gpaValue.removeAttribute("required");

  switch (reportTypeID) {
    case "1":
      classSelectBox.style.display = "block";
      reportClassID.setAttribute("required", "true");
      break;

    case "2":
      programSelectBox.style.display = "block";
      reportProgramID.setAttribute("required", "true");
      break;

    case "3":
      departmentSelectBox.style.display = "block";
      reportDepartmentID.setAttribute("required", "true");
      break;

    case "4":
      campusSelectBox.style.display = "block";
      reportCampusID.setAttribute("required", "true");
      break;

    default:
      break;
  }

  if (reportTypeID !== "") {
    ntaLevelSelectBox.style.display = "block";
    bestStudentTypeSelectBox.style.display = "block";
    yosSelectBox.style.display = "block";
    genderSelectBox.style.display = "block";
    awardClassificationSelectBox.style.display = "block";
    gpaMeasurementSelectBox.style.display = "block";
  } else {
    ntaLevelSelectBox.style.display = "none";
    bestStudentTypeSelectBox.style.display = "none";
    yosSelectBox.style.display = "none";
    genderSelectBox.style.display = "none";
    awardClassificationSelectBox.style.display = "none";
    gpaMeasurementSelectBox.style.display = "none";
    gpaInputBox.style.display = "none";
  }
}

function getMeasurementType() {
  let measurementTypeID = document.getElementById("gpaMeasurementTypeID").value;
  let gpaInputBox = document.getElementById("gpaInputBox");
  let gpaValue = document.getElementById("gpaValue");

  gpaValue.removeAttribute("required");

  if (measurementTypeID !== "") {
    gpaInputBox.style.display = "block";
    gpaValue.setAttribute("required", "true");
    gpaValue.setAttribute("step", "0.01");
  } else {
    gpaInputBox.style.display = "none";
    gpaValue.removeAttribute("required");
    gpaValue.value = "";
  }
}

function resetAllFields(except) {
  const fields = [
    "reportClassID",
    "reportModifiedClassID",
    "reportClassName",
    "reportProgramID",
    "reportModifiedProgramID",
    "reportProgramName",
    "reportDepartmentID",
    "reportModifiedDepartmentID",
    "reportDepartmentName",
    "reportCampusID",
    "reportModifiedCampusID",
    "reportCampusName",
  ];

  fields.forEach((field) => {
    if (!field.includes(except)) {
      document.getElementById(field).value = "";
    }
  });
}

function getNTALevelData() {
  let [id = 0, name = ""] = document
    .getElementById("reportNTALevelID")
    .value.trim()
    .split(",");
  document.getElementById("reportModifiedNTALevelID").value = id;
  document.getElementById("reportNTALevelName").value = name;
}

function getClassID() {
  let [id = 0, name = ""] = document
    .getElementById("reportClassID")
    .value.trim()
    .split(",");
  document.getElementById("reportModifiedClassID").value = id;
  document.getElementById("reportClassName").value = name;
  resetAllFields("Class");
}

function getProgramData() {
  let [id = 0, name = ""] = document
    .getElementById("reportProgramID")
    .value.trim()
    .split(",");
  document.getElementById("reportModifiedProgramID").value = id;
  document.getElementById("reportProgramName").value = name;
  resetAllFields("Program");
}

function getDepartmentData() {
  let [id = 0, name = ""] = document
    .getElementById("reportDepartmentID")
    .value.trim()
    .split(",");
  document.getElementById("reportModifiedDepartmentID").value = id;
  document.getElementById("reportDepartmentName").value = name;
  resetAllFields("Department");
}

function getCampusData() {
  let [id = 0, name = ""] = document
    .getElementById("reportCampusID")
    .value.trim()
    .split(",");
  document.getElementById("reportModifiedCampusID").value = id;
  document.getElementById("reportCampusName").value = name;
  resetAllFields("Campus");
}

function getReportBestStudentType() {
  let bestStudentReportType = document.getElementById(
    "reportBestStudentType"
  ).value;
  if (bestStudentReportType == "2") {
    document.getElementById("isReportOveral").value = true;
  } else {
    document.getElementById("isReportOveral").value = false;
  }

  if (bestStudentReportType == "3") {
    document.getElementById("isReportNotOveral").value = true;
  } else {
    document.getElementById("isReportNotOveral").value = false;
  }
}
