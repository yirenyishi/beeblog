{{define "nav"}}
<div class='nav-container'>
    <nav>
        <ul class="layui-nav layui-bg-molv" lay-filter="">
            <p>云悦</p>
            <li class="layui-nav-item">
                <a href="/blog/new"><i class="layui-icon">&#xe642;</i>写文章</a>
            </li>
            <li class="layui-nav-item">
                <a href="/">首页 </a>
            </li>
            <li class="layui-nav-item">
                <a href="/blogs">博客 </a>
            </li>
            <li class="layui-nav-item">
                <a href="/note" target="_blank">笔记</a>
            </li>
            <li class="layui-nav-item">
                <a href="/map">地图</a>
            </li>
            <li class="layui-nav-item">
                <a href="/us">福利</a>
            </li>
            <li class="layui-nav-item">
                <a href="/me/info">个人中心</span></a>
            </li>
            {{if .IsLogin }}
            <li class="layui-nav-item">
                <a href=""><img src="{{.HeadImg}}" class="layui-nav-img"></a>
                <dl class="layui-nav-child">
                    <dd><a href="/me/info">个人中心</a></dd>
                    <dd><a href="/note">我的笔记</a></dd>
                    <dd><a href="/logout">退出登录</a></dd>
                </dl>
            </li>
            {{else}}
            <li class="layui-nav-item">
                <a href="/login">登录 <span class="layui-badge-dot"></span></a>
            </li>
            <li class="layui-nav-item">
                <a href="/regist">注册</a>
            </li>
            {{end}}
        </ul>
    </nav>
</div>
<script>
    //注意：导航 依赖 element 模块，否则无法进行功能性操作
    layui.use('element', function(){
        var element = layui.element;

        //…
    });
</script>
{{end}}