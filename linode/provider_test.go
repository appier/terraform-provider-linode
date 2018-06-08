package linode

import (
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
)

var testProvider *schema.Provider

func init() {
	testProvider = Provider().(*schema.Provider)
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}
