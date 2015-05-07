package main

import (
	"fmt"
	"go-mpi/MPI"
	"os"
	"unsafe"
)

const (
	BENCHMARK = "MPI%s Bandwidth Test"
	HEADER    = "# " + BENCHMARK + "\n"

	LOOP_LARGE         = 20
	SKIP_LARGE         = 2
	LARGE_MESSAGE_SIZE = 8192
	WINDOW_SIZE_LARGE  = 64
)

var (
	s_buf_original []byte        = make([]byte, MYBUFSIZE, MYBUFSIZE)
	r_buf_original []byte        = make([]byte, MYBUFSIZE, MYBUFSIZE)
	request        []MPI.Request = make([]MPI.Request, MAX_REQ_NUM, MAX_REQ_NUM)
	reqstat        []MPI.Status  = make([]MPI.Status, MAX_REQ_NUM, MAX_REQ_NUM)
)

func main() {

	//fmt.Printf("len(s_buf_original)=%d\ncap(s_buf_original)=%d\n&s_buf_original[0]=%d\n",
	//	len(s_buf_original), cap(s_buf_original), &s_buf_original[0])

	var myid, numprocs int
	var s_buf, r_buf []byte
	var t_start, t_end float64 = 0.0, 0.0
	var skip, loop, window_size int = 10, 100, 64

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
			window_size = WINDOW_SIZE_LARGE
		}

		MPI.Barrier(MPI.COMM_WORLD)

		if myid == 0 {
			for i := 0; i < loop+skip; i++ {
				if i == skip {
					t_start = MPI.Wtime()
				}

				for j := 0; j < window_size; j++ {
					MPI.Isend(s_buf, size, MPI.CHAR, 1, 100, MPI.COMM_WORLD, &request[j])
				}
				tmp_stat, _ := MPI.Waitall(request[:window_size])
				copy(reqstat[:window_size], tmp_stat) // should be optimized
				reqstat[0], _ = MPI.Recv(r_buf, 4, MPI.CHAR, 1, 101, MPI.COMM_WORLD)
			}

			t_end = MPI.Wtime()

		} else if myid == 1 {
			for i := 0; i < loop+skip; i++ {
				for j := 0; j < window_size; j++ {
					MPI.Irecv(r_buf, size, MPI.CHAR, 0, 100, MPI.COMM_WORLD, &request[j])
				}
				tmp_stat, _ := MPI.Waitall(request[:window_size])
				copy(reqstat[:window_size], tmp_stat)
				MPI.Send(s_buf, 4, MPI.CHAR, 0, 101, MPI.COMM_WORLD)
			}
		}

		if myid == 0 {
			t := t_end - t_start
			var bw float64 = float64(size) / 1e6 * float64(loop) * float64(window_size) / t
			fmt.Fprintf(os.Stdout, "%-*d%*.*f\n", 10, size, FIELD_WIDTH,
				FLOAT_PRECISION, bw)
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
		fmt.Printf("%-*s%*s\n", 10, "# Size", FIELD_WIDTH, "Bandwidth (MB/s)")
	}
}

func touch_data(sbuf, rbuf []byte, rank, size int) {
	for i := 0; i < size; i++ {
		sbuf[i] = 'a'
		rbuf[i] = 'b'
	}
}
