package middleware

import "gin-vue-template/pkg/logic"

type Middleware struct {
	logic *logic.Logic
}

func NewMiddleware(logic *logic.Logic) *Middleware {
	return &Middleware{logic: logic}
}
