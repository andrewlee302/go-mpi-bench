#ifndef BENCHMARK_HEADER
    #include "benchmark.h"
#endif

char s_buf_original[MYBUFSIZE];
char r_buf_original[MYBUFSIZE];

void *
align_buffer (void * ptr, unsigned long align_size)
{
    return (void *)(((unsigned long)ptr + (align_size - 1)) / align_size *
            align_size);
}

int
allocate_memory (char ** sbuf, char ** rbuf, int rank)
{
    unsigned long align_size = getpagesize();

    assert(align_size <= MAX_ALIGNMENT);

    *sbuf = align_buffer(s_buf_original, align_size);
    *rbuf = align_buffer(r_buf_original, align_size);
    return 0;
}

void
touch_data (void * sbuf, void * rbuf, int rank, size_t size)
{
    memset(sbuf, 'a', size);
    memset(rbuf, 'b', size);
}