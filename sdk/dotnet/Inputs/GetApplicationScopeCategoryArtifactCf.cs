// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aquasec.Inputs
{

    public sealed class GetApplicationScopeCategoryArtifactCfArgs : Pulumi.InvokeArgs
    {
        [Input("expression", required: true)]
        public string Expression { get; set; } = null!;

        [Input("variables")]
        private List<Inputs.GetApplicationScopeCategoryArtifactCfVariableArgs>? _variables;
        public List<Inputs.GetApplicationScopeCategoryArtifactCfVariableArgs> Variables
        {
            get => _variables ?? (_variables = new List<Inputs.GetApplicationScopeCategoryArtifactCfVariableArgs>());
            set => _variables = value;
        }

        public GetApplicationScopeCategoryArtifactCfArgs()
        {
        }
    }
}
