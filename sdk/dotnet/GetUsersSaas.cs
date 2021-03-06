// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aquasec
{
    public static class GetUsersSaas
    {
        /// <summary>
        /// The data source `aquasec.getUsersSaas` provides a method to query all saas users within the Aqua users management. The fields returned from this query are detailed in the Schema section below.
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
        ///         var users = Output.Create(Aquasec.GetUsers.InvokeAsync());
        ///         this.FirstUserEmail = data.Aquasec_users_saas.Users.Users[0].Email;
        ///     }
        /// 
        ///     [Output("firstUserEmail")]
        ///     public Output&lt;string&gt; FirstUserEmail { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetUsersSaasResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetUsersSaasResult>("aquasec:index/getUsersSaas:getUsersSaas", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetUsersSaasResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetUsersSaasUserResult> Users;

        [OutputConstructor]
        private GetUsersSaasResult(
            string id,

            ImmutableArray<Outputs.GetUsersSaasUserResult> users)
        {
            Id = id;
            Users = users;
        }
    }
}
