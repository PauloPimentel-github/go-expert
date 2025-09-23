package main

import (
	"context"
	"fmt"
)

type ctxKey string

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, ctxKey("token"), "password")
	bookHotel(ctx, "Hotel")
}

func bookHotel(ctx context.Context, name string) {
	token := ctx.Value(ctxKey("token"))
	fmt.Println(token)
}
