package kubernetes

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccKubernetesManifest(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() {},
		Providers: testAccProviders,
		//CheckDestroy: testAccCheckConsulKeysDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccKubernetesManifestConfig,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
	})
}

const testAccKubernetesManifestConfig = `
provider "kubernetes" {
  	endpoint = "https://127.0.0.1:6443"
}
resource "kubernetes_manifest" "myapp" {
	endpoint = "https://127.0.0.1:6443"

}`
