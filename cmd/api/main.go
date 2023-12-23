package main

import (
	"context"
	"fmt"
	"time"

	"github.com/RussellLuo/orchestrator"
	"github.com/RussellLuo/orchestrator/builtin"
)

func main() {
	task := builtin.NewSerial("get_debt_from_user").Timeout(3*time.Second).Tasks(
		builtin.NewHTTP("get_user").Timeout(2*time.Second).Get(
			"http://localhost:4445/users/${input.userId}",
		),
		builtin.NewHTTP("get_debt").Timeout(2*time.Second).Get(
			"http://localhost:4444/debts/${get_user.body.debtId}",
		),
	)

	input := orchestrator.NewInput(map[string]any{"userId": "c644e258-2916-4543-9fbd-2c1ac6ad486d"})
	output, err := task.Execute(context.Background(), input)
	if err != nil {
		fmt.Println(err)
		return
	}

	body := output["body"].(map[string]any)
	fmt.Println(body)

}
