/**
 * Created by Rick on 16/7/3.
 */


$(document).ready(function () {
    $(".submit").click(function () {
        var $articletitle=$(".la-article-title").val();
        var $categorytId=$(".la-article-type").val();
        var $articlecontent=$(".markdown-editor").val();

        var $summary=(String)($(".la-markdown-preview").html()).replace(" ","").replace("\\S|\r\n","").substr(0,200);
        //alert($summary);
        var $author="";
        var $articleid=$(".hidden-article-id").html();
        var $uniqid=$(".hidden-article-uniqid").html();
        $.postJSON("/admin/modify",{
            articleTitle:$articletitle,
            articleContent:$articlecontent,
            categoryId:$categorytId,
            author:$author,
            summary:$summary,
            uniqid:$uniqid,
            articleid:$articleid
        },function (data) {
            if(data.result=="true") {
                alert("Article modify successed!");
                window.location.href=location.href;
            }else if(data.result=="false"){
                alert("Article modify failure!");
            }else if(data.result=="error"){
                alert("Server Error!");
            }

        })
    })
});

