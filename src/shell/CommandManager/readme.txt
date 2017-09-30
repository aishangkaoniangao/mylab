[描述]
这个小工具，用来管理多个启动shell命令，比如php，或者python或者shell本身。
如果进程只需要单个，可以直接使用crontab；如果是多个，也可以用ctontab.但是这里有个问题：比如需要20个进程，那么需要在crontab中写入20条数据，很难管理和计数。这个小工具就是用在这个场景，管理较多的shell命令。

[使用方法]
1，在conf文件夹下定义自己的meta文件。
比如：
run_num='ps aux | grep file.php | grep -v grep | wc -l'
CMD='php project/file.php'
PROCESS_NUM=10
2，执行automake.sh，用来生成所需要的执行shell文件
命令：cd $tools_root; bash automake.sh;
3，执行start_command.sh，用来启动脚本执行
命令：cd $tools_root; bash start_command.sh