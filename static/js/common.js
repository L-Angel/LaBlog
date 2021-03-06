/**
 * Created by Rick on 16/5/27.
 */


$(document).ready(function () {
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

    $.postJSON = function(url, args, callback) {
        var xsrf, xsrflist;
        xsrf = $.cookie("_xsrf");
        xsrflist = xsrf.split("|");
        args._xsrf = base64_decode(xsrflist[0]);
        $.ajax({url: url, data: $.param(args), dataType: "text", type: "POST",
            success: function(response) {
                callback(eval("(" + response + ")"));
            }});
    };

    $.deleteJSON = function(url, args, callback) {
        var xsrf, xsrflist;
        xsrf = $.cookie("_xsrf");
        xsrflist = xsrf.split("|");
        args._xsrf = base64_decode(xsrflist[0]);
       // alert(args._xsrf);
        $.ajax({url: url,
            beforeSend:function (request) {
                request.setRequestHeader("X-XSRFToken",args._xsrf)
            },
            dataType: "text", type: "DELETE",
            success: function(response) {
                callback(eval("(" + response + ")"));
            }});
    };
    $.putJSON = function(url, args, callback) {
        var xsrf, xsrflist;
        xsrf = $.cookie("_xsrf");
        xsrflist = xsrf.split("|");
        args._xsrf = base64_decode(xsrflist[0]);
       // alert(args._xsrf);
        $.ajax({url: url,
              data: $.param(args),
            dataType: "text", type: "PUT",
            success: function(response) {
                callback(eval("(" + response + ")"));
            }});
    };

});