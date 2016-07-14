/**
 * Created by Rick on 16/7/14.
 */
$(document).ready(function () {
    var uploader = Qiniu.uploader({
        runtimes: 'html5,flash,html4',
        browse_button: 'pickfiles',
        container: 'container',
        drop_element: 'container',
        max_file_size: '100mb',
        flash_swf_url: '../../js/common/Moxie.swf',
        dragdrop: true,
        chunk_size: '4mb',
        uptoken:'Vx9QXeqQsLR_-NmpiRV6lDmTUBcT-N20q48K1IM4:Y7eB64FnS730YQDiSIW5-4yiuxI=:eyJzY29wZSI6ImxhYmxvZyIsImRlYWRsaW5lIjoxNDY4NDc5MjM0fQ==',
        // uptoken_url: $('#uptoken_url').val(),  //当然建议这种通过url的方式获取token
        domain: "https://o9a877aer.bkt.clouddn.com",
        auto_start: false,
        init: {
            'FilesAdded': function(up, files) {

                plupload.each(files, function(file) {

                });
                $(".file-upload-tips").html("文件添加完毕.");
            },
            'BeforeUpload': function(up, file) {
                /*
                var progress = new FileProgress(file, 'fsUploadProgress');
                var chunk_size = plupload.parseSize(this.getOption('chunk_size'));
                if (up.runtime === 'html5' && chunk_size) {
                    progress.setChunkProgess(chunk_size);
                }
                */
            },
            'UploadProgress': function(up, file) {
                /*
                var progress = new FileProgress(file, 'fsUploadProgress');
                var chunk_size = plupload.parseSize(this.getOption('chunk_size'));

                progress.setProgress(file.percent + "%", file.speed, chunk_size);
                */
                $(".file-upload-tips").html(file.percent+"%",file.speed,'4mb');

            },
            'UploadComplete': function() {
               // $('#success').show();
                $(".file-upload-tips").html("文件上传完毕.");

            },
            'FileUploaded': function(up, file, info) {
              //  var progress = new FileProgress(file, 'fsUploadProgress');
               // progress.setComplete(up, info);
                var result=$.parseJSON(info);

                $(".file-upload-tips").html(up+result+"文件已经上传.");

            },
            'Error': function(up, err, errTip) {
                /*
                $('table').show();
                var progress = new FileProgress(err.file, 'fsUploadProgress');
                progress.setError();
                progress.setStatus(errTip);
                */
                $(".file-upload-tips").html(errTip);

            }
        }
    });

    uploader.bind('FileUploaded', function() {
        console.log('hello man,a file is uploaded');
    });

    $('#up_load').on('click', function(){
        uploader.start();
    });
    $('#stop_load').on('click', function(){
        uploader.stop();
    });


});
/*
//引入Plupload 、qiniu.js后
var uploader = new Qiniu.uploader({
    runtimes: 'html5,flash,html4',    //上传模式,依次退化
    browse_button: 'file-upload-btn',       //上传选择的点选按钮，**必需**
    uptoken_url: '/token',            //Ajax请求upToken的Url，**强烈建议设置**（服务端提供）
    // uptoken : '', //若未指定uptoken_url,则必须指定 uptoken ,uptoken由其他程序生成
    // unique_names: true, // 默认 false，key为文件名。若开启该选项，SDK为自动生成上传成功后的key（文件名）。
    // save_key: true,   // 默认 false。若在服务端生成uptoken的上传策略中指定了 `sava_key`，则开启，SDK会忽略对key的处理
    domain: 'https://portal.qiniu.com/bucket/lablog',   //bucket 域名，下载资源时用到，**必需**
    get_new_uptoken: false,  //设置上传文件的时候是否每次都重新获取新的token
    container: 'container',           //上传区域DOM ID，默认是browser_button的父元素，
    max_file_size: '100mb',           //最大文件体积限制
    flash_swf_url: '/static/js/common/Moxie.swf',  //引入flash,相对路径
    max_retries: 3,                   //上传失败最大重试次数
    dragdrop: false,                   //开启可拖曳上传
   // drop_element: 'container',        //拖曳上传区域元素的ID，拖曳文件或文件夹后可触发上传
    chunk_size: '4mb',                //分块上传时，每片的体积
    auto_start: true,                 //选择文件后自动上传，若关闭需要自己绑定事件触发上传
    init: {
        'FilesAdded': function(up, files) {
            plupload.each(files, function(file) {
                // 文件添加进队列后,处理相关的事情
            });
        },
        'BeforeUpload': function(up, file) {
            // 每个文件上传前,处理相关的事情
        },
        'UploadProgress': function(up, file) {
            // 每个文件上传时,处理相关的事情
        },
        'FileUploaded': function(up, file, info) {
            // 每个文件上传成功后,处理相关的事情
            // 其中 info 是文件上传成功后，服务端返回的json，形式如
            // {
            //    "hash": "Fh8xVqod2MQ1mocfI4S4KpRL6D98",
            //    "key": "gogopher.jpg"
            //  }
            // 参考http://developer.qiniu.com/docs/v6/api/overview/up/response/simple-response.html

            // var domain = up.getOption('domain');
            // var res = parseJSON(info);
            // var sourceLink = domain + res.key; 获取上传成功后的文件的Url
        },
        'Error': function(up, err, errTip) {
            //上传出错时,处理相关的事情
        },
        'UploadComplete': function() {
            //队列文件处理完毕后,处理相关的事情
        },
        'Key': function(up, file) {
            // 若想在前端对每个文件的key进行个性化处理，可以配置该函数
            // 该配置必须要在 unique_names: false , save_key: false 时才生效

            var key = "";
            // do something with key here
            return key
        }
    }
});

// domain 为七牛空间（bucket)对应的域名，选择某个空间后，可通过"空间设置->基本设置->域名设置"查看获取

// uploader 为一个plupload对象，继承了所有plupload的方法，参考http://plupload.com/docs
*/