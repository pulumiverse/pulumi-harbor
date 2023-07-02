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
    /// <summary>
    /// ## Example Usage
    /// </summary>
    [HarborResourceType("harbor:index/configSystem:ConfigSystem")]
    public partial class ConfigSystem : global::Pulumi.CustomResource
    {
        [Output("projectCreationRestriction")]
        public Output<string?> ProjectCreationRestriction { get; private set; } = null!;

        [Output("readOnly")]
        public Output<bool?> ReadOnly { get; private set; } = null!;

        [Output("robotNamePrefix")]
        public Output<string?> RobotNamePrefix { get; private set; } = null!;

        [Output("robotTokenExpiration")]
        public Output<int?> RobotTokenExpiration { get; private set; } = null!;

        [Output("scannerSkipUpdatePulltime")]
        public Output<bool?> ScannerSkipUpdatePulltime { get; private set; } = null!;


        /// <summary>
        /// Create a ConfigSystem resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConfigSystem(string name, ConfigSystemArgs? args = null, CustomResourceOptions? options = null)
            : base("harbor:index/configSystem:ConfigSystem", name, args ?? new ConfigSystemArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConfigSystem(string name, Input<string> id, ConfigSystemState? state = null, CustomResourceOptions? options = null)
            : base("harbor:index/configSystem:ConfigSystem", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse/pulumi-harbor",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ConfigSystem resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConfigSystem Get(string name, Input<string> id, ConfigSystemState? state = null, CustomResourceOptions? options = null)
        {
            return new ConfigSystem(name, id, state, options);
        }
    }

    public sealed class ConfigSystemArgs : global::Pulumi.ResourceArgs
    {
        [Input("projectCreationRestriction")]
        public Input<string>? ProjectCreationRestriction { get; set; }

        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        [Input("robotNamePrefix")]
        public Input<string>? RobotNamePrefix { get; set; }

        [Input("robotTokenExpiration")]
        public Input<int>? RobotTokenExpiration { get; set; }

        [Input("scannerSkipUpdatePulltime")]
        public Input<bool>? ScannerSkipUpdatePulltime { get; set; }

        public ConfigSystemArgs()
        {
        }
        public static new ConfigSystemArgs Empty => new ConfigSystemArgs();
    }

    public sealed class ConfigSystemState : global::Pulumi.ResourceArgs
    {
        [Input("projectCreationRestriction")]
        public Input<string>? ProjectCreationRestriction { get; set; }

        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        [Input("robotNamePrefix")]
        public Input<string>? RobotNamePrefix { get; set; }

        [Input("robotTokenExpiration")]
        public Input<int>? RobotTokenExpiration { get; set; }

        [Input("scannerSkipUpdatePulltime")]
        public Input<bool>? ScannerSkipUpdatePulltime { get; set; }

        public ConfigSystemState()
        {
        }
        public static new ConfigSystemState Empty => new ConfigSystemState();
    }
}
