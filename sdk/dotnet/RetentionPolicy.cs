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
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import harbor:index/retentionPolicy:RetentionPolicy main /retentions/10
    /// ```
    /// </summary>
    [HarborResourceType("harbor:index/retentionPolicy:RetentionPolicy")]
    public partial class RetentionPolicy : global::Pulumi.CustomResource
    {
        [Output("rules")]
        public Output<ImmutableArray<Outputs.RetentionPolicyRule>> Rules { get; private set; } = null!;

        /// <summary>
        /// The schedule of when you would like the policy to run. This can be `Hourly`, `Daily`, `Weekly` or can be a custom cron string.
        /// </summary>
        [Output("schedule")]
        public Output<string?> Schedule { get; private set; } = null!;

        /// <summary>
        /// The project id of which you would like to apply this policy.
        /// </summary>
        [Output("scope")]
        public Output<string> Scope { get; private set; } = null!;


        /// <summary>
        /// Create a RetentionPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RetentionPolicy(string name, RetentionPolicyArgs args, CustomResourceOptions? options = null)
            : base("harbor:index/retentionPolicy:RetentionPolicy", name, args ?? new RetentionPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RetentionPolicy(string name, Input<string> id, RetentionPolicyState? state = null, CustomResourceOptions? options = null)
            : base("harbor:index/retentionPolicy:RetentionPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RetentionPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RetentionPolicy Get(string name, Input<string> id, RetentionPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new RetentionPolicy(name, id, state, options);
        }
    }

    public sealed class RetentionPolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("rules", required: true)]
        private InputList<Inputs.RetentionPolicyRuleArgs>? _rules;
        public InputList<Inputs.RetentionPolicyRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.RetentionPolicyRuleArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// The schedule of when you would like the policy to run. This can be `Hourly`, `Daily`, `Weekly` or can be a custom cron string.
        /// </summary>
        [Input("schedule")]
        public Input<string>? Schedule { get; set; }

        /// <summary>
        /// The project id of which you would like to apply this policy.
        /// </summary>
        [Input("scope", required: true)]
        public Input<string> Scope { get; set; } = null!;

        public RetentionPolicyArgs()
        {
        }
        public static new RetentionPolicyArgs Empty => new RetentionPolicyArgs();
    }

    public sealed class RetentionPolicyState : global::Pulumi.ResourceArgs
    {
        [Input("rules")]
        private InputList<Inputs.RetentionPolicyRuleGetArgs>? _rules;
        public InputList<Inputs.RetentionPolicyRuleGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.RetentionPolicyRuleGetArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// The schedule of when you would like the policy to run. This can be `Hourly`, `Daily`, `Weekly` or can be a custom cron string.
        /// </summary>
        [Input("schedule")]
        public Input<string>? Schedule { get; set; }

        /// <summary>
        /// The project id of which you would like to apply this policy.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        public RetentionPolicyState()
        {
        }
        public static new RetentionPolicyState Empty => new RetentionPolicyState();
    }
}
