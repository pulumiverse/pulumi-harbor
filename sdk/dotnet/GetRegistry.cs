// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Harbor
{
    public static class GetRegistry
    {
        /// <summary>
        /// ## Example Usage
        /// 
        /// ```terraform
        /// data "harbor_registry" "main" {
        ///   name          = "test_docker_harbor"
        /// }
        /// 
        /// output "harbor_registry_id" {
        ///   value   = data.harbor_registry.main.id
        /// }
        /// ```
        /// </summary>
        public static Task<GetRegistryResult> InvokeAsync(GetRegistryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRegistryResult>("harbor:index/getRegistry:getRegistry", args ?? new GetRegistryArgs(), options.WithDefaults());

        /// <summary>
        /// ## Example Usage
        /// 
        /// ```terraform
        /// data "harbor_registry" "main" {
        ///   name          = "test_docker_harbor"
        /// }
        /// 
        /// output "harbor_registry_id" {
        ///   value   = data.harbor_registry.main.id
        /// }
        /// ```
        /// </summary>
        public static Output<GetRegistryResult> Invoke(GetRegistryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRegistryResult>("harbor:index/getRegistry:getRegistry", args ?? new GetRegistryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRegistryArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the register.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetRegistryArgs()
        {
        }
        public static new GetRegistryArgs Empty => new GetRegistryArgs();
    }

    public sealed class GetRegistryInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the register.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetRegistryInvokeArgs()
        {
        }
        public static new GetRegistryInvokeArgs Empty => new GetRegistryInvokeArgs();
    }


    [OutputType]
    public sealed class GetRegistryResult
    {
        /// <summary>
        /// The description of the external container register.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// If the certificate of the external container register can be verified.
        /// </summary>
        public readonly bool Insecure;
        /// <summary>
        /// The name of the register.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The id of the register within harbor.
        /// </summary>
        public readonly int RegistryId;
        /// <summary>
        /// The health status of the external container register
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The type of the provider type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The url endpoint for the external container register
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private GetRegistryResult(
            string description,

            string id,

            bool insecure,

            string name,

            int registryId,

            string status,

            string type,

            string url)
        {
            Description = description;
            Id = id;
            Insecure = insecure;
            Name = name;
            RegistryId = registryId;
            Status = status;
            Type = type;
            Url = url;
        }
    }
}
