package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-harbor/v3/go/harbor"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		harbor.NewProject(ctx, "project", &harbor.ProjectArgs{
			Name: pulumi.String("pulumi-harbor"),
		})
		return nil
	})
}
