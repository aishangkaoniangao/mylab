# introduce
��ƪ������һ������mysql��demo����demo��չʾ��mysql����ɾ�Ĳ���Щ���ò���

# How to build
+ 1, go get -v github.com/ziutek/mymysql/...
+ 2, go build db.go

# Demo action
+ 1, insert
+ 2, update
+ 3, delete
+ 4, get

# env
+ 1, db����ϢҪ�޸ĳ��Լ���mysql����
+ 2, database��code��ʹ�õĿ�����blog��Ҫ�޸ĳ��Լ��Ŀ���
+ 3��table��code��ʹ�õı�����user

# table
```
CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '����id',
  `username` varchar(200) NOT NULL DEFAULT '' COMMENT '�û���',
  `password` varchar(40) NOT NULL DEFAULT '' COMMENT '����',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '����ʱ��',
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '�޸�ʱ��',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='�û���'
```