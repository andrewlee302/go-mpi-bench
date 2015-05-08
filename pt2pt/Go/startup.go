package main

import (
	"fmt"
	"go-mpi/MPI"
	"os"
	"time"
)

func main() {

	var myid int
	var t_start time.Time
	var t float64
	t_start = time.Now()
	MPI.Init(&os.Args)
	t = float64(time.Now().Sub(t_start).Nanoseconds()) / 1e3
	myid, _ = MPI.Comm_rank(MPI.COMM_WORLD)
	if myid == 0 {
		fmt.Fprintf(os.Stdout, "%*.*f\n", FIELD_WIDTH, FLOAT_PRECISION, t)
	}
	MPI.Finalize()
	os.Exit(0)
}
