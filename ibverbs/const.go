//go:build linux

package ibverbs

//#cgo linux LDFLAGS: -libverbs
//#include <infiniband/verbs.h>
import "C"
import "errors"

// access flag
const (
	IBV_ACCESS_LOCAL_WRITE   = C.IBV_ACCESS_LOCAL_WRITE
	IBV_ACCESS_REMOTE_WRITE  = C.IBV_ACCESS_REMOTE_WRITE
	IBV_ACCESS_REMOTE_READ   = C.IBV_ACCESS_REMOTE_READ
	IBV_ACCESS_REMOTE_ATOMIC = C.IBV_ACCESS_REMOTE_ATOMIC
)

// qp type
const (
	IBV_QPT_RC = C.IBV_QPT_RC
)

// wr action
const (
	IBV_WR_SEND          = C.IBV_WR_SEND
	IBV_WR_SEND_WITH_IMM = C.IBV_WR_SEND_WITH_IMM
	IBV_WR_RDMA_WRITE    = C.IBV_WR_RDMA_WRITE
	IBV_WR_RDMA_READ     = C.IBV_WR_RDMA_READ

	IBV_SEND_SIGNALED = C.IBV_SEND_SIGNALED
	IBV_SEND_INLINE   = C.IBV_SEND_INLINE
)

var QPClosedErr = errors.New("qp already closed")
