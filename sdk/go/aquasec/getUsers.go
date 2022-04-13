// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aquasec

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The data source `getUsers` provides a method to query all users within the Aqua users database. The fields returned from this query are detailed in the Schema section below.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aquasec/sdk/go/aquasec"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		users, err := aquasec.GetUsers(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("firstUserName", users.Users[0].Name)
// 		return nil
// 	})
// }
// ```
func GetUsers(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetUsersResult, error) {
	var rv GetUsersResult
	err := ctx.Invoke("aquasec:index/getUsers:getUsers", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getUsers.
type GetUsersResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id    string         `pulumi:"id"`
	Users []GetUsersUser `pulumi:"users"`
}