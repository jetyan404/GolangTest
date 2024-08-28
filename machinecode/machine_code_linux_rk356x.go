//go:build linux && rk356x

package machinecode

import (
	"strings"

	"github.com/jamesits/go-rkvendorstorage"
)

// GetVersion 获取当前机器码版本信息
func GetVersion() string {
	return "0.0.0.0"
}

// GetMachineCode 获取当前机器码
func GetMachineCode() string {
	_, data, _ := rkvendorstorage.Read(rkvendorstorage.IDVendorSN)
	return string(data)
}

// GetMachineCodeList 获取机器码列表
func GetMachineCodeList() []MachineCodeInfo {
	// 定义机器码信息列表
	var machineCodeInfoList []MachineCodeInfo
	// 获取当前机器码版本
	// 追加到列表
	machineCodeInfoList = append(machineCodeInfoList, MachineCodeInfo{
		Version: strings.ReplaceAll(GetVersion(), ".", ""),
		Value:   GetMachineCode(),
	})

	// 返回机器码信息列表
	return machineCodeInfoList
}
