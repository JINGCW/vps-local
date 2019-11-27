#!/bin/bash

port=8080
target_pid=`sudo lsof -i :$port | awk '{print $2}'|awk 'NR==2{print}'`

echo running pid is: $target_pid
if [ -n "$target_pid" ];then
  sudo kill -9 $target_pid
fi

sudo nohup go run up_down_files.go >/dev/null 2>&1 &

echo Succeed!!!
