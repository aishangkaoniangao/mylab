[����]
���С���ߣ���������������shell�������php������python����shell����
�������ֻ��Ҫ����������ֱ��ʹ��crontab������Ƕ����Ҳ������ctontab.���������и����⣺������Ҫ20�����̣���ô��Ҫ��crontab��д��20�����ݣ����ѹ���ͼ��������С���߾��������������������϶��shell���

[ʹ�÷���]
1����conf�ļ����¶����Լ���meta�ļ���
���磺
run_num='ps aux | grep file.php | grep -v grep | wc -l'
CMD='php project/file.php'
PROCESS_NUM=10
2��ִ��automake.sh��������������Ҫ��ִ��shell�ļ�
���cd $tools_root; bash automake.sh;
3��ִ��start_command.sh�����������ű�ִ��
���cd $tools_root; bash start_command.sh