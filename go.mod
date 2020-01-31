module github.com/lolocompany/terraform-provider-lolo

go 1.13

//replace github.com/lolocompany/lolo-sdk-go => ../lolo-sdk-go

require (
	github.com/hashicorp/terraform-plugin-sdk v1.6.0
	github.com/lolocompany/lolo-sdk-go v0.0.0-20200131132157-3fa56764affe
)
