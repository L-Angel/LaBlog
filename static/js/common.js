/**
 * Created by Rick on 16/5/27.
 */


$(function () {
    var $clientHeight = $(window).height();
    $(".left-nav").height($clientHeight);
    $(".container").height($clientHeight);
       window.onresize=function() {
        var $clientHeight = $(window).height();
        $(".left-nav").height($clientHeight);
        $(".container").height($clientHeight);
    };

});