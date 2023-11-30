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
				elementValue = element.getAttribute('aria-label') || element.textContent;
			}
			elementValues.push(elementValue);
		});
		elementValues;
`
