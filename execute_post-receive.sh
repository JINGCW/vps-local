#!/bin/bash

echo current dir is: $PWD

./launch.sh
if [ $? -ne 0 ];then
  chmod +x launch.sh
  ./launch.sh
fi
