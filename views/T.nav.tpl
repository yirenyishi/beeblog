{{define "nav"}}
<div class='nav-container'>
    <nav class="navbar navbar-default">
        <div class="container-fluid">
            <div class="navbar-header">
                <a class="navbar-brand" href="/"> 云悦 </a>
            </div>
            <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
                <ul class="nav navbar-nav">
                    <li {{if .IsHome}} class="active" {{end}}>
                        <a href="/">首页 </a>
                    </li>
                    <li {{if .IsBlog}} class="active" {{end}}>
                        <a href="/blogs">博客 </a>
                    </li>
                    <li>
                        <a href="/note" target="_blank">笔记</a>
                    </li>
                    <li {{if .IsMap}} class="active" {{end}}>
                        <a href="/map">地图</a>
                    </li>
                    <li {{if .IsUs}} class="active" {{end}}>
                        <a href="/us">福利</a>
                    </li>
                </ul>
            {{/*<form class="navbar-form navbar-left" role="search">*/}}
            {{/*<div class="form-group">*/}}
            {{/*<input type="text" class="form-control" placeholder="Search">*/}}
            {{/*</div>*/}}
            {{/*<button type="submit" class="btn btn-default">Submit</button>*/}}
            {{/*</form>*/}}
                <ul class="nav navbar-nav navbar-right">
                {{if .IsLogin }}
                    <li><a href="/me/blog" style="padding: 0;"><img src="{{.HeadImg}}" alt="头像" class="img-circle"></a>
                    </li>
                {{else}}
                    <li><a href="/login">登录</a></li>
                    <li><a href="/regist">注册</a></li>
                {{end}}
                {{if .IsLogin }}
                    <li class="dropdown">
                        <a href="/me/blog" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-expanded="false">
                        {{.NickName}}<span class="caret"></span></a>
                        <ul class="dropdown-menu" role="menu">
                            <li><a href="/me/blog">个人中心</a></li>
                            <li><a href="/note" target="_blank">我的笔记</a></li>
                            <li><a href="/u/{{.UserId}}">我的首页</a></li>
                            <li class="divider"></li>
                            <li><a href="/logout">安全退出</a></li>
                        </ul>
                    </li>
                {{end}}
                </ul>
            </div><!-- /.navbar-collapse -->
        </div><!-- /.container-fluid -->
    </nav>
</div>
{{end}}