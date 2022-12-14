// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getOrganization(args: GetOrganizationArgs, opts?: pulumi.InvokeOptions): Promise<GetOrganizationResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("sentry:index/getOrganization:getOrganization", {
        "slug": args.slug,
    }, opts);
}

/**
 * A collection of arguments for invoking getOrganization.
 */
export interface GetOrganizationArgs {
    slug: string;
}

/**
 * A collection of values returned by getOrganization.
 */
export interface GetOrganizationResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly internalId: string;
    readonly name: string;
    readonly slug: string;
}

export function getOrganizationOutput(args: GetOrganizationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOrganizationResult> {
    return pulumi.output(args).apply(a => getOrganization(a, opts))
}

/**
 * A collection of arguments for invoking getOrganization.
 */
export interface GetOrganizationOutputArgs {
    slug: pulumi.Input<string>;
}
