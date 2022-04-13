// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aquasec.Inputs
{

    public sealed class ImageAssuranceChecksPerformedGetArgs : Pulumi.ResourceArgs
    {
        [Input("assuranceType")]
        public Input<string>? AssuranceType { get; set; }

        [Input("blocking")]
        public Input<bool>? Blocking { get; set; }

        [Input("control")]
        public Input<string>? Control { get; set; }

        /// <summary>
        /// If DTA was skipped.
        /// </summary>
        [Input("dtaSkipped")]
        public Input<bool>? DtaSkipped { get; set; }

        /// <summary>
        /// The reason why DTA was skipped.
        /// </summary>
        [Input("dtaSkippedReason")]
        public Input<string>? DtaSkippedReason { get; set; }

        [Input("failed")]
        public Input<bool>? Failed { get; set; }

        [Input("policyName")]
        public Input<string>? PolicyName { get; set; }

        public ImageAssuranceChecksPerformedGetArgs()
        {
        }
    }
}
