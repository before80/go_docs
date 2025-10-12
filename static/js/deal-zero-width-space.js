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

    function adjustTocHeight() {
        const toc = document.querySelector('.td-sidebar-toc');
        const mainContent = document.querySelector('main');
        const tdToc = toc.querySelector(".td-toc")
    
        if (toc && mainContent) {
            // const mainHeight = mainContent.scrollHeight;
            const viewportHeight = window.innerHeight;
            const tdTocHeight = tdToc.scrollHeight;
            console.log("tdTocHeight=",tdTocHeight)
            // 156为.td-toc元素上面一些元素的高度，12是padding-top的高度，24是padding-bottom的高度
            let tocH = tdTocHeight + 156 + 12 + 24 + 1;

            // 150为footer元素的高度,64是头部菜单的高度
            if (tocH  < viewportHeight - 150 - 64 ) {
                toc.style.maxHeight = tocH + 'px';
                toc.style.overflowY = "hidden"
            } else {
                toc.style.maxHeight = tocH + 10 + 'px';
                toc.style.overflowY = "auto"
            }
        }
    }

    // 在页面加载和调整大小时调用
    window.addEventListener('load', adjustTocHeight);
    window.addEventListener('resize', adjustTocHeight);
})();