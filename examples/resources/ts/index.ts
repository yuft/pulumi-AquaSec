import * as pulumi from "@pulumi/pulumi";
import * as aquasec from "@pulumi/aquasec"

const group = new aquasec.Group("IacGroup", {
    name: "IacGroup"
});