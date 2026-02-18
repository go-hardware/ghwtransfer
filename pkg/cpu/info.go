package cpu

import (
	"fmt"

	"github.com/samber/lo"
)

type Info struct {
	CountProcs int
}

func (i *Info) Print() string {
	fmt.Printf("processors: %d\n", i.CountProcs)
	procs := make([]int, i.CountProcs)
	procStrs := lo.Map(procs, func(x int, _ int) string {
		return fmt.Sprintf("%d", x)
	})
	fmt.Printf("procs: %v", procStrs)
}
