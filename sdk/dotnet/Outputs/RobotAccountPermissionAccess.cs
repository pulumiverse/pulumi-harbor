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
    public sealed class RobotAccountPermissionAccess
    {
        public readonly string Action;
        public readonly string? Effect;
        public readonly string Resource;

        [OutputConstructor]
        private RobotAccountPermissionAccess(
            string action,

            string? effect,

            string resource)
        {
            Action = action;
            Effect = effect;
            Resource = resource;
        }
    }
}
