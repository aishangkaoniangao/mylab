# introduce
这篇文章是一个关于mysql的demo。该demo仅展示了mysql的增删改查这些常用操作

# How to build
+ 1, go get -v github.com/ziutek/mymysql/...
+ 2, go build db.go

# Demo action
+ 1, insert
+ 2, update
+ 3, delete
+ 4, get

# env
+ 1, db：信息要修改成自己的mysql变量
+ 2, database：code中使用的库名是blog，要修改成自己的库名
+ 3，table：code中使用的表名是user

# table
```
CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `username` varchar(200) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(40) NOT NULL DEFAULT '' COMMENT '密码',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='用户表'
```