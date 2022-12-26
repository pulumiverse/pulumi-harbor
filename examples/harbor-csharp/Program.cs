using System.Collections.Generic;
using Pulumi;
using Pulumiverse.Harbor;

return await Deployment.RunAsync(() =>
{
   var registry = new Registry("registry", new RegistryArgs
   {
      ProviderName= "docker-hub",
      EndpointUrl="https://hub.docker.com",
      Name= "pulumi-harbor",
   });
   var project = new Project("project", new ProjectArgs
   {
      RegistryId= registry.RegistryId,
      Name= "pulumi-harbor",
      Public= "true" 
   });
});
