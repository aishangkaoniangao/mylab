<?php
$srcfile = "thumb/word-1.jpg";
$dstfile = "thumb/word.jpg";

$src_width = 0;
$src_height = 0;

list($width, $height, $type, $attr) = getImageSize($srcfile);
$src_width = $width;
$src_height = $height;

$text = "2010-10-09";

$fnt = "zkt.ttf"; //字体所用的TrueType字库

//加载文件
$src_photo = imageCreateFromJpeg($srcfile);

//创建字体的颜色值  int
$color = imageColorAllocate($src_photo, 255, 255, 255);

$x = intval($src_width/2);
$y = intval($src_height/2);

//给image贴文字
//          (image, font size, angle, x, y, int color, string fontfile, string text)
imageTtfText($src_photo, 13, 0, $x, $y, $color, $fnt, $text);

imageJpeg($src_photo, $dstfile, 100);
imageDestroy($src_photo);
