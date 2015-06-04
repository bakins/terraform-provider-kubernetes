package kubernetes

import (
	"testing"

	"github.com/hashicorp/terraform/config"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"kubernetes": testAccProvider,
	}

	// Use the demo address for the acceptance tests
	testAccProvider.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		conf := &Config{Endpoint: "http://127.0.0.1:8080"}
		return conf.Client()
	}
}

func TestResourceProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestResourceProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func TestResourceProvider_Configure(t *testing.T) {
	rp := Provider()

	raw := map[string]interface{}{
		"endpoint": "http://127.0.0.1:8080",
	}

	rawConfig, err := config.NewRawConfig(raw)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	err = rp.Configure(terraform.NewResourceConfig(rawConfig))
	if err != nil {
		t.Fatalf("err: %s", err)
	}
}
