#!/bin/bash

#part1: 基础操作

#1.1: if条件判断
a=0
b="0"
if [ $a -eq $b ] #eq ne lt le gt ge
then
    echo "$a=$b"
else
    echo "$a!=$b"
fi

#1.2: 比较操作符: http://blog.csdn.net/ithomer/article/details/6836382
if [ 100 = 0 ]
then
    echo "0=0"
fi

if [ -z "abc" ]
then
    echo "len is 0"
else
    echo "len is not 0"
fi


#1.3: for-while循环
for i in "1" "2" "3"
do
    echo $i;
    echo "";
done

for ((i=0;i<5;i++))
do
    echo $i
done
for ((i=0;i<5;i+=2))
do
    echo $i
done

i=0
while :
do
    i=`expr $i + 1`
    if [ $i = 100 ]
    then
        echo $i
        break
    fi
done
echo $i

#1.4: 数组下标
a=("a" "b" "c"  "d")

for((i=0;i<3;i++))
do
    echo ${a[i]}
done

#1.5: case操作
num=1
case $num in
    1) echo "num is 1";
       echo "num is 1, too";;
    2) echo "num is 2";;
esac

#part2: 字符串操作
X="abc"
echo '$X' #not interpret $X
echo "$X" #interpret $X to 0
echo ''
echo "\""
echo ${#X} #get string length
echo ${X:1:1} #substring: string,index,n substract string from index to index+n
echo ${X:0-1} #substring: 获取从倒数第N个字符开始到结尾的字符串 ${string:0-n}

#part3: function 函数操作
function compare()
{
    if [[ $1 -gt 100 ]]
    then
        echo "over"
        return
    fi
    echo $1
    echo $2
    echo "compare"
}
compare 101 2
#函数定义要在执行之前 顺序不能弄反