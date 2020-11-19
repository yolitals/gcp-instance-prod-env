package test

import (
	"crypto/tls"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"
)


func TestCreation(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"project_name" : "wwcode-terraform-admin",
		},
	}
	terraform.Init(t, terraformOptions)
	_, err:= terraform.ApplyAndIdempotentE(t, terraformOptions)
	// Methods ended on E allow you to handle the error
	if err != nil {
		t.Fatal(err)
	}
}
func TestHttpService(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"project_name" : "wwcode-terraform-admin",
		},
	}
	// Get app_url from terraform outputs
	app_url:= terraform.Output(t, terraformOptions, "app_url")
	
	// Setup a TLS configuration to submit with the helper, a blank struct is acceptable
	tlsConfig := tls.Config{}

	// It can take a minute or so for the Instance to boot up, so retry a few times
	maxRetries := 5
	timeBetweenRetries := 5 * time.Second

	// Verify that we get back a 200 OK with the expected instanceText
	http_helper.HTTPDoWithRetry(t, "GET", app_url, nil, nil, 200 , maxRetries, timeBetweenRetries, &tlsConfig)

}
func TestDestroy(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"project_name" : "wwcode-terraform-admin",
		},
	}
	terraform.Destroy(t, terraformOptions)

}