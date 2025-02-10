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
    public static class GetRobotAccounts
    {
        /// <summary>
        /// ## Example Usage
        /// 
        /// ```terraform
        /// data "harbor_robot_accounts" "example" {
        ///   name = "example-robot"
        /// }
        /// 
        /// output "robot_account_ids" {
        ///   value = [data.harbor_robot_accounts.example.robot_accounts.*.id]
        /// }
        /// ```
        /// </summary>
        public static Task<GetRobotAccountsResult> InvokeAsync(GetRobotAccountsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRobotAccountsResult>("harbor:index/getRobotAccounts:getRobotAccounts", args ?? new GetRobotAccountsArgs(), options.WithDefaults());

        /// <summary>
        /// ## Example Usage
        /// 
        /// ```terraform
        /// data "harbor_robot_accounts" "example" {
        ///   name = "example-robot"
        /// }
        /// 
        /// output "robot_account_ids" {
        ///   value = [data.harbor_robot_accounts.example.robot_accounts.*.id]
        /// }
        /// ```
        /// </summary>
        public static Output<GetRobotAccountsResult> Invoke(GetRobotAccountsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRobotAccountsResult>("harbor:index/getRobotAccounts:getRobotAccounts", args ?? new GetRobotAccountsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// ## Example Usage
        /// 
        /// ```terraform
        /// data "harbor_robot_accounts" "example" {
        ///   name = "example-robot"
        /// }
        /// 
        /// output "robot_account_ids" {
        ///   value = [data.harbor_robot_accounts.example.robot_accounts.*.id]
        /// }
        /// ```
        /// </summary>
        public static Output<GetRobotAccountsResult> Invoke(GetRobotAccountsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetRobotAccountsResult>("harbor:index/getRobotAccounts:getRobotAccounts", args ?? new GetRobotAccountsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRobotAccountsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Level of the robot account, currently either `system` or `project`. Default is `system`.
        /// </summary>
        [Input("level")]
        public string? Level { get; set; }

        /// <summary>
        /// The name of the robot account to filter by.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The id of the project within harbor.
        /// </summary>
        [Input("projectId")]
        public int? ProjectId { get; set; }

        public GetRobotAccountsArgs()
        {
        }
        public static new GetRobotAccountsArgs Empty => new GetRobotAccountsArgs();
    }

    public sealed class GetRobotAccountsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Level of the robot account, currently either `system` or `project`. Default is `system`.
        /// </summary>
        [Input("level")]
        public Input<string>? Level { get; set; }

        /// <summary>
        /// The name of the robot account to filter by.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The id of the project within harbor.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        public GetRobotAccountsInvokeArgs()
        {
        }
        public static new GetRobotAccountsInvokeArgs Empty => new GetRobotAccountsInvokeArgs();
    }


    [OutputType]
    public sealed class GetRobotAccountsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Level of the robot account, currently either `system` or `project`. Default is `system`.
        /// </summary>
        public readonly string? Level;
        /// <summary>
        /// The name of the robot account to filter by.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The id of the project within harbor.
        /// </summary>
        public readonly int? ProjectId;
        public readonly ImmutableArray<Outputs.GetRobotAccountsRobotAccountResult> RobotAccounts;

        [OutputConstructor]
        private GetRobotAccountsResult(
            string id,

            string? level,

            string? name,

            int? projectId,

            ImmutableArray<Outputs.GetRobotAccountsRobotAccountResult> robotAccounts)
        {
            Id = id;
            Level = level;
            Name = name;
            ProjectId = projectId;
            RobotAccounts = robotAccounts;
        }
    }
}
