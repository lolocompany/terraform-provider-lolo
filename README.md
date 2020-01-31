# terraform-provider-lolo
Lolo Terraform Provider

## Example use

```
# Configure the Lolo provider
provider "lolo" {
  api_key = "${var.lolo_api_key}"
}

# Create a new app
resource "lolo_app" "app" {
  name = "My app"
  description = "Managed by Terraform"
  
  deployment {
    runtime_id = "lolo-eu"
    replicas = 1
  }
}
```
