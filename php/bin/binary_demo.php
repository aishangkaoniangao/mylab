<?php
//二进制转字符串，字符串转二进制
//字符串本身就可以当为二进制来处理
$str = "abcd";
echo bin2hex($str)."\n";//61626364

//int整型转二进制，二进制转int整型
$integer = 1000;
$int_bin = pack("N", $integer);
echo bin2hex($int_bin);
$int_int = unpack("N", $int_bin);//return array(1=>1000)
var_dump($int_int)."\n";

//二进制拼接
//1*int+4*char+1*int+10*char
$bin = '';
$pid = 12345;
$bin .= pack("N", $pid);
$c1 = "abcd";
$bin .= $c1;
$rand = rand(1, 1000);
$bin .= pack("N", $rand);
$c2 = "faklfjekalrjkler";
$bin .= substr($c2, 0, 10);
var_dump($pid, $c1, $rand, $c2);
echo bin2hex($bin)."\n";die;
//解释N1pid:前8B按照N来解析，解析成无符号32位整型，解析后的值是该整型，解析后的key是pid
//解释/是分隔符，用来划分各个数据单元
//C4str表示：这一部分的4B按照C来解析，解析后的值是C，解析后的key是str
$result = unpack("N1pid/C4str/N1rand/C10str2", $bin);
//var_dump($result);

$r_pid = $r_c1 = $r_rand = $r_c2 = '';
$r_pid = $result["pid"];
$r_rand = $result["rand"];
$r_c1 = "";
$i = 1;
while ($i <= 4) {
    $r_c1 .= chr($result["str" . $i]);
    $i++;
}
$i = 1;
while ($i <= 10) {
    $r_c2 .= chr($result["str2" . $i]);
    $i++;
}
var_dump($r_pid, $r_c1, $r_rand, $r_c2);

//直接用pack方式太别扭，不如直接拼接好用
//pack对于字符串的拼接处理的不太好，不能直接用，会把单个字符拆成8B长度，其实只需要1B
$bin2 = pack("N*C*N*C*", $pid, ord("c"), ord("d"), $rand, substr($c2, 0, 10));
echo bin2hex($bin2)."\n";
