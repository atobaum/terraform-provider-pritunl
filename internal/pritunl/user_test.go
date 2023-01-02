package pritunl_test

import (
	"encoding/json"
	"testing"

	"github.com/atobaum/terraform-provider-pritunl/internal/pritunl"
	"github.com/stretchr/testify/assert"
)

func TestPinMarshal(t *testing.T) {
	p := pritunl.Pin{IsSet: true, Value: ""}

	res, err := p.MarshalJSON()

	assert.Equal(t, nil, err)
	assert.Equal(t, "true", string(res))

	p = pritunl.Pin{IsSet: true, Value: "pin_ex"}

	res, err = p.MarshalJSON()

	assert.Equal(t, nil, err)
	assert.Equal(t, "\"pin_ex\"", string(res))

	p = pritunl.Pin{IsSet: true, Value: "unknown"}

	res, err = p.MarshalJSON()

	assert.Equal(t, nil, err)
	assert.Equal(t, "true", string(res))

	p = pritunl.Pin{IsSet: false, Value: "pin_ex"}

	res, err = p.MarshalJSON()

	assert.Error(t, err)

	p = pritunl.Pin{IsSet: false, Value: ""}

	res, err = p.MarshalJSON()

	assert.Equal(t, nil, err)
	assert.Equal(t, "false", string(res))
}

func TestIntialValue(t *testing.T) {
	p := pritunl.Pin{}

	assert.Equal(t, false, p.IsSet)
	assert.Equal(t, "", p.Value)
}

func TestPinUnmarshal(t *testing.T) {
	p := pritunl.Pin{}

	err := json.Unmarshal([]byte("\"123123\""), &p)

	assert.Equal(t, nil, err)
	assert.Equal(t, true, p.IsSet)
	assert.Equal(t, "123123", p.Value)

	p = pritunl.Pin{}

	err = json.Unmarshal([]byte("true"), &p)

	assert.Equal(t, nil, err)
	assert.Equal(t, true, p.IsSet)
	assert.Equal(t, "unknown", p.Value)

	p = pritunl.Pin{}

	err = json.Unmarshal([]byte("false"), &p)

	assert.Equal(t, nil, err)
	assert.Equal(t, false, p.IsSet)
	assert.Equal(t, "", p.Value)
}
