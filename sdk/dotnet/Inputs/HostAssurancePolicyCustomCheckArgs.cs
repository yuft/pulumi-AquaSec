// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aquasec.Inputs
{

    public sealed class HostAssurancePolicyCustomCheckArgs : Pulumi.ResourceArgs
    {
        [Input("author")]
        public Input<string>? Author { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("engine")]
        public Input<string>? Engine { get; set; }

        [Input("lastModified")]
        public Input<int>? LastModified { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        [Input("scriptId")]
        public Input<string>? ScriptId { get; set; }

        [Input("severity")]
        public Input<string>? Severity { get; set; }

        [Input("snippet")]
        public Input<string>? Snippet { get; set; }

        public HostAssurancePolicyCustomCheckArgs()
        {
        }
    }
}
