// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDatasourceLoadBalancerProtocols_basic(t *testing.T) {
	providers := testAccProviders
	config := testProviderConfig() + `
	data "oci_load_balancer_protocols" "t" {
		compartment_id = "${var.compartment_id}"
	}`

	resourceName := "data.oci_load_balancer_protocols.t"

	resource.UnitTest(t, resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 providers,
		Steps: []resource.TestStep{
			{
				ImportState:       true,
				ImportStateVerify: true,
				Config:            config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "protocols.#"),
					resource.TestCheckResourceAttrSet(resourceName, "protocols.0.name"),
				),
			},
		},
	})
}
