package main

import (
	"fmt"
	"go-mpi/MPI"
	"os"
	"unsafe"
)

const (
	BENCHMARK = "MPI%s Latency Test"
	HEADER    = "# " + BENCHMARK + "\n"

	LOOP_LARGE         = 100
	SKIP_LARGE         = 10
	LARGE_MESSAGE_SIZE = 8192
)

var (
	s_buf_original []byte = make([]byte, MYBUFSIZE, MYBUFSIZE)
	r_buf_original []byte = make([]byte, MYBUFSIZE, MYBUFSIZE)
)

func main() {

	//fmt.Printf("len(s_buf_original)=%d\ncap(s_buf_original)=%d\n&s_buf_original[0]=%d\n",
	//	len(s_buf_original), cap(s_buf_original), &s_buf_original[0])

	var myid, numprocs int
	var reqstat MPI.Status
	var s_buf, r_buf []byte
	var t_start, t_end float64 = 0.0, 0.0
	var skip, loop int = 1000, 10000

	MPI.Init(&os.Args)
	numprocs, _ = MPI.Comm_size(MPI.COMM_WORLD)
	myid, _ = MPI.Comm_rank(MPI.COMM_WORLD)

	if numprocs != 2 {
		if myid == 0 {
			fmt.Fprintf(os.Stderr, "This test requires exactly two processes\n")
		}
		MPI.Finalize()
		os.Exit(1)
	}

	if allocate_memory(&s_buf, &r_buf, myid) {
		/* Error allocating memory */
		MPI.Finalize()
		os.Exit(1)
	}
	print_header(myid)

	/* Latency test */
	for size := 1; size <= MAX_MSG_SIZE; size *= 2 {
		touch_data(s_buf, r_buf, myid, size)

		if size > LARGE_MESSAGE_SIZE {
			loop = LOOP_LARGE
			skip = SKIP_LARGE
		}

		MPI.Barrier(MPI.COMM_WORLD)

		if myid == 0 {
			for i := 0; i < loop+skip; i++ {
				if i == skip {
					t_start = MPI.Wtime()
				}

				MPI.Send(s_buf, size, MPI.CHAR, 1, 1, MPI.COMM_WORLD)
				reqstat, _ = MPI.Recv(r_buf, size, MPI.CHAR, 1, 1, MPI.COMM_WORLD)
			}

			t_end = MPI.Wtime()
		} else if myid == 1 {
			for i := 0; i < loop+skip; i++ {
				reqstat, _ = MPI.Recv(r_buf, size, MPI.CHAR, 0, 1, MPI.COMM_WORLD)
				MPI.Send(s_buf, size, MPI.CHAR, 0, 1, MPI.COMM_WORLD)
			}
		}

		if myid == 0 {
			var latency float64 = (t_end - t_start) * 1e6 / (2.0 * float64(loop))
			fmt.Fprintf(os.Stdout, "%-*d%*.*f\n", 10, size, FIELD_WIDTH,
				FLOAT_PRECISION, latency)
		}
	}
	_ = reqstat
	MPI.Finalize()
	os.Exit(0)
}

func align_buffer(buf_original []byte, align_size int) []byte {
	buf_start := int(uintptr((unsafe.Pointer(&buf_original[0]))))
	align_start := (buf_start + align_size - 1) / align_size * align_size
	//fmt.Printf("\nalign_size=%d\nbuf_start=%d\nalign_start=%d\n", align_size, buf_start, align_start)
	return buf_original[(align_start - buf_start):]
}

func allocate_memory(s_buf, r_buf *[]byte, myid int) bool {
	var align_size int = os.Getpagesize()
	*s_buf = align_buffer(s_buf_original, align_size)
	*r_buf = align_buffer(r_buf_original, align_size)
	return false
}

func print_header(myid int) {
	if 0 == myid {
		fmt.Printf(HEADER, "")
		fmt.Printf("%-*s%*s\n", 10, "# Size", FIELD_WIDTH, "Latency (us)")
	}
}

func touch_data(sbuf, rbuf []byte, rank, size int) {
	for i := 0; i < size; i++ {
		sbuf[i] = 'a'
		rbuf[i] = 'b'
	}
}
