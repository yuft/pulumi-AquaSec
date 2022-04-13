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
    public sealed class GetUsersUserResult
    {
        public readonly string Email;
        public readonly bool FirstTime;
        public readonly bool IsSuper;
        public readonly string Name;
        public readonly string Plan;
        public readonly string Role;
        public readonly ImmutableArray<string> Roles;
        public readonly string Type;
        public readonly bool UiAccess;
        public readonly string UserId;

        [OutputConstructor]
        private GetUsersUserResult(
            string email,

            bool firstTime,

            bool isSuper,

            string name,

            string plan,

            string role,

            ImmutableArray<string> roles,

            string type,

            bool uiAccess,

            string userId)
        {
            Email = email;
            FirstTime = firstTime;
            IsSuper = isSuper;
            Name = name;
            Plan = plan;
            Role = role;
            Roles = roles;
            Type = type;
            UiAccess = uiAccess;
            UserId = userId;
        }
    }
}