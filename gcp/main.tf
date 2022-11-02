module "vm_dev" {
  source             = "git::https://github.com/yolitals/gcp-terraform.git//gcp/modules/compute_instance"
  project_name       = var.project_name
  region             = var.region
  instance_name      = var.instance_name
  machine_type       = var.machine_type
  boot_image         = var.boot_image
  ssh_key            = var.ssh_key
  target_tags        = var.target_tags
  ssh_private_key    = var.ssh_private_key
  json_credential    = var.json_credential
  firewall_rule_name = var.firewall_rule_name
  local_exec         = "ansible-playbook  -e 'host_key_checking=False' ./docker.yml -i inventory -u ubuntu --private-key ${var.ssh_private_key}"
}
