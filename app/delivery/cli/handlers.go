package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kabbachokk/golang-test-3/app"
	"github.com/spf13/cobra"
)

type cliHandlers struct {
	uc app.UseCase
	rc *cobra.Command // rootCmd
}

func NewCliHandlers(
	uc app.UseCase,
	rc *cobra.Command,
) *cliHandlers {
	return &cliHandlers{uc, rc}
}

func (p *cliHandlers) RootHandler(cmd *cobra.Command, args []string) {
	str := strings.Split(strings.Trim(args[0], "[]"), ",")
	ids := make([]int, len(str))
	for i, s := range str {
		ids[i], _ = strconv.Atoi(s)
	}

	res, err := p.uc.GetOrderRacksByOrderID(ids)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf("\n=+=+=+=")
	fmt.Printf("\nСтраница сборки заказов %s\n", strings.Join(args, " "))

	cur := ""
	for _, v := range res {
		if cur != v.PrimaryRack {
			cur = v.PrimaryRack
			fmt.Printf("\n===Стеллаж %s", v.PrimaryRack)
		}
		fmt.Printf("\n%s (id=%d)", v.ProductName, v.ProductID)
		fmt.Printf("\nзаказ %d, %dшт", v.OrderID, v.Qty)

		if v.SecondaryRacks != "" {
			fmt.Printf("\nдоп стеллаж: %s", v.SecondaryRacks)
		}

		fmt.Printf("\n")
	}
}
