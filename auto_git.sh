#!/bin/bash

echo `git status`

git add .
git commit -m "$1"

git push live master
