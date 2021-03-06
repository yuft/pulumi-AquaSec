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
    public sealed class ApplicationScopeCategory
    {
        public readonly ImmutableArray<Outputs.ApplicationScopeCategoryArtifact> Artifacts;
        public readonly ImmutableArray<Outputs.ApplicationScopeCategoryEntityScope> EntityScopes;
        public readonly ImmutableArray<Outputs.ApplicationScopeCategoryInfrastructure> Infrastructures;
        public readonly ImmutableArray<Outputs.ApplicationScopeCategoryWorkload> Workloads;

        [OutputConstructor]
        private ApplicationScopeCategory(
            ImmutableArray<Outputs.ApplicationScopeCategoryArtifact> artifacts,

            ImmutableArray<Outputs.ApplicationScopeCategoryEntityScope> entityScopes,

            ImmutableArray<Outputs.ApplicationScopeCategoryInfrastructure> infrastructures,

            ImmutableArray<Outputs.ApplicationScopeCategoryWorkload> workloads)
        {
            Artifacts = artifacts;
            EntityScopes = entityScopes;
            Infrastructures = infrastructures;
            Workloads = workloads;
        }
    }
}
