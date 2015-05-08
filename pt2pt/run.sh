#!/usr/bin/env bash
if [[ $# < 2 || "$1" != "local" && "$1" != "non-local" ]]; then
    echo "./run.sh [mode] [hosts]"
    echo "mode: local non-local"
    exit 1
fi
if [[ ! -d "$(dirname $0)/log" ]]; then
    mkdir $(dirname $0)/log
fi
mode=$1
hosts=$2
mpiexec -hosts ${hosts} -n 2 C/c-bandwidth   | tee log/${mode}-c-bd.log
mpiexec -hosts ${hosts} -n 2 Go/go-bandwidth | tee log/${mode}-go-bd.log
mpiexec -hosts ${hosts} -n 2 C/c-latency     | tee log/${mode}-c-ly.log
mpiexec -hosts ${hosts} -n 2 Go/go-latency   | tee log/${mode}-go-ly.log
