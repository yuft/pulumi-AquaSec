// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aquasec
{
    public static class GetPermissionsSets
    {
        /// <summary>
        /// The data source `aquasec.PermissionsSets` provides a method to query all permissions within the Aqua CSPMThe fields returned from this query are detailed in the Schema section below.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Linq;
        /// using Pulumi;
        /// using Aquasec = Pulumi.Aquasec;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var testpermissionsset = Output.Create(Aquasec.GetPermissionsSets.InvokeAsync());
        ///         this.PermissionsSets = testpermissionsset;
        ///         this.PermissionsSetsNames = 
        ///         {
        ///             testpermissionsset,
        ///         }.Select(__item =&gt; 
        ///         {
        ///             __item.Apply(obj =&gt; obj.PermissionsSets),
        ///         }.Select(__item =&gt; __item?.Name).ToList()).ToList();
        ///     }
        /// 
        ///     [Output("permissionsSets")]
        ///     public Output&lt;string&gt; PermissionsSets { get; set; }
        ///     [Output("permissionsSetsNames")]
        ///     public Output&lt;string&gt; PermissionsSetsNames { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetPermissionsSetsResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPermissionsSetsResult>("aquasec:index/getPermissionsSets:getPermissionsSets", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetPermissionsSetsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetPermissionsSetsPermissionsSetResult> PermissionsSets;

        [OutputConstructor]
        private GetPermissionsSetsResult(
            string id,

            ImmutableArray<Outputs.GetPermissionsSetsPermissionsSetResult> permissionsSets)
        {
            Id = id;
            PermissionsSets = permissionsSets;
        }
    }
}
