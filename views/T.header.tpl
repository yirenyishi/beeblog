{{define "header"}}
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
{{if .IsDDesc }}
        {{else}}
    <meta name="description" content="云悦,个人随笔是一个面向IT技术人员,提供个人平时工作总结和在线记录学习笔记,个人技术博客,在线云笔记,码农笔录,最新的技术博客,www.aiprose.com">
{{end}}
    <link rel="stylesheet" href="//g.alicdn.com/sui/sui3/0.0.18/css/sui.min.css">
    <script type="text/javascript" src="//g.alicdn.com/sj/lib/jquery/dist/jquery.min.js"></script>
{{/*<script type="text/javascript" src="//g.alicdn.com/sui/sui3/0.0.18/js/sui.min.js"></script>*/}}
    <script type="text/javascript" src="/static/js/sui.js"></script>
    <script type="text/javascript" src="/static/js/layer.js"></script>
    <link type="text/css" rel="styleSheet" href="/static/css/common.css"/>
{{end}}