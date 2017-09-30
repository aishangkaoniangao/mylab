#/bin/bash

########################################
# 检查参数数量是否正确
# check_args_num function_name expect_num achieved_num
########################################
function check_args_num()
{
    if [ $# -ne 3 ] 
    then
        echo "function check_args_num  expect 3 args, but achieved $? args.";
        exit 1;
    fi
    local func_name=$1;
    local expect_num=$2;
    local achieve_num=$3;
    if [ $expect_num -ne $achieve_num ]
    then
        echo "function $func_name expect $expect_num args, but achieved $achieve_num args.";
        exit 1;
    fi
}

