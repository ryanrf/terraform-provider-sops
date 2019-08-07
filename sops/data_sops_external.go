package sops

import (
	"io/ioutil"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceExternal() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceExternalRead,

		Schema: map[string]*schema.Schema{
			"input_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"source": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"data": &schema.Schema{
				Type:     schema.TypeMap,
				Computed: true,
			},
			"raw": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceExternalRead(d *schema.ResourceData, meta interface{}) error {
	source := d.Get("source").(string)
	content, err := ioutil.ReadAll(strings.NewReader(source))
	if err != nil {
		return err
	}

	format := d.Get("input_type").(string)
	return readData(content, format, d)
}
