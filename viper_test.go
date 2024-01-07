package main

import (
  "fmt"
  "testing"

  "go.uber.org/fx"
)

func TestV1(t *testing.T) {
  fx.New(
    fx.Supply(
      fx.Annotate(
        "ssss",
        fx.ResultTags(`group:"a1"`),
      ),
    ), fx.Supply(
      fx.Annotate(
        "ssss1",
        fx.ResultTags(`group:"a1"`),
      ),
    ),
    fx.Invoke(
      fx.Annotate(
        func(name ...string) {
          fmt.Println(name)
        },
        fx.ParamTags(`group:"a1"`),
      ),
    ),
  ).Run()
}

func TestV2(t *testing.T) {
  fx.New(
    fx.Supply("accccccccccccca"),
    fx.Supply(
      fx.Annotate(
        "ssss",
        fx.ResultTags(`group:"a1"`),
      ),
    ), fx.Supply(
      fx.Annotate(
        "ssss1",
        fx.ResultTags(`group:"a1"`),
      ),
    ),
    fx.Invoke(
      fx.Annotate(
        func(some string, name ...string) {
          fmt.Println(name)
        },
        fx.ParamTags(`optional:"true"group:"a1"`),

      ),
    ),
  ).Run()
}
