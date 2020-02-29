#!/bin/sh -l

echo "Hello $1"
pwd
ls
git status
git branch
find /
time=$(date)
echo ::set-output name=time::$time
