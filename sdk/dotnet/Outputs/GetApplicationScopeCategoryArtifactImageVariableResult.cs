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
    public sealed class GetApplicationScopeCategoryArtifactImageVariableResult
    {
        public readonly string? Attribute;
        public readonly string? Value;

        [OutputConstructor]
        private GetApplicationScopeCategoryArtifactImageVariableResult(
            string? attribute,

            string? value)
        {
            Attribute = attribute;
            Value = value;
        }
    }
}
