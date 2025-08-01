// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Harbor
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import harbor:index/registry:Registry main /registries/7
    /// ```
    /// </summary>
    [HarborResourceType("harbor:index/registry:Registry")]
    public partial class Registry : global::Pulumi.CustomResource
    {
        [Output("accessId")]
        public Output<string?> AccessId { get; private set; } = null!;

        [Output("accessSecret")]
        public Output<string?> AccessSecret { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The url endpoint for the external container register ie `"https://hub.docker.com"`
        /// </summary>
        [Output("endpointUrl")]
        public Output<string> EndpointUrl { get; private set; } = null!;

        [Output("insecure")]
        public Output<bool?> Insecure { get; private set; } = null!;

        /// <summary>
        /// The name of the register.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the provider.
        /// </summary>
        [Output("providerName")]
        public Output<string> ProviderName { get; private set; } = null!;

        [Output("registryId")]
        public Output<int> RegistryId { get; private set; } = null!;

        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a Registry resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Registry(string name, RegistryArgs args, CustomResourceOptions? options = null)
            : base("harbor:index/registry:Registry", name, args ?? new RegistryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Registry(string name, Input<string> id, RegistryState? state = null, CustomResourceOptions? options = null)
            : base("harbor:index/registry:Registry", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-harbor",
                AdditionalSecretOutputs =
                {
                    "accessSecret",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Registry resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Registry Get(string name, Input<string> id, RegistryState? state = null, CustomResourceOptions? options = null)
        {
            return new Registry(name, id, state, options);
        }
    }

    public sealed class RegistryArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessId")]
        public Input<string>? AccessId { get; set; }

        [Input("accessSecret")]
        private Input<string>? _accessSecret;
        public Input<string>? AccessSecret
        {
            get => _accessSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accessSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The url endpoint for the external container register ie `"https://hub.docker.com"`
        /// </summary>
        [Input("endpointUrl", required: true)]
        public Input<string> EndpointUrl { get; set; } = null!;

        [Input("insecure")]
        public Input<bool>? Insecure { get; set; }

        /// <summary>
        /// The name of the register.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the provider.
        /// </summary>
        [Input("providerName", required: true)]
        public Input<string> ProviderName { get; set; } = null!;

        public RegistryArgs()
        {
        }
        public static new RegistryArgs Empty => new RegistryArgs();
    }

    public sealed class RegistryState : global::Pulumi.ResourceArgs
    {
        [Input("accessId")]
        public Input<string>? AccessId { get; set; }

        [Input("accessSecret")]
        private Input<string>? _accessSecret;
        public Input<string>? AccessSecret
        {
            get => _accessSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accessSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The url endpoint for the external container register ie `"https://hub.docker.com"`
        /// </summary>
        [Input("endpointUrl")]
        public Input<string>? EndpointUrl { get; set; }

        [Input("insecure")]
        public Input<bool>? Insecure { get; set; }

        /// <summary>
        /// The name of the register.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the provider.
        /// </summary>
        [Input("providerName")]
        public Input<string>? ProviderName { get; set; }

        [Input("registryId")]
        public Input<int>? RegistryId { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        public RegistryState()
        {
        }
        public static new RegistryState Empty => new RegistryState();
    }
}
