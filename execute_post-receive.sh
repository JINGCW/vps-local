#!/bin/bash

echo current dir is: $PWD

sudo bash launch.sh
if [ $? -ne 0 ];then
  chmod +x launch.sh
  sudo bash launch.sh
fi
