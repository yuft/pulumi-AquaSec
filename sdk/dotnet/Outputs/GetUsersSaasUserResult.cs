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
    public sealed class GetUsersSaasUserResult
    {
        public readonly bool AccountAdmin;
        public readonly bool Confirmed;
        public readonly string Created;
        public readonly ImmutableArray<string> CspRoles;
        public readonly bool Dashboard;
        public readonly string Email;
        public readonly ImmutableArray<Outputs.GetUsersSaasUserGroupResult> Groups;
        public readonly ImmutableArray<Outputs.GetUsersSaasUserLoginResult> Logins;
        public readonly bool Multiaccount;
        public readonly bool PasswordReset;
        public readonly string Provider;
        public readonly bool SendAnnouncements;
        public readonly bool SendNewPlugins;
        public readonly bool SendNewRisks;
        public readonly bool SendScanResults;
        public readonly string UserId;

        [OutputConstructor]
        private GetUsersSaasUserResult(
            bool accountAdmin,

            bool confirmed,

            string created,

            ImmutableArray<string> cspRoles,

            bool dashboard,

            string email,

            ImmutableArray<Outputs.GetUsersSaasUserGroupResult> groups,

            ImmutableArray<Outputs.GetUsersSaasUserLoginResult> logins,

            bool multiaccount,

            bool passwordReset,

            string provider,

            bool sendAnnouncements,

            bool sendNewPlugins,

            bool sendNewRisks,

            bool sendScanResults,

            string userId)
        {
            AccountAdmin = accountAdmin;
            Confirmed = confirmed;
            Created = created;
            CspRoles = cspRoles;
            Dashboard = dashboard;
            Email = email;
            Groups = groups;
            Logins = logins;
            Multiaccount = multiaccount;
            PasswordReset = passwordReset;
            Provider = provider;
            SendAnnouncements = sendAnnouncements;
            SendNewPlugins = sendNewPlugins;
            SendNewRisks = sendNewRisks;
            SendScanResults = sendScanResults;
            UserId = userId;
        }
    }
}