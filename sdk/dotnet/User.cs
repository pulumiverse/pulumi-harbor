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
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Harbor = Pulumiverse.Harbor;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var main = new Harbor.User("main", new()
    ///     {
    ///         Email = "john@smith.com",
    ///         FullName = "John Smith",
    ///         Password = "Password12345!",
    ///         Username = "john",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// An internal user harbor user can be imported using the `user id` eg, `
    /// 
    /// ```sh
    ///  $ pulumi import harbor:index/user:User main /users/19
    /// ```
    /// 
    ///  `
    /// </summary>
    [HarborResourceType("harbor:index/user:User")]
    public partial class User : global::Pulumi.CustomResource
    {
        [Output("admin")]
        public Output<bool?> Admin { get; private set; } = null!;

        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        [Output("email")]
        public Output<string> Email { get; private set; } = null!;

        [Output("fullName")]
        public Output<string> FullName { get; private set; } = null!;

        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a User resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public User(string name, UserArgs args, CustomResourceOptions? options = null)
            : base("harbor:index/user:User", name, args ?? new UserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private User(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
            : base("harbor:index/user:User", name, state, MakeResourceOptions(options, id))
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
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing User resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static User Get(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
        {
            return new User(name, id, state, options);
        }
    }

    public sealed class UserArgs : global::Pulumi.ResourceArgs
    {
        [Input("admin")]
        public Input<bool>? Admin { get; set; }

        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("email", required: true)]
        public Input<string> Email { get; set; } = null!;

        [Input("fullName", required: true)]
        public Input<string> FullName { get; set; } = null!;

        [Input("password", required: true)]
        private Input<string>? _password;
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public UserArgs()
        {
        }
        public static new UserArgs Empty => new UserArgs();
    }

    public sealed class UserState : global::Pulumi.ResourceArgs
    {
        [Input("admin")]
        public Input<bool>? Admin { get; set; }

        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("email")]
        public Input<string>? Email { get; set; }

        [Input("fullName")]
        public Input<string>? FullName { get; set; }

        [Input("password")]
        private Input<string>? _password;
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("username")]
        public Input<string>? Username { get; set; }

        public UserState()
        {
        }
        public static new UserState Empty => new UserState();
    }
}
