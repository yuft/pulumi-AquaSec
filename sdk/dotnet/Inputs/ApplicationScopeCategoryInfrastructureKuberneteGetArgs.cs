// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aquasec.Inputs
{

    public sealed class ApplicationScopeCategoryInfrastructureKuberneteGetArgs : Pulumi.ResourceArgs
    {
        [Input("expression")]
        public Input<string>? Expression { get; set; }

        [Input("variables")]
        private InputList<Inputs.ApplicationScopeCategoryInfrastructureKuberneteVariableGetArgs>? _variables;
        public InputList<Inputs.ApplicationScopeCategoryInfrastructureKuberneteVariableGetArgs> Variables
        {
            get => _variables ?? (_variables = new InputList<Inputs.ApplicationScopeCategoryInfrastructureKuberneteVariableGetArgs>());
            set => _variables = value;
        }

        public ApplicationScopeCategoryInfrastructureKuberneteGetArgs()
        {
        }
    }
}
