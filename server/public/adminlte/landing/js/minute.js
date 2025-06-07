$("#btnCreateMinute").on("click", function () {
    $("#minuteCreate").css("display", "block")
    $("#btnCreateMinute").css("display", "none")
    $("#btnCloseMinute").css("display", "none")
});

$("#btnCreateCancel").on("click", function () {
    $("#minuteCreate").css("display", "none")
    $("#btnCreateMinute").css("display", "block")
    $("#btnCloseMinute").css("display", "block")
});

$("#btnCloseMinute").on("click", function () {
    $("#minuteClose").css("display", "block")
    $("#btnCloseMinute").css("display", "none")
    $("#btnCreateMinute").css("display", "none")
});

$("#btnCloseCancel").on("click", function () {
    $("#minuteClose").css("display", "none")
    $("#btnCloseMinute").css("display", "block")
    $("#btnCreateMinute").css("display", "block")
});


$("#btnStats").on("click", function () {
    $("#dvStats").css("display", "block")
    $("#btnStats").css("display", "none")
    $("#btnCancel").css("display", "block")

});

$("#btnCancel").on("click", function () {
    $("#dvStats").css("display", "none")
    $("#btnStats").css("display", "block")
    $("#btnCancel").css("display", "none")

});


// Makes sure user writes closing remarks when closing minutes
const textarea = document.getElementById('minuteClosingRemarks');
const btnSubmitRemarks = document.getElementById('btnSubmitRemarks');

textarea.addEventListener('input', function () {
    if (textarea.value.trim() === '') {
        btnSubmitRemarks.disabled = true;
    } else {
        btnSubmitRemarks.disabled = false;
    }
});


// Makes sure user selects officer and write minute when replying to the minute
const minuteDesc = document.getElementById('minuteDesc');
const selectOfficer = document.getElementById('selectOfficer');
const btnSubmit = document.getElementById('btnSubmit');

minuteDesc.addEventListener('input', checkForm);
selectOfficer.addEventListener('change', checkForm);

function checkForm() {
    if (minuteDesc.value.trim() !== '' && selectOfficer.value !== '') {
        btnSubmit.disabled = false;
    } else {
        btnSubmit.disabled = true;
    }
}