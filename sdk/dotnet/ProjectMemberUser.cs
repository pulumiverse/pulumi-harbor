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
    /// Harbor project member user can be imported using the `project id` and `member id` eg, `&lt;break&gt;&lt;break&gt;```sh&lt;break&gt; $ pulumi import harbor:index/projectMemberUser:ProjectMemberUser main /projects/10/members/200 &lt;break&gt;```&lt;break&gt;&lt;break&gt;`
    /// </summary>
    [HarborResourceType("harbor:index/projectMemberUser:ProjectMemberUser")]
    public partial class ProjectMemberUser : global::Pulumi.CustomResource
    {
        [Output("memberId")]
        public Output<int> MemberId { get; private set; } = null!;

        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        [Output("userName")]
        public Output<string> UserName { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectMemberUser resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectMemberUser(string name, ProjectMemberUserArgs args, CustomResourceOptions? options = null)
            : base("harbor:index/projectMemberUser:ProjectMemberUser", name, args ?? new ProjectMemberUserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectMemberUser(string name, Input<string> id, ProjectMemberUserState? state = null, CustomResourceOptions? options = null)
            : base("harbor:index/projectMemberUser:ProjectMemberUser", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectMemberUser resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectMemberUser Get(string name, Input<string> id, ProjectMemberUserState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectMemberUser(name, id, state, options);
        }
    }

    public sealed class ProjectMemberUserArgs : global::Pulumi.ResourceArgs
    {
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        public ProjectMemberUserArgs()
        {
        }
        public static new ProjectMemberUserArgs Empty => new ProjectMemberUserArgs();
    }

    public sealed class ProjectMemberUserState : global::Pulumi.ResourceArgs
    {
        [Input("memberId")]
        public Input<int>? MemberId { get; set; }

        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        [Input("role")]
        public Input<string>? Role { get; set; }

        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public ProjectMemberUserState()
        {
        }
        public static new ProjectMemberUserState Empty => new ProjectMemberUserState();
    }
}
