package env

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/opencost/opencost/pkg/log"
	"github.com/opencost/opencost/pkg/util/timeutil"
)

const (
	AWSAccessKeyIDEnvVar     = "AWS_ACCESS_KEY_ID"
	AWSAccessKeySecretEnvVar = "AWS_SECRET_ACCESS_KEY"
	AWSClusterIDEnvVar       = "AWS_CLUSTER_ID"
	AWSPricingURL            = "AWS_PRICING_URL"

	AlibabaAccessKeyIDEnvVar     = "ALIBABA_ACCESS_KEY_ID"
	AlibabaAccessKeySecretEnvVar = "ALIBABA_SECRET_ACCESS_KEY"

	AzureOfferIDEnvVar        = "AZURE_OFFER_ID"
	AzureBillingAccountEnvVar = "AZURE_BILLING_ACCOUNT"

	KubecostNamespaceEnvVar        = "KUBECOST_NAMESPACE"
	PodNameEnvVar                  = "POD_NAME"
	ClusterIDEnvVar                = "CLUSTER_ID"
	ClusterProfileEnvVar           = "CLUSTER_PROFILE"
	PrometheusServerEndpointEnvVar = "PROMETHEUS_SERVER_ENDPOINT"
	MaxQueryConcurrencyEnvVar      = "MAX_QUERY_CONCURRENCY"
	QueryLoggingFileEnvVar         = "QUERY_LOGGING_FILE"
	RemoteEnabledEnvVar            = "REMOTE_WRITE_ENABLED"
	RemotePWEnvVar                 = "REMOTE_WRITE_PASSWORD"
	SQLAddressEnvVar               = "SQL_ADDRESS"
	UseCSVProviderEnvVar           = "USE_CSV_PROVIDER"
	CSVRegionEnvVar                = "CSV_REGION"
	CSVEndpointEnvVar              = "CSV_ENDPOINT"
	CSVPathEnvVar                  = "CSV_PATH"
	ConfigPathEnvVar               = "CONFIG_PATH"
	CloudProviderAPIKeyEnvVar      = "CLOUD_PROVIDER_API_KEY"
	DisableAggregateCostModelCache = "DISABLE_AGGREGATE_COST_MODEL_CACHE"

	EmitPodAnnotationsMetricEnvVar       = "EMIT_POD_ANNOTATIONS_METRIC"
	EmitNamespaceAnnotationsMetricEnvVar = "EMIT_NAMESPACE_ANNOTATIONS_METRIC"

	EmitKsmV1MetricsEnvVar = "EMIT_KSM_V1_METRICS"
	EmitKsmV1MetricsOnly   = "EMIT_KSM_V1_METRICS_ONLY"

	ThanosEnabledEnvVar      = "THANOS_ENABLED"
	ThanosQueryUrlEnvVar     = "THANOS_QUERY_URL"
	ThanosOffsetEnvVar       = "THANOS_QUERY_OFFSET"
	ThanosMaxSourceResEnvVar = "THANOS_MAX_SOURCE_RESOLUTION"

	LogCollectionEnabledEnvVar    = "LOG_COLLECTION_ENABLED"
	ProductAnalyticsEnabledEnvVar = "PRODUCT_ANALYTICS_ENABLED"
	ErrorReportingEnabledEnvVar   = "ERROR_REPORTING_ENABLED"
	ValuesReportingEnabledEnvVar  = "VALUES_REPORTING_ENABLED"

	DBBasicAuthUsername = "DB_BASIC_AUTH_USERNAME"
	DBBasicAuthPassword = "DB_BASIC_AUTH_PW"
	DBBearerToken       = "DB_BEARER_TOKEN"

	MultiClusterBasicAuthUsername = "MC_BASIC_AUTH_USERNAME"
	MultiClusterBasicAuthPassword = "MC_BASIC_AUTH_PW"
	MultiClusterBearerToken       = "MC_BEARER_TOKEN"

	InsecureSkipVerify = "INSECURE_SKIP_VERIFY"

	KubeConfigPathEnvVar = "KUBECONFIG_PATH"

	UTCOffsetEnvVar                  = "UTC_OFFSET"
	CurrentClusterIdFilterEnabledVar = "CURRENT_CLUSTER_ID_FILTER_ENABLED"

	CacheWarmingEnabledEnvVar            = "CACHE_WARMING_ENABLED"
	ETLEnabledEnvVar                     = "ETL_ENABLED"
	ETLMaxPrometheusQueryDurationMinutes = "ETL_MAX_PROMETHEUS_QUERY_DURATION_MINUTES"
	ETLResolutionSeconds                 = "ETL_RESOLUTION_SECONDS"
	LegacyExternalAPIDisabledVar         = "LEGACY_EXTERNAL_API_DISABLED"

	PromClusterIDLabelEnvVar = "PROM_CLUSTER_ID_LABEL"

	PricingConfigmapName  = "PRICING_CONFIGMAP_NAME"
	MetricsConfigmapName  = "METRICS_CONFIGMAP_NAME"
	KubecostJobNameEnvVar = "KUBECOST_JOB_NAME"

	KubecostConfigBucketEnvVar    = "KUBECOST_CONFIG_BUCKET"
	ClusterInfoFileEnabledEnvVar  = "CLUSTER_INFO_FILE_ENABLED"
	ClusterCacheFileEnabledEnvVar = "CLUSTER_CACHE_FILE_ENABLED"

	PrometheusQueryOffsetEnvVar                 = "PROMETHEUS_QUERY_OFFSET"
	PrometheusRetryOnRateLimitResponseEnvVar    = "PROMETHEUS_RETRY_ON_RATE_LIMIT"
	PrometheusRetryOnRateLimitMaxRetriesEnvVar  = "PROMETHEUS_RETRY_ON_RATE_LIMIT_MAX_RETRIES"
	PrometheusRetryOnRateLimitDefaultWaitEnvVar = "PROMETHEUS_RETRY_ON_RATE_LIMIT_DEFAULT_WAIT"

	PrometheusHeaderXScopeOrgIdEnvVar = "PROMETHEUS_HEADER_X_SCOPE_ORGID"

	IngestPodUIDEnvVar = "INGEST_POD_UID"

	ETLReadOnlyMode = "ETL_READ_ONLY"

	AllocationNodeLabelsEnabled     = "ALLOCATION_NODE_LABELS_ENABLED"
	AllocationNodeLabelsIncludeList = "ALLOCATION_NODE_LABELS_INCLUDE_LIST"

	regionOverrideList = "REGION_OVERRIDE_LIST"

	ExportCSVFile       = "EXPORT_CSV_FILE"
	ExportCSVLabelsList = "EXPORT_CSV_LABELS_LIST"
	ExportCSVLabelsAll  = "EXPORT_CSV_LABELS_ALL"
)

const DefaultConfigMountPath = "/var/configs"

var offsetRegex = regexp.MustCompile(`^(\+|-)(\d\d):(\d\d)$`)

func IsETLReadOnlyMode() bool {
	return GetBool(ETLReadOnlyMode, false)
}

func GetExportCSVFile() string {
	return Get(ExportCSVFile, "")
}

func GetExportCSVLabelsAll() bool {
	return GetBool(ExportCSVLabelsAll, false)
}

func GetExportCSVLabelsList() []string {
	return GetList(ExportCSVLabelsList, ",")
}

// GetKubecostConfigBucket returns a file location for a mounted bucket configuration which is used to store
// a subset of kubecost configurations that require sharing via remote storage.
func GetKubecostConfigBucket() string {
	return Get(KubecostConfigBucketEnvVar, "")
}

// IsClusterInfoFileEnabled returns true if the cluster info is read from a file or pulled from the local
// cloud provider and kubernetes.
func IsClusterInfoFileEnabled() bool {
	return GetBool(ClusterInfoFileEnabledEnvVar, false)
}

// IsClusterCacheFileEnabled returns true if the kubernetes cluster data is read from a file or pulled from the local
// kubernetes API.
func IsClusterCacheFileEnabled() bool {
	return GetBool(ClusterCacheFileEnabledEnvVar, false)
}

// IsPrometheusRetryOnRateLimitResponse will attempt to retry if a 429 response is received OR a 400 with a body containing
// ThrottleException (common in AWS services like AMP)
func IsPrometheusRetryOnRateLimitResponse() bool {
	return GetBool(PrometheusRetryOnRateLimitResponseEnvVar, true)
}

// GetPrometheusRetryOnRateLimitMaxRetries returns the maximum number of retries that should be attempted prior to failing.
// Only used if IsPrometheusRetryOnRateLimitResponse() is true.
func GetPrometheusRetryOnRateLimitMaxRetries() int {
	return GetInt(PrometheusRetryOnRateLimitMaxRetriesEnvVar, 5)
}

// GetPrometheusRetryOnRateLimitDefaultWait returns the default wait time for a retriable rate limit response without a
// Retry-After header.
func GetPrometheusRetryOnRateLimitDefaultWait() time.Duration {
	return GetDuration(PrometheusRetryOnRateLimitDefaultWaitEnvVar, 100*time.Millisecond)
}

// GetPrometheusHeaderXScopeOrgId returns the default value for X-Scope-OrgID header used for requests in Mimir/Cortex-Tenant API.
// To use Mimir(or Cortex-Tenant) instead of Prometheus add variable from cluster settings:
// "PROMETHEUS_HEADER_X_SCOPE_ORGID": "my-cluster-name"
// Then set Prometheus URL to prometheus API endpoint:
// "PROMETHEUS_SERVER_ENDPOINT": "http://mimir-url/prometheus/"
func GetPrometheusHeaderXScopeOrgId() string {
	return Get(PrometheusHeaderXScopeOrgIdEnvVar, "")
}

// GetPrometheusQueryOffset returns the time.Duration to offset all prometheus queries by. NOTE: This env var is applied
// to all non-range queries made via our query context. This should only be applied when there is a significant delay in
// data arriving in the target prom db. For example, if supplying a thanos or cortex querier for the prometheus server, using
// a 3h offset will ensure that current time = current time - 3h.
//
// This offset is NOT the same as the GetThanosOffset() option, as that is only applied to queries made specifically targeting
// thanos. This offset is applied globally.
func GetPrometheusQueryOffset() time.Duration {
	offset := Get(PrometheusQueryOffsetEnvVar, "")
	if offset == "" {
		return 0
	}

	dur, err := timeutil.ParseDuration(offset)
	if err != nil {
		return 0
	}
	return dur
}

func GetPricingConfigmapName() string {
	return Get(PricingConfigmapName, "pricing-configs")
}

func GetMetricsConfigmapName() string {
	return Get(MetricsConfigmapName, "metrics-config")
}

// IsEmitNamespaceAnnotationsMetric returns true if cost-model is configured to emit the kube_namespace_annotations metric
// containing the namespace annotations
func IsEmitNamespaceAnnotationsMetric() bool {
	return GetBool(EmitNamespaceAnnotationsMetricEnvVar, false)
}

// IsEmitPodAnnotationsMetric returns true if cost-model is configured to emit the kube_pod_annotations metric containing
// pod annotations.
func IsEmitPodAnnotationsMetric() bool {
	return GetBool(EmitPodAnnotationsMetricEnvVar, false)
}

// IsEmitKsmV1Metrics returns true if cost-model is configured to emit all necessary KSM v1
// metrics that were removed in KSM v2
func IsEmitKsmV1Metrics() bool {
	return GetBool(EmitKsmV1MetricsEnvVar, true)
}

func IsEmitKsmV1MetricsOnly() bool {
	return GetBool(EmitKsmV1MetricsOnly, false)
}

// GetAWSAccessKeyID returns the environment variable value for AWSAccessKeyIDEnvVar which represents
// the AWS access key for authentication
func GetAWSAccessKeyID() string {
	awsAccessKeyID := Get(AWSAccessKeyIDEnvVar, "")
	// If the sample nil service key name is set, zero it out so that it is not
	// misinterpreted as a real service key.
	if awsAccessKeyID == "AKIXXX" {
		awsAccessKeyID = ""
	}
	return awsAccessKeyID
}

// GetAWSAccessKeySecret returns the environment variable value for AWSAccessKeySecretEnvVar which represents
// the AWS access key secret for authentication
func GetAWSAccessKeySecret() string {
	return Get(AWSAccessKeySecretEnvVar, "")
}

// GetAWSClusterID returns the environment variable value for AWSClusterIDEnvVar which represents
// an AWS specific cluster identifier.
func GetAWSClusterID() string {
	return Get(AWSClusterIDEnvVar, "")
}

// GetAWSPricingURL returns an optional alternative URL to fetch AWS pricing data from; for use in airgapped environments
func GetAWSPricingURL() string {
	return Get(AWSPricingURL, "")
}

// GetAlibabaAccessKeyID returns the environment variable value for AlibabaAccessKeyIDEnvVar which represents
// the Alibaba access key for authentication
func GetAlibabaAccessKeyID() string {
	return Get(AlibabaAccessKeyIDEnvVar, "")
}

// GetAlibabaAccessKeySecret returns the environment variable value for AlibabaAccessKeySecretEnvVar which represents
// the Alibaba access key secret for authentication
func GetAlibabaAccessKeySecret() string {
	return Get(AlibabaAccessKeySecretEnvVar, "")
}

// GetAzureOfferID returns the environment variable value for AzureOfferIDEnvVar which represents
// the Azure offer ID for determining prices.
func GetAzureOfferID() string {
	return Get(AzureOfferIDEnvVar, "")
}

// GetAzureBillingAccount returns the environment variable value for
// AzureBillingAccountEnvVar which represents the Azure billing
// account for determining prices. If this is specified
// customer-specific prices will be downloaded from the consumption
// price sheet API.
func GetAzureBillingAccount() string {
	return Get(AzureBillingAccountEnvVar, "")
}

// GetKubecostNamespace returns the environment variable value for KubecostNamespaceEnvVar which
// represents the namespace the cost model exists in.
func GetKubecostNamespace() string {
	return Get(KubecostNamespaceEnvVar, "kubecost")
}

// GetPodName returns the name of the current running pod. If this environment variable is not set,
// empty string is returned.
func GetPodName() string {
	return Get(PodNameEnvVar, "")
}

// GetClusterProfile returns the environment variable value for ClusterProfileEnvVar which
// represents the cluster profile configured for
func GetClusterProfile() string {
	return Get(ClusterProfileEnvVar, "development")
}

// GetClusterID returns the environment variable value for ClusterIDEnvVar which represents the
// configurable identifier used for multi-cluster metric emission.
func GetClusterID() string {
	return Get(ClusterIDEnvVar, "")
}

// GetPromClusterFilter returns environment variable value CurrentClusterIdFilterEnabledVar which
// represents additional prometheus filter for all metrics for current cluster id
func GetPromClusterFilter() string {
	if GetBool(CurrentClusterIdFilterEnabledVar, false) {
		return fmt.Sprintf("%s=\"%s\"", GetPromClusterLabel(), GetClusterID())
	}
	return ""
}

// GetPrometheusServerEndpoint returns the environment variable value for PrometheusServerEndpointEnvVar which
// represents the prometheus server endpoint used to execute prometheus queries.
func GetPrometheusServerEndpoint() string {
	return Get(PrometheusServerEndpointEnvVar, "")
}

func GetInsecureSkipVerify() bool {
	return GetBool(InsecureSkipVerify, false)
}

// IsAggregateCostModelCacheDisabled returns the environment variable value for DisableAggregateCostModelCache which
// will inform the aggregator on whether to load cached data. Defaults to false
func IsAggregateCostModelCacheDisabled() bool {
	return GetBool(DisableAggregateCostModelCache, false)
}

// IsRemoteEnabled returns the environment variable value for RemoteEnabledEnvVar which represents whether
// or not remote write is enabled for prometheus for use with SQL backed persistent storage.
func IsRemoteEnabled() bool {
	return GetBool(RemoteEnabledEnvVar, false)
}

// GetRemotePW returns the environment variable value for RemotePWEnvVar which represents the remote
// persistent storage password.
func GetRemotePW() string {
	return Get(RemotePWEnvVar, "")
}

// GetSQLAddress returns the environment variable value for SQLAddressEnvVar which represents the SQL
// database address used with remote persistent storage.
func GetSQLAddress() string {
	return Get(SQLAddressEnvVar, "")
}

// IsUseCSVProvider returns the environment variable value for UseCSVProviderEnvVar which represents
// whether or not the use of a CSV cost provider is enabled.
func IsUseCSVProvider() bool {
	return GetBool(UseCSVProviderEnvVar, false)
}

// GetCSVRegion returns the environment variable value for CSVRegionEnvVar which represents the
// region configured for a CSV provider.
func GetCSVRegion() string {
	return Get(CSVRegionEnvVar, "")
}

// GetCSVEndpoint returns the environment variable value for CSVEndpointEnvVar which represents the
// endpoint configured for a S3 CSV provider another than AWS S3.
func GetCSVEndpoint() string {
	return Get(CSVEndpointEnvVar, "")
}

// GetCSVPath returns the environment variable value for CSVPathEnvVar which represents the key path
// configured for a CSV provider.
func GetCSVPath() string {
	return Get(CSVPathEnvVar, "")
}

// GetCostAnalyzerVolumeMountPath is an alias of GetConfigPath, which returns the mount path for the
// Cost Analyzer volume, which stores configs, persistent data, etc.
func GetCostAnalyzerVolumeMountPath() string {
	return GetConfigPathWithDefault(DefaultConfigMountPath)
}

// GetConfigPath returns the environment variable value for ConfigPathEnvVar which represents the cost
// model configuration path
func GetConfigPathWithDefault(defaultValue string) string {
	return Get(ConfigPathEnvVar, defaultValue)
}

// GetCloudProviderAPI returns the environment variable value for CloudProviderAPIEnvVar which represents
// the API key provided for the cloud provider.
func GetCloudProviderAPIKey() string {
	return Get(CloudProviderAPIKeyEnvVar, "")
}

// IsThanosEnabled returns the environment variable value for ThanosEnabledEnvVar which represents whether
// or not thanos is enabled.
func IsThanosEnabled() bool {
	return GetBool(ThanosEnabledEnvVar, false)
}

// GetThanosQueryUrl returns the environment variable value for ThanosQueryUrlEnvVar which represents the
// target query endpoint for hitting thanos.
func GetThanosQueryUrl() string {
	return Get(ThanosQueryUrlEnvVar, "")
}

// GetThanosOffset returns the environment variable value for ThanosOffsetEnvVar which represents the total
// amount of time to offset all queries made to thanos.
func GetThanosOffset() string {
	return Get(ThanosOffsetEnvVar, "3h")
}

// GetThanosMaxSourceResolution returns the environment variable value for ThanosMaxSourceResEnvVar which represents
// the max source resolution to use when querying thanos.
func GetThanosMaxSourceResolution() string {
	res := Get(ThanosMaxSourceResEnvVar, "raw")

	switch res {
	case "raw":
		return "0s"
	case "0s":
		fallthrough
	case "5m":
		fallthrough
	case "1h":
		return res
	default:
		return "0s"
	}
}

// IsLogCollectionEnabled returns the environment variable value for LogCollectionEnabledEnvVar which represents
// whether or not log collection has been enabled for kubecost deployments.
func IsLogCollectionEnabled() bool {
	return GetBool(LogCollectionEnabledEnvVar, true)
}

// IsProductAnalyticsEnabled returns the environment variable value for ProductAnalyticsEnabledEnvVar
func IsProductAnalyticsEnabled() bool {
	return GetBool(ProductAnalyticsEnabledEnvVar, true)
}

// IsErrorReportingEnabled returns the environment variable value for ErrorReportingEnabledEnvVar
func IsErrorReportingEnabled() bool {
	return GetBool(ErrorReportingEnabledEnvVar, true)
}

// IsValuesReportingEnabled returns the environment variable value for ValuesReportingEnabledEnvVar
func IsValuesReportingEnabled() bool {
	return GetBool(ValuesReportingEnabledEnvVar, true)
}

// GetMaxQueryConcurrency returns the environment variable value for MaxQueryConcurrencyEnvVar
func GetMaxQueryConcurrency() int {
	return GetInt(MaxQueryConcurrencyEnvVar, 5)
}

// GetQueryLoggingFile returns a file location if query logging is enabled. Otherwise, empty string
func GetQueryLoggingFile() string {
	return Get(QueryLoggingFileEnvVar, "")
}

func GetDBBasicAuthUsername() string {
	return Get(DBBasicAuthUsername, "")
}

func GetDBBasicAuthUserPassword() string {
	return Get(DBBasicAuthPassword, "")

}

func GetDBBearerToken() string {
	return Get(DBBearerToken, "")
}

// GetMultiClusterBasicAuthUsername returns the environment variable value for MultiClusterBasicAuthUsername
func GetMultiClusterBasicAuthUsername() string {
	return Get(MultiClusterBasicAuthUsername, "")
}

// GetMultiClusterBasicAuthPassword returns the environment variable value for MultiClusterBasicAuthPassword
func GetMultiClusterBasicAuthPassword() string {
	return Get(MultiClusterBasicAuthPassword, "")
}

func GetMultiClusterBearerToken() string {
	return Get(MultiClusterBearerToken, "")
}

// GetKubeConfigPath returns the environment variable value for KubeConfigPathEnvVar
func GetKubeConfigPath() string {
	return Get(KubeConfigPathEnvVar, "")
}

// GetUTCOffset returns the environment variable value for UTCOffset
func GetUTCOffset() string {
	return Get(UTCOffsetEnvVar, "")
}

// GetParsedUTCOffset returns the duration of the configured UTC offset
func GetParsedUTCOffset() time.Duration {
	offset := time.Duration(0)

	if offsetStr := GetUTCOffset(); offsetStr != "" {
		match := offsetRegex.FindStringSubmatch(offsetStr)
		if match == nil {
			log.Warnf("Illegal UTC offset: %s", offsetStr)
			return offset
		}

		sig := 1
		if match[1] == "-" {
			sig = -1
		}

		hrs64, _ := strconv.ParseInt(match[2], 10, 64)
		hrs := sig * int(hrs64)

		mins64, _ := strconv.ParseInt(match[3], 10, 64)
		mins := sig * int(mins64)

		offset = time.Duration(hrs)*time.Hour + time.Duration(mins)
	}

	return offset
}

// GetKubecostJobName returns the environment variable value for KubecostJobNameEnvVar
func GetKubecostJobName() string {
	return Get(KubecostJobNameEnvVar, "kubecost")
}

func IsCacheWarmingEnabled() bool {
	return GetBool(CacheWarmingEnabledEnvVar, true)
}

func IsETLEnabled() bool {
	return GetBool(ETLEnabledEnvVar, true)
}

func GetETLMaxPrometheusQueryDuration() time.Duration {
	dayMins := 60 * 24
	mins := time.Duration(GetInt64(ETLMaxPrometheusQueryDurationMinutes, int64(dayMins)))
	return mins * time.Minute
}

// GetETLResolution determines the resolution of ETL queries. The smaller the
// duration, the higher the resolution; the higher the resolution, the more
// accurate the query results, but the more computationally expensive.
func GetETLResolution() time.Duration {
	// Use the configured ETL resolution, or default to
	// 5m (i.e. 300s)
	secs := time.Duration(GetInt64(ETLResolutionSeconds, 300))
	return secs * time.Second
}

func LegacyExternalCostsAPIDisabled() bool {
	return GetBool(LegacyExternalAPIDisabledVar, false)
}

// GetPromClusterLabel returns the environment variable value for PromClusterIDLabel
func GetPromClusterLabel() string {
	return Get(PromClusterIDLabelEnvVar, "cluster_id")
}

// IsIngestingPodUID returns the env variable from ingestPodUID, which alters the
// contents of podKeys in Allocation
func IsIngestingPodUID() bool {
	return GetBool(IngestPodUIDEnvVar, false)
}

func GetAllocationNodeLabelsEnabled() bool {
	return GetBool(AllocationNodeLabelsEnabled, true)
}

var defaultAllocationNodeLabelsIncludeList []string = []string{
	"cloud.google.com/gke-nodepool",
	"eks.amazonaws.com/nodegroup",
	"kubernetes.azure.com/agentpool",
	"node.kubernetes.io/instance-type",
	"topology.kubernetes.io/region",
	"topology.kubernetes.io/zone",
}

func GetAllocationNodeLabelsIncludeList() []string {
	// If node labels are not enabled, return an empty list.
	if !GetAllocationNodeLabelsEnabled() {
		return []string{}
	}

	list := GetList(AllocationNodeLabelsIncludeList, ",")

	// If node labels are enabled, but the white list is empty, use defaults.
	if len(list) == 0 {
		return defaultAllocationNodeLabelsIncludeList
	}

	return list
}

func GetRegionOverrideList() []string {
	regionList := GetList(regionOverrideList, ",")

	if regionList == nil {
		return []string{}
	}

	return regionList
}
