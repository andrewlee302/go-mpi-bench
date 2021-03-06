package main

const (
	MAX_REQ_NUM       = 1000
	FIELD_WIDTH       = 20
	FLOAT_PRECISION   = 2
	MESSAGE_ALIGNMENT = 64
	MAX_ALIGNMENT     = 65536
	MAX_MSG_SIZE      = (1 << 22)
	MYBUFSIZE         = MAX_MSG_SIZE + MAX_ALIGNMENT
)
