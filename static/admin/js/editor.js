/**
 * Created by Rick on 16/6/15.
 */
$(document).ready(function () {
  $(".submit").click(function () {
      var $articletile=$(".la-article-title").val();
      var $categorytId=$(".a-article-type").val();
      var $articlecontent=$(".markdown-editor").val();
      var $author="";
      alert($articlecontent)
      $.postJSON("/admin/editor",{
          articleTitle:$articletile,
          articleContent:$articlecontent,
          author:$author
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
