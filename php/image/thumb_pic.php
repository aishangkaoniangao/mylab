<?php
$thumb_file = "/tmp/b.jpg";
$thumb_w = 256;
$thumb_h = 144;
$thumb_photo = imageCreateTrueColor($thumb_w, $thumb_h);

$srcfile = "thumb/20151124-084619-97f91256ac8eac01525bc36d80a111ff7996d4bf-5-1280-720.jpg";
list($src_w, $src_h) = getImageSize($srcfile);
$src_photo = imageCreateFromJpeg($srcfile);
imageCopyResampled($thumb_photo, $src_photo, 0, 0, 0, 0, $thumb_w, $thumb_h, $src_w, $src_h);

imageJpeg($thumb_photo, $thumb_file, 100);
