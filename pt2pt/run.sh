#!/usr/bin/env bash

mpiexec -n 2 C/c-bandwidth  | tee log/c-bd.log
mpiexec -n 2 Go/go-bandwidth | tee log/go-bd.log
mpiexec -n 2 C/c-latency    | tee log/c-lt.log
mpiexec -n 2 Go/go-latency   | tee log/go-lt.log
