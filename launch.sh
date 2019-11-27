#!/bin/bash

port=8080
target_pid=lsof -i :$port | awk '{print $2}'|awk 'NR==2{print}'

if [ -n "$pid" ];then
  kill -9 $pid
fi

nohup go run up_down_files.go >/dev/null 2>&1 &

echo Succeed!!!
