package mso

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"mso": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	// We will use this function later on to make sure our test environment is valid.
	// For example, you can make sure here that some environment variables are set.
	if v := os.Getenv("MSO_USERNAME"); v == "" {
		t.Fatal("username variable must be set for acceptance tests")
	}

	if v := os.Getenv("MSO_PASSWORD"); v == "" {

		t.Fatal("password variable must be set for acceptance tests")
	}
	if v := os.Getenv("MSO_URL"); v == "" {
		t.Fatal("url variable must be set for acceptance tests")
	}

}
