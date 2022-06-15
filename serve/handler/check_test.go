package handler

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestCheck(t *testing.T) {
	type args struct {
		context *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Check(tt.args.context)
		})
	}
}
