# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetApplicationScopeResult',
    'AwaitableGetApplicationScopeResult',
    'get_application_scope',
    'get_application_scope_output',
]

@pulumi.output_type
class GetApplicationScopeResult:
    """
    A collection of values returned by getApplicationScope.
    """
    def __init__(__self__, author=None, categories=None, description=None, id=None, name=None, owner_email=None):
        if author and not isinstance(author, str):
            raise TypeError("Expected argument 'author' to be a str")
        pulumi.set(__self__, "author", author)
        if categories and not isinstance(categories, list):
            raise TypeError("Expected argument 'categories' to be a list")
        pulumi.set(__self__, "categories", categories)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if owner_email and not isinstance(owner_email, str):
            raise TypeError("Expected argument 'owner_email' to be a str")
        pulumi.set(__self__, "owner_email", owner_email)

    @property
    @pulumi.getter
    def author(self) -> str:
        return pulumi.get(self, "author")

    @property
    @pulumi.getter
    def categories(self) -> Optional[Sequence['outputs.GetApplicationScopeCategoryResult']]:
        return pulumi.get(self, "categories")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ownerEmail")
    def owner_email(self) -> Optional[str]:
        return pulumi.get(self, "owner_email")


class AwaitableGetApplicationScopeResult(GetApplicationScopeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApplicationScopeResult(
            author=self.author,
            categories=self.categories,
            description=self.description,
            id=self.id,
            name=self.name,
            owner_email=self.owner_email)


def get_application_scope(categories: Optional[Sequence[pulumi.InputType['GetApplicationScopeCategoryArgs']]] = None,
                          description: Optional[str] = None,
                          name: Optional[str] = None,
                          owner_email: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApplicationScopeResult:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_aquasec as aquasec

    default = aquasec.get_application_scope(name="Global")
    pulumi.export("scopes", default)
    ```
    """
    __args__ = dict()
    __args__['categories'] = categories
    __args__['description'] = description
    __args__['name'] = name
    __args__['ownerEmail'] = owner_email
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aquasec:index/getApplicationScope:getApplicationScope', __args__, opts=opts, typ=GetApplicationScopeResult).value

    return AwaitableGetApplicationScopeResult(
        author=__ret__.author,
        categories=__ret__.categories,
        description=__ret__.description,
        id=__ret__.id,
        name=__ret__.name,
        owner_email=__ret__.owner_email)


@_utilities.lift_output_func(get_application_scope)
def get_application_scope_output(categories: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetApplicationScopeCategoryArgs']]]]] = None,
                                 description: Optional[pulumi.Input[Optional[str]]] = None,
                                 name: Optional[pulumi.Input[str]] = None,
                                 owner_email: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetApplicationScopeResult]:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_aquasec as aquasec

    default = aquasec.get_application_scope(name="Global")
    pulumi.export("scopes", default)
    ```
    """
    ...
