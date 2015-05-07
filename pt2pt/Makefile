install: c-latency c-bandwidth go-latency go-bandwidth

c-latency:
	cd C && mpicc latency.c bench_util.c -o c-latency
	
c-bandwidth:
	cd C && mpicc bandwidth.c bench_util.c -o c-bandwidth

go-latency: Go-MPI
	cd Go && go build -o go-latency latency.go bench_util.go

go-bandwidth:
	cd Go && go build -o go-bandwidth bandwidth.go bench_util.go

Go-MPI:
	go install go-mpi/MPI

clean:
	cd C && rm c-latency c-bandwidth
	cd Go && rm go-latency go-bandwidth
