#include <mpi.h>
#include <unistd.h>
#include <stdlib.h>
#include <stdio.h>
#include <math.h>
#include <assert.h>
#include <string.h>
#include <sys/time.h>

#define MAX_REQ_NUM 1000

#define FIELD_WIDTH 20
#define FLOAT_PRECISION 2

#define MESSAGE_ALIGNMENT 64
#define MAX_ALIGNMENT 65536
#define MAX_MSG_SIZE (1<<22)
#define MYBUFSIZE (MAX_MSG_SIZE + MAX_ALIGNMENT)

int allocate_memory (char **sbuf, char **rbuf, int rank);
void print_header (int rank);
void touch_data (void *sbuf, void *rbuf, int rank, size_t size);