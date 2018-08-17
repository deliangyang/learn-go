package main

import (
	"os"
	"fmt"
	"strconv"
	"io/ioutil"
	"bytes"
)

func main() {
	slist := getAllPid()

	fmt.Printf("%8s %9s %s\n", "PID", "SWAP", "COMMAND")
	var total int64
	for _, info := range slist {
		fmt.Printf("%8d %9s %s\n", info.Pid, FormatSize(info.Size), info.Comm)
		total += info.Size
	}
	fmt.Printf("Total: %8s\n", FormatSize(total))
}

type Info struct {
	Pid  int
	Size int64
	Comm string
}

func getAllPid() (list [] Info) {
	f, _ := os.Open("/proc")
	defer f.Close()

	names, err := f.Readdirnames(0)
	if err != nil {
		panic(err)
	}

	for _, name := range names {
		pid, err := strconv.Atoi(name)
		if err != nil {
			continue
		}
		info, err := getPidInfo(pid)
		if err != nil || info.Size == 0 {
			continue
		}
		list = append(list, info)
	}
	return
}

var (
	nullBytes  = []byte{0x0}
	emptyBytes = []byte(" ")
)

func getPidInfo(pid int) (info Info, err error) {
	info.Pid = pid
	var bs []byte
	bs, err = ioutil.ReadFile(fmt.Sprintf("/proc/%d/cmdline", pid))
	if err != nil {
		panic(err)
	}
	if bytes.HasPrefix(bs, nullBytes) {
		bs = bs[:len(bs) - 1]
	}

	info.Comm = string(bytes.Replace(bs, nullBytes, emptyBytes, -1))
	bs, err = ioutil.ReadFile(fmt.Sprintf("/proc/%d/smaps", pid))
	if err != nil {
		return
	}
	var total int64
	for _, line := range bytes.Split(bs, []byte("\n")) {
		if bytes.HasPrefix(line, []byte("Size:")) {
			start := bytes.IndexAny(line, "0123456789")
			end := bytes.Index(line[start:], []byte(" "))
			size, err := strconv.ParseInt(string(line[start:start+end]), 10, 0)
			if err != nil {
				continue
			}
			total += size
		}
	}

	info.Size = total * 1024
	return
}

var units = []string{"", "K", "M", "G", "T"}

func FormatSize(s int64) string {
	unit := 0
	f := float64(s)
	for unit < len(units) && f > 1100.0 {
		f /= 1024.0
		unit++
	}
	if unit == 0 {
		return fmt.Sprintf("%dB", int64(f))
	} else {
		return fmt.Sprintf("%.1f%siB", f, units[unit])
	}
}