package opalib

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/open-policy-agent/opa/rego"
)

// 存储权限控制策略
// 查询权限控制
func Check() error {
	r := rego.New(
		rego.Query("x = data.example.allow"),
		rego.Load([]string{"./example.rego"}, nil))

	ctx := context.Background()
	query, err := r.PrepareForEval(ctx)
	if err != nil {
		// handle error
		return err
	}

	bs, err := os.ReadFile("./input.json")
	if err != nil {
		// handle error
		return err
	}

	var input interface{}

	if err := json.Unmarshal(bs, &input); err != nil {
		// handle error
		return err
	}

	rs, err := query.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		// handle error
		return err
	}

	log.Printf("query Eval %+v", rs)
	return nil
}
