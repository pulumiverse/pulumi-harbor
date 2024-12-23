// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Harbor.Outputs
{

    [OutputType]
    public sealed class GetProjectMemberUsersProjectMemberUserResult
    {
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The id of the project within harbor.
        /// </summary>
        public readonly string ProjectId;
        public readonly string Role;
        public readonly string UserName;

        [OutputConstructor]
        private GetProjectMemberUsersProjectMemberUserResult(
            string id,

            string projectId,

            string role,

            string userName)
        {
            Id = id;
            ProjectId = projectId;
            Role = role;
            UserName = userName;
        }
    }
}
