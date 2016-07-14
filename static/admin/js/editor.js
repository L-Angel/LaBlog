/**
 * Created by Rick on 16/6/15.
 */
$(document).ready(function () {
  $(".submit").click(function () {
      var $articletitle=$(".la-article-title").val();
      var $categorytId=$(".la-article-type").val();
      var $articlecontent=$(".markdown-editor").val();

      var $summary=(String)($(".la-markdown-preview").html()).replace(" ","").replace("\\S|\r\n","").substr(0,200);
      //alert($summary);
      var $author="";
      $.postJSON("/admin/editor",{
          articleTitle:$articletitle,
          articleContent:$articlecontent,
          categoryId:$categorytId,
          author:$author,
          summary:$summary
      },function (data) {
          if(data.result=="true") {
              alert("Article add successed!");
              window.location.href=location.href;
          }else if(data.result=="false"){
              alert("Article add failure!");
          }else if(data.result=="error"){
              alert("Server Error!");
          }

      })
  })
});
