
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
