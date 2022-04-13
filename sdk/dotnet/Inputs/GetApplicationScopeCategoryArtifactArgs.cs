// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aquasec.Inputs
{

    public sealed class GetApplicationScopeCategoryArtifactInputArgs : Pulumi.ResourceArgs
    {
        [Input("cfs")]
        private InputList<Inputs.GetApplicationScopeCategoryArtifactCfInputArgs>? _cfs;
        public InputList<Inputs.GetApplicationScopeCategoryArtifactCfInputArgs> Cfs
        {
            get => _cfs ?? (_cfs = new InputList<Inputs.GetApplicationScopeCategoryArtifactCfInputArgs>());
            set => _cfs = value;
        }

        [Input("functions")]
        private InputList<Inputs.GetApplicationScopeCategoryArtifactFunctionInputArgs>? _functions;
        public InputList<Inputs.GetApplicationScopeCategoryArtifactFunctionInputArgs> Functions
        {
            get => _functions ?? (_functions = new InputList<Inputs.GetApplicationScopeCategoryArtifactFunctionInputArgs>());
            set => _functions = value;
        }

        [Input("images")]
        private InputList<Inputs.GetApplicationScopeCategoryArtifactImageInputArgs>? _images;
        public InputList<Inputs.GetApplicationScopeCategoryArtifactImageInputArgs> Images
        {
            get => _images ?? (_images = new InputList<Inputs.GetApplicationScopeCategoryArtifactImageInputArgs>());
            set => _images = value;
        }

        public GetApplicationScopeCategoryArtifactInputArgs()
        {
        }
    }
}