// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aquasec.Inputs
{

    public sealed class UserSaasLoginGetArgs : Pulumi.ResourceArgs
    {
        [Input("created")]
        public Input<string>? Created { get; set; }

        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        [Input("userId")]
        public Input<int>? UserId { get; set; }

        public UserSaasLoginGetArgs()
        {
        }
    }
}