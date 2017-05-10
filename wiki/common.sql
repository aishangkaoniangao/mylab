#统计商品id pid出现次数最高的那两个pid和出现的次数
select count(1) as counter, pid from product_price group by pid order by counter desc limit 0,2;
