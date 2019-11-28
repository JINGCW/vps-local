#!/bin/bash

if [ ! -n "$1" ];then
  echo you have not input a message! 
else
  git add .
  git commit -m "$1"
  git push live master
fi

