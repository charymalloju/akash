package base_test

import (
	"testing"

	"github.com/ovrclk/photon/types/base"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBytes_JSON(t *testing.T) {
	obj := new(base.Bytes)
	js := []byte(`"F8"`)

	require.NoError(t, obj.UnmarshalJSON(js))
	assert.Equal(t, (*obj)[0], uint8(0xF8))
	assert.Len(t, *obj, 1)

	js_, err := obj.MarshalJSON()
	require.NoError(t, err)
	assert.Equal(t, js, js_)
}

func TestBytes_String(t *testing.T) {
	obj := new(base.Bytes)
	str := "f8"

	require.NoError(t, obj.DecodeString(str))
	assert.Equal(t, (*obj)[0], uint8(0xf8))
	assert.Len(t, *obj, 1)

	str_ := obj.EncodeString()
	assert.Equal(t, str, str_)
}
