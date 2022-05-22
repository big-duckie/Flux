package utils

import (
	"fmt"
	"math"
)

type ByteUnit int

const (
	B   ByteUnit = 1
	KiB ByteUnit = 10
	MiB ByteUnit = 20
	GiB ByteUnit = 30
	TiB ByteUnit = 40
	PiB ByteUnit = 50
	EiB ByteUnit = 60
)

func ByteUnitParse(str string) ByteUnit {
	switch str {
	case "KiB":
		return KiB
	case "MiB":
		return MiB
	case "GiB":
		return GiB
	}
	return B
}

func (unit ByteUnit) StringLong() string {
	switch unit {
	case KiB:
		return "Kibibyte"
	case MiB:
		return "Mebibyte"
	case GiB:
		return "Gibibyte"
	}
	return "Byte"
}

func (unit ByteUnit) StringLongPlural() string {
	return unit.StringLong() + "s"
}

func (unit ByteUnit) StringShort() string {
	switch unit {
	case KiB:
		return "KiB"
	case MiB:
		return "MiB"
	case GiB:
		return "GiB"
	}
	return "B"
}

func (unit ByteUnit) Format(size int64) string {
	return fmt.Sprintf("%.2f %s", float64(size)/math.Pow(2, float64(unit)), unit.StringShort())
}
