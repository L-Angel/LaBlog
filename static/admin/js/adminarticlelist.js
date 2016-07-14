/**
 * Created by Rick on 16/6/19.
 */

$(document).ready(function () {
    $("body").delegate(".row-operate-rel","click",function () {
        $articleid=$(this).parent().parent().parent().children().eq(0).text();
        $text=$(this).text();
        $release=false;
        if($text==="发布"){
            $release=true;
        }else if($text==="撤销发布"){
            $release=false;
        }
        $.putJSON("/admin/articlelist",{
            release:$release,
            aid:$articleid
        },function (data,textStatus) {
            //alert(data.result);
            if(data.result==="true"){
                alert("success.");
                window.location.href=location.href;
            }else{
                alert("failure.");
            }
        })
       // $.postJSON("/admin/");
    });
    $("body").delegate(".row-operate-del","click",function () {
        $articleid=$(this).parent().parent().parent().children().eq(0).text();
       // alert("*"+$articleid+"*");
        $.deleteJSON("/admin/articlelist/"+$articleid,{
        },function (data,textStatus) {
            if(data.result==="true"){
                alert("Delete success.");
                window.location.href=location.href;
            }else{
                alert("Delete failure."+data.msg);
            }
        })

    });
});