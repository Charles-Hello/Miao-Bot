package util

import (
	"fmt"
	"github.com/pkg/errors"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

type Server struct {
	Os   Os   `json:"os"`
	Cpu  Cpu  `json:"cpu"`
	Ram  Ram  `json:"ram"`
	Disk Disk `json:"disk"`
}

type Os struct {
	GOOS         string `json:"goos"`
	NumCPU       int    `json:"numCpu"`
	Compiler     string `json:"compiler"`
	GoVersion    string `json:"goVersion"`
	NumGoroutine int    `json:"numGoroutine"`
}

type Cpu struct {
	Cpus  []float64 `json:"cpus"`
	Cores int       `json:"cores"`
}

type Ram struct {
	UsedMB      int `json:"usedMb"`
	TotalMB     int `json:"totalMb"`
	UsedPercent int `json:"usedPercent"`
}

type Disk struct {
	UsedMB      int `json:"usedMb"`
	UsedGB      int `json:"usedGb"`
	TotalMB     int `json:"totalMb"`
	TotalGB     int `json:"totalGb"`
	UsedPercent int `json:"usedPercent"`
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: InitCPU
//@description: OS信息
//@return: o Os, err error

func InitOS() (o Os) {
	o.GOOS = runtime.GOOS
	o.NumCPU = runtime.NumCPU()
	o.Compiler = runtime.Compiler
	o.GoVersion = runtime.Version()
	o.NumGoroutine = runtime.NumGoroutine()
	return o
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: InitCPU
//@description: CPU信息
//@return: c Cpu, err error

func InitCPU() (c Cpu, err error) {
	if cores, err := cpu.Counts(false); err != nil {
		return c, err
	} else {
		c.Cores = cores
	}
	if cpus, err := cpu.Percent(time.Duration(200)*time.Millisecond, true); err != nil {
		return c, err
	} else {
		c.Cpus = cpus
	}
	return c, nil
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: InitRAM
//@description: RAM信息
//@return: r Ram, err error

func InitRAM() (r Ram, err error) {
	if u, err := mem.VirtualMemory(); err != nil {
		return r, err
	} else {
		r.UsedMB = int(u.Used) / MB
		r.TotalMB = int(u.Total) / MB
		r.UsedPercent = int(u.UsedPercent)
	}
	return r, nil
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: InitDisk
//@description: 硬盘信息
//@return: d Disk, err error

func InitDisk() (d Disk, err error) {
	if u, err := disk.Usage("/"); err != nil {
		return d, err
	} else {
		d.UsedMB = int(u.Used) / MB
		d.UsedGB = int(u.Used) / GB
		d.TotalMB = int(u.Total) / MB
		d.TotalGB = int(u.Total) / GB
		d.UsedPercent = int(u.UsedPercent)
	}
	return d, nil
}

var s Server

func GetServerInfo() string {
	str1 := fmt.Sprintf("运行环境 >>>\n操作系统:%s\nCPU数量:%d\n编译器:%s\nGolang版本:%s\n\n硬盘信息 >>>\n总容量:%dGB\n已使用:%dGB\n已使用百分比:%d",
		s.Os.GOOS, s.Os.NumCPU, s.Os.Compiler, s.Os.GoVersion, s.Disk.TotalGB, s.Disk.UsedGB, s.Disk.UsedPercent)
	str2 := fmt.Sprintf("内存信息 >>>\n总内存:%dMB\n已使用:%dMB\n已使用百分比:%d",
		s.Ram.TotalMB, s.Ram.UsedMB, s.Ram.UsedPercent)
	return str1 + "%\n\n" + str2 + "%"
}

func InitServerInfo() {

	s.Os = InitOS()

	initCPU, err := InitCPU()
	if err != nil {
		fmt.Printf("original err:%T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace:\n %+v\n",err)  // %+v 可以在打印的时候打印完整的堆栈信息
	}
	s.Cpu = initCPU

	ram, err := InitRAM()
	if err != nil {
		fmt.Printf("original err:%T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace:\n %+v\n",err)  // %+v 可以在打印的时候打印完整的堆栈信息
	}
	s.Ram = ram

	initDisk, err := InitDisk()
	if err != nil {
		fmt.Printf("original err:%T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace:\n %+v\n",err)  // %+v 可以在打印的时候打印完整的堆栈信息
	}
	s.Disk = initDisk
}
