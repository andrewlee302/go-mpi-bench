#ifndef BENCHMARK_HEADER
	#include "benchmark.h"
#endif

#define BENCHMARK "MPI%s Bandwidth Test"
#define HEADER "# " BENCHMARK "\n"

#define LOOP_LARGE  20
#define WINDOW_SIZE_LARGE  64
#define SKIP_LARGE  2
#define LARGE_MESSAGE_SIZE  8192

char s_buf_original[MYBUFSIZE];
char r_buf_original[MYBUFSIZE];

MPI_Request request[MAX_REQ_NUM];
MPI_Status  reqstat[MAX_REQ_NUM];

int allocate_memory (char **sbuf, char **rbuf, int rank);
void print_header (int rank);
void touch_data (void *sbuf, void *rbuf, int rank, size_t size);

int
main (int argc, char *argv[])
{
	int myid, numprocs, i, j;
    int size;
    char *s_buf, *r_buf;
    double t_start = 0.0, t_end = 0.0, t = 0.0;
    int skip = 10;
    int loop = 100;
    int window_size = 64;

    MPI_Init(&argc, &argv);
    MPI_Comm_size(MPI_COMM_WORLD, &numprocs);
    MPI_Comm_rank(MPI_COMM_WORLD, &myid);

    if(numprocs != 2) {
        if(myid == 0) {
            fprintf(stderr, "This test requires exactly two processes\n");
        }

        MPI_Finalize();
        exit(EXIT_FAILURE);
    }

    if (allocate_memory(&s_buf, &r_buf, myid)) {
        /* Error allocating memory */
        MPI_Finalize();
        exit(EXIT_FAILURE);
    }

    print_header(myid);

    /* Bandwidth test */
    for(size = 1; size <= MAX_MSG_SIZE; size *= 2) {
        touch_data(s_buf, r_buf, myid, size);

        if(size > LARGE_MESSAGE_SIZE) {
            loop = LOOP_LARGE;
            skip = SKIP_LARGE;
            window_size = WINDOW_SIZE_LARGE;
        }

        if(myid == 0) {
            for(i = 0; i < loop + skip; i++) {
                if(i == skip) {
                    t_start = MPI_Wtime();
                }

                for(j = 0; j < window_size; j++) {
                    MPI_Isend(s_buf, size, MPI_CHAR, 1, 100, MPI_COMM_WORLD,
                            request + j);
                }

                MPI_Waitall(window_size, request, reqstat);
                MPI_Recv(r_buf, 4, MPI_CHAR, 1, 101, MPI_COMM_WORLD,
                        &reqstat[0]);
            }

            t_end = MPI_Wtime();
            t = t_end - t_start;
        }

        else if(myid == 1) {
            for(i = 0; i < loop + skip; i++) {
                for(j = 0; j < window_size; j++) {
                    MPI_Irecv(r_buf, size, MPI_CHAR, 0, 100, MPI_COMM_WORLD,
                            request + j);
                }

                MPI_Waitall(window_size, request, reqstat);
                MPI_Send(s_buf, 4, MPI_CHAR, 0, 101, MPI_COMM_WORLD);
            }
        }

        if(myid == 0) {
            double tmp = size / 1e6 * loop * window_size; // MB magnitude

            fprintf(stdout, "%-*d%*.*f\n", 10, size, FIELD_WIDTH,
                    FLOAT_PRECISION, tmp / t);
            fflush(stdout);
        }
    }

    MPI_Finalize();
    return EXIT_SUCCESS;
}

void
print_header (int rank)
{
    if (0 == rank) {
        printf(HEADER, "");
        printf("%-*s%*s\n", 10, "# Size", FIELD_WIDTH, "Bandwidth (MB/s)");
        fflush(stdout);
    }
}