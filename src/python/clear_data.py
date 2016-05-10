#!/usr/bin/python

#DONOT-DELETE: very import for free space
import os , sys, getpass, time

print time.strftime("%Y-%m-%d %H:%M:%S");

#/home/hdp-jia/tmp/v0001/0001/src/20160226164519.0046606.rank.0001/20160226164519.0046606.rank.0001
dir="/home/hdp-jia/tmp/v0001/0001/src/";

cur_ts = time.time()
cur_ts -= 86400 * 25;
time = time.strftime("%Y%m%d", time.gmtime(cur_ts))
file = dir+time;
delete_file = file + "*";
delete_cmd =  "rm -rf " + delete_file;
print delete_cmd, "\n";
os.system(delete_cmd);
