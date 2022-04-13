// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aquasec.Inputs
{

    public sealed class ApplicationScopeCategoryWorkloadGetArgs : Pulumi.ResourceArgs
    {
        [Input("cfs")]
        private InputList<Inputs.ApplicationScopeCategoryWorkloadCfGetArgs>? _cfs;
        public InputList<Inputs.ApplicationScopeCategoryWorkloadCfGetArgs> Cfs
        {
            get => _cfs ?? (_cfs = new InputList<Inputs.ApplicationScopeCategoryWorkloadCfGetArgs>());
            set => _cfs = value;
        }

        [Input("kubernetes")]
        private InputList<Inputs.ApplicationScopeCategoryWorkloadKuberneteGetArgs>? _kubernetes;
        public InputList<Inputs.ApplicationScopeCategoryWorkloadKuberneteGetArgs> Kubernetes
        {
            get => _kubernetes ?? (_kubernetes = new InputList<Inputs.ApplicationScopeCategoryWorkloadKuberneteGetArgs>());
            set => _kubernetes = value;
        }

        [Input("os")]
        private InputList<Inputs.ApplicationScopeCategoryWorkloadOGetArgs>? _os;
        public InputList<Inputs.ApplicationScopeCategoryWorkloadOGetArgs> Os
        {
            get => _os ?? (_os = new InputList<Inputs.ApplicationScopeCategoryWorkloadOGetArgs>());
            set => _os = value;
        }

        public ApplicationScopeCategoryWorkloadGetArgs()
        {
        }
    }
}
