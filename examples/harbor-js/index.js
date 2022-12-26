"use strict";
const harbor = require("@pulumiverse/harbor");

const registry = new harbor.Registry("registry", {
    providerName: "docker-hub",
    endpointUrl: "https://hub.docker.com",
    name: "pulumi-harbor"
})

const project = new harbor.Project("project", {
    name: "pulumi-harbor",
    registryId: registry.registryId,
    public: "true",
})
