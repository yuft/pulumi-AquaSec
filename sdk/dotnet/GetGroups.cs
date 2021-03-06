// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aquasec
{
    public static class GetGroups
    {
        /// <summary>
        /// The data source `aquasec.getGroups` provides a method to query all groups within the Aqua CSPMgroup database. The fields returned from this query are detailed in the Schema section below.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aquasec = Pulumi.Aquasec;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var groups = Output.Create(Aquasec.GetGroups.InvokeAsync());
        ///         this.FirstGroupName = groups.Apply(groups =&gt; groups.Groups?[0]?.Name);
        ///     }
        /// 
        ///     [Output("firstGroupName")]
        ///     public Output&lt;string&gt; FirstGroupName { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetGroupsResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGroupsResult>("aquasec:index/getGroups:getGroups", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetGroupsResult
    {
        public readonly ImmutableArray<Outputs.GetGroupsGroupResult> Groups;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetGroupsResult(
            ImmutableArray<Outputs.GetGroupsGroupResult> groups,

            string id)
        {
            Groups = groups;
            Id = id;
        }
    }
}
