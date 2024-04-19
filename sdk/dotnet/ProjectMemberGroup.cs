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
    /// $ pulumi import harbor:index/projectMemberGroup:ProjectMemberGroup main /projects/10/members/200
    /// ```
    /// </summary>
    [HarborResourceType("harbor:index/projectMemberGroup:ProjectMemberGroup")]
    public partial class ProjectMemberGroup : global::Pulumi.CustomResource
    {
        [Output("groupId")]
        public Output<int?> GroupId { get; private set; } = null!;

        [Output("groupName")]
        public Output<string?> GroupName { get; private set; } = null!;

        [Output("ldapGroupDn")]
        public Output<string?> LdapGroupDn { get; private set; } = null!;

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
        /// The group type.  Can be set to `"ldap"`, `"internal"` or `"oidc"`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectMemberGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectMemberGroup(string name, ProjectMemberGroupArgs args, CustomResourceOptions? options = null)
            : base("harbor:index/projectMemberGroup:ProjectMemberGroup", name, args ?? new ProjectMemberGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectMemberGroup(string name, Input<string> id, ProjectMemberGroupState? state = null, CustomResourceOptions? options = null)
            : base("harbor:index/projectMemberGroup:ProjectMemberGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectMemberGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectMemberGroup Get(string name, Input<string> id, ProjectMemberGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectMemberGroup(name, id, state, options);
        }
    }

    public sealed class ProjectMemberGroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("groupId")]
        public Input<int>? GroupId { get; set; }

        [Input("groupName")]
        public Input<string>? GroupName { get; set; }

        [Input("ldapGroupDn")]
        public Input<string>? LdapGroupDn { get; set; }

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
        /// The group type.  Can be set to `"ldap"`, `"internal"` or `"oidc"`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ProjectMemberGroupArgs()
        {
        }
        public static new ProjectMemberGroupArgs Empty => new ProjectMemberGroupArgs();
    }

    public sealed class ProjectMemberGroupState : global::Pulumi.ResourceArgs
    {
        [Input("groupId")]
        public Input<int>? GroupId { get; set; }

        [Input("groupName")]
        public Input<string>? GroupName { get; set; }

        [Input("ldapGroupDn")]
        public Input<string>? LdapGroupDn { get; set; }

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
        /// The group type.  Can be set to `"ldap"`, `"internal"` or `"oidc"`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ProjectMemberGroupState()
        {
        }
        public static new ProjectMemberGroupState Empty => new ProjectMemberGroupState();
    }
}
