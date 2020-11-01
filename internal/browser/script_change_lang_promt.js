var firstLang = false;
document.querySelector("#mainBlock > div > div.row.setLang > div.form-group.col-xs-6.col-md-6.form-inline.source-sets > div:nth-child(4) > button").click();
var item = document.querySelector("#mainBlock > div > div.row.setLang > div.form-group.col-xs-6.col-md-6.form-inline.source-sets > div.btn-group.bootstrap-select.no-pop.open > div > ul");
if(!item) {
    document.querySelector("#mainBlock > div > div.row.setLang > div.form-group.col-xs-6.col-md-6.form-inline.source-sets > div:nth-child(4) > button").click();
    item = document.querySelector("#mainBlock > div > div.row.setLang > div.form-group.col-xs-6.col-md-6.form-inline.source-sets > div.btn-group.bootstrap-select.no-pop.open > div > ul");
}
for(var y = 0; y < 2; y++ ) {
    for ( var i = 0; i < item.childElementCount; i++) {
        var langName = item.childNodes[i].firstElementChild.firstElementChild.innerText;
        if (langName == LANG_1) {
            item.childNodes[i].firstElementChild.click();
            firstLang = true;
            break;
        }
    }
    if(!firstLang) {
        document.querySelector("#bRevLang").click();
    }else{
        break
    }
}
document.querySelector("#mainBlock > div > div.row.setLang > div.form-group.col-xs-6.col-md-6.form-inline.result-sets > div.btn-group.bootstrap-select.no-pop > button").click();
var item = document.querySelector("#mainBlock > div > div.row.setLang > div.form-group.col-xs-6.col-md-6.form-inline.result-sets > div.btn-group.bootstrap-select.no-pop.open > div > ul");
if(!item) {
    document.querySelector("#mainBlock > div > div.row.setLang > div.form-group.col-xs-6.col-md-6.form-inline.result-sets > div.btn-group.bootstrap-select.no-pop > button").click();
    item = document.querySelector("#mainBlock > div > div.row.setLang > div.form-group.col-xs-6.col-md-6.form-inline.result-sets > div.btn-group.bootstrap-select.no-pop.open > div > ul");
}
for ( var i = 0; i < item.childElementCount; i++) {
    var langName = item.childNodes[i].firstElementChild.firstElementChild.innerText;
    if (langName == LANG_2) {
        item.childNodes[i].firstElementChild.click();
        break;
    }
}
