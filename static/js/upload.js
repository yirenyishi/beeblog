function DragImgUpload(id,options) {
    this.me = $(id);
    var defaultOpt = {
        boxWidth:'120px',
        boxHeight:'auto'
    }
    this.preview = $('<div id="preview"><img src="/static/img/upload.png" class="img-responsive"  style="width: 100%;height: auto;" alt="" title=""> </div>');
    this.opts=$.extend(true, defaultOpt,{
    }, options);
    this.init();
    this.callback = this.opts.callback;
}

//定义原型方法
DragImgUpload.prototype = {
    init:function () {
        this.me.append(this.preview);
        this.me.append(this.fileupload);
        this.cssInit();
        this.eventClickInit();
    },
    cssInit:function () {
        this.me.css({
            'width':this.opts.boxWidth,
            'height':this.opts.boxHeight,
            'border':'1px solid #cccccc',
            'padding':'10px',
            'cursor':'pointer'
        })
        this.preview.css({
            'height':'100%',
            'overflow':'hidden'
        })

    },
    onDragover:function (e) {
        e.stopPropagation();
        e.preventDefault();
        e.dataTransfer.dropEffect = 'copy';
    },
    onDrop:function (e) {
        var self = this;
        e.stopPropagation();
        e.preventDefault();
        var fileList = e.dataTransfer.files; //获取文件对象
        // do something upload
        if(fileList.length == 0){
            return false;
        }
        //检测文件是不是图片
        if(fileList[0].type.indexOf('image') === -1){
            alert("您拖的不是图片！");
            return false;
        }

        //拖拉图片到浏览器，可以实现预览功能
        var img = window.URL.createObjectURL(fileList[0]);
        var filename = fileList[0].name; //图片名称
        var filesize = Math.floor((fileList[0].size)/1024);
        if(filesize>500){
            alert("上传大小不能超过500K.");
            return false;
        }

        self.me.find("img").attr("src",img);
        self.me.find("img").attr("title",filename);
        if(this.callback){
            this.callback(fileList);
        }
    },
    eventClickInit:function () {
        var self = this;
        this.me.unbind().click(function () {
            self.createImageUploadDialog();
        })
        var dp = this.me[0];
        dp.addEventListener('dragover', function(e) {
            self.onDragover(e);
        });
        dp.addEventListener("drop", function(e) {
            self.onDrop(e);
        });


    },
    onChangeUploadFile:function () {
        var fileInput = this.fileInput;
        var files = fileInput.files;
        var file = files[0];
        var img = window.URL.createObjectURL(file);
        var filename = file.name;
        this.me.find("img").attr("src",img);
        this.me.find("img").attr("title",filename);
        if(this.callback){
            this.callback(files);
        }
    },
    createImageUploadDialog:function () {
        var fileInput = this.fileInput;
        if (!fileInput) {
            //创建临时input元素
            fileInput = document.createElement('input');
            //设置input type为文件类型
            fileInput.type = 'file';
            //设置文件name
            fileInput.name = 'ime-images';
            //允许上传多个文件
            fileInput.multiple = true;
            fileInput.onchange  = this.onChangeUploadFile.bind(this);
            this.fileInput = fileInput;
        }
        //触发点击input点击事件，弹出选择文件对话框
        fileInput.click();
    }




}