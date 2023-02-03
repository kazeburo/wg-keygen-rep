package keypair

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPair(t *testing.T) {
	p := NewPair("test")
	assert.Equal(t, p.Priv, "mIbQgYhMfWWaL+qgxVrQFaO/TxsrC4Is0V1sFbDwCkg=", "private key should be equal")
	assert.Equal(t, p.Pub, "2dXFeXHu+wheOrr3paSmzbgYXzAQVYPNsJrY9hiG7GU=", "public key should be equal")
}
