#!/usr/bin/env bash

if [[ $# < 2 ]];then
	echo "usage: ./startup.sh [hosts] [app]"
    exit 1
fi

hosts=$1
app=$2

MAX_PROC_NUM=56
i=1
echo "# MPI Startup Overhead Test"
printf "%-*s%*s\n" 10 "# procNum" 20 "Overhead (us)"
while [[ $i -le $MAX_PROC_NUM ]];do
    printf "%-*d" 10 $i
	mpiexec -hosts $hosts -n $i $app 
    #sleep 5 
    i=$(($i*2))
done 

if [[ $i -ne $MAX_PROC_NUM ]];then
    i=$MAX_PROC_NUM
    printf "%-*d" 10 $i
	mpiexec -hosts $hosts -n $i $app 
fi

