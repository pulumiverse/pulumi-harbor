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
    public static class GetUsers
    {
        /// <summary>
        /// ## Example Usage
        /// 
        /// ```terraform
        /// data "harbor_users" "example" {
        ///   username = "example-user"
        /// }
        /// 
        /// output "users_ids" {
        ///   value = [data.harbor_users.example.users.*.id]
        /// }
        /// ```
        /// </summary>
        public static Task<GetUsersResult> InvokeAsync(GetUsersArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUsersResult>("harbor:index/getUsers:getUsers", args ?? new GetUsersArgs(), options.WithDefaults());

        /// <summary>
        /// ## Example Usage
        /// 
        /// ```terraform
        /// data "harbor_users" "example" {
        ///   username = "example-user"
        /// }
        /// 
        /// output "users_ids" {
        ///   value = [data.harbor_users.example.users.*.id]
        /// }
        /// ```
        /// </summary>
        public static Output<GetUsersResult> Invoke(GetUsersInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUsersResult>("harbor:index/getUsers:getUsers", args ?? new GetUsersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUsersArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The email of the user to filter by.
        /// </summary>
        [Input("email")]
        public string? Email { get; set; }

        /// <summary>
        /// The name of the user to filter by.
        /// </summary>
        [Input("username")]
        public string? Username { get; set; }

        public GetUsersArgs()
        {
        }
        public static new GetUsersArgs Empty => new GetUsersArgs();
    }

    public sealed class GetUsersInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The email of the user to filter by.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// The name of the user to filter by.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public GetUsersInvokeArgs()
        {
        }
        public static new GetUsersInvokeArgs Empty => new GetUsersInvokeArgs();
    }


    [OutputType]
    public sealed class GetUsersResult
    {
        /// <summary>
        /// The email of the user to filter by.
        /// </summary>
        public readonly string? Email;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the user to filter by.
        /// </summary>
        public readonly string? Username;
        public readonly ImmutableArray<Outputs.GetUsersUserResult> Users;

        [OutputConstructor]
        private GetUsersResult(
            string? email,

            string id,

            string? username,

            ImmutableArray<Outputs.GetUsersUserResult> users)
        {
            Email = email;
            Id = id;
            Username = username;
            Users = users;
        }
    }
}
