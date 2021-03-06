package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMISVPCDefaultRoutingTableDataSource_basic(t *testing.T) {
	node := "data.ibm_is_vpc_default_routing_table.def_route_table"
	vpcname := fmt.Sprintf("tf-vpcname-%d", acctest.RandIntRange(100, 200))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPCDefaultRoutingTableDataSourceConfig(vpcname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "id"),
					resource.TestCheckResourceAttrSet(node, "name"),
					resource.TestCheckResourceAttrSet(node, "lifecycle_state"),
				),
			},
		},
	})
}

func testAccCheckIBMISVPCDefaultRoutingTableDataSourceConfig(vpcname string) string {
	return fmt.Sprintf(`

	resource "ibm_is_vpc" "test_vpc" {
  		name = "%s"
	}
	
	data "ibm_is_vpc_default_routing_table" "def_route_table" {
		vpc = ibm_is_vpc.test_vpc.id
	}
	`, vpcname)
}
