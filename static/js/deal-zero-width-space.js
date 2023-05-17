(function(){
    "use strict";
    var articleEle = document.querySelector('div.td-content')
    var nodeList = articleEle.querySelectorAll('p')
    var reg = /^\u200b/i
    for(var node of nodeList) {
        if (reg.test(node.textContent)) {
            node.style.textIndent = "2em"
        }
    }
})();