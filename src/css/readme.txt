width: #block的宽度
height: #block的高度
background-color: #背景颜色

#文本相关
font-size: #文字大小
font-family: #字体
color:    #文本颜色 不是font-color
background: #fa2f2f; #背景颜色
text-align: center; #文字位置(居左，居中，居右)
text-indent: 5em; #文字缩进
更多：http://www.w3school.com.cn/css/css_text.asp

#块相关
padding: 上右下左
border:
margin: 上右下左

#定位
position: relative|absolute 相对|绝对定位
top:      顶部偏移
left:     左部偏移
float:    right|left #浮动位置
display: none(不显示) block(块级元素) inline(行内元素) inline-block(行内块) #显示方式 
display: http://www.w3school.com.cn/cssref/pr_class_display.asp
vertical-align: top; #两个div顶部对齐使用该方案

#div相关特效
border-radius: 3px; #边角是圆角矩形，矩形半径是3px

overflow-x: 横向的块展示溢出处理 visible, hidden, scroll等
应用于块级元素，替换元素，表单元格 (参考：http://www.cnblogs.com/xiaohuochai/p/5289653.html)

overflow-y:纵向的块展示溢出处理


[js-css网址导航]

css的几个关键概念：
display： http://www.w3school.com.cn/cssref/pr_class_display.asp
white-space：http://www.w3school.com.cn/cssref/pr_text_white-space.asp
float：http://www.w3school.com.cn/cssref/pr_class_float.asp
css速查：http://www.w3school.com.cn/cssref/index.asp
隐藏：http://blog.sina.com.cn/s/blog_693d183d0100p4fk.html
disabled和readonly：http://blog.csdn.net/ligang2585116/article/details/44921967
返回上一页：http://blog.sina.com.cn/s/blog_54eeb5d90100cls3.html
字符串操作：http://www.cnblogs.com/xuebin/articles/1296837.html
jquery速查：http://www.php100.com/manual/jquery/

[js学习笔记]
字符串：
a="abcdabcd"
获取长度：a.length
获取某个字符：a[0]（获取0号元素）
循环获取字符：
for (var i=0;i<a.length;i++) {
console.log(a[i]);
}
for (i in a) {console.log(a[i]);}


数组：
获取数组长度：var a=[1,2,3,4]; a.length
判断key是否存在：if (a[key] == null) {...}
删除key：delete a.key
新增key: a.key=value
修改key: a.key = value
遍历数组：for (i in a) {console.log(a[i]);}

1，引入外部文件
<script type="text/javascript" src="public.js"></script>

2，获取各个类型的object
获取id类型：$('#id')
获取class类型：$('.class')
获取input-name元素：$('input[name="mobile"]')

3，url请求
$.post(url, data, function(e){})

4，dom元素获取，修改
$(selector).attr('src', 'src');
$(selector).val([value]);

html获取和修改
$(selector).html();
$(selector).html('new html');

5，input框失去焦点
$(selector).change([data], fn);
input框切换到其它地方：$(selector).blur();

6，图片相关
<img src="data:image/png;base64,这里放字符"/>

7 JSON相关
object=JSON.parse(string); 字符串转数组

##TODO
#复杂的语法
.class .class1{} //表示同时有两种类选择器时 执行此css样式 顺序不区分先后
ul li:last-child{} //ul li标签的最后一个元素执行此css样式 有先后顺序
.active-li, .active-li:hover //active-li的类执行此css样式；同时active-li的hover动作也执行此动作


#前端代码片段
//$("#pub_result").html(result.data);
//$("#pub_result").show();
//document.getElementById('pub_result').scrollIntoView();


=====================================================
模态窗口：bootstrap自带的，效果非常棒
相关链接：http://www.ziqiangxuetang.com/bootstrap/bootstrap-modal-plugin.html

<div id="myABC"></div>
<!-- 模态框（Modal） -->
<div class="modal fade" id="myModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close"
                        data-dismiss="modal" aria-hidden="true">
                    &times;
                </button>
                <h4 class="modal-title" id="myModalLabel">
                    发布结果
                </h4>
            </div>
            <div class="modal-body" id="modal_body">
            </div>
            <div class="modal-footer">
                <button type="button" class="btn btn-primary"
                        data-dismiss="modal">关闭
                </button>
            </div>
        </div><!-- /.modal-content -->
    </div><!-- /.modal -->
</div>
=====================================================