/**
 * Created by Rick on 16/5/27.
 */


$(function () {
    var $clientWidth = $(window).width();
    if ($clientWidth < 700) {

        var $clientHeight = $(window).height();
        $(".left-nav").height($clientHeight );
        $(".container").height($clientHeight - 55);
        window.onresize = function () {
            var $clientHeight = $(window).height();
            $(".left-nav").height($clientHeight);
            $(".container").height($clientHeight - 55);
        };
    }else{
        var $clientHeight = $(window).height();
        $(".left-nav").height($clientHeight);
        $(".container").height($clientHeight);
        window.onresize = function () {
            var $clientHeight = $(window).height();
            $(".left-nav").height($clientHeight);
            $(".container").height($clientHeight);
        };
    }



});