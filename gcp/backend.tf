terraform {
  backend "gcs" {
    bucket = "wwcode-terraform-admin"
    prefix = "terraform/prod-infra"
  }
}
