package kubernetes

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceKubernetesReplicationController() *schema.Resource {
	return &schema.Resource{
		Create: resourceKubernetesReplicationControllerCreate,
		Read:   resourceKubernetesReplicationControllerRead,
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

func resourceKubernetesReplicationControllerCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] resourceKubernetesReplicationControllerCreate")
	return nil
}

func resourceKubernetesReplicationControllerRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] resourceKubernetesReplicationControllerRead")
	return nil
}
