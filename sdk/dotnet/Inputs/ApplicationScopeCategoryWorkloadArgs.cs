// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aquasec.Inputs
{

    public sealed class ApplicationScopeCategoryWorkloadArgs : Pulumi.ResourceArgs
    {
        [Input("cfs")]
        private InputList<Inputs.ApplicationScopeCategoryWorkloadCfArgs>? _cfs;
        public InputList<Inputs.ApplicationScopeCategoryWorkloadCfArgs> Cfs
        {
            get => _cfs ?? (_cfs = new InputList<Inputs.ApplicationScopeCategoryWorkloadCfArgs>());
            set => _cfs = value;
        }

        [Input("kubernetes")]
        private InputList<Inputs.ApplicationScopeCategoryWorkloadKuberneteArgs>? _kubernetes;
        public InputList<Inputs.ApplicationScopeCategoryWorkloadKuberneteArgs> Kubernetes
        {
            get => _kubernetes ?? (_kubernetes = new InputList<Inputs.ApplicationScopeCategoryWorkloadKuberneteArgs>());
            set => _kubernetes = value;
        }

        [Input("os")]
        private InputList<Inputs.ApplicationScopeCategoryWorkloadOArgs>? _os;
        public InputList<Inputs.ApplicationScopeCategoryWorkloadOArgs> Os
        {
            get => _os ?? (_os = new InputList<Inputs.ApplicationScopeCategoryWorkloadOArgs>());
            set => _os = value;
        }

        public ApplicationScopeCategoryWorkloadArgs()
        {
        }
    }
}
