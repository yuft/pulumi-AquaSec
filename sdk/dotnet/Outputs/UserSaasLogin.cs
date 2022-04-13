// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aquasec.Outputs
{

    [OutputType]
    public sealed class UserSaasLogin
    {
        public readonly string? Created;
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly int? Id;
        public readonly string? IpAddress;
        public readonly int? UserId;

        [OutputConstructor]
        private UserSaasLogin(
            string? created,

            int? id,

            string? ipAddress,

            int? userId)
        {
            Created = created;
            Id = id;
            IpAddress = ipAddress;
            UserId = userId;
        }
    }
}
