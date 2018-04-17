/*
 * Copyright (c) 2018, 奶爸<1@5.nu>
 * All rights reserved.
 */

package readability

import (
	"testing"
)

func TestParse(t *testing.T) {
	var htmlStr string
	{
	}
	htmlStr = `
<!DOCTYPE HTML>
<!--suppress ALL -->
<html class="no-js bg" lang="zh-cmn-Hans">
<head>
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1" />
    <meta charset="UTF-8">
    <!--IE 8浏览器的页面渲染方式-->
    <meta http-equiv="X-UA-Compatible" content="IE=edge, chrome=1">
    <!--默认使用极速内核：针对国内浏览器产商-->
    <meta name="renderer" content="webkit">
    <!--chrome Android 地址栏颜色-->
        <meta name="theme-color" content="#3a3f51" />
    
    <meta http-equiv="x-dns-prefetch-control" content="on">

    <title>Golang区别于bytes.Buffer的io.Pipe - 终身学习</title>
            <link rel="icon" type="image/ico" href="/usr/uploads/fav.ico">
        <meta name="description" content="因为终身学习博客博主的 GoCD 开源持续交付平台的缘故，碰到了一个问题怎么实时获取 SSH 的执行结果并展示给用户？方案一 session.Run终身学习博客博主反复的看了几遍 golang ..." />
<meta name="keywords" content="golang" />
<meta name="generator" content="Typecho 1.1/17.10.30" />
<meta name="template" content="handsome" />
<link rel="pingback" href="https://www.lifelonglearning.cc/action/xmlrpc" />
<link rel="EditURI" type="application/rsd+xml" title="RSD" href="https://www.lifelonglearning.cc/action/xmlrpc?rsd" />
<link rel="wlwmanifest" type="application/wlwmanifest+xml" href="https://www.lifelonglearning.cc/action/xmlrpc?wlw" />
<link rel="alternate" type="application/rss+xml" title="Golang区别于bytes.Buffer的io.Pipe &raquo; 终身学习 &raquo; RSS 2.0" href="https://www.lifelonglearning.cc/feed/p100_bytes-buffer-and-io-pipe.html" />
<link rel="alternate" type="application/rdf+xml" title="Golang区别于bytes.Buffer的io.Pipe &raquo; 终身学习 &raquo; RSS 1.0" href="https://www.lifelonglearning.cc/feed/rss/p100_bytes-buffer-and-io-pipe.html" />
<link rel="alternate" type="application/atom+xml" title="Golang区别于bytes.Buffer的io.Pipe &raquo; 终身学习 &raquo; ATOM 1.0" href="https://www.lifelonglearning.cc/feed/atom/p100_bytes-buffer-and-io-pipe.html" />

        <!-- 第三方CDN加载CSS -->
        <link href="https://cdn.bootcss.com/bootstrap/3.3.4/css/bootstrap.min.css" rel="stylesheet">


    <!-- 本地css静态资源 -->
    <link rel="stylesheet" href="https://www.lifelonglearning.cc/usr/themes/handsome/assets/css/function.min.css?v=4.2.12018022601" type="text/css" />
    <link rel="stylesheet" href="https://www.lifelonglearning.cc/usr/themes/handsome/assets/css/handsome.min.css?v=4.2.12018022601" type="text/css" />

    <!--主题组件css文件加载-->
    <link rel="stylesheet" href="https://www.lifelonglearning.cc/usr/themes/handsome/assets/css/features/jquery.fancybox.min.css?v=4.2.12018022601" type="text/css" />


    <!--引入英文字体文件-->
        <link rel="stylesheet" href="https://www.lifelonglearning.cc/usr/themes/handsome/assets/css/font.css?v=4.2.12018022601" type="text/css" />
    
    <style type="text/css">
        
        html.bg{
                   background-image:
               -moz-radial-gradient(-20% 140%, ellipse ,  rgba(235,167,171,.6) 30%,rgba(255,255,227,0) 50%),
               -moz-radial-gradient(60% 40%,ellipse,   #d9e3e5 10%,rgba(44,70,76,.0) 60%),
               -moz-linear-gradient(-45deg,  rgba(62,70,92,.8) -10%,rgba(220,230,200,.8) 80% )
               ;
           background-image:
               -o-radial-gradient(-20% 140%, ellipse ,  rgba(235,167,171,.6) 30%,rgba(255,255,227,0) 50%),
               -o-radial-gradient(60% 40%,ellipse,   #d9e3e5 10%,rgba(44,70,76,.0) 60%),
               -o-linear-gradient(-45deg,  rgba(62,70,92,.8) -10%,rgba(220,230,200,.8) 80% )
               ;
           background-image:
               -ms-radial-gradient(-20% 140%, ellipse ,  rgba(235,167,171,.6) 30%,rgba(255,255,227,0) 50%),
               -ms-radial-gradient(60% 40%,ellipse,   #d9e3e5 10%,rgba(44,70,76,.0) 60%),
               -ms-linear-gradient(-45deg,  rgba(62,70,92,.8) -10%,rgba(220,230,200,.8) 80% )
               ;
           background-image:
               -webkit-radial-gradient(-20% 140%, ellipse ,  rgba(235,167,171,.6) 30%,rgba(255,255,227,0) 50%),
               -webkit-radial-gradient(60% 40%,ellipse,   #d9e3e5 10%,rgba(44,70,76,.0) 60%),
               -webkit-linear-gradient(-45deg,  rgba(62,70,92,.8) -10%,rgba(220,230,200,.8) 80% )
               ;
        }
            </style>

    <!--全站jquery-->
    <script src="https://cdn.bootcss.com/jquery/2.1.4/jquery.min.js"></script>

    <!--网站统计代码-->
    <script type="text/javascript">
    (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
})(window,document,'script','https://www.google-analytics.com/analytics.js','ga');

ga('create', 'UA-111315498-1', 'auto');
ga('send', 'pageview');    </script>

</head>

<body id="body">
    <div id="alllayout" class="app app-aside-fixed container app-header-fixed ">    <!-- header -->
        <header id="header" class="app-header navbar" role="menu">
      <!-- navbar header（交集处） -->
        <div class="navbar-header bg-info dker">
        <button class="pull-right visible-xs dk" ui-toggle-class="show animated animated-lento fadeIn" target=".navbar-collapse">
          <i class="fa fa-gear text-lg"></i>
        </button>
        <button class="pull-right visible-xs" ui-toggle-class="off-screen animated" target=".app-aside" ui-scroll="app">
          <i class="fa fa-menu text-lg"></i>
        </button>
        <!-- brand -->
        <a href="https://www.lifelonglearning.cc/" class="navbar-brand text-lt">
          <i class="iconfont icon-bokeyuan"></i>
          <span class="hidden-folded m-l-xs">终身学习</span>
        </a>
        <!-- / brand -->
      </div>
      <!-- / navbar header -->

      <!-- navbar collapse（顶部导航栏） -->
    <div class="collapse pos-rlt navbar-collapse box-shadow bg-info dker">        <!-- buttons -->
        <div class="nav navbar-nav hidden-xs">
          <a href="#" class="btn no-shadow navbar-btn" ui-toggle-class="app-aside-folded" target=".app">
            <i class="fa fa-dedent text icon-fw"></i>
            <i class="fa fa-indent icon-fw text-active"></i>
          </a>
          <a href="#" class="btn no-shadow navbar-btn" ui-toggle-class="show" target="#aside-user">
            <i class="iconfont icon-user icon-fw"></i>
          </a>
        </div>
        <!-- / buttons -->


        <!-- search form -->
        <form id="searchform" class="navbar-form navbar-form-sm navbar-left shift" method="post" role="search">
          <div class="form-group">
            <div class="input-group rounded bg-light">
              <input id="search_input" type="search" name="s" class="transparent rounded form-control input-sm no-border padder" required placeholder="输入关键词搜索…">
              <span id="search_submit" class="transparent input-group-btn">
                  <button  type="submit" class="transparent btn btn-sm"><i class="fa fa-search"></i></button>
              </span>
            </div>
          </div>
        </form>
          <a href="" style="display: none" id="searchUrl"></a>
        <!-- / search form -->
                <ul class="nav navbar-nav navbar-right">
                        <li class="music-box hidden-xs hidden-sm">
                <div id="skPlayer"></div>
            </li>
            <li class="dropdown "><a class="skPlayer-list-switch dropdown-toggle"><i class="fa fa-headphones"></i><span class="visible-xs-inline"></span></a></li>
                                                          <!--登录管理-->
          <li class="dropdown" id="easyLogin">
            <a onclick="return false" data-toggle="dropdown" class="dropdown-toggle clear" data-toggle="dropdown">
                          <span class="thumb-sm avatar pull-right m-t-n-sm m-b-n-sm m-l-sm">
                <img src="https://cdn.v2ex.com/gravatar/fd1a4203245e430e56fb71977db714d4?s=40&r=G&d=">
                <i class="on md b-white bottom"></i>
              </span>
              <span class="hidden-sm hidden-md">奶爸</span>
                          <b class="caret"></b><!--下三角符号-->
            </a>
            <!-- dropdown(已经登录) -->
                        <ul class="dropdown-menu animated fadeInRight" id="Logged-in">
              <li class="wrapper b-b m-b-sm bg-light m-t-n-xs">
                <div>
                                                <p>晚上好，奶爸.</p>
                              </div>
                <div class="progress progress-xs m-b-none dker">
                  <div class="progress-bar progress-bar-info" data-toggle="tooltip" data-original-title="时间已经度过75.00%" style="width: 75.00%"></div>
                </div>
              </li>
              <!--文章RSS订阅-->
              <li>
                <a target="_blank" href="https://www.lifelonglearning.cc/feed/">
                  <i style="position: relative;width: 30px;margin: -11px -10px;margin-right: 0px;overflow: hidden;line-height: 30px;text-align: center;" class="fa fa-rss"></i><span>文章RSS</span>
                </a>
              </li>
              <!--评论RSS订阅-->
              <li>
                <a target="_blank" href="https://www.lifelonglearning.cc/feed/comments/"><i style="position: relative;width: 30px;margin: -11px -10px;margin-right: 0px;overflow: hidden;line-height: 30px;text-align: center;" class="fa fa-rss-square"></i><span>评论RSS</span></a>
              </li>
              <!--后台管理(登录时候才会显示)-->
                            <li>
                <a target="_blank" href="https://www.lifelonglearning.cc/admin/"><i style="position: relative;width: 30px;margin: -11px -10px;margin-right: 0px;overflow: hidden;line-height: 30px;text-align: center;" class="fa fa-cogs"></i><span>后台管理</span></a>
              </li>
              
              <li class="divider"></li>
              <li>
                <a id="sign_out" no-pjax href="https://www.lifelonglearning.cc/action/logout">退出</a>
              </li>
            </ul>
            <!-- / dropdown(已经登录) -->
                    </li>
          <!--/登录管理-->
                    </ul>
      </div>
      <!-- / navbar collapse -->
  </header>
    <!-- / header -->
	<!-- aside -->
	    <!--选择侧边栏的颜色-->
  <aside id="aside" class="app-aside hidden-xs bg-light dker b-r">  <!--<aside>-->
      <div class="aside-wrap">
        <div class="navi-wrap">
          <!-- user -->
          <div class="clearfix hidden-xs text-center hide  show" id="aside-user">
            <div class="dropdown wrapper">
                          <a href="https://www.lifelonglearning.cc">
                            <span class="thumb-lg w-auto-folded avatar m-t-sm">
                  <img src="/usr/uploads/avatar.jpg" class="img-full">
                </span>
              </a>
              <a href="#" data-toggle="dropdown" class="dropdown-toggle hidden-folded">
                <span class="clear">
                  <span class="block m-t-sm">
                    <strong class="font-bold text-lt">奶爸</strong>
                    <b class="caret"></b>
                  </span>
                  <span class="text-muted text-xs block">Lifelong Learner</span>
                </span>
              </a>
              <!-- dropdown -->
              <ul class="dropdown-menu animated fadeInRight w hidden-folded">
                <li class="wrapper b-b m-b-sm bg-info m-t-n-xs">
                  <span class="arrow top hidden-folded arrow-info"></span>
                  <div>
                                                <p>晚上好，注意早点休息</p>
                                  </div>
                  <div class="progress progress-xs m-b-none dker">
                    <div class="progress-bar bg-white" data-toggle="tooltip" data-original-title="时间已经度过75.00%" style="width: 75.00%"></div>
                  </div>
                </li>
              <!--文章RSS订阅-->
              <li>
                <a href="https://www.lifelonglearning.cc/feed/" data-toggle="tooltip" title="订阅文章 Feed 源">
                  <i style="position: relative;width: 30px;margin: -11px -10px;margin-right: 0px;overflow: hidden;line-height: 30px;text-align: center;" class="fa fa-rss" ></i><span>文章RSS</span>
                </a>
              </li>
              <!--评论RSS订阅-->
              <li>
                <a href="https://www.lifelonglearning.cc/feed/comments/" data-toggle="tooltip" title="订阅评论 Feed 源"><i style="position: relative;width: 30px;margin: -11px -10px;margin-right: 0px;overflow: hidden;line-height: 30px;text-align: center;" class="fa fa-rss-square" ></i><span>评论RSS</span></a>
              </li>
              <li class="divider"></li>
                            <!--后台管理-->
              <li>
                <a target="_blank" href="https://www.lifelonglearning.cc/admin/"><i style="position: relative;width: 30px;margin: -11px -10px;margin-right: 0px;overflow: hidden;line-height: 30px;text-align: center;" class="fa fa-cogs"></i><span>后台管理</span></a>
              </li>
                            </ul>
              <!-- / dropdown -->
            </div>
            <div class="line dk hidden-folded"></div>
          </div>
          <!-- / user -->

          <!-- nav -->
          <nav ui-nav class="navi clearfix">
            <ul class="nav">
             <!--index-->
              <li class="hidden-folded padder m-t m-b-sm text-muted text-xs">
                <span>导航</span>
              </li>
                                          <!--主页-->
              <li>
                <a href="https://www.lifelonglearning.cc/" class="auto">
                  <i class="iconfont icon-zhuye icon text-md"></i>
                  <span>首页</span>
                </a>
              </li>
              <!-- /主页 -->
                                          <li class="line dk"></li>
			<!--Components-->
              <li class="hidden-folded padder m-t m-b-sm text-muted text-xs">
                <span>组成</span>
              </li>
              <!--分类category-->
              <li>
                <a class="auto">
                  <span class="pull-right text-muted">
                    <i class="fa icon-fw fa-angle-right text"></i>
                    <i class="fa icon-fw fa-angle-down text-active"></i>
                  </span>
                  <i class="glyphicon glyphicon-th"></i>
                  <span>分类</span>
                </a>
                <ul class="nav nav-sub dk">
                  <li class="nav-sub-header">
                    <a data-no-instant>
                      <span>分类</span>
                    </a>
                  </li><!--不会显示出来-->
                  <!--循环输出分类-->
                <li><a href="https://www.lifelonglearning.cc/c1_develop/"><span>Develop</span></a></li><li><a href="https://www.lifelonglearning.cc/c2_life/"><span>Life</span></a></li><li><a href="https://www.lifelonglearning.cc/c3_machine_learning/"><span>Machine Learning</span></a></li><li><a href="https://www.lifelonglearning.cc/c21_devops/"><span>DevOps</span></a></li><li><a href="https://www.lifelonglearning.cc/c55_manual/"><span>Manual</span></a></li>                </ul>
              </li>
              <!--独立页面pages-->
              <li>
                <a class="auto">
                  <span class="pull-right text-muted">
                    <i class="fa icon-fw fa-angle-right text"></i>
                    <i class="fa icon-fw fa-angle-down text-active"></i>
                  </span>
                  <i class="glyphicon glyphicon-file"></i>
                  <span>页面</span>
                </a>
                <ul class="nav nav-sub dk">
                  <li class="nav-sub-header">
                    <a data-no-instant>
                      <span>页面</span>
                    </a>
                  </li><!--这个字段不会被显示出来-->
                  <!--循环输出独立页面-->
                                                         <li><a href="https://www.lifelonglearning.cc/archive.html"><span>归档</span></a></li>
                                       <li><a href="https://www.lifelonglearning.cc/guestbook.html"><span>留言</span></a></li>
                                       <li><a href="https://www.lifelonglearning.cc/nav.html"><span>导航</span></a></li>
                                       <li><a href="https://www.lifelonglearning.cc/about.html"><span>关于</span></a></li>
                                   </ul>
              </li>
              <!--友情链接-->
              <li>
                <a class="auto">
                  <span class="pull-right text-muted">
                    <i class="fa icon-fw fa-angle-right text"></i>
                    <i class="fa icon-fw fa-angle-down text-active"></i>
                  </span>
                  <i class="iconfont icon-links"></i>
                  <span>友链</span>
                </a>
                <ul class="nav nav-sub dk">
                  <li class="nav-sub-header">
                    <a data-no-instant>
                      <span>友链</span>
                    </a>
                  </li>
                  <!--使用links插件，输出全站友链-->
                                 </ul>
              </li>

            </ul>
          </nav>
          <!-- nav -->
        </div><!--.navi-wrap-->
      </div><!--.aside-wrap-->
  </aside>
<!-- content -->
<div id="content" class="app-content">
	<!-- / aside -->

<!-- <div id="content" class="app-content"> -->
        <div id="loadingbar" class="butterbar hide">
            <span class="bar"></span>
        </div>
   <a class="off-screen-toggle hide"></a>
   <main class="app-content-body ">
    <div class="hbox hbox-auto-xs hbox-auto-sm">

	<!-- 解析测试 -->
	<div> foo <br> bar <br> <br> abc<br> </div>
	<div> <font style="xxx">Hello</font> </div>
	<noscript>test</noscript>
	<style>.xxx{}</style>
	<!-- 解析测试 -->

    <!--文章-->
     <div class="col center-part">
    <!--标题下的一排功能信息图标：作者/时间/浏览次数/评论数/分类-->
      
        <header id="small_widgets" class="bg-light lter b-b wrapper-md">
             <h1 class="entry-title m-n font-thin h3 text-black l-h">Golang区别于bytes.Buffer的io.Pipe
                     <a class="superscript" href="https://www.lifelonglearning.cc/admin/write-post.php?cid=100" target="_blank"><i class="fa fa-edit" aria-hidden="true"></i></a>
                     </h1>       <!--文章标题下面的小部件-->
       <ul class="entry-meta text-muted list-inline m-b-none small post-head-icon">
       <!--作者-->
        <li class="meta-author"><i class="fa fa-user" aria-hidden="true"></i><span class="sr-only">博主：</span> <a class="meta-value" href="https://www.lifelonglearning.cc/author/1/" rel="author"> 奶爸</a></li>
        <!--发布时间-->
        <li class="meta-date"><i class="fa fa-clock-o" aria-hidden="true"></i>&nbsp;<span class="sr-only">发布时间：</span><time class="meta-value">2018 年 04 月 09 日</time></li>
        <!--浏览数-->
        <li class="meta-views"><i class="fa fa-eye" aria-hidden="true"></i>&nbsp;<span class="meta-value">17&nbsp;次浏览</span></li>
                   <!--评论数-->
        <li class="meta-comments"><i class="iconfont icon-comments-o" aria-hidden="true"></i>&nbsp;<a
                    class="meta-value" href="#comments">&nbsp;暂无评论</a></li>
                   <!--分类-->
        <li class="meta-categories"><i class="fa fa-tags" aria-hidden="true"></i> <span class="sr-only">分类：</span> <span class="meta-value"><a href="https://www.lifelonglearning.cc/c1_develop/">Develop</a></span></li>
       </ul>
      </header>
      <div class="wrapper-md" id="post-panel">
	   <ol class="breadcrumb bg-white b-a" itemscope=""><li>
                 <a href="https://www.lifelonglearning.cc/" itemprop="breadcrumb" title="返回首页" data-toggle="tooltip"><i class="fa fa-home" aria-hidden="true"></i>&nbsp;首页</a>
             </li><li class="active">正文&nbsp;&nbsp;</li>
              <div style="float:right;">
   分享到：
   <style>
   i.iconfont.icon-qzone:after {
    padding: 0 0px 0 5px;
    color: #ccc;
    content: "/\00a0";
    }
   </style>
   <a href="http://sns.qzone.qq.com/cgi-bin/qzshare/cgi_qzshare_onekey?url=https://www.lifelonglearning.cc/p100_bytes-buffer-and-io-pipe.html&title=Golang区别于bytes.Buffer的io.Pipe&site=https://www.lifelonglearning.cc/" itemprop="breadcrumb" target="_blank" title="" data-toggle="tooltip" data-original-title="分享到QQ空间" onclick="window.open(this.href, 'qzone-share', 'width=550,height=335');return false;"><i style ="font-size:15px;" class="iconfont icon-qzone" aria-hidden="true"></i></a>
   <a href="http://service.weibo.com/share/share.php?url=https://www.lifelonglearning.cc/p100_bytes-buffer-and-io-pipe.html&title=Golang区别于bytes.Buffer的io.Pipe" target="_blank" itemprop="breadcrumb" title="" data-toggle="tooltip" data-original-title="分享到微博" onclick="window.open(this.href, 'weibo-share', 'width=550,height=335');return false;"><i style ="font-size:15px;" class="fa fa-weibo" aria-hidden="true"></i></a>
  </div></ol>       <!--博客文章样式 begin with .blog-post-->
       <div id="postpage" class="blog-post">
        <article class="panel">
        <!--文章页面的头图-->
        <div class="entry-thumbnail" aria-hidden="true"><div class="item-thumb" style="background-image: url(https://www.lifelonglearning.cc/usr/themes/handsome/usr/img/sj/4.jpg)"></div></div>         <!--文章内容-->
         <div id="post-content" class="wrapper-lg">
          <div class="entry-content l-h-2x">
          <p>因为终身学习博客博主的 GoCD 开源持续交付平台的缘故，碰到了一个问题</p><blockquote><p>怎么实时获取 SSH 的执行结果并展示给用户？</p></blockquote><h3>方案一 session.Run</h3><p>终身学习博客博主反复的看了几遍 golang 的 ssh 包和 io.Writer 的相关实现。开始使用的方案是将 <code>ssh.Session.Stdout</code> 重定向至 <code>bytes.Buffer</code>，使用 <code>ssh.Run</code> 来远程执行脚本，然后使用 buffer 的 <code>String</code> 方法获取脚本的执行结果。<br>这个方案有一个比较大的弊端：<strong>无法实时看到运行中发生的事情，可控性比较差</strong>。而且我也用过一些其他的 CI、CD 工具，基本上日志都是可以即时查看的，所以我改变了方案。</p><h3>方案二 session.Start + bytes.Buffer</h3><p>这一次使用 <code>session.Start</code> 配合 <code>session.Wait</code> 执行脚本并同步不断读取远程输出。测试过程中发现命令执行完很久了 <code>bytes.Buffer</code> 中的数据还是没有完全 get 到，是缓存问题。<br>所以我费劲找了一个无缓存的 Writer：<code>io.Pipe</code>。</p><h3>方案三 io.Pipe</h3><p>换成 <code>io.Pipe</code> 之后远程输入可以即时读取了，我也将实时输出展示到了 Web 管理界面。</p>          </div>
                          <!--文章的页脚部件：打赏和其他信息的输出-->
             
             <div class="show-foot">
                 <div class="notebook">
                     <i class="fa fa-clock-o"></i>
                     <span>最后修改：2018 年 04 月 09 日 01 : 06  AM</span>
                 </div>
                 <div class="copyright" data-toggle="tooltip" data-html="true" data-original-title="转载请联系作者获得授权，并注明转载地址"><span>© 著作权归作者所有</span>
                 </div>
             </div>
                                  <!--/文章的页脚部件：打赏和其他信息的输出-->
         </div>
        </article>
       </div>
       <!--上一篇&下一篇-->
       <nav class="m-t-lg m-b-lg">
        <ul class="pager">
        <li class="next"> <a href="https://www.lifelonglearning.cc/p98_golang-http-request-not-redirect.html" title="Golang 获取 HTTP 请求的 redirect URL" data-toggle="tooltip"> 下一篇 </a></li>           </ul>
       </nav>
       <!--评论-->
        
            <style>
            textarea#comment{
                background-image: url('https://ww4.sinaimg.cn/large/a15b4afegy1fkjxbstyknj20vy0b0my0');
                background-size: contain;
                background-repeat: no-repeat;
                background-position: right;
                background-color: #fafdff;
            }
        </style>
    
    
    <div id="comments">
        
                    <!--评论列表-->
                    
        <!--如果允许评论，会出现评论框和个人信息的填写-->
                    <div id="respond-post-100" class="respond comment-respond">

                <h4 id="reply-title" class="comment-reply-title m-t-lg m-b">发表评论                    <small class="cancel-comment-reply">
                        <a id="cancel-comment-reply-link" href="https://www.lifelonglearning.cc/p100_bytes-buffer-and-io-pipe.html#respond-post-100" rel="nofollow" style="display:none" onclick="return TypechoComment.cancelReply();">取消回复</a>                    </small>
                </h4>
                <form id="comment_form" method="post" action="https://www.lifelonglearning.cc/p100_bytes-buffer-and-io-pipe.html/comment"  class="comment-form" role="form">
                    <div class="comment-form-comment form-group">
                        <label for="comment">评论                            <span class="required text-danger">*</span></label>
                        <textarea id="comment" class="textarea form-control OwO-textarea" name="text" rows="5" placeholder="说点什么吧……" onkeydown="if(event.ctrlKey&&event.keyCode==13){document.getElementById('submit').click();return false};"></textarea>
                        <div class="OwO"></div>
                    </div>
                    <!--判断是否登录-->
                                            <p id="welcomeInfo">欢迎&nbsp;<a data-no-intant target="blank" href="https://www.lifelonglearning.cc/admin/profile.php">奶爸</a>&nbsp;归来！&nbsp;<a href="https://www.lifelonglearning.cc/action/logout" id="logoutIn" title="Logout">退出&raquo;</a></p>
                                            <!--提交按钮-->
                        <div class="form-group">
                            <button type="submit" name="submit" id="submit" class="submit btn btn-success padder-lg">
                                <span class="text">发表评论</span>
                                <span class="text-active">提交中...</span>
                            </button>
                            <i class="animate-spin fa fa-spinner hide" id="spin"></i>
                            <input type="hidden" name="comment_post_ID" value="448" id="comment_post_ID">
                            <input type="hidden" name="comment_parent" id="comment_parent" value="0">
                        </div>
                </form>
            </div>
        
            </div>


      </div>
     </div>
     <!--文章右侧边栏开始-->
           <aside class="asideBar col w-md bg-white-only b-l bg-auto no-border-xs" role="complementary">
     <div id="sidebar">
      <section id="tabs-4" class="widget widget_tabs clear">
       <div class="nav-tabs-alt no-js-hide">
        <ul class="nav nav-tabs nav-justified" role="tablist">
         <li class="active" role="presentation"> <a href="#widget-tabs-4-hots" role="tab" aria-controls="widget-tabs-4-hots" aria-expanded="true" data-toggle="tab"> <i class="glyphicon glyphicon-fire text-md text-muted wrapper-sm" aria-hidden="true"></i> <span class="sr-only">热门文章</span> </a></li>
                     <li role="presentation"> <a href="#widget-tabs-4-comments" role="tab" aria-controls="widget-tabs-4-comments" aria-expanded="false" data-toggle="tab"> <i class="glyphicon glyphicon-comment text-md text-muted wrapper-sm" aria-hidden="true"></i> <span class="sr-only">最新评论</span> </a></li>
                     <li role="presentation"> <a href="#widget-tabs-4-random" role="tab" aria-controls="widget-tabs-4-random" aria-expanded="false" data-toggle="tab"> <i class="glyphicon glyphicon-transfer text-md text-muted wrapper-sm" aria-hidden="true"></i> <span class="sr-only">随机文章</span> </a></li>
        </ul>
       </div>
       <div class="tab-content">
       <!--热门文章-->
        <div id="widget-tabs-4-hots" class="tab-pane wrapper-md active" role="tabpanel">
         <h3 class="widget-title m-t-none text-md">热门文章</h3>
         <ul class="list-group no-bg no-borders pull-in m-b-none">
          <li class="list-group-item">
                <a href="https://www.lifelonglearning.cc/p82_interview.html" class="pull-left thumb-sm m-r"><img style="height: 40px!important;width: 40px!important;" src="https://www.lifelonglearning.cc/usr/themes/handsome/usr/img/sj2/14.jpg" class="img-circle"></a>
                <div class="clear">
                    <h4 class="h5 l-h"> <a href="https://www.lifelonglearning.cc/p82_interview.html" title="野生（初级）程序员面试经"> 野生（初级）程序员面试经 </a></h4>
                    <small class="text-muted post-head-icon>
                    <span class="meta-views"> <i class="iconfont icon-comments-o" aria-hidden="true"></i> <span class="sr-only">评论数：</span> <span class="meta-value">3</span>
                    </span>
                    <span class="meta-date m-l-sm"> <i class="fa fa-eye" aria-hidden="true"></i> <span class="sr-only">浏览次数:</span> <span class="meta-value">226</span>
                    </span>
                    </small>
                    </div>
            </li><li class="list-group-item">
                <a href="https://www.lifelonglearning.cc/p39_morvanpython.html" class="pull-left thumb-sm m-r"><img style="height: 40px!important;width: 40px!important;" src="https://www.lifelonglearning.cc/usr/themes/handsome/usr/img/sj2/4.jpg" class="img-circle"></a>
                <div class="clear">
                    <h4 class="h5 l-h"> <a href="https://www.lifelonglearning.cc/p39_morvanpython.html" title="莫烦Python:机器学习入门教程网站"> 莫烦Python:机器学习入门教程网站 </a></h4>
                    <small class="text-muted post-head-icon>
                    <span class="meta-views"> <i class="iconfont icon-comments-o" aria-hidden="true"></i> <span class="sr-only">评论数：</span> <span class="meta-value">1</span>
                    </span>
                    <span class="meta-date m-l-sm"> <i class="fa fa-eye" aria-hidden="true"></i> <span class="sr-only">浏览次数:</span> <span class="meta-value">158</span>
                    </span>
                    </small>
                    </div>
            </li><li class="list-group-item">
                <a href="https://www.lifelonglearning.cc/p12_typecho.html" class="pull-left thumb-sm m-r"><img style="height: 40px!important;width: 40px!important;" src="https://www.lifelonglearning.cc/usr/themes/handsome/usr/img/sj2/10.jpg" class="img-circle"></a>
                <div class="clear">
                    <h4 class="h5 l-h"> <a href="https://www.lifelonglearning.cc/p12_typecho.html" title="Typecho优化指南"> Typecho优化指南 </a></h4>
                    <small class="text-muted post-head-icon>
                    <span class="meta-views"> <i class="iconfont icon-comments-o" aria-hidden="true"></i> <span class="sr-only">评论数：</span> <span class="meta-value">0</span>
                    </span>
                    <span class="meta-date m-l-sm"> <i class="fa fa-eye" aria-hidden="true"></i> <span class="sr-only">浏览次数:</span> <span class="meta-value">125</span>
                    </span>
                    </small>
                    </div>
            </li><li class="list-group-item">
                <a href="https://www.lifelonglearning.cc/p57_coffee-miner.html" class="pull-left thumb-sm m-r"><img style="height: 40px!important;width: 40px!important;" src="https://www.lifelonglearning.cc/usr/themes/handsome/usr/img/sj2/8.jpg" class="img-circle"></a>
                <div class="clear">
                    <h4 class="h5 l-h"> <a href="https://www.lifelonglearning.cc/p57_coffee-miner.html" title="劫持WiFi网络接入设备挖矿指南"> 劫持WiFi网络接入设备挖矿指南 </a></h4>
                    <small class="text-muted post-head-icon>
                    <span class="meta-views"> <i class="iconfont icon-comments-o" aria-hidden="true"></i> <span class="sr-only">评论数：</span> <span class="meta-value">0</span>
                    </span>
                    <span class="meta-date m-l-sm"> <i class="fa fa-eye" aria-hidden="true"></i> <span class="sr-only">浏览次数:</span> <span class="meta-value">205</span>
                    </span>
                    </small>
                    </div>
            </li><li class="list-group-item">
                <a href="https://www.lifelonglearning.cc/p47_gogs-pre-receive.html" class="pull-left thumb-sm m-r"><img style="height: 40px!important;width: 40px!important;" src="https://www.lifelonglearning.cc/usr/themes/handsome/usr/img/sj2/13.jpg" class="img-circle"></a>
                <div class="clear">
                    <h4 class="h5 l-h"> <a href="https://www.lifelonglearning.cc/p47_gogs-pre-receive.html" title="Gogs迁移出错：remote hooks No such file or directory"> Gogs迁移出错：remote hooks No such file or directory </a></h4>
                    <small class="text-muted post-head-icon>
                    <span class="meta-views"> <i class="iconfont icon-comments-o" aria-hidden="true"></i> <span class="sr-only">评论数：</span> <span class="meta-value">0</span>
                    </span>
                    <span class="meta-date m-l-sm"> <i class="fa fa-eye" aria-hidden="true"></i> <span class="sr-only">浏览次数:</span> <span class="meta-value">111</span>
                    </span>
                    </small>
                    </div>
            </li>         </ul>
        </div>
                   <!--最新评论-->
        <div id="widget-tabs-4-comments" class="tab-pane wrapper-md no-js-show" role="tabpanel">
         <h3 class="widget-title m-t-none text-md">最新评论</h3>
         <ul class="list-group no-borders pull-in auto m-b-none">
                              <li class="list-group-item">

              <a href="https://www.lifelonglearning.cc/p82_interview.html#comment-15" class="pull-left thumb-sm avatar m-r">
                  <img nogallery src="https://q.qlogo.cn/g?b=qq&nk=1150009140&s=40" class="avatar-40 photo img-circle" style="height:40px!important; width: 40px!important;">              </a>
              <a href="https://www.lifelonglearning.cc/p82_interview.html#comment-15" class="text-muted">
                  <i class="fa fa-comment-o pull-right m-t-sm text-sm" title="详情" aria-hidden="true" data-toggle="tooltip" data-placement="auto left"></i>
                  <span class="sr-only">评论详情</span>
              </a>
              <div class="clear">
                  <div class="text-ellipsis">
                      <a href="https://www.lifelonglearning.cc/p82_interview.html#comment-15" title="你猜猜看"> 你猜猜看 </a>
                  </div>
                  <small class="text-muted">
                      <span>
                          ::aru:despise:: 呵                      </span>
                  </small>
              </div>
          </li>
                    <li class="list-group-item">

              <a href="https://www.lifelonglearning.cc/p82_interview.html#comment-13" class="pull-left thumb-sm avatar m-r">
                  <img nogallery src="https://q.qlogo.cn/g?b=qq&nk=1150009140&s=40" class="avatar-40 photo img-circle" style="height:40px!important; width: 40px!important;">              </a>
              <a href="https://www.lifelonglearning.cc/p82_interview.html#comment-13" class="text-muted">
                  <i class="fa fa-comment-o pull-right m-t-sm text-sm" title="详情" aria-hidden="true" data-toggle="tooltip" data-placement="auto left"></i>
                  <span class="sr-only">评论详情</span>
              </a>
              <div class="clear">
                  <div class="text-ellipsis">
                      <a href="https://www.lifelonglearning.cc/p82_interview.html#comment-13" title="你猜猜看"> 你猜猜看 </a>
                  </div>
                  <small class="text-muted">
                      <span>
                          说得对(☆ω☆)                      </span>
                  </small>
              </div>
          </li>
                    <li class="list-group-item">

              <a href="https://www.lifelonglearning.cc/p39_morvanpython.html#comment-12" class="pull-left thumb-sm avatar m-r">
                  <img nogallery src="https://cdn.v2ex.com/gravatar/9e543b9d68c191fdc484c3bbe9f953a4?s=40&r=G&d=" class="avatar-40 photo img-circle" style="height:40px!important; width: 40px!important;">              </a>
              <a href="https://www.lifelonglearning.cc/p39_morvanpython.html#comment-12" class="text-muted">
                  <i class="fa fa-comment-o pull-right m-t-sm text-sm" title="详情" aria-hidden="true" data-toggle="tooltip" data-placement="auto left"></i>
                  <span class="sr-only">评论详情</span>
              </a>
              <div class="clear">
                  <div class="text-ellipsis">
                      <a href="https://www.lifelonglearning.cc/p39_morvanpython.html#comment-12" title="友人C"> 友人C </a>
                  </div>
                  <small class="text-muted">
                      <span>
                          mark，也打算学下Python了                      </span>
                  </small>
              </div>
          </li>
                   </ul>
        </div>
                   <!--随机文章-->
        <div id="widget-tabs-4-random" class="tab-pane wrapper-md no-js-show" role="tabpanel">
            <h3 class="widget-title m-t-none text-md">随机文章</h3>
            <ul class="list-group no-bg no-borders pull-in">
            <li class="list-group-item">
                <a href="https://www.lifelonglearning.cc/p26_duclicity.html" class="pull-left thumb-sm m-r"><img style="height: 40px!important;width: 40px!important;" src="https://www.lifelonglearning.cc/usr/themes/handsome/usr/img/sj2/14.jpg" class="img-circle"></a>
                <div class="clear">
                    <h4 class="h5 l-h"> <a href="https://www.lifelonglearning.cc/p26_duclicity.html" title="DUCPLICITY 中文文档"> DUCPLICITY 中文文档 </a></h4>
                    <small class="text-muted post-head-icon">
                    <span class="meta-views"> <i class="iconfont icon-comments-o" aria-hidden="true"></i> <span class="sr-only">评论数：</span> <span class="meta-value">0</span>
                    </span>
                    <span class="meta-date m-l-sm"> <i class="fa fa-eye" aria-hidden="true"></i> <span class="sr-only">浏览次数:</span> <span class="meta-value">115</span>
                    </span>
                    </small>
                    </div>
            </li><li class="list-group-item">
                <a href="https://www.lifelonglearning.cc/p79_raspberrypi-nas-downloader.html" class="pull-left thumb-sm m-r"><img style="height: 40px!important;width: 40px!important;" src="https://www.lifelonglearning.cc/usr/themes/handsome/usr/img/sj2/4.jpg" class="img-circle"></a>
                <div class="clear">
                    <h4 class="h5 l-h"> <a href="https://www.lifelonglearning.cc/p79_raspberrypi-nas-downloader.html" title="树莓派 Aria2 搭建 NAS 远程下载"> 树莓派 Aria2 搭建 NAS 远程下载 </a></h4>
                    <small class="text-muted post-head-icon">
                    <span class="meta-views"> <i class="iconfont icon-comments-o" aria-hidden="true"></i> <span class="sr-only">评论数：</span> <span class="meta-value">0</span>
                    </span>
                    <span class="meta-date m-l-sm"> <i class="fa fa-eye" aria-hidden="true"></i> <span class="sr-only">浏览次数:</span> <span class="meta-value">84</span>
                    </span>
                    </small>
                    </div>
            </li><li class="list-group-item">
                <a href="https://www.lifelonglearning.cc/p39_morvanpython.html" class="pull-left thumb-sm m-r"><img style="height: 40px!important;width: 40px!important;" src="https://www.lifelonglearning.cc/usr/themes/handsome/usr/img/sj2/10.jpg" class="img-circle"></a>
                <div class="clear">
                    <h4 class="h5 l-h"> <a href="https://www.lifelonglearning.cc/p39_morvanpython.html" title="莫烦Python:机器学习入门教程网站"> 莫烦Python:机器学习入门教程网站 </a></h4>
                    <small class="text-muted post-head-icon">
                    <span class="meta-views"> <i class="iconfont icon-comments-o" aria-hidden="true"></i> <span class="sr-only">评论数：</span> <span class="meta-value">1</span>
                    </span>
                    <span class="meta-date m-l-sm"> <i class="fa fa-eye" aria-hidden="true"></i> <span class="sr-only">浏览次数:</span> <span class="meta-value">158</span>
                    </span>
                    </small>
                    </div>
            </li><li class="list-group-item">
                <a href="https://www.lifelonglearning.cc/p64_git-syncing-a-fork.html" class="pull-left thumb-sm m-r"><img style="height: 40px!important;width: 40px!important;" src="https://www.lifelonglearning.cc/usr/themes/handsome/usr/img/sj2/8.jpg" class="img-circle"></a>
                <div class="clear">
                    <h4 class="h5 l-h"> <a href="https://www.lifelonglearning.cc/p64_git-syncing-a-fork.html" title="Git同步Fork的仓库"> Git同步Fork的仓库 </a></h4>
                    <small class="text-muted post-head-icon">
                    <span class="meta-views"> <i class="iconfont icon-comments-o" aria-hidden="true"></i> <span class="sr-only">评论数：</span> <span class="meta-value">0</span>
                    </span>
                    <span class="meta-date m-l-sm"> <i class="fa fa-eye" aria-hidden="true"></i> <span class="sr-only">浏览次数:</span> <span class="meta-value">73</span>
                    </span>
                    </small>
                    </div>
            </li><li class="list-group-item">
                <a href="https://www.lifelonglearning.cc/p50_webminer.html" class="pull-left thumb-sm m-r"><img style="height: 40px!important;width: 40px!important;" src="https://www.lifelonglearning.cc/usr/themes/handsome/usr/img/sj2/13.jpg" class="img-circle"></a>
                <div class="clear">
                    <h4 class="h5 l-h"> <a href="https://www.lifelonglearning.cc/p50_webminer.html" title="浏览器挖矿指南"> 浏览器挖矿指南 </a></h4>
                    <small class="text-muted post-head-icon">
                    <span class="meta-views"> <i class="iconfont icon-comments-o" aria-hidden="true"></i> <span class="sr-only">评论数：</span> <span class="meta-value">0</span>
                    </span>
                    <span class="meta-date m-l-sm"> <i class="fa fa-eye" aria-hidden="true"></i> <span class="sr-only">浏览次数:</span> <span class="meta-value">238</span>
                    </span>
                    </small>
                    </div>
            </li>            </ul>
        </div>
       </div>
      </section>
      <!--循环输出分类列表-->
      <section id="categories-2" class="widget widget_categories wrapper-md clear">
       <h3 class="widget-title m-t-none text-md">分类</h3>
       <ul class="list-group">
        <li class="list-group-item"> <a href="https://www.lifelonglearning.cc/c1_develop/"> <span class="badge pull-right">12</span>Develop</a></li><li class="list-group-item"> <a href="https://www.lifelonglearning.cc/c2_life/"> <span class="badge pull-right">4</span>Life</a></li><li class="list-group-item"> <a href="https://www.lifelonglearning.cc/c3_machine_learning/"> <span class="badge pull-right">1</span>Machine Learning</a></li><li class="list-group-item"> <a href="https://www.lifelonglearning.cc/c21_devops/"> <span class="badge pull-right">6</span>DevOps</a></li><li class="list-group-item"> <a href="https://www.lifelonglearning.cc/c55_manual/"> <span class="badge pull-right">1</span>Manual</a></li>       </ul>
      </section>
                  <!--在文章页面输出目录，在其他页面输出标签云-->
                              <section id="tag_toc" class="widget widget_categories wrapper-md clear">
                  <h3 class="widget-title m-t-none text-md">文章目录</h3>
                  <div class="tags l-h-2x">
                      <div id="toc"></div>
                  </div>
              </section>
                  </div>
     </aside>
     <!--文章右侧边栏结束-->
    </div>
   </main>

    <!-- footer -->
	<script type="text/javascript">
(function () {
    window.TypechoComment = {
        dom : function (id) {
            return document.getElementById(id);
        },
    
        create : function (tag, attr) {
            var el = document.createElement(tag);
        
            for (var key in attr) {
                el.setAttribute(key, attr[key]);
            }
        
            return el;
        },

        reply : function (cid, coid) {
            var comment = this.dom(cid), parent = comment.parentNode,
                response = this.dom('respond-post-100'), input = this.dom('comment-parent'),
                form = 'form' == response.tagName ? response : response.getElementsByTagName('form')[0],
                textarea = response.getElementsByTagName('textarea')[0];

            if (null == input) {
                input = this.create('input', {
                    'type' : 'hidden',
                    'name' : 'parent',
                    'id'   : 'comment-parent'
                });

                form.appendChild(input);
            }

            input.setAttribute('value', coid);

            if (null == this.dom('comment-form-place-holder')) {
                var holder = this.create('div', {
                    'id' : 'comment-form-place-holder'
                });

                response.parentNode.insertBefore(holder, response);
            }

            comment.appendChild(response);
            this.dom('cancel-comment-reply-link').style.display = '';

            if (null != textarea && 'text' == textarea.name) {
                textarea.focus();
            }

            return false;
        },

        cancelReply : function () {
            var response = this.dom('respond-post-100'),
            holder = this.dom('comment-form-place-holder'), input = this.dom('comment-parent');

            if (null != input) {
                input.parentNode.removeChild(input);
            }

            if (null == holder) {
                return true;
            }

            this.dom('cancel-comment-reply-link').style.display = 'none';
            holder.parentNode.insertBefore(response, holder);
            return false;
        }
    };
})();
</script>
<script type="text/javascript">
var registCommentEvent = function() {
    var event = document.addEventListener ? {
        add: 'addEventListener',
        focus: 'focus',
        load: 'DOMContentLoaded'
    } : {
        add: 'attachEvent',
        focus: 'onfocus',
        load: 'onload'
    };
    var r = document.getElementById('respond-post-100');
        
    if (null != r) {
        var forms = r.getElementsByTagName('form');
        if (forms.length > 0) {
            var f = forms[0], textarea = f.getElementsByTagName('textarea')[0], added = false;

            if (null != textarea && 'text' == textarea.name) {
                textarea[event.add](event.focus, function () {
                    if (!added) {
                        var input = document.createElement('input');
                        input.type = 'hidden';
                        input.name = '_';
                            input.value = (function () {
    var _S8T = '425'//'Hz'
+//'IX8'
'5'+'6a4'//'K0m'
+'2'//'op'
+'5c'//'C'
+'e96'//'8A'
+'197'//'Uu7'
+//'ca'
'ca'+//'Y7'
'cd'+''///*'nC'*/'nC'
+'86'//'m'
+//'Y'
'2'+/* 'yAT'//'yAT' */''+/* 'Cs'//'Cs' */''+//'k4'
'1c6'+'79'//'i'
+//'ufh'
'f76'+//'zvR'
'a'+//'Ap'
'85', _4lfY1Q6 = [[16,18]];
    
    for (var i = 0; i < _4lfY1Q6.length; i ++) {
        _S8T = _S8T.substring(0, _4lfY1Q6[i][0]) + _S8T.substring(_4lfY1Q6[i][1]);
    }

    return _S8T;
})();
                    
                        f.appendChild(input);
                        
                        input = document.createElement('input');
                        input.type = 'hidden';
                        input.name = 'checkReferer';
                        input.value = 'false';
                        
                        f.appendChild(input);
                        

                        added = true;
                    }
                });
            }
        }
    }
};
</script></div>
<!-- /content -->
  <footer id="footer" class="app-footer" role="footer">
    <div class="wrapper b-t bg-light">
      <span class="pull-right hidden-xs">
            Powered by <a target="blank" href="http://www.typecho.org">Typecho</a>&nbsp;|&nbsp;Theme by <a target="blank" href="https://www.ihewro.com/archives/489/">handsome</a>
      </span>
      &copy;&nbsp;2018 Copyright&nbsp;奶爸    </div>
      <!--希望保留版权信息，谢谢-->
            <script type="text/template" id="tmpl-customizer">
          <div class="settings panel panel-default setting_body_panel" aria-hidden="true">
              <button class="btn btn-default no-shadow pos-abt" data-toggle="tooltip" data-placement="left" data-original-title="设置" data-toggle-class=".settings=active, .settings-icon=animate-spin">
                  <i class="fa fa-gear settings-icon"></i>
              </button>
              <div class="panel-heading">
                  <button class="pull-right btn btn-xs btn-rounded btn-danger" name="reset" data-toggle="tooltip" data-placement="top" data-original-title="恢复默认值" >重置</button>
                  设置		</div>
              <div class="setting_body">
              <div class="panel-body">
                  <# for ( var keys = _.keys( data.sections.settings ), i = 0, name; keys.length > i; ++i ) { #>
                      <div<# if ( i !== ( keys.length - 1 ) ) print( ' class="m-b-sm"' ); #>>
                          <label class="i-switch bg-info pull-right">
                              <input type="checkbox" name="{{ keys[i] }}" value="1"<# if ( data.defaults[keys[i]] ) print( ' checked="checked"' ); #> />
                                  <i></i>
                          </label>
                          {{ data.sections.settings[keys[i]] }}
              </div>
              <# } #>
          </div>
          <div class="wrapper b-t b-light bg-light lter r-b">
              <div class="row row-sm">
                  <div class="col-xs-6">
                      <#
                              _.each( data.sections.colors, function( color, i ) {
                              var newColumnBefore = ( i % 7 ) === 6;
                              #>
                          <label class="i-checks block<# if ( !newColumnBefore ) print( ' m-b-sm' ); #>">
                              <input type="radio" name="color" value="{{ i }}"<# if ( data.defaults['color'] === i ) print( ' checked="checked"' ); #> />
                                  <span class="block bg-light clearfix pos-rlt">
								<span class="active pos-abt w-full h-full bg-black-opacity text-center">
									<i class="fa fa-check text-md text-white m-t-xs"></i>
								</span>
								<b class="{{ color.navbarHeader }} header"></b>
								<b class="{{ color.navbarCollapse }} header"></b>
								<b class="{{ color.aside.replace( ' b-r', '' ) }}"></b>
							</span>
                          </label>
                          <#
                                  if ( newColumnBefore && ( i + 1 ) < data.sections.colors.length )
                          print( '</div><div class="col-xs-6">' );
                              } );
                              #>
                          </div>
                  </div>
              </div>
          </div>
          </div>
      </script>
      
      <div class="topButton panel panel-default">
          <button id="goToTop" class="btn btn-default no-shadow pos-abt hide">
              <i class="fa fa-chevron-circle-up" aria-hidden="true"></i>
          </button>
      </div>
  </footer>
  </div><!--end of .app app-header-fixed-->


    <!--定义全局变量-->
    <script type="text/javascript">
        window['LocalConst'] = {
            COMMENT_NAME_INFO: '必须填写昵称或姓名',
            COMMENT_EMAIL_INFO: '必须填写电子邮箱地址',
            COMMENT_EMAIL_LEGAL_INFO: '邮箱地址不合法',
            COMMENT_CONTENT_INFO: '必须填写评论内容',
            COMMENT_SUBMIT_ERROR: '提交失败，请重试！',
            COMMENT_CONTENT_LEGAL_INFO: '提交失败,您的输入内容不符合规则！',

            LOGIN_USERNAME_INFO: '必须填写用户名',
            LOGIN_PASSWORD_INFO: '请填写密码',
            LOGIN_SUBMIT_ERROR: '登录失败，请重新登录',
            LOGIN_SUBMIT_INFO: '用户名或者密码错误，请重试',
            LOGIN_SUBMIT_SUCCESS: '登录成功',
            LOGOUT_SUCCESS_REFRESH: '退出成功，正在刷新当前页面',

            LOGOUT_ERROR: '退出失败，请重试',
            LOGOUT_SUCCESS: '退出成功',

            SUBMIT_PASSWORD_INFO: '密码错误，请重试',
            COMMENT_TITLE: '评论通知',
            LOGIN_TITLE: '登录通知',
            ChANGYAN_APP_KEY: '',
            CHANGYAN_CONF: '',

            COMMENT_SYSTEM: '0',
            COMMENT_SYSTEM_ROOT: '0',
            COMMENT_SYSTEM_CHANGYAN: '1',
            COMMENT_SYSTEM_OTHERS: '2',
            EMOJI: '表情',
            IS_PJAX: '1',
            IS_PAJX_COMMENT: '1',
            BASE_SCRIPT_URL: 'https://www.lifelonglearning.cc/usr/themes/handsome/',
            THEME_COLOR: '9',
            THEME_HEADER_FIX: '1',
            THEME_ASIDE_FIX: '1',
            THEME_ASIDE_FOLDED: '',
            THEME_ASIDE_DOCK: '',
            THEME_CONTAINER_BOX: '1',
            THEME_HIGHLIGHT_CODE: '1',
            THEME_TOC: '1',
            TOC_TITLE: '文章目录',
            HEADER_FIX: '固定头部',
            ASIDE_FIX: '固定导航',
            ASIDE_FOLDED: '折叠导航',
            ASIDE_DOCK: '置顶导航',
            CONTAINER_BOX: '盒子模型',
            OFF_SCROLL_HEIGHT: '50',
            COMMENT_REJECT_PLACEHOLDER: '居然什么也不说，哼',
            COMMENT_PLACEHOLDER: '说点什么吧……',
            SHOW_SETTING_BUTTON: '1',
            THEME_VERSION: '4.2.12018022601'
        };

        window.paceOptions = {
            document: false,
            ajax: false,
            restartOnRequestAfter: false,
        }
    </script>



<!--CDN加载-->
<script src="https://cdn.bootcss.com/bootstrap/3.3.4/js/bootstrap.min.js"></script>


    <script src="https://www.lifelonglearning.cc/usr/themes/handsome/assets/js/features/jquery.pjax.min.js" type="text/javascript"></script>
    <script>
        $(document).pjax('a[href^="https://www.lifelonglearning.cc/"]:not(a[target="_blank"], a[no-pjax])', {
            container: '#content',
            fragment: '#content',
            timeout: 8000
        }).on('pjax:click', function() {
            window['Page'].doPJAXClickAction();
                        $("#loadingbar").attr("class","butterbar active");
            

                        $('body,html').animate({scrollTop:0},);
            
        }).on('pjax:complete', function() {
            window['Page'].doPJAXCompleteAction();

                        $("#loadingbar").attr("class","butterbar hide");
            
                                })
    </script>


<!--主题组件js加载-->

<!--pjax动画组件-->

<!--maxJax公式组件-->

<!--平滑滚动组件-->
<!--lightgallery必备组件-->
<script src="https://www.lifelonglearning.cc/usr/themes/handsome/assets/js/features/jquery.fancybox.min.js?v=4.2.12018022601"></script>

<!--component/comments.php 页面必需js（只有选择了原生评论的时候才会加载）-->
    <script src="https://www.lifelonglearning.cc/usr/themes/handsome/assets/js/features/OwO.min.js?v=4.2.12018022601"></script>
<!--component/comments.php 必需js结束-->

<!--全局播放器组件-->
<script src="https://www.lifelonglearning.cc/usr/themes/handsome/assets/js/music.min.js?v=4.2.12018022601"></script>
<script>
        var player = new skPlayer({"autoplay":false,"listshow":false,"mode":"listloop","music":{"type":"cloud","source":"1424171551","media":"tencent"}});
    //player.play();
</script>

<!--主题组件js加载结束-->

<!--主题核心js-->
    <script src="https://www.lifelonglearning.cc/usr/themes/handsome/assets/js/function.min.js?v=4.2.12018022601"></script>
    <script src="https://www.lifelonglearning.cc/usr/themes/handsome/assets/js/core.min.js?v=4.2.12018022601"></script>

    <script type="text/javascript">
            </script>

</body><!--#body end-->
</html><!--html end-->
  	<!-- / footer -->
`
	Parse(htmlStr, Option{})
}
