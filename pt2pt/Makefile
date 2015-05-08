install: c-latency c-bandwidth c-startup go-latency go-bandwidth go-startup

c-latency:
	cd C && mpicc latency.c bench_util.c -o c-latency
	
c-bandwidth:
	cd C && mpicc bandwidth.c bench_util.c -o c-bandwidth

c-startup:
	cd C && mpicc startup.c bench_util.c -o c-startup

go-latency: Go-MPI
	cd Go && go build -o go-latency latency.go bench_util.go

go-bandwidth:
	cd Go && go build -o go-bandwidth bandwidth.go bench_util.go

go-startup:
	cd Go && go build -o go-startup startup.go bench_util.go

Go-MPI:
	go install go-mpi/MPI

clean:
	cd C && rm c-latency c-bandwidth c-startup
	cd Go && rm go-latency go-bandwidth go-startup
