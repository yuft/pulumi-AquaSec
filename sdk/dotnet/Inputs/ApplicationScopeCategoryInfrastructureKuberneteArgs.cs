// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aquasec.Inputs
{

    public sealed class ApplicationScopeCategoryInfrastructureKuberneteArgs : Pulumi.ResourceArgs
    {
        [Input("expression")]
        public Input<string>? Expression { get; set; }

        [Input("variables")]
        private InputList<Inputs.ApplicationScopeCategoryInfrastructureKuberneteVariableArgs>? _variables;
        public InputList<Inputs.ApplicationScopeCategoryInfrastructureKuberneteVariableArgs> Variables
        {
            get => _variables ?? (_variables = new InputList<Inputs.ApplicationScopeCategoryInfrastructureKuberneteVariableArgs>());
            set => _variables = value;
        }

        public ApplicationScopeCategoryInfrastructureKuberneteArgs()
        {
        }
    }
}
