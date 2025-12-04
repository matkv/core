package config

type Device string

const (
	Desktop Device = "desktop"
	Laptop  Device = "laptop"
	WSL     Device = "wsl"
)

func (d Device) IsValidDevice() bool {
	switch d {
	case Desktop, Laptop, WSL:
		return true
	default:
		return false
	}
}
