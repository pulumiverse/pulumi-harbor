import pulumiverse_harbor as harbor
import pulumi

registry = harbor.Registry("registry", name="pulumi-harbor",
                           endpoint_url="https://harbor.pulumi.com",
                           provider_name="docker-hub")

project = harbor.Project("project", name="pulumi-harbor",
                         registry_id=registry.registry_id,
                         public="true")
