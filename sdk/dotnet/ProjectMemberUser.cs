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
    /// $ pulumi import harbor:index/projectMemberUser:ProjectMemberUser main /projects/10/members/200
    /// ```
    /// </summary>
    [HarborResourceType("harbor:index/projectMemberUser:ProjectMemberUser")]
    public partial class ProjectMemberUser : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The member id of the member.
        /// </summary>
        [Output("memberId")]
        public Output<int> MemberId { get; private set; } = null!;

        /// <summary>
        /// The project id of the project that the entity will have access to.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// The permissions that the entity will be granted.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        /// <summary>
        /// The name of the member entity.
        /// </summary>
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
        /// <summary>
        /// The project id of the project that the entity will have access to.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// The permissions that the entity will be granted.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        /// <summary>
        /// The name of the member entity.
        /// </summary>
        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        public ProjectMemberUserArgs()
        {
        }
        public static new ProjectMemberUserArgs Empty => new ProjectMemberUserArgs();
    }

    public sealed class ProjectMemberUserState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The member id of the member.
        /// </summary>
        [Input("memberId")]
        public Input<int>? MemberId { get; set; }

        /// <summary>
        /// The project id of the project that the entity will have access to.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// The permissions that the entity will be granted.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// The name of the member entity.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public ProjectMemberUserState()
        {
        }
        public static new ProjectMemberUserState Empty => new ProjectMemberUserState();
    }
}
