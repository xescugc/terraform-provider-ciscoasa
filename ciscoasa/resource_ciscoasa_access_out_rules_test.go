package ciscoasa

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/xanzy/go-ciscoasa/ciscoasa"
)

func TestAccCiscoASAAccessOutRules_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCiscsoASAAccessOutRulesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCiscoASAAccessOutRules_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASAAccessOutRulesExists("ciscoasa_access_out_rules.foo"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.#", "3"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.0.source", "192.168.10.0/23"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.0.destination", "192.168.12.0/23"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.0.destination_service", "icmp/0"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.0.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.0.permit", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.1.source", "192.168.10.5/32"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.1.destination", "192.168.15.0/25"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.1.destination_service", "tcp/443"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.1.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.1.permit", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.2.source", "192.168.10.0/24"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.2.source_service", "udp"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.2.destination", "192.168.15.6/32"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.2.destination_service", "udp/53"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.2.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.2.permit", "true"),
				),
			},
		},
	})
}

func TestAccCiscoASAAccessOutRules_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCiscsoASAAccessOutRulesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCiscoASAAccessOutRules_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASAAccessOutRulesExists("ciscoasa_access_out_rules.foo"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.#", "3"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.0.source", "192.168.10.0/23"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.0.destination", "192.168.12.0/23"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.0.destination_service", "icmp/0"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.0.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.0.permit", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.1.source", "192.168.10.5/32"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.1.destination", "192.168.15.0/25"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.1.destination_service", "tcp/443"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.1.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.1.permit", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.2.source", "192.168.10.0/24"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.2.source_service", "udp"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.2.destination", "192.168.15.6/32"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.2.destination_service", "udp/53"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.2.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.2.permit", "true"),
				),
			},

			{
				Config: testAccCiscoASAAccessOutRules_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCiscoASAAccessOutRulesExists("ciscoasa_access_out_rules.foo"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.0.source", "192.168.10.0/23"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.0.destination", "192.168.12.0/24"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.0.destination_service", "icmp/8"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.0.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.0.permit", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.2.source", "192.168.12.0/24"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.2.source_service", "tcp"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.2.destination", "192.168.15.16/32"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.2.destination_service", "tcp/53"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.2.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.2.permit", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.1.source", "192.168.10.0/24"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.1.destination", "192.168.15.0/25"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.1.destination_service", "tcp/443"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.1.active", "true"),
					resource.TestCheckResourceAttr(
						"ciscoasa_access_out_rules.foo", "rule.1.permit", "true"),
				),
			},
		},
	})
}

func testAccCheckCiscoASAAccessOutRulesExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No interface ID is set")
		}

		ca := testAccProvider.Meta().(*ciscoasa.Client)
		l, err := ca.Access.ListAccessOutRules(rs.Primary.ID)

		if err != nil {
			return err
		}

		if len(l.Items) == 0 {
			return fmt.Errorf("No rules for interface %s found", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckCiscsoASAAccessOutRulesDestroy(s *terraform.State) error {
	ca := testAccProvider.Meta().(*ciscoasa.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ciscoasa_access_out_rules" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No interface ID is set")
		}

		_, err := ca.Access.ListAccessOutRules(rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Rules for interface %s still exist", rs.Primary.ID)
		}

	}

	return nil
}

var testAccCiscoASAAccessOutRules_basic = fmt.Sprintf(`
resource "ciscoasa_access_out_rules" "foo" {
  interface = "%s"
  rule {
    source = "192.168.10.5/32"
    destination = "192.168.15.0/25"
    destination_service = "tcp/443"
  }
  rule {
    source = "192.168.10.0/24"
	source_service = "udp"
    destination = "192.168.15.6/32"
    destination_service = "udp/53"
  }
  rule {
    source = "192.168.10.0/23"
    destination = "192.168.12.0/23"
    destination_service = "icmp/0"
  }
}`,
	CISCOASA_INTERFACE_NAME)

var testAccCiscoASAAccessOutRules_update = fmt.Sprintf(`
resource "ciscoasa_access_out_rules" "foo" {
  interface = "%s"
  rule {
    source = "192.168.10.0/24"
    destination = "192.168.15.0/25"
    destination_service = "tcp/443"
  }
  rule {
    source = "192.168.12.0/24"
	source_service = "tcp"
    destination = "192.168.15.16/32"
    destination_service = "tcp/53"
  }
  rule {
    source = "192.168.10.0/23"
    destination = "192.168.12.0/24"
    destination_service = "icmp/8"
  }
}`,
	CISCOASA_INTERFACE_NAME)
