package pritunl_test

import (
	"testing"

	"github.com/atobaum/terraform-provider-pritunl/internal/pritunl"
	"github.com/stretchr/testify/assert"
)

func TestPinMarshal(t *testing.T) {
	p := pritunl.Pin{IsSet: true, Value: ""}

	res, err := p.MarshalJSON()

	assert.Equal(t, err, nil)
	assert.Equal(t, string(res), "true")

	p = pritunl.Pin{IsSet: true, Value: "pin_ex"}

	res, err = p.MarshalJSON()

	assert.Equal(t, err, nil)
	assert.Equal(t, string(res), "\"pin_ex\"")

	p = pritunl.Pin{IsSet: false, Value: "pin_ex"}

	res, err = p.MarshalJSON()

	assert.Error(t, err)

	p = pritunl.Pin{IsSet: false, Value: ""}

	res, err = p.MarshalJSON()

	assert.Equal(t, err, nil)
	assert.Equal(t, string(res), "false")
}
