package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-harbor/v3/go/harbor"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		registry, err := harbor.NewRegistry(ctx, "registry", &harbor.RegistryArgs{
			ProviderName: pulumi.String("docker-hub"),
			EndpointUrl:  pulumi.String("https://hub.docker.com"),
			Name:         pulumi.String("pulumi-harbor"),
		})
		if err != nil {
			return err
		}

		_, err = harbor.NewProject(ctx, "project", &harbor.ProjectArgs{
			Name:       pulumi.String("pulumi-harbor"),
			Public:     pulumi.String("true"),
			RegistryId: registry.RegistryId,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
