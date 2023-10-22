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
    public sealed class GetGroupsGroupResult
    {
        /// <summary>
        /// The name of the group to filter by.
        /// </summary>
        public readonly string GroupName;
        /// <summary>
        /// The type of the group.
        /// </summary>
        public readonly int GroupType;
        /// <summary>
        /// The ID of the group.
        /// </summary>
        public readonly int Id;
        /// <summary>
        /// The LDAP group DN to filter by.
        /// </summary>
        public readonly string LdapGroupDn;

        [OutputConstructor]
        private GetGroupsGroupResult(
            string groupName,

            int groupType,

            int id,

            string ldapGroupDn)
        {
            GroupName = groupName;
            GroupType = groupType;
            Id = id;
            LdapGroupDn = ldapGroupDn;
        }
    }
}
