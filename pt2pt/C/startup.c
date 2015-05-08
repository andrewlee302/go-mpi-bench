#ifndef BENCHMARK_HEADER
	#include "benchmark.h"
#endif

int
main (int argc, char *argv[])
{
	int myid;
	struct timeval t_start;
    struct timeval t_end;
    double t = 0.0;
    gettimeofday(&t_start, NULL);
    MPI_Init(&argc, &argv);
    gettimeofday(&t_end, NULL);
    t = 1e6 * (t_end.tv_sec - t_start.tv_sec) + t_end.tv_usec- t_start.tv_usec;
    MPI_Comm_rank(MPI_COMM_WORLD, &myid);
    if(myid == 0)
    	fprintf(stdout, "%*.*f\n", FIELD_WIDTH, FLOAT_PRECISION, t);
    fflush(stdout);
    MPI_Finalize();
    exit(EXIT_SUCCESS);
}