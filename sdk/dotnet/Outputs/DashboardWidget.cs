// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sentry.Outputs
{

    [OutputType]
    public sealed class DashboardWidget
    {
        public readonly string DisplayType;
        public readonly string? Id;
        public readonly string? Interval;
        public readonly Outputs.DashboardWidgetLayout Layout;
        public readonly int? Limit;
        public readonly ImmutableArray<Outputs.DashboardWidgetQuery> Queries;
        public readonly string Title;
        public readonly string? WidgetType;

        [OutputConstructor]
        private DashboardWidget(
            string displayType,

            string? id,

            string? interval,

            Outputs.DashboardWidgetLayout layout,

            int? limit,

            ImmutableArray<Outputs.DashboardWidgetQuery> queries,

            string title,

            string? widgetType)
        {
            DisplayType = displayType;
            Id = id;
            Interval = interval;
            Layout = layout;
            Limit = limit;
            Queries = queries;
            Title = title;
            WidgetType = widgetType;
        }
    }
}
