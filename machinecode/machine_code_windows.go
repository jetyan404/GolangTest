//go:build windows

package machinecode

/*
 #cgo CFLAGS: -I${SRCDIR}/lib/windows-include
 #cgo windows,386 LDFLAGS: -L${SRCDIR}/lib/windows-386 -lyoyo_activate
 #cgo windows,amd64 LDFLAGS: -L${SRCDIR}/lib/windows-amd64 -lyoyo_activate
 #cgo LDFLAGS: -lstdc++ -lgcc -lsetupapi
 #include "libyoyo_activate.h"
*/
import "C"
import (
	"sort"
	"strings"
)

// GetVersion 获取当前机器码版本信息
func GetVersion() string {
	version := C.get_activate_version()
	return C.GoString(version)
}

// GetMachineCode 获取当前机器码
func GetMachineCode() string {
	machineCode := C.yoyo_machine_code()
	return C.GoString(machineCode)
}

// GetMachineCodeList 获取机器码列表
func GetMachineCodeList() []MachineCodeInfo {
	// 获取机器码列表字符串
	machineCodeInfoStr := C.GoString(C.get_sn_query())
	// 以#号为分隔符拆解字符串
	machineCodeInfoStrList := strings.Split(machineCodeInfoStr, "#")
	if len(machineCodeInfoStrList) == 0 {
		return nil
	}
	// 定义机器码信息列表
	var machineCodeInfoList []MachineCodeInfo
	// 遍历
	for _, item := range machineCodeInfoStrList {
		if len(item) > 0 && strings.Contains(item, "|") {
			// 以|为分隔符拆解
			typeAndMachineCode := strings.Split(item, "|")
			if len(typeAndMachineCode) == 2 {
				machineCodeInfoList = append(machineCodeInfoList, MachineCodeInfo{
					Version: typeAndMachineCode[0],
					Value:   typeAndMachineCode[1],
				})
			}
		}
	}

	// 获取当前机器码版本
	// 追加到列表
	machineCodeInfoList = append(machineCodeInfoList, MachineCodeInfo{
		Version: strings.ReplaceAll(GetVersion(), ".", ""),
		Value:   GetMachineCode(),
	})

	// 按版本排序
	sort.Slice(machineCodeInfoList, func(i, j int) bool {
		return machineCodeInfoList[i].Version > machineCodeInfoList[j].Version
	})

	// 返回机器码信息列表
	return machineCodeInfoList
}
