// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aquasec.Inputs
{

    public sealed class ApplicationScopeCategoryArtifactArgs : Pulumi.ResourceArgs
    {
        [Input("cfs")]
        private InputList<Inputs.ApplicationScopeCategoryArtifactCfArgs>? _cfs;
        public InputList<Inputs.ApplicationScopeCategoryArtifactCfArgs> Cfs
        {
            get => _cfs ?? (_cfs = new InputList<Inputs.ApplicationScopeCategoryArtifactCfArgs>());
            set => _cfs = value;
        }

        [Input("functions")]
        private InputList<Inputs.ApplicationScopeCategoryArtifactFunctionArgs>? _functions;
        public InputList<Inputs.ApplicationScopeCategoryArtifactFunctionArgs> Functions
        {
            get => _functions ?? (_functions = new InputList<Inputs.ApplicationScopeCategoryArtifactFunctionArgs>());
            set => _functions = value;
        }

        [Input("images")]
        private InputList<Inputs.ApplicationScopeCategoryArtifactImageArgs>? _images;
        public InputList<Inputs.ApplicationScopeCategoryArtifactImageArgs> Images
        {
            get => _images ?? (_images = new InputList<Inputs.ApplicationScopeCategoryArtifactImageArgs>());
            set => _images = value;
        }

        public ApplicationScopeCategoryArtifactArgs()
        {
        }
    }
}