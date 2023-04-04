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
    /// ## # Resource: harbor.RobotAccount
    /// 
    /// Harbor supports different levels of robot accounts. Currently `system` and `project` level robot accounts are supported.
    /// 
    /// ## Example Usage
    /// ### System Level
    /// Introduced in harbor 2.2.0, system level robot accounts can have basically [all available permissions](https://github.com/goharbor/harbor/blob/-/src/common/rbac/const.go) in harbor and are not dependent on a single project.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Harbor = Pulumiverse.Harbor;
    /// using Random = Pulumi.Random;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var password = new Random.RandomPassword("password", new()
    ///     {
    ///         Length = 12,
    ///         Special = false,
    ///     });
    /// 
    ///     var main = new Harbor.Project("main");
    /// 
    ///     var system = new Harbor.RobotAccount("system", new()
    ///     {
    ///         Description = "system level robot account",
    ///         Level = "system",
    ///         Secret = resource.Random_password.Password.Result,
    ///         Permissions = new[]
    ///         {
    ///             new Harbor.Inputs.RobotAccountPermissionArgs
    ///             {
    ///                 Accesses = new[]
    ///                 {
    ///                     new Harbor.Inputs.RobotAccountPermissionAccessArgs
    ///                     {
    ///                         Action = "create",
    ///                         Resource = "labels",
    ///                     },
    ///                 },
    ///                 Kind = "system",
    ///                 Namespace = "/",
    ///             },
    ///             new Harbor.Inputs.RobotAccountPermissionArgs
    ///             {
    ///                 Accesses = new[]
    ///                 {
    ///                     new Harbor.Inputs.RobotAccountPermissionAccessArgs
    ///                     {
    ///                         Action = "push",
    ///                         Resource = "repository",
    ///                     },
    ///                     new Harbor.Inputs.RobotAccountPermissionAccessArgs
    ///                     {
    ///                         Action = "read",
    ///                         Resource = "helm-chart",
    ///                     },
    ///                     new Harbor.Inputs.RobotAccountPermissionAccessArgs
    ///                     {
    ///                         Action = "read",
    ///                         Resource = "helm-chart-version",
    ///                     },
    ///                 },
    ///                 Kind = "project",
    ///                 Namespace = main.Name,
    ///             },
    ///             new Harbor.Inputs.RobotAccountPermissionArgs
    ///             {
    ///                 Accesses = new[]
    ///                 {
    ///                     new Harbor.Inputs.RobotAccountPermissionAccessArgs
    ///                     {
    ///                         Action = "pull",
    ///                         Resource = "repository",
    ///                     },
    ///                 },
    ///                 Kind = "project",
    ///                 Namespace = "*",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// The above example, creates a system level robot account with permissions to
    /// - permission to create labels on system level
    /// - pull repository across all projects
    /// - push repository to project "my-project-name"
    /// - read helm-chart and helm-chart-version in project "my-project-name"
    /// ### Project Level
    /// 
    /// Other than system level robot accounts, project level robot accounts can interact on project level only.
    /// The [available permissions](https://github.com/goharbor/harbor/blob/-/src/common/rbac/const.go) are mostly the same as for system level robots.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Harbor = Pulumiverse.Harbor;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var main = new Harbor.Project("main");
    /// 
    ///     var project = new Harbor.RobotAccount("project", new()
    ///     {
    ///         Description = "project level robot account",
    ///         Level = "project",
    ///         Permissions = new[]
    ///         {
    ///             new Harbor.Inputs.RobotAccountPermissionArgs
    ///             {
    ///                 Accesses = new[]
    ///                 {
    ///                     new Harbor.Inputs.RobotAccountPermissionAccessArgs
    ///                     {
    ///                         Action = "pull",
    ///                         Resource = "repository",
    ///                     },
    ///                     new Harbor.Inputs.RobotAccountPermissionAccessArgs
    ///                     {
    ///                         Action = "push",
    ///                         Resource = "repository",
    ///                     },
    ///                 },
    ///                 Kind = "project",
    ///                 Namespace = main.Name,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// The above example creates a project level robot account with permissions to
    /// - pull repository on project "main"
    /// - push repository on project "main"
    /// </summary>
    [HarborResourceType("harbor:index/robotAccount:RobotAccount")]
    public partial class RobotAccount : global::Pulumi.CustomResource
    {
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("disable")]
        public Output<bool?> Disable { get; private set; } = null!;

        [Output("duration")]
        public Output<int?> Duration { get; private set; } = null!;

        [Output("fullName")]
        public Output<string> FullName { get; private set; } = null!;

        [Output("level")]
        public Output<string> Level { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("permissions")]
        public Output<ImmutableArray<Outputs.RobotAccountPermission>> Permissions { get; private set; } = null!;

        [Output("robotId")]
        public Output<string> RobotId { get; private set; } = null!;

        [Output("secret")]
        public Output<string> Secret { get; private set; } = null!;


        /// <summary>
        /// Create a RobotAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RobotAccount(string name, RobotAccountArgs args, CustomResourceOptions? options = null)
            : base("harbor:index/robotAccount:RobotAccount", name, args ?? new RobotAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RobotAccount(string name, Input<string> id, RobotAccountState? state = null, CustomResourceOptions? options = null)
            : base("harbor:index/robotAccount:RobotAccount", name, state, MakeResourceOptions(options, id))
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
                    "secret",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RobotAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RobotAccount Get(string name, Input<string> id, RobotAccountState? state = null, CustomResourceOptions? options = null)
        {
            return new RobotAccount(name, id, state, options);
        }
    }

    public sealed class RobotAccountArgs : global::Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("disable")]
        public Input<bool>? Disable { get; set; }

        [Input("duration")]
        public Input<int>? Duration { get; set; }

        [Input("level", required: true)]
        public Input<string> Level { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("permissions", required: true)]
        private InputList<Inputs.RobotAccountPermissionArgs>? _permissions;
        public InputList<Inputs.RobotAccountPermissionArgs> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<Inputs.RobotAccountPermissionArgs>());
            set => _permissions = value;
        }

        [Input("secret")]
        private Input<string>? _secret;
        public Input<string>? Secret
        {
            get => _secret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public RobotAccountArgs()
        {
        }
        public static new RobotAccountArgs Empty => new RobotAccountArgs();
    }

    public sealed class RobotAccountState : global::Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("disable")]
        public Input<bool>? Disable { get; set; }

        [Input("duration")]
        public Input<int>? Duration { get; set; }

        [Input("fullName")]
        public Input<string>? FullName { get; set; }

        [Input("level")]
        public Input<string>? Level { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("permissions")]
        private InputList<Inputs.RobotAccountPermissionGetArgs>? _permissions;
        public InputList<Inputs.RobotAccountPermissionGetArgs> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<Inputs.RobotAccountPermissionGetArgs>());
            set => _permissions = value;
        }

        [Input("robotId")]
        public Input<string>? RobotId { get; set; }

        [Input("secret")]
        private Input<string>? _secret;
        public Input<string>? Secret
        {
            get => _secret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public RobotAccountState()
        {
        }
        public static new RobotAccountState Empty => new RobotAccountState();
    }
}
