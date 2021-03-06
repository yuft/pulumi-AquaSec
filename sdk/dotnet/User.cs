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
    /// The `aquasec.User` resource manages your users within Aqua.
    /// 
    /// The users created must have at least one Role that is already present within Aqua.
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
    ///         var iaC = new Aquasec.User("iaC", new Aquasec.UserArgs
    ///         {
    ///             UserId = "IaC",
    ///             Password = @var.Password,
    ///             Roles = 
    ///             {
    ///                 "infrastructure",
    ///             },
    ///             Email = "infrastructure@example.com",
    ///             FirstTime = true,
    ///         });
    ///         // Display name for this user
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [AquasecResourceType("aquasec:index/user:User")]
    public partial class User : Pulumi.CustomResource
    {
        [Output("email")]
        public Output<string?> Email { get; private set; } = null!;

        [Output("firstTime")]
        public Output<bool?> FirstTime { get; private set; } = null!;

        [Output("isSuper")]
        public Output<bool> IsSuper { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        [Output("passwordConfirm")]
        public Output<string?> PasswordConfirm { get; private set; } = null!;

        [Output("plan")]
        public Output<string> Plan { get; private set; } = null!;

        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        [Output("roles")]
        public Output<ImmutableArray<string>> Roles { get; private set; } = null!;

        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        [Output("uiAccess")]
        public Output<bool> UiAccess { get; private set; } = null!;

        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a User resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public User(string name, UserArgs args, CustomResourceOptions? options = null)
            : base("aquasec:index/user:User", name, args ?? new UserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private User(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
            : base("aquasec:index/user:User", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing User resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static User Get(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
        {
            return new User(name, id, state, options);
        }
    }

    public sealed class UserArgs : Pulumi.ResourceArgs
    {
        [Input("email")]
        public Input<string>? Email { get; set; }

        [Input("firstTime")]
        public Input<bool>? FirstTime { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        [Input("passwordConfirm")]
        public Input<string>? PasswordConfirm { get; set; }

        [Input("roles", required: true)]
        private InputList<string>? _roles;
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public UserArgs()
        {
        }
    }

    public sealed class UserState : Pulumi.ResourceArgs
    {
        [Input("email")]
        public Input<string>? Email { get; set; }

        [Input("firstTime")]
        public Input<bool>? FirstTime { get; set; }

        [Input("isSuper")]
        public Input<bool>? IsSuper { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("passwordConfirm")]
        public Input<string>? PasswordConfirm { get; set; }

        [Input("plan")]
        public Input<string>? Plan { get; set; }

        [Input("role")]
        public Input<string>? Role { get; set; }

        [Input("roles")]
        private InputList<string>? _roles;
        public InputList<string> Roles
        {
            get => _roles ?? (_roles = new InputList<string>());
            set => _roles = value;
        }

        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("uiAccess")]
        public Input<bool>? UiAccess { get; set; }

        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public UserState()
        {
        }
    }
}
