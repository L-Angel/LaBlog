/**
 * Created by Rick on 16/5/27.
 */
$(function () {
    var show=true;
    $("body").delegate("header","click",function () {
        if(show===false){
            $(".left-nav").hide(100);
            $(".mobile-top-people").show(100);
            $(".mobile-top-info").show(100);
            show=true;
        }else{
            $(".left-nav").show(100);
            $(".mobile-top-people").hide(100);
            $(".mobile-top-info").hide(100);

            show=false;
        }
    })

});