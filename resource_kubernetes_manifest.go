package kubernetes

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceKubernetesManifest() *schema.Resource {
	return &schema.Resource{
		Create: resourceKubernetesManifestCreate,
		Read:   resourceKubernetesManifestRead,
		Schema: map[string]*schema.Schema{
			"endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceKubernetesManifestCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] resourceKubernetesManifestCreate")
	return nil
}

func resourceKubernetesManifestRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] resourceKubernetesManifestRead")
	return nil
}
