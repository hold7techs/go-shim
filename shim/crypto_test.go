package shim

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignatureSha256(t *testing.T) {
	type args struct {
		k string
		v []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{"secret_key", []byte("hello_world")}, "20f54ae6995c7982c6b5e9dea18ea723085bdcc1f55db9964d33540860db604b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Sha256Signature(tt.args.k, tt.args.v), "Sha256Signature(%v, %v)", tt.args.k, tt.args.v)
		})
	}
}
