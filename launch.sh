#!/bin/bash

port=8080
target_pid=`lsof -i :$port | awk '{print $2}'|awk 'NR==2{print}'`

echo running pid is: $target_pid
if [ -n "$target_pid" ];then
  kill -9 $target_pid
fi

nohup go run up_down_files.go >/dev/null 2>&1 &

echo Succeed!!!
