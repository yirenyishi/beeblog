{{define "header"}}
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
{{if .IsDDesc }}
        {{else}}
    <meta name="description" content="码农随笔,个人随笔是一个面向IT技术人员,提供个人平时工作总结和在线记录学习笔记,个人技术博客,在线云笔记,码农笔录,最新的技术博客,www.aiprose.com">
{{end}}
    <script type="text/javascript" src="/static/js/checkm.js"></script>
    <script type="text/javascript" src="https://cdn.bootcss.com/jquery/3.4.0/jquery.min.js"></script>
    <script type="text/javascript" src="/static/js/layer.js"></script>
    <link rel="stylesheet" href="/static/layui/css/layui.css">
    <script type="text/javascript" src="/static/layui/layui.js"></script>
<!--    <link type="text/css" rel="styleSheet" href="/static/css/common.min.css"/>-->
    <link type="text/css" rel="styleSheet" href="/static/css/common.css"/>
{{end}}