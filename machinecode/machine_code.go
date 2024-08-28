package machinecode

import "fmt"

// MachineCodeInfo 机器码信息
type MachineCodeInfo struct {
	Version string // 机器码版本
	Value   string // 机器码值
}

func (p *MachineCodeInfo) Clone() *MachineCodeInfo {
	if p == nil {
		return nil
	}
	return &MachineCodeInfo{
		Version: p.Version,
		Value:   p.Value,
	}
}

func (p *MachineCodeInfo) String() string {
	if p == nil {
		return "nil"
	}
	return fmt.Sprintf("(%s:%s)", p.Version, p.Value)
}
