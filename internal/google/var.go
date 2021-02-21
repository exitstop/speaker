package google

var JS_SET_TYPE string = `
var obj = document.querySelector("#yDmH0d > c-wiz > div > div.WFnNle > c-wiz > div.OlSOob > c-wiz > div.ccvoYb > div.AxqVh > div.OPPzxe > c-wiz.rm1UF.UnxENd > span > span > div > textarea");

obj.value = "%s";

var event = new Event('input', {
	'bubbles': true,
	'cancelable': true
});
obj.dispatchEvent(event);
`

var JS_SET_EVENT_GOOGLE string = `
console.log("JS_SET_EVENT_GOOGLE");
window.MyGlobalVar = 0;
window.GlobalText = "";

function locationHashChanged() {
	console.log("----004-----");
	window.MyGlobalVar = 1;
}

window.onhashchange = locationHashChanged;

var proxied = window.XMLHttpRequest.prototype.open;
window.XMLHttpRequest.prototype.open = function() {
	this.addEventListener('load', function() {
			console.log("JS_SET_EVENT_GOOGLE:001");
			if (this.responseText.indexOf("generic") >= 0) {
				console.log("JS_SET_EVENT_GOOGLE:002");
				if (this.responseText.indexOf("[[[") >= 0) {
					console.log("JS_SET_EVENT_GOOGLE:003");
					window.GlobalText = '' + this.responseText;
					console.log('load('+MyGlobalVar+'): ' + this.responseText);
					window.MyGlobalVar = 2;
				}
			}
	});
	return proxied.apply(this, arguments);
};
`

var JS_GET_TEXT_GOOGLE string = `window.GlobalText`

var JS_CLEAR_VAR string = `window.GlobalText = ""`
