import * as harbor from '@pulumiverse/harbor';

let registry = new harbor.Registry('registry', {
    providerName: "docker-hub",
    endpointUrl: "https://hub.docker.com",
    name: "pulumi-harbor"
});

let project = new harbor.Project('project', {
    name: "pulumi-harbor",
    registryId: registry.registryId,
    public: "true",
});
