package action

import (
	"fmt"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"strconv"
)

//	 公共方法 -- 获取内存使用情况
func getMemSize() (rep string) {
	v, _ := mem.VirtualMemory()

	fmt.Printf("        Mem       : %v GB  Free: %v MB Usage:%f%%\n", v.Total/4/1024/1024/1024/1024, v.Free/4/1024/1024/1024, v.UsedPercent)
	//
	rep = strconv.FormatFloat(v.UsedPercent, 'f', 2, 64)

	//
	return rep
}

//	 公共方法 -- 获取 CPU 使用情况
func getCpuSize() (rep string) {

	d, _ := disk.Usage("/")

	rep = strconv.FormatFloat(d.UsedPercent, 'f', 2, 64)

	return rep
}
