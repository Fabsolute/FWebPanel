package math

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Format(s uint64) string {
	sizes := []string{"B", "K", "M", "G", "T", "P", "E"}
	return MakeHumanReadable(s, 1000, sizes)
}

func Deformat(s string) uint64 {
	sizes := []string{"B", "K", "M", "G", "T", "P", "E"}
	return ParseBytes(s, 1000, sizes)
}

func MakeHumanReadable(s uint64, base float64, sizes []string) string {
	if s < 10 {
		return fmt.Sprintf("%d B", s)
	}
	e := math.Floor(Logn(float64(s), base))
	suffix := sizes[int(e)]
	val := math.Floor(float64(s)/math.Pow(base, e)*10+0.5) / 10
	return fmt.Sprintf("%.1f %s", val, suffix)
}

func ParseBytes(s string, base float64, sizes []string) uint64 {
	index := 0
	value := s
	for i, size := range sizes {
		s = strings.TrimSpace(s)
		if strings.HasSuffix(s, size) {
			value = strings.TrimSpace(strings.TrimSuffix(s, size))
			index = i
			break
		}
	}
	f, _ := strconv.ParseFloat(value, 64)
	return uint64(f * math.Pow(base, float64(index)))
}
