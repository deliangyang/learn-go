package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
	"unsafe"
)

type RedisConn struct {
	RedisConn net.Conn
}

func (conn RedisConn) Command(cmd []byte) ([]byte, error) {
	if _, err := conn.RedisConn.Write(cmd); err != nil {
		panic(err)
	}
	reader := bufio.NewReader(conn.RedisConn)
	buffer, _, err := reader.ReadLine()
	if err != nil {
		return nil, err
	}
	return buffer, nil
}

func main() {

	conn, err := net.Dial("tcp", "192.168.1.18:6380")
	if err != nil {
		panic(err)
	}
	redis := RedisConn{
		RedisConn: conn,
	}

	defer conn.Close()

	if result, err := redis.Command([]byte("Ping\n")); err == nil {
		fmt.Println(string(result))
	}

	if result, err := redis.Command([]byte("set test 2\n")); err == nil {
		fmt.Println(string(result))
	}

	if result, err := redis.Command([]byte("get test\n")); err == nil {
		fmt.Println(string(result))
	}

	a := 3
	var b *int
	b = &a
	*b = 4
	fmt.Println("b:", *b)
	fmt.Println("a:", a)
	// Output:
	//  +PONG
	var c interface{}
	c = 3
	fmt.Println("b:", c.(int))

	// 任意地址可以转化为pointer
	d := unsafe.Pointer(&a)
	fmt.Println(d)
	fmt.Println("uintptr:", uintptr(d))

	cc := unsafe.Pointer(uintptr(unsafe.Pointer(&a)) + 2)
	fmt.Println(cc)
	fmt.Println("uintptr:", uintptr(cc))

	var bytePool = sync.Pool{
		New: func() interface{} {
			b := make([]byte, 1024)
			return b
		},
	}

	for i := 0; i < 10; i++ {
		bytePool.Put([]byte(fmt.Sprintf("x:%d", i)))
	}

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i:", i, string(bytePool.Get().([]byte)))
		}(i)
	}
}
