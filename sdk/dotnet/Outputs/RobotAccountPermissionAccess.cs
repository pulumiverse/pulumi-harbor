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
        /// <summary>
        /// Eg. `push`, `pull`, `read`, etc. Check [available actions](https://github.com/goharbor/harbor/blob/-/src/common/rbac/const.go).
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// Either `allow` or `deny`. Defaults to `allow`.
        /// </summary>
        public readonly string? Effect;
        /// <summary>
        /// Eg. `repository`, `labels`, etc. Check [available resources](https://github.com/goharbor/harbor/blob/-/src/common/rbac/const.go).
        /// </summary>
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
