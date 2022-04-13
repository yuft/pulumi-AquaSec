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
    public sealed class GetRolesRoleResult
    {
        public readonly string Description;
        public readonly string Name;
        public readonly string Permission;
        public readonly ImmutableArray<string> Scopes;
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetRolesRoleResult(
            string description,

            string name,

            string permission,

            ImmutableArray<string> scopes,

            string updatedAt)
        {
            Description = description;
            Name = name;
            Permission = permission;
            Scopes = scopes;
            UpdatedAt = updatedAt;
        }
    }
}