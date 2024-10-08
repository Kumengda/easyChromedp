package jsCode

const GetAllOnclickUrl = `
var clickableElements = document.querySelectorAll('a, button, input[type="button"], input[type="submit"], input[type="image"], [role="button"], [role="link"], [onclick], [tabindex], [contenteditable="true"],[contenteditable=""], [contenteditable="inherit"]');
		var elementValues = [];
		clickableElements.forEach(function(element) {
			var elementValue = "";
			if (element.tagName.toLowerCase() === 'a') {
				elementValue = element.getAttribute('href');
			} else if (element.tagName.toLowerCase() === 'input' && (element.type === 'button' || element.type === 'submit')) {
				elementValue = element.value;
			} else if (element.tagName.toLowerCase() === 'input' && element.type === 'image') {
				elementValue = element.alt;
			} else if (element.tagName.toLowerCase() === 'button') {
				elementValue = element.textContent;
			} else {
				elementValue = element.getAttribute('aria-label');
			}
			if (elementValue!==null&&elementValue!==""){
				elementValues.push(elementValue);
			}
		});
		elementValues;
`
const ParseFrom = `
var formList = [];

var forms = document.querySelectorAll('form');

forms.forEach(function(form) {
    var action = form.getAttribute('action');
    var enctype = form.getAttribute('enctype');
    var method = form.getAttribute('method') || 'GET';
    var formData = [];

    var inputs = form.querySelectorAll('input, textarea, select');
    var formDataList= [];
    inputs.forEach(function(input) {
        var name = input.getAttribute('name');
        var type=input.getAttribute('type');
        var value = input.value;
        var oneData = {};
        oneData.name = name;
        oneData.value=value;
        oneData.type=type;
		oneData.enctype = enctype;
        formDataList.push(oneData);
    });

    var formInfo = {
        action: action,
        method: method,
        formData: formDataList
    };

    formList.push(formInfo);
});
formList;

`
