
#filter-test
run_num=`ps aux | grep filter.php | grep -v grep | wc -l`
CMD="php ${PROJECT_ROOT}/filter.php"
PROCESS_NUM=5

function  run()
{
    if [ $run_num -lt $PROCESS_NUM ]
    then
        for((i=$run_num;i<$PROCESS_NUM;i++))
        do
            {
                echo $CMD
                eval $CMD
            }
        done
    fi
}

run
echo "threads run"

run_num=0
