# Go-MPI-Bench
[Go-MPI](https://github.com/JohannWeging/go-mpi) is a go-binding MPI interface, Go-MPI-bench is a benchmark for the evaluation of it. It contains:

* Point-to-point Latency
* Network Bandwidth
* Strat Overhead

## Go-Binding MPI Datatype
* BYTE                   
* CHAR                  
* CHARACTER             
* UNSIGNED_CHAR         
* SIGNED_CHAR           
* SHORT                 
* UNSIGNED_SHORT        
* INT                   
* INTEGER               
* UNSIGNED              
* LONG                  
* UNSIGNED_LONG         
* LONG_LONG_INT         
* UNSIGNED_LONG_LONG_INT
* FLOAT                 
* REAL                  
* DOUBLE                
* DOUBLE_PRECISION      
* COMPLEX               
* DOUBLE_COMPLEX   

## Go-Binding MPI Programming Interface(Only list common interface)
package MPI
* MPI_INIT(): func Init(argv *[]string) int
* MPI_FINALIZE(): func Finalize() int
* MPI_COMM_SIZE(comm, size): func Comm_size(comm Comm) (int, int)
* MPI_COMM_RANK(comm, rank): func Comm_rank(comm Comm) (int, int)
* MPI_SEND(buf, count, datatype, dest, tag, comm): func Send(buffer interface{},
	count int,
	dataType Datatype,
	dest int,
	tag int,
	comm Comm) int
* MPI_RECV (buf, count, datatype, source, tag, comm, status): func Recv(buffer interface{},
	count int,
	dataType Datatype,
	source int,
	tag int,
	comm Comm) (Status, int)
