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
    // 获取表单的提交连接、提交方法和提交参数
    var action = form.getAttribute('action');
    var enctype = form.getAttribute('enctype');
    var method = form.getAttribute('method') || 'GET'; // 默认为 GET 方法
    var formData = [];

    // 获取表单中的每个字段的名称和值
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
        if (type==="file") {
            oneData.enctype=enctype; 
        }
        formDataList.push(oneData);
    });

    // 将表单信息放入列表
    var formInfo = {
        action: action,
        method: method,
        formData: formDataList
    };

    formList.push(formInfo);
});
formList;

`
