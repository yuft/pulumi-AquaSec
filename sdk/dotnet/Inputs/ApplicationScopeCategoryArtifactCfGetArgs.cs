// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aquasec.Inputs
{

    public sealed class ApplicationScopeCategoryArtifactCfGetArgs : Pulumi.ResourceArgs
    {
        [Input("expression")]
        public Input<string>? Expression { get; set; }

        [Input("variables")]
        private InputList<Inputs.ApplicationScopeCategoryArtifactCfVariableGetArgs>? _variables;
        public InputList<Inputs.ApplicationScopeCategoryArtifactCfVariableGetArgs> Variables
        {
            get => _variables ?? (_variables = new InputList<Inputs.ApplicationScopeCategoryArtifactCfVariableGetArgs>());
            set => _variables = value;
        }

        public ApplicationScopeCategoryArtifactCfGetArgs()
        {
        }
    }
}
