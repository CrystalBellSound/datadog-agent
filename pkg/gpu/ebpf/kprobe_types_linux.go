// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs -- -I ../../network/ebpf/c -I ../../ebpf/c -fsigned-char kprobe_types.go

package ebpf

type CudaEventType uint32
type CudaEventHeader struct {
	Pid_tgid  uint64
	Stream_id uint64
	Ktime_ns  uint64
	Type      uint32
	Pad       uint32
	Cgroup    [129]byte
	Pad_cgo_0 [7]byte
}

type CudaKernelLaunch struct {
	Header          CudaEventHeader
	Kernel_addr     uint64
	Shared_mem_size uint64
	Grid_size       Dim3
	Block_size      Dim3
}
type Dim3 struct {
	X uint32
	Y uint32
	Z uint32
}

type CudaSync struct {
	Header CudaEventHeader
}

type CudaMemEvent struct {
	Header    CudaEventHeader
	Size      uint64
	Addr      uint64
	Type      uint32
	Pad_cgo_0 [4]byte
}
type CudaMemEventType uint32

type CudaSetDeviceEvent struct {
	Header    CudaEventHeader
	Device    int32
	Pad_cgo_0 [4]byte
}

type CudaEventKey struct {
	Event     uint64
	Pid       uint32
	Pad_cgo_0 [4]byte
}
type CudaEventValue struct {
	Stream          uint64
	Access_ktime_ns uint64
}

const CudaEventTypeKernelLaunch CudaEventType = 0x0
const CudaEventTypeMemory CudaEventType = 0x1
const CudaEventTypeSync CudaEventType = 0x2
const CudaEventTypeSetDevice CudaEventType = 0x3

const CudaMemAlloc = 0x0
const CudaMemFree = 0x1

const SizeofCudaKernelLaunch = 0xd0
const SizeofCudaMemEvent = 0xc0
const SizeofCudaEventHeader = 0xa8
const SizeofCudaSync = 0xa8
const SizeofCudaSetDeviceEvent = 0xb0
