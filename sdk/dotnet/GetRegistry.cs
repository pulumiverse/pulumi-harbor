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
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Harbor = Pulumi.Harbor;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var main = Harbor.GetRegistry.Invoke(new()
        ///     {
        ///         Name = "test_docker_harbor",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["harborRegistryId"] = main.Apply(getRegistryResult =&gt; getRegistryResult.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRegistryResult> InvokeAsync(GetRegistryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRegistryResult>("harbor:index/getRegistry:getRegistry", args ?? new GetRegistryArgs(), options.WithDefaults());

        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Harbor = Pulumi.Harbor;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var main = Harbor.GetRegistry.Invoke(new()
        ///     {
        ///         Name = "test_docker_harbor",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["harborRegistryId"] = main.Apply(getRegistryResult =&gt; getRegistryResult.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRegistryResult> Invoke(GetRegistryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRegistryResult>("harbor:index/getRegistry:getRegistry", args ?? new GetRegistryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRegistryArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetRegistryArgs()
        {
        }
        public static new GetRegistryArgs Empty => new GetRegistryArgs();
    }

    public sealed class GetRegistryInvokeArgs : global::Pulumi.InvokeArgs
    {
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
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Insecure;
        public readonly string Name;
        public readonly int RegistryId;
        public readonly string Status;
        public readonly string Type;
        public readonly string Url;

        [OutputConstructor]
        private GetRegistryResult(
            string description,

            string id,

            string insecure,

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
