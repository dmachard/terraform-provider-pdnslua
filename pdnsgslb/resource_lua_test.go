package pdnsgslb

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPdnsgslbLua_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPdnsgslbDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckPdnsgslbLuaConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPdnsgslbLuaExists("pdnsgslb_lua.testlua"),
				),
			},
		},
	})
}

func testAccCheckPdnsgslbDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "pdnsgslb_lua" {
			continue
		}

		recordId := rs.Primary.ID

		_, err := c.doDelete(recordId)
		if err != nil {
			return err
		}
	}

	return nil
}

func testAccCheckPdnsgslbLuaExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No record id set")
		}

		return nil
	}
}

const testAccCheckPdnsgslbLuaConfig_basic = `
resource "pdnsgslb_lua" "testlua" {
	zone = "test.internal."
	name = "testlua"
	record {
	  rrtype = "TXT"
	  ttl = 30
	  snippet = "os.date()"
	}
}`
