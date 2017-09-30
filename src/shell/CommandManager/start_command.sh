#!/bin/bash

PHP="/usr/local/php/bin/php"; 
TOOLS_ROOT=`pwd`; #工具根路径
PROJECT_ROOT=$TOOLS_ROOT'/../'; #项目根路径

date +%F/%T

. $TOOLS_ROOT/utils.sh
for file in $TOOLS_ROOT/conf/*
do
    if [ ${file:0-8} = '.meta.sh' ]
    then
        echo $file
        . $file
    fi
done
