package example

import (
	"fmt"
	"github.com/leonpatmore/helm-direct-use-plugin/internal/simple"
)

func Handle() {
	fmt.Print(simple.Add(1, 3))
}
