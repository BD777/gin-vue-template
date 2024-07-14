package main

import (
	"gin-vue-template/pkg/wire"
)

func main() {
	s := wire.InitializeService()

	s.HTTPServer.ListenAndServe()
}
