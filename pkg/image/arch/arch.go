package arch

import (
	"runtime"
	"fmt"
)

type KnownArchType string

const (
	Unknown KnownArchType = ""
	Arm     KnownArchType = "arm"
	ArmBE   KnownArchType = "armbe"
	Arm64   KnownArchType = "arm64"
	Arm64BE KnownArchType = "arm64be"
)

var knownValues = map[KnownArchType]string{
	Arm:     string(Arm),
	ArmBE:   string(ArmBE),
	Arm64:   string(Arm64),
	Arm64BE: string(Arm64BE),
}

func Values() []string {
	var values []string
	for _, value := range knownValues {
		values = append(values, value)
	}
	return values
}

func (arch KnownArchType) Valid() bool {
	_, ok := knownValues[arch]
	return ok
}

func (arch KnownArchType) IsNative() bool {
	fmt.Println("arch: ", arch)
	fmt.Println("runtime.GOARCH: ", runtime.GOARCH)
	return string(arch) == runtime.GOARCH
}
