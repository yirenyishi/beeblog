{{define "memenu"}}
<div class="me-menu">
    <div style="height: 50px;line-height: 50px">
        <a href="/u/{{.User.Id}}" target="_blank"><img src="{{.HeadImg}}" alt="头像" class="img-circle"></a>
        <a href="/u/{{.User.Id}}" target="_blank" style="margin-left: 15px;font-size: 18px;text-decoration: none">{{.User.NickName}}</a>
    </div>
    <hr style="height:1px;border:none;border-top:1px solid #EEE;margin: 6px;"/>
    <div style="display: flex;height: 30px;line-height: 30px;">
        <div style="display: inline-block;width: 100px">
            <span>文章: </span>
            <span>{{.User.BlogCount}}</span>
        </div>
        <div style="display: inline-block;flex: 1">
            <span>访问: </span>
            <span>{{.User.BlogBrowes}}</span>
        </div>
    </div>
    <div style="display: flex;height: 30px;line-height: 30px;">
        <div style="display: inline-block;width: 100px">
            <span>评论: </span>
            <span>{{.User.BlogComment}}</span>
        </div>
        <div style="display: inline-block;flex: 1">
            <span>喜欢: </span>
            <span>{{.User.BlogLike}}</span>
        </div>
    </div>
    <hr style="height:1px;border:none;border-top:1px solid #EEE;margin: 6px;"/>
    <p {{if .IsMeBlog}}class="active"{{end}}><span></span><a href="/me/blog">我的博客</a></p>
    <p {{if .IsMeNote}}class="active"{{end}}><span></span><a href="/me/note">笔记文件夹</a></p>
    <p {{if .IsMeLike}}class="active"{{end}}><span></span><a href="/me/like">我的收藏</a></p>
    <p {{if .IsMeInfo}}class="active"{{end}}><span></span><a href="/me/info">我的资料</a></p>
</div>
{{end}}