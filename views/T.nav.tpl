{{define "nav"}}
<div class='nav-container'>
    <div class="sui-navbar">
        <div class="navbar-inner">
            <a href="/" class="sui-brand">个人随笔</a>
            <ul class="sui-nav">
                <li {{if .IsHome}}class="active"{{end}}><a href="/">首页</a></li>
                <li {{if .IsNote}}class="active"{{end}}><a href="/note">笔记</a></li>
                <li {{if .IsMap}}class="active"{{end}}><a href="/map">地图</a></li>
            </ul>
            <form class="sui-form sui-form pull-left">
                <input type="text" placeholder="宝贝/店铺名称...">
                <button class="sui-btn">搜索</button>
            </form>
            <ul class="sui-nav pull-right">
                <li><a href="#">个人中心</a></li>
                <li><a href="#">帮助</a></li>
            </ul>
        </div>
    </div>
</div>
{{end}}