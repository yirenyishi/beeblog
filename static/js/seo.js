$(function () {
    /*百度收录*/
    var bp = document.createElement('script');
    var curProtocol = window.location.protocol.split(':')[0];
    if (curProtocol === 'https') {
        bp.src = 'https://zz.bdstatic.com/linksubmit/push.js';
    }
    else {
        bp.src = 'http://push.zhanzhang.baidu.com/push.js';
    }
    var s = document.getElementsByTagName("script")[0];
    s.parentNode.insertBefore(bp, s);

    /*360收录*/
    var qihuo = document.createElement('script');
    if (curProtocol === 'https') {
        qihuo.src = 'https://jspassport.ssl.qhimg.com/11.0.1.js?d33b14200fa89b7ecd3780341cd234c';
    }
    else {
        qihuo.src = 'http://js.passport.qihucdn.com/11.0.1.js?d33b14200fa89b7ecd3780341cd234c4';
    }
    qihuo.id = "sozz"
    var s = document.getElementsByTagName("script")[1];
    s.parentNode.insertBefore(qihuo, s);
})