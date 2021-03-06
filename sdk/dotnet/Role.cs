// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aquasec
{
    /// <summary>
    /// The `aquasec.Role` resource manages your roles within Aqua.
    /// 
    /// The roles created must have permission set and at least one Role Application Scope that is already present within Aqua.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aquasec = Pulumi.Aquasec;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var iaC = new Aquasec.Role("iaC", new Aquasec.RoleArgs
    ///         {
    ///             Description = "RoleIaC",
    ///             Permission = "PermissionIaC",
    ///             RoleName = "RoleIaC",
    ///             Scopes = 
    ///             {
    ///                 "Global",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [AquasecResourceType("aquasec:index/role:Role")]
    public partial class Role : Pulumi.CustomResource
    {
        [Output("author")]
        public Output<string> Author { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("permission")]
        public Output<string> Permission { get; private set; } = null!;

        [Output("roleName")]
        public Output<string> RoleName { get; private set; } = null!;

        [Output("scopes")]
        public Output<ImmutableArray<string>> Scopes { get; private set; } = null!;

        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a Role resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Role(string name, RoleArgs args, CustomResourceOptions? options = null)
            : base("aquasec:index/role:Role", name, args ?? new RoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Role(string name, Input<string> id, RoleState? state = null, CustomResourceOptions? options = null)
            : base("aquasec:index/role:Role", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Role resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Role Get(string name, Input<string> id, RoleState? state = null, CustomResourceOptions? options = null)
        {
            return new Role(name, id, state, options);
        }
    }

    public sealed class RoleArgs : Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("permission", required: true)]
        public Input<string> Permission { get; set; } = null!;

        [Input("roleName", required: true)]
        public Input<string> RoleName { get; set; } = null!;

        [Input("scopes", required: true)]
        private InputList<string>? _scopes;
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public RoleArgs()
        {
        }
    }

    public sealed class RoleState : Pulumi.ResourceArgs
    {
        [Input("author")]
        public Input<string>? Author { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("permission")]
        public Input<string>? Permission { get; set; }

        [Input("roleName")]
        public Input<string>? RoleName { get; set; }

        [Input("scopes")]
        private InputList<string>? _scopes;
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public RoleState()
        {
        }
    }
}
