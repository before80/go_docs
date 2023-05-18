# poll

https://pkg.go.dev/internal/poll@go1.20.1



Package poll supports non-blocking I/O on file descriptors with polling. This supports I/O operations that block only a goroutine, not a thread. This is used by the net and os packages. It uses a poller built into the runtime, with support from the runtime scheduler.











  
  


  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  

## 常量 ¶

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/poll/hook_cloexec.go;l=12)

```
var Accept4Func func(int, int) (int, syscall.Sockaddr, error) = syscall.Accept4
```

Accept4Func is used to hook the accept4 call.

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/poll/hook_unix.go;l=15)

```
var AcceptFunc func(int) (int, syscall.Sockaddr, error) = syscall.Accept
```

AcceptFunc is used to hook the accept call.

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/poll/hook_unix.go;l=12)

```
var CloseFunc func(int) error = syscall.Close
```

CloseFunc is used to hook the close call.

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/poll/fd.go;l=51)

```
var ErrDeadlineExceeded error = &DeadlineExceededError{}
```

ErrDeadlineExceeded is returned for an expired deadline. This is exported by the os package as os.ErrDeadlineExceeded.

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/poll/fd.go;l=35)

```
var ErrFileClosing = errors.New("use of closed file")
```

ErrFileClosing is returned when a file descriptor is used after it has been closed.

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/poll/fd.go;l=31)

```
var ErrNetClosing = errNetClosing{}
```

ErrNetClosing is returned when a network descriptor is used after it has been closed.

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/poll/fd.go;l=39)

```
var ErrNoDeadline = errors.New("file type does not support deadline")
```

ErrNoDeadline is returned when a request is made to set a deadline on a file type that does not use the poller.

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/poll/fd.go;l=66)

```
var ErrNotPollable = errors.New("not pollable")
```

ErrNotPollable is returned when the file or socket is not suitable for event notification.

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/poll/fd.go;l=83)

```
var TestHookDidWritev = func(wrote int) {}
```

TestHookDidWritev is a hook for testing writev.

## 函数

#### func CopyFileRange [¶](https://pkg.go.dev/internal/poll@go1.20.1#CopyFileRange)added in go1.15

```
func CopyFileRange(dst, src *FD, remain int64) (written int64, handled bool, err error)
```

CopyFileRange copies at most remain bytes of data from src to dst, using the copy_file_range system call. dst and src must refer to regular files.

#### func DupCloseOnExec [¶](https://pkg.go.dev/internal/poll@go1.20.1#DupCloseOnExec)added in go1.11

```
func DupCloseOnExec(fd int) (int, string, error)
```

DupCloseOnExec dups fd and marks it close-on-exec.

#### func IsPollDescriptor [¶](https://pkg.go.dev/internal/poll@go1.20.1#IsPollDescriptor)added in go1.12

```
func IsPollDescriptor(fd uintptr) bool
```

IsPollDescriptor reports whether fd is the descriptor being used by the poller. This is only used for testing.

#### func SendFile [¶](https://pkg.go.dev/internal/poll@go1.20.1#SendFile)

```
func SendFile(dstFD *FD, src int, remain int64) (int64, error, bool)
```

SendFile wraps the sendfile system call.

#### func Splice [¶](https://pkg.go.dev/internal/poll@go1.20.1#Splice)added in go1.11

```
func Splice(dst, src *FD, remain int64) (written int64, handled bool, sc string, err error)
```

Splice transfers at most remain bytes of data from src to dst, using the splice system call to minimize copies of data from and to userspace.

Splice gets a pipe buffer from the pool or creates a new one if needed, to serve as a buffer for the data transfer. src and dst must both be stream-oriented sockets.

If err != nil, sc is the system call which caused the error.

## 类型

### type DeadlineExceededError [¶](https://pkg.go.dev/internal/poll@go1.20.1#DeadlineExceededError)added in go1.15

```
type DeadlineExceededError struct{}
```

DeadlineExceededError is returned for an expired deadline.

#### (*DeadlineExceededError) Error [¶](https://pkg.go.dev/internal/poll@go1.20.1#DeadlineExceededError.Error)added in go1.15

```
func (e *DeadlineExceededError) Error() string
```

Implement the net.Error interface. The string is "i/o timeout" because that is what was returned by earlier Go versions. Changing it may break programs that match on error strings.

#### (*DeadlineExceededError) Temporary [¶](https://pkg.go.dev/internal/poll@go1.20.1#DeadlineExceededError.Temporary)added in go1.15

```
func (e *DeadlineExceededError) Temporary() bool
```

#### (*DeadlineExceededError) Timeout [¶](https://pkg.go.dev/internal/poll@go1.20.1#DeadlineExceededError.Timeout)added in go1.15

```
func (e *DeadlineExceededError) Timeout() bool
```

### type FD [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD)

```
type FD struct {

	// System file descriptor. Immutable until Close.
	Sysfd int

	// Whether this is a streaming descriptor, as opposed to a
	// packet-based descriptor like a UDP socket. Immutable.
	IsStream bool

	// Whether a zero byte read indicates EOF. This is false for a
	// message based socket connection.
	ZeroReadIsEOF bool
	// contains filtered or unexported fields
}
```

FD is a file descriptor. The net and os packages use this type as a field of a larger type representing a network connection or OS file.

#### (*FD) Accept [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.Accept)

```
func (fd *FD) Accept() (int, syscall.Sockaddr, string, error)
```

Accept wraps the accept network call.

#### (*FD) Close [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.Close)

```
func (fd *FD) Close() error
```

Close closes the FD. The underlying file descriptor is closed by the destroy method when there are no remaining references.

#### (*FD) Dup [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.Dup)added in go1.11

```
func (fd *FD) Dup() (int, string, error)
```

Dup duplicates the file descriptor.

#### (*FD) Fchdir [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.Fchdir)

```
func (fd *FD) Fchdir() error
```

Fchdir wraps syscall.Fchdir.

#### (*FD) Fchmod [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.Fchmod)

```
func (fd *FD) Fchmod(mode uint32) error
```

Fchmod wraps syscall.Fchmod.

#### (*FD) Fchown [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.Fchown)

```
func (fd *FD) Fchown(uid, gid int) error
```

Fchown wraps syscall.Fchown.

#### (*FD) Fstat [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.Fstat)

```
func (fd *FD) Fstat(s *syscall.Stat_t) error
```

Fstat wraps syscall.Fstat

#### (*FD) Fsync [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.Fsync)

```
func (fd *FD) Fsync() error
```

Fsync wraps syscall.Fsync.

#### (*FD) Ftruncate [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.Ftruncate)

```
func (fd *FD) Ftruncate(size int64) error
```

Ftruncate wraps syscall.Ftruncate.

#### (*FD) Init [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.Init)

```
func (fd *FD) Init(net string, pollable bool) error
```

Init initializes the FD. The Sysfd field should already be set. This can be called multiple times on a single FD. The net argument is a network name from the net package (e.g., "tcp"), or "file". Set pollable to true if fd should be managed by runtime netpoll.

#### (*FD) Pread [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.Pread)

```
func (fd *FD) Pread(p []byte, off int64) (int, error)
```

Pread wraps the pread system call.

#### (*FD) Pwrite [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.Pwrite)

```
func (fd *FD) Pwrite(p []byte, off int64) (int, error)
```

Pwrite wraps the pwrite system call.

#### (*FD) RawControl [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.RawControl)

```
func (fd *FD) RawControl(f func(uintptr)) error
```

RawControl invokes the user-defined function f for a non-IO operation.

#### (*FD) RawRead [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.RawRead)

```
func (fd *FD) RawRead(f func(uintptr) bool) error
```

RawRead invokes the user-defined function f for a read operation.

#### (*FD) RawWrite [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.RawWrite)

```
func (fd *FD) RawWrite(f func(uintptr) bool) error
```

RawWrite invokes the user-defined function f for a write operation.

#### (*FD) Read [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.Read)

```
func (fd *FD) Read(p []byte) (int, error)
```

Read implements io.Reader.

#### (*FD) ReadDirent [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.ReadDirent)

```
func (fd *FD) ReadDirent(buf []byte) (int, error)
```

ReadDirent wraps syscall.ReadDirent. We treat this like an ordinary system call rather than a call that tries to fill the buffer.

#### (*FD) ReadFrom [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.ReadFrom)

```
func (fd *FD) ReadFrom(p []byte) (int, syscall.Sockaddr, error)
```

ReadFrom wraps the recvfrom network call.

#### (*FD) ReadFromInet4 [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.ReadFromInet4)added in go1.18

```
func (fd *FD) ReadFromInet4(p []byte, from *syscall.SockaddrInet4) (int, error)
```

ReadFromInet4 wraps the recvfrom network call for IPv4.

#### (*FD) ReadFromInet6 [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.ReadFromInet6)added in go1.18

```
func (fd *FD) ReadFromInet6(p []byte, from *syscall.SockaddrInet6) (int, error)
```

ReadFromInet6 wraps the recvfrom network call for IPv6.

#### (*FD) ReadMsg [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.ReadMsg)

```
func (fd *FD) ReadMsg(p []byte, oob []byte, flags int) (int, int, int, syscall.Sockaddr, error)
```

ReadMsg wraps the recvmsg network call.

#### (*FD) ReadMsgInet4 [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.ReadMsgInet4)added in go1.18

```
func (fd *FD) ReadMsgInet4(p []byte, oob []byte, flags int, sa4 *syscall.SockaddrInet4) (int, int, int, error)
```

ReadMsgInet4 is ReadMsg, but specialized for syscall.SockaddrInet4.

#### (*FD) ReadMsgInet6 [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.ReadMsgInet6)added in go1.18

```
func (fd *FD) ReadMsgInet6(p []byte, oob []byte, flags int, sa6 *syscall.SockaddrInet6) (int, int, int, error)
```

ReadMsgInet6 is ReadMsg, but specialized for syscall.SockaddrInet6.

#### (*FD) Seek [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.Seek)

```
func (fd *FD) Seek(offset int64, whence int) (int64, error)
```

Seek wraps syscall.Seek.

#### (*FD) SetBlocking [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.SetBlocking)added in go1.10

```
func (fd *FD) SetBlocking() error
```

SetBlocking puts the file into blocking mode.

#### (*FD) SetDeadline [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.SetDeadline)

```
func (fd *FD) SetDeadline(t time.Time) error
```

SetDeadline sets the read and write deadlines associated with fd.

#### (*FD) SetReadDeadline [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.SetReadDeadline)

```
func (fd *FD) SetReadDeadline(t time.Time) error
```

SetReadDeadline sets the read deadline associated with fd.

#### (*FD) SetWriteDeadline [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.SetWriteDeadline)

```
func (fd *FD) SetWriteDeadline(t time.Time) error
```

SetWriteDeadline sets the write deadline associated with fd.

#### (*FD) SetsockoptByte [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.SetsockoptByte)

```
func (fd *FD) SetsockoptByte(level, name int, arg byte) error
```

SetsockoptByte wraps the setsockopt network call with a byte argument.

#### (*FD) SetsockoptIPMreq [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.SetsockoptIPMreq)

```
func (fd *FD) SetsockoptIPMreq(level, name int, mreq *syscall.IPMreq) error
```

SetsockoptIPMreq wraps the setsockopt network call with an IPMreq argument.

#### (*FD) SetsockoptIPMreqn [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.SetsockoptIPMreqn)

```
func (fd *FD) SetsockoptIPMreqn(level, name int, mreq *syscall.IPMreqn) error
```

SetsockoptIPMreqn wraps the setsockopt network call with an IPMreqn argument.

#### (*FD) SetsockoptIPv6Mreq [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.SetsockoptIPv6Mreq)

```
func (fd *FD) SetsockoptIPv6Mreq(level, name int, mreq *syscall.IPv6Mreq) error
```

SetsockoptIPv6Mreq wraps the setsockopt network call with an IPv6Mreq argument.

#### (*FD) SetsockoptInet4Addr [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.SetsockoptInet4Addr)

```
func (fd *FD) SetsockoptInet4Addr(level, name int, arg [4]byte) error
```

SetsockoptInet4Addr wraps the setsockopt network call with an IPv4 address.

#### (*FD) SetsockoptInt [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.SetsockoptInt)

```
func (fd *FD) SetsockoptInt(level, name, arg int) error
```

SetsockoptInt wraps the setsockopt network call with an int argument.

#### (*FD) SetsockoptLinger [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.SetsockoptLinger)

```
func (fd *FD) SetsockoptLinger(level, name int, l *syscall.Linger) error
```

SetsockoptLinger wraps the setsockopt network call with a Linger argument.

#### (*FD) Shutdown [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.Shutdown)

```
func (fd *FD) Shutdown(how int) error
```

Shutdown wraps syscall.Shutdown.

#### (*FD) WaitWrite [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.WaitWrite)

```
func (fd *FD) WaitWrite() error
```

WaitWrite waits until data can be read from fd.

#### (*FD) Write [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.Write)

```
func (fd *FD) Write(p []byte) (int, error)
```

Write implements io.Writer.

#### (*FD) WriteMsg [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.WriteMsg)

```
func (fd *FD) WriteMsg(p []byte, oob []byte, sa syscall.Sockaddr) (int, int, error)
```

WriteMsg wraps the sendmsg network call.

#### (*FD) WriteMsgInet4 [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.WriteMsgInet4)added in go1.18

```
func (fd *FD) WriteMsgInet4(p []byte, oob []byte, sa *syscall.SockaddrInet4) (int, int, error)
```

WriteMsgInet4 is WriteMsg specialized for syscall.SockaddrInet4.

#### (*FD) WriteMsgInet6 [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.WriteMsgInet6)added in go1.18

```
func (fd *FD) WriteMsgInet6(p []byte, oob []byte, sa *syscall.SockaddrInet6) (int, int, error)
```

WriteMsgInet6 is WriteMsg specialized for syscall.SockaddrInet6.

#### (*FD) WriteOnce [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.WriteOnce)added in go1.10

```
func (fd *FD) WriteOnce(p []byte) (int, error)
```

WriteOnce is for testing only. It makes a single write call.

#### (*FD) WriteTo [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.WriteTo)

```
func (fd *FD) WriteTo(p []byte, sa syscall.Sockaddr) (int, error)
```

WriteTo wraps the sendto network call.

#### (*FD) WriteToInet4 [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.WriteToInet4)added in go1.18

```
func (fd *FD) WriteToInet4(p []byte, sa *syscall.SockaddrInet4) (int, error)
```

WriteToInet4 wraps the sendto network call for IPv4 addresses.

#### (*FD) WriteToInet6 [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.WriteToInet6)added in go1.18

```
func (fd *FD) WriteToInet6(p []byte, sa *syscall.SockaddrInet6) (int, error)
```

WriteToInet6 wraps the sendto network call for IPv6 addresses.

#### (*FD) Writev [¶](https://pkg.go.dev/internal/poll@go1.20.1#FD.Writev)

```
func (fd *FD) Writev(v *[][]byte) (int64, error)
```

Writev wraps the writev system call.