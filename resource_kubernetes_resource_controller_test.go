package kubernetes

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccKubernetesReplicationController(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() {},
		Providers: testAccProviders,
		//CheckDestroy: testAccCheckConsulKeysDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccKubernetesReplicationControllerConfig,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
	})
}

const testAccKubernetesReplicationControllerConfig = `
provider "kubernetes" {
  	endpoint = "https://127.0.0.1:6443"
}
resource "kubernetes_replication_controller" "myapp" {
	endpoint = "https://127.0.0.1:6443"

}`
