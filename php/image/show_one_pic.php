<?php
//把多张图片合成一张图片 1做缩略图 2加时间
$sn = "36020092068";
$srcdir = dirname(__FILE__)."/data/$sn/"; //源文件夹
$dst_file = "/tmp/$sn.jpg"; //目标文件

$row_photo_num = 5; //横排的图片数目
$x_photo_i = 0; //横向计数
$y_photo_i = 0; //纵向计数
$dst_x = 0; //目标图片的x坐标
$dst_y = 0; //目标图片的y坐标

$width = 1305; //目标图片的宽度
$height = 636; //目标图片的高度
$dst_photo = imagecreatetruecolor($width, $height); //目标图片

$thumb_w = 256; //缩略图宽度
$thumb_h = 144; //缩略图高度

$x_delta = 5; //横向两张图片间隔
$y_delta = 15; //纵向两张图片间隔

$fnt = "zkt.ttf"; //字体所用的TrueType字库

$dp = opendir($srcdir);

while (($file = readdir($dp)) !== false) {
    if ($file == "." ||  $file == "..") {
        continue;
    }

    $text = getDateFromFile($file);

    if (($x_photo_i % $row_photo_num) == 0 && $x_photo_i != 0) {
        $dst_y += $thumb_h + $y_delta;
        $dst_x = 0;
    }

    $x_photo_i++;

    //贴thumb图
    $srcfile = $srcdir . $file;
    $add_photo = imageCreateFromJpeg($srcfile);
    list($add_w, $add_h) = getImageSize($srcfile);
    if ($add_w > $thumb_w) {
        $thumb_photo = imageCreateTrueColor($thumb_w, $thumb_h);
        imageCopyResampled($thumb_photo, $add_photo, 0, 0, 0, 0, $thumb_w, $thumb_h, $add_w, $add_h);
        $add_photo = $thumb_photo;
    }
    imageCopy($dst_photo, $add_photo, $dst_x, $dst_y, $src_x=0, $src_y=0, $src_w=$thumb_w, $src_h=$thumb_h);

    //贴文字
    $text_x = $dst_x + intval($thumb_w / 2) - 20;
    $text_y = $dst_y + intval($thumb_h) + 12;
    
    $color = imageColorAllocate($dst_photo, 255, 255, 255); //文字颜色 
    //贴文字    (image, font size, angle, x, y, int color, string fontfile, string text)
    imageTtfText($dst_photo, 12, 0, $text_x, $text_y, $color, $fnt, $text);

    $dst_x += $thumb_w + $x_delta;
}
imagejpeg($dst_photo, $dst_file, 100);
closedir($dp);

function getDateFromFile($file)
{
    $parts = explode("-", $file);
    $date = $parts[1];
    $unixtime = strtotime($date);
    return date("Y-m-d", $unixtime);
}
