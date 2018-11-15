{{define "header"}}
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="">
    <title>个人随笔</title>
    <link href="http://g.alicdn.com/sj/dpl/1.5.1/css/sui.min.css" rel="stylesheet">
    <script type="text/javascript" src="http://g.alicdn.com/sj/lib/jquery/dist/jquery.min.js"></script>
    <script type="text/javascript" src="http://g.alicdn.com/sj/dpl/1.5.1/js/sui.min.js"></script>
</head>
<body>
<div class='container'>
    <div class="sui-navbar">
        <div class="navbar-inner">
            <a href="/" class="sui-brand">个人随笔</a>
            <ul class="sui-nav">
                <li class="active"><a href="/">首页</a></li>
                <li><a href="#">笔记</a></li>
                <li><a href="#">地图</a></li>
            </ul>
            <form class="sui-form sui-form pull-right">
                <input type="text" placeholder="宝贝/店铺名称...">
                <button class="sui-btn">搜索</button>
            </form>
        </div>
    </div>
</div>
{{end}}