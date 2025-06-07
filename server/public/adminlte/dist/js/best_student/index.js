(function () {
  $("#viewOveralBestStudentTab").html(
    '<div class="loadingio-spinner-double-ring-xymjggaeym"><div class="ldio-r2hchxjaheg"><div></div><div></div><div><div></div></div><div><div></div></div></div></div>'
  );
  let request = $.post("/soma/best-student/list", { is_overal: true });
  request.done(function (e) {
    $("#viewOveralBestStudentTab").html(e);
  });
  request.always(function (jqXHR) {});
})();

$('#nav-tab a[href="#nav-view-overal-best-student').on("click", function (e) {
  e.preventDefault();
  $("#viewOveralBestStudentTab").html(
    '<div class="loadingio-spinner-double-ring-xymjggaeym"><div class="ldio-r2hchxjaheg"><div></div><div></div><div><div></div></div><div><div></div></div></div></div>'
  );
  let request = $.post("/soma/best-student/list", { is_overal: true });
  request.done(function (e) {
    $("#viewOveralBestStudentTab").html(e);
  });
  request.always(function (jqXHR) {});
});

$('#nav-tab a[href="#nav-view-level-best-student').on("click", function (e) {
  e.preventDefault();
  $("#viewLevelBestStudentTab").html(
    '<div class="loadingio-spinner-double-ring-xymjggaeym"><div class="ldio-r2hchxjaheg"><div></div><div></div><div><div></div></div><div><div></div></div></div></div>'
  );
  let request = $.post("/soma/best-student/list", { is_not_overal: true });
  request.done(function (e) {
    $("#viewLevelBestStudentTab").html(e);
  });
  request.always(function (jqXHR) {});
});

function generateBestStudent() {
  let generateBox = document.getElementById("generateBox");
  generateBox.style.display = "block";
}

function validateClass() {
  let classData = document.getElementById("classID").value.trim();
  let awardTypeSelectBox = document.getElementById("awardTypeSelect");
  let awardTypeSelect = document.getElementById("awardTypeID");
  let classID = document.getElementById("modifiedClassID");
  let isEligibleToGraduate = document.getElementById("isEligibleToGraduate");
  let programTypeID = 0;

  if (classData !== "") {
    awardTypeSelectBox.style.display = "block";
    classID.value = classData.split(",")[0];
    programTypeID = classData.split(",")[2];
    isEligibleToGraduate.value = classData.split(",")[3];

    document.getElementById("displayUniqueNTALevel").style.display = "none";

    //reset award list type selection
    awardTypeSelect.selectedIndex = 0;
    $("#awardTypeID").trigger("change");
  } else {
    awardTypeSelectBox.style.display = "none";
    classID.value = 0;
  }

  if (programTypeID == 2) {
    let newOption = document.createElement("option");
    newOption.value = "3";
    newOption.text = "Overal Per Class Year of Study";

    let optionExists = Array.from(awardTypeSelect.options).find(
      (option) => option.value === "3"
    );
    if (!optionExists) {
      awardTypeSelect.add(newOption);
    }
  } else {
    let optionToRemove = Array.from(awardTypeSelect.options).find(
      (option) => option.value === "3"
    );
    if (optionToRemove) {
      awardTypeSelect.removeChild(optionToRemove);
    }
  }
}

function getBestStudentType() {
  let awardType = document.getElementById("awardTypeID");
  let classSelect = document.getElementById("classID");
  let classID = document.getElementById("modifiedClassID").value;
  let isEligible = document.getElementById("isEligibleToGraduate").value;

  let isOveral = document.getElementById("isOveral");
  let downloadButton = document.getElementById("downloadButton");
  let awardTypeSelect = document.getElementById("awardTypeSelect");
  let displayUniqueNTALevel = document.getElementById("displayUniqueNTALevel");

  // Reset all UI elements initially
  isOveral.value = "false";
  downloadButton.style.display = "none";
  displayUniqueNTALevel.style.display = "none";

  if (awardType.value === "1") {
    isOveral.value = "true";
    downloadButton.style.display = "block";

    if (isEligible === "false") {
      alert(
        "Sorry, this service is for graduating classes only. Please select a graduating class."
      );

      awardTypeSelect.style.display = "none";

      resetSelection(classSelect, awardType);
      return;
    }
  } else if (awardType.value === "2") {
    getUniqueClassNTALevel(classID);
    displayUniqueNTALevel.style.display = "block";
  } else if (awardType.value === "3") {
    getAllNTALevelYOS(classID);
    displayUniqueNTALevel.style.display = "block";
  }
}

// Function to reset class and awardType selections
function resetSelection(classSelect, awardType) {
  classSelect.selectedIndex = 0;
  awardType.selectedIndex = 0;

  $("#classID").trigger("change");
  $("#awardTypeID").trigger("change");
}

function getUniqueClassNTALevel(classID) {
  $("#class-unique-nta-level").html(
    '<div class="loadingio-spinner-double-ring-xymjggaeym"><div class="ldio-r2hchxjaheg"><div></div><div></div><div><div></div></div><div><div></div></div></div></div>'
  );
  let request = $.post("/soma/nta-level/unique-exist-by-class", {
    class_id: classID,
    csrf: "{{ .csrf }}",
  });
  request.done(function (e) {
    $("#class-unique-nta-level").html(e);
  });
  request.fail(function () {
    $("#class-unique-nta-level").html(
      '<p class="text-danger text-center">Failed to load data. Please try again later.</p>'
    );
  });
  request.always(function (jqXHR) {});
}

function getAllNTALevelYOS(classID) {
  $("#class-unique-nta-level").html(
    '<div class="loadingio-spinner-double-ring-xymjggaeym"><div class="ldio-r2hchxjaheg"><div></div><div></div><div><div></div></div><div><div></div></div></div></div>'
  );
  let request = $.post("/soma/nta-level/exist-by-class", {
    class_id: classID,
    csrf: "{{ .csrf }}",
  });
  request.done(function (e) {
    $("#class-unique-nta-level").html(e);
  });
  request.fail(function () {
    $("#class-unique-nta-level").html(
      '<p class="text-danger text-center">Failed to load data. Please try again later.</p>'
    );
  });
  request.always(function (jqXHR) {});
}
