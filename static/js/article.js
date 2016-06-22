/**
 * Created by Rick on 16/6/21.
 */
$(document).ready(function () {
    /**
    * 分享到qq空间
    *
    * */
    $(".article-share-item:eq(0)").click(function () {
        var $articletitle=$(":header:eq(1)").html();
        var href="http://sns.qzone.qq.com/cgi-bin/qzshare/cgi_qzshare_onekey?url="
        +encodeURIComponent(location.href)
            +"&title="+encodeURIComponent($articletitle);
        location.href=href;

    });
    /**
     *
     * 分享到朋友圈
     *
     * */
    /*
    $(".article-share-item:eq(1)").click(function () {
        var $articletitle=$(":header:eq(1)").html();
        var $href=encodeURIComponent(location.href);
        wx.onMenuShareTimeline({
            title: $articletitle, // 分享标题
            link: $href, // 分享链接
            imgUrl: '', // 分享图标
            success: function () {
                // 用户确认分享后执行的回调函数
                alert(1);
            },
            cancel: function () {
                // 用户取消分享后执行的回调函数
            }
        });
        /*
        var href="http://sns.qzone.qq.com/cgi-bin/qzshare/cgi_qzshare_onekey?url="
            +encodeURIComponent(location.href)
            +"&title="+encodeURIComponent($articletitle);
        location.href=href;
        */
/*
    });
*/

});