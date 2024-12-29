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
    public sealed class GetUsersUserResult
    {
        public readonly bool Admin;
        public readonly string Comment;
        /// <summary>
        /// The email of the user to filter by.
        /// </summary>
        public readonly string Email;
        public readonly string FullName;
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the user to filter by.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private GetUsersUserResult(
            bool admin,

            string comment,

            string email,

            string fullName,

            string id,

            string username)
        {
            Admin = admin;
            Comment = comment;
            Email = email;
            FullName = fullName;
            Id = id;
            Username = username;
        }
    }
}