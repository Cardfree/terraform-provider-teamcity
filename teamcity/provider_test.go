package teamcity

import (
	//"log"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"teamcity": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("TEAMCITY_URL"); v == "" {
		t.Fatal("TEAMCITY_URL must be set for acceptance tests")
	}
	if v := os.Getenv("TEAMCITY_USERNAME"); v == "" {
		t.Fatal("TEAMCITY_USERNAME must be set for acceptance tests")
	}
	if v := os.Getenv("TEAMCITY_PASSWORD"); v == "" {
		t.Fatal("TEAMCITY_PASSWORD must be set for acceptance tests")
	}
}
