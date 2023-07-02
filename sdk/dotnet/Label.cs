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
    /// * Create a global label within harbor
    /// 
    /// * Creates a label for project
    /// 
    /// ## Import
    /// 
    /// Harbor label can be imported using the `label id` eg, `&lt;break&gt;&lt;break&gt;```sh&lt;break&gt; $ pulumi import harbor:index/label:Label main /labels/1 &lt;break&gt;```&lt;break&gt;&lt;break&gt;`
    /// </summary>
    [HarborResourceType("harbor:index/label:Label")]
    public partial class Label : global::Pulumi.CustomResource
    {
        [Output("color")]
        public Output<string?> Color { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("projectId")]
        public Output<string?> ProjectId { get; private set; } = null!;

        [Output("scope")]
        public Output<string> Scope { get; private set; } = null!;


        /// <summary>
        /// Create a Label resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Label(string name, LabelArgs? args = null, CustomResourceOptions? options = null)
            : base("harbor:index/label:Label", name, args ?? new LabelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Label(string name, Input<string> id, LabelState? state = null, CustomResourceOptions? options = null)
            : base("harbor:index/label:Label", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Label resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Label Get(string name, Input<string> id, LabelState? state = null, CustomResourceOptions? options = null)
        {
            return new Label(name, id, state, options);
        }
    }

    public sealed class LabelArgs : global::Pulumi.ResourceArgs
    {
        [Input("color")]
        public Input<string>? Color { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        public LabelArgs()
        {
        }
        public static new LabelArgs Empty => new LabelArgs();
    }

    public sealed class LabelState : global::Pulumi.ResourceArgs
    {
        [Input("color")]
        public Input<string>? Color { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("scope")]
        public Input<string>? Scope { get; set; }

        public LabelState()
        {
        }
        public static new LabelState Empty => new LabelState();
    }
}
