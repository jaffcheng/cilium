// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/bootstrap/v2/bootstrap.proto

package v2

import (
	fmt "fmt"
	v23 "github.com/cilium/cilium/pkg/envoy/envoy/api/v2"
	auth "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/auth"
	core "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"
	v2 "github.com/cilium/cilium/pkg/envoy/envoy/config/metrics/v2"
	v2alpha "github.com/cilium/cilium/pkg/envoy/envoy/config/overload/v2alpha"
	v22 "github.com/cilium/cilium/pkg/envoy/envoy/config/ratelimit/v2"
	v21 "github.com/cilium/cilium/pkg/envoy/envoy/config/trace/v2"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/lyft/protoc-gen-validate/validate"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Bootstrap :ref:`configuration overview <config_overview_v2_bootstrap>`.
type Bootstrap struct {
	// Node identity to present to the management server and for instance
	// identification purposes (e.g. in generated headers).
	Node *core.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// Statically specified resources.
	StaticResources *Bootstrap_StaticResources `protobuf:"bytes,2,opt,name=static_resources,json=staticResources,proto3" json:"static_resources,omitempty"`
	// xDS configuration sources.
	DynamicResources *Bootstrap_DynamicResources `protobuf:"bytes,3,opt,name=dynamic_resources,json=dynamicResources,proto3" json:"dynamic_resources,omitempty"`
	// Configuration for the cluster manager which owns all upstream clusters
	// within the server.
	ClusterManager *ClusterManager `protobuf:"bytes,4,opt,name=cluster_manager,json=clusterManager,proto3" json:"cluster_manager,omitempty"`
	// Health discovery service config option.
	// (:ref:`core.ApiConfigSource <envoy_api_msg_core.ApiConfigSource>`)
	HdsConfig *core.ApiConfigSource `protobuf:"bytes,14,opt,name=hds_config,json=hdsConfig,proto3" json:"hds_config,omitempty"`
	// Optional file system path to search for startup flag files.
	FlagsPath string `protobuf:"bytes,5,opt,name=flags_path,json=flagsPath,proto3" json:"flags_path,omitempty"`
	// Optional set of stats sinks.
	StatsSinks []*v2.StatsSink `protobuf:"bytes,6,rep,name=stats_sinks,json=statsSinks,proto3" json:"stats_sinks,omitempty"`
	// Configuration for internal processing of stats.
	StatsConfig *v2.StatsConfig `protobuf:"bytes,13,opt,name=stats_config,json=statsConfig,proto3" json:"stats_config,omitempty"`
	// Optional duration between flushes to configured stats sinks. For
	// performance reasons Envoy latches counters and only flushes counters and
	// gauges at a periodic interval. If not specified the default is 5000ms (5
	// seconds).
	StatsFlushInterval *duration.Duration `protobuf:"bytes,7,opt,name=stats_flush_interval,json=statsFlushInterval,proto3" json:"stats_flush_interval,omitempty"`
	// Optional watchdog configuration.
	Watchdog *Watchdog `protobuf:"bytes,8,opt,name=watchdog,proto3" json:"watchdog,omitempty"`
	// Configuration for an external tracing provider. If not specified, no
	// tracing will be performed.
	Tracing *v21.Tracing `protobuf:"bytes,9,opt,name=tracing,proto3" json:"tracing,omitempty"`
	// Configuration for an external rate limit service provider. If not
	// specified, any calls to the rate limit service will immediately return
	// success.
	RateLimitService *v22.RateLimitServiceConfig `protobuf:"bytes,10,opt,name=rate_limit_service,json=rateLimitService,proto3" json:"rate_limit_service,omitempty"`
	// Configuration for the runtime configuration provider. If not specified, a
	// “null” provider will be used which will result in all defaults being used.
	Runtime *Runtime `protobuf:"bytes,11,opt,name=runtime,proto3" json:"runtime,omitempty"`
	// Configuration for the local administration HTTP server.
	Admin *Admin `protobuf:"bytes,12,opt,name=admin,proto3" json:"admin,omitempty"`
	// Optional overload manager configuration.
	OverloadManager      *v2alpha.OverloadManager `protobuf:"bytes,15,opt,name=overload_manager,json=overloadManager,proto3" json:"overload_manager,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *Bootstrap) Reset()         { *m = Bootstrap{} }
func (m *Bootstrap) String() string { return proto.CompactTextString(m) }
func (*Bootstrap) ProtoMessage()    {}
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1197defdf9c5e6a, []int{0}
}

func (m *Bootstrap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bootstrap.Unmarshal(m, b)
}
func (m *Bootstrap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bootstrap.Marshal(b, m, deterministic)
}
func (m *Bootstrap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bootstrap.Merge(m, src)
}
func (m *Bootstrap) XXX_Size() int {
	return xxx_messageInfo_Bootstrap.Size(m)
}
func (m *Bootstrap) XXX_DiscardUnknown() {
	xxx_messageInfo_Bootstrap.DiscardUnknown(m)
}

var xxx_messageInfo_Bootstrap proto.InternalMessageInfo

func (m *Bootstrap) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *Bootstrap) GetStaticResources() *Bootstrap_StaticResources {
	if m != nil {
		return m.StaticResources
	}
	return nil
}

func (m *Bootstrap) GetDynamicResources() *Bootstrap_DynamicResources {
	if m != nil {
		return m.DynamicResources
	}
	return nil
}

func (m *Bootstrap) GetClusterManager() *ClusterManager {
	if m != nil {
		return m.ClusterManager
	}
	return nil
}

func (m *Bootstrap) GetHdsConfig() *core.ApiConfigSource {
	if m != nil {
		return m.HdsConfig
	}
	return nil
}

func (m *Bootstrap) GetFlagsPath() string {
	if m != nil {
		return m.FlagsPath
	}
	return ""
}

func (m *Bootstrap) GetStatsSinks() []*v2.StatsSink {
	if m != nil {
		return m.StatsSinks
	}
	return nil
}

func (m *Bootstrap) GetStatsConfig() *v2.StatsConfig {
	if m != nil {
		return m.StatsConfig
	}
	return nil
}

func (m *Bootstrap) GetStatsFlushInterval() *duration.Duration {
	if m != nil {
		return m.StatsFlushInterval
	}
	return nil
}

func (m *Bootstrap) GetWatchdog() *Watchdog {
	if m != nil {
		return m.Watchdog
	}
	return nil
}

func (m *Bootstrap) GetTracing() *v21.Tracing {
	if m != nil {
		return m.Tracing
	}
	return nil
}

func (m *Bootstrap) GetRateLimitService() *v22.RateLimitServiceConfig {
	if m != nil {
		return m.RateLimitService
	}
	return nil
}

func (m *Bootstrap) GetRuntime() *Runtime {
	if m != nil {
		return m.Runtime
	}
	return nil
}

func (m *Bootstrap) GetAdmin() *Admin {
	if m != nil {
		return m.Admin
	}
	return nil
}

func (m *Bootstrap) GetOverloadManager() *v2alpha.OverloadManager {
	if m != nil {
		return m.OverloadManager
	}
	return nil
}

type Bootstrap_StaticResources struct {
	// Static :ref:`Listeners <envoy_api_msg_Listener>`. These listeners are
	// available regardless of LDS configuration.
	Listeners []*v23.Listener `protobuf:"bytes,1,rep,name=listeners,proto3" json:"listeners,omitempty"`
	// If a network based configuration source is specified for :ref:`cds_config
	// <envoy_api_field_config.bootstrap.v2.Bootstrap.DynamicResources.cds_config>`, it's necessary
	// to have some initial cluster definitions available to allow Envoy to know
	// how to speak to the management server. These cluster definitions may not
	// use :ref:`EDS <arch_overview_dynamic_config_eds>` (i.e. they should be static
	// IP or DNS-based).
	Clusters []*v23.Cluster `protobuf:"bytes,2,rep,name=clusters,proto3" json:"clusters,omitempty"`
	// These static secrets can be used by :ref:`SdsSecretConfig
	// <envoy_api_msg_auth.SdsSecretConfig>`
	Secrets              []*auth.Secret `protobuf:"bytes,3,rep,name=secrets,proto3" json:"secrets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Bootstrap_StaticResources) Reset()         { *m = Bootstrap_StaticResources{} }
func (m *Bootstrap_StaticResources) String() string { return proto.CompactTextString(m) }
func (*Bootstrap_StaticResources) ProtoMessage()    {}
func (*Bootstrap_StaticResources) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1197defdf9c5e6a, []int{0, 0}
}

func (m *Bootstrap_StaticResources) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bootstrap_StaticResources.Unmarshal(m, b)
}
func (m *Bootstrap_StaticResources) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bootstrap_StaticResources.Marshal(b, m, deterministic)
}
func (m *Bootstrap_StaticResources) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bootstrap_StaticResources.Merge(m, src)
}
func (m *Bootstrap_StaticResources) XXX_Size() int {
	return xxx_messageInfo_Bootstrap_StaticResources.Size(m)
}
func (m *Bootstrap_StaticResources) XXX_DiscardUnknown() {
	xxx_messageInfo_Bootstrap_StaticResources.DiscardUnknown(m)
}

var xxx_messageInfo_Bootstrap_StaticResources proto.InternalMessageInfo

func (m *Bootstrap_StaticResources) GetListeners() []*v23.Listener {
	if m != nil {
		return m.Listeners
	}
	return nil
}

func (m *Bootstrap_StaticResources) GetClusters() []*v23.Cluster {
	if m != nil {
		return m.Clusters
	}
	return nil
}

func (m *Bootstrap_StaticResources) GetSecrets() []*auth.Secret {
	if m != nil {
		return m.Secrets
	}
	return nil
}

type Bootstrap_DynamicResources struct {
	// All :ref:`Listeners <envoy_api_msg_Listener>` are provided by a single
	// :ref:`LDS <arch_overview_dynamic_config_lds>` configuration source.
	LdsConfig *core.ConfigSource `protobuf:"bytes,1,opt,name=lds_config,json=ldsConfig,proto3" json:"lds_config,omitempty"`
	// All post-bootstrap :ref:`Cluster <envoy_api_msg_Cluster>` definitions are
	// provided by a single :ref:`CDS <arch_overview_dynamic_config_cds>`
	// configuration source.
	CdsConfig *core.ConfigSource `protobuf:"bytes,2,opt,name=cds_config,json=cdsConfig,proto3" json:"cds_config,omitempty"`
	// A single :ref:`ADS <config_overview_v2_ads>` source may be optionally
	// specified. This must have :ref:`api_type
	// <envoy_api_field_core.ApiConfigSource.api_type>` :ref:`GRPC
	// <envoy_api_enum_value_core.ApiConfigSource.ApiType.GRPC>`. Only
	// :ref:`ConfigSources <envoy_api_msg_core.ConfigSource>` that have
	// the :ref:`ads <envoy_api_field_core.ConfigSource.ads>` field set will be
	// streamed on the ADS channel.
	AdsConfig *core.ApiConfigSource `protobuf:"bytes,3,opt,name=ads_config,json=adsConfig,proto3" json:"ads_config,omitempty"`
	// [#not-implemented-hide:] Hide from docs.
	DeprecatedV1         *Bootstrap_DynamicResources_DeprecatedV1 `protobuf:"bytes,4,opt,name=deprecated_v1,json=deprecatedV1,proto3" json:"deprecated_v1,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *Bootstrap_DynamicResources) Reset()         { *m = Bootstrap_DynamicResources{} }
func (m *Bootstrap_DynamicResources) String() string { return proto.CompactTextString(m) }
func (*Bootstrap_DynamicResources) ProtoMessage()    {}
func (*Bootstrap_DynamicResources) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1197defdf9c5e6a, []int{0, 1}
}

func (m *Bootstrap_DynamicResources) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bootstrap_DynamicResources.Unmarshal(m, b)
}
func (m *Bootstrap_DynamicResources) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bootstrap_DynamicResources.Marshal(b, m, deterministic)
}
func (m *Bootstrap_DynamicResources) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bootstrap_DynamicResources.Merge(m, src)
}
func (m *Bootstrap_DynamicResources) XXX_Size() int {
	return xxx_messageInfo_Bootstrap_DynamicResources.Size(m)
}
func (m *Bootstrap_DynamicResources) XXX_DiscardUnknown() {
	xxx_messageInfo_Bootstrap_DynamicResources.DiscardUnknown(m)
}

var xxx_messageInfo_Bootstrap_DynamicResources proto.InternalMessageInfo

func (m *Bootstrap_DynamicResources) GetLdsConfig() *core.ConfigSource {
	if m != nil {
		return m.LdsConfig
	}
	return nil
}

func (m *Bootstrap_DynamicResources) GetCdsConfig() *core.ConfigSource {
	if m != nil {
		return m.CdsConfig
	}
	return nil
}

func (m *Bootstrap_DynamicResources) GetAdsConfig() *core.ApiConfigSource {
	if m != nil {
		return m.AdsConfig
	}
	return nil
}

// Deprecated: Do not use.
func (m *Bootstrap_DynamicResources) GetDeprecatedV1() *Bootstrap_DynamicResources_DeprecatedV1 {
	if m != nil {
		return m.DeprecatedV1
	}
	return nil
}

// [#not-implemented-hide:] Hide from docs.
type Bootstrap_DynamicResources_DeprecatedV1 struct {
	// Deprecated service discovery service config for when using API v1 REST.
	SdsConfig            *core.ConfigSource `protobuf:"bytes,1,opt,name=sds_config,json=sdsConfig,proto3" json:"sds_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Bootstrap_DynamicResources_DeprecatedV1) Reset() {
	*m = Bootstrap_DynamicResources_DeprecatedV1{}
}
func (m *Bootstrap_DynamicResources_DeprecatedV1) String() string { return proto.CompactTextString(m) }
func (*Bootstrap_DynamicResources_DeprecatedV1) ProtoMessage()    {}
func (*Bootstrap_DynamicResources_DeprecatedV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1197defdf9c5e6a, []int{0, 1, 0}
}

func (m *Bootstrap_DynamicResources_DeprecatedV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bootstrap_DynamicResources_DeprecatedV1.Unmarshal(m, b)
}
func (m *Bootstrap_DynamicResources_DeprecatedV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bootstrap_DynamicResources_DeprecatedV1.Marshal(b, m, deterministic)
}
func (m *Bootstrap_DynamicResources_DeprecatedV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bootstrap_DynamicResources_DeprecatedV1.Merge(m, src)
}
func (m *Bootstrap_DynamicResources_DeprecatedV1) XXX_Size() int {
	return xxx_messageInfo_Bootstrap_DynamicResources_DeprecatedV1.Size(m)
}
func (m *Bootstrap_DynamicResources_DeprecatedV1) XXX_DiscardUnknown() {
	xxx_messageInfo_Bootstrap_DynamicResources_DeprecatedV1.DiscardUnknown(m)
}

var xxx_messageInfo_Bootstrap_DynamicResources_DeprecatedV1 proto.InternalMessageInfo

func (m *Bootstrap_DynamicResources_DeprecatedV1) GetSdsConfig() *core.ConfigSource {
	if m != nil {
		return m.SdsConfig
	}
	return nil
}

// Administration interface :ref:`operations documentation
// <operations_admin_interface>`.
type Admin struct {
	// The path to write the access log for the administration server. If no
	// access log is desired specify ‘/dev/null’. This is only required if
	// :ref:`address <envoy_api_field_config.bootstrap.v2.Admin.address>` is set.
	AccessLogPath string `protobuf:"bytes,1,opt,name=access_log_path,json=accessLogPath,proto3" json:"access_log_path,omitempty"`
	// The cpu profiler output path for the administration server. If no profile
	// path is specified, the default is ‘/var/log/envoy/envoy.prof’.
	ProfilePath string `protobuf:"bytes,2,opt,name=profile_path,json=profilePath,proto3" json:"profile_path,omitempty"`
	// The TCP address that the administration server will listen on.
	// If not specified, Envoy will not start an administration server.
	Address              *core.Address `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Admin) Reset()         { *m = Admin{} }
func (m *Admin) String() string { return proto.CompactTextString(m) }
func (*Admin) ProtoMessage()    {}
func (*Admin) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1197defdf9c5e6a, []int{1}
}

func (m *Admin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Admin.Unmarshal(m, b)
}
func (m *Admin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Admin.Marshal(b, m, deterministic)
}
func (m *Admin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Admin.Merge(m, src)
}
func (m *Admin) XXX_Size() int {
	return xxx_messageInfo_Admin.Size(m)
}
func (m *Admin) XXX_DiscardUnknown() {
	xxx_messageInfo_Admin.DiscardUnknown(m)
}

var xxx_messageInfo_Admin proto.InternalMessageInfo

func (m *Admin) GetAccessLogPath() string {
	if m != nil {
		return m.AccessLogPath
	}
	return ""
}

func (m *Admin) GetProfilePath() string {
	if m != nil {
		return m.ProfilePath
	}
	return ""
}

func (m *Admin) GetAddress() *core.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

// Cluster manager :ref:`architecture overview <arch_overview_cluster_manager>`.
type ClusterManager struct {
	// Name of the local cluster (i.e., the cluster that owns the Envoy running
	// this configuration). In order to enable :ref:`zone aware routing
	// <arch_overview_load_balancing_zone_aware_routing>` this option must be set.
	// If *local_cluster_name* is defined then :ref:`clusters
	// <envoy_api_msg_Cluster>` must be defined in the :ref:`Bootstrap
	// static cluster resources
	// <envoy_api_field_config.bootstrap.v2.Bootstrap.StaticResources.clusters>`. This is unrelated to
	// the :option:`--service-cluster` option which does not `affect zone aware
	// routing <https://github.com/envoyproxy/envoy/issues/774>`_.
	LocalClusterName string `protobuf:"bytes,1,opt,name=local_cluster_name,json=localClusterName,proto3" json:"local_cluster_name,omitempty"`
	// Optional global configuration for outlier detection.
	OutlierDetection *ClusterManager_OutlierDetection `protobuf:"bytes,2,opt,name=outlier_detection,json=outlierDetection,proto3" json:"outlier_detection,omitempty"`
	// Optional configuration used to bind newly established upstream connections.
	// This may be overridden on a per-cluster basis by upstream_bind_config in the cds_config.
	UpstreamBindConfig *core.BindConfig `protobuf:"bytes,3,opt,name=upstream_bind_config,json=upstreamBindConfig,proto3" json:"upstream_bind_config,omitempty"`
	// A management server endpoint to stream load stats to via
	// *StreamLoadStats*. This must have :ref:`api_type
	// <envoy_api_field_core.ApiConfigSource.api_type>` :ref:`GRPC
	// <envoy_api_enum_value_core.ApiConfigSource.ApiType.GRPC>`.
	LoadStatsConfig      *core.ApiConfigSource `protobuf:"bytes,4,opt,name=load_stats_config,json=loadStatsConfig,proto3" json:"load_stats_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ClusterManager) Reset()         { *m = ClusterManager{} }
func (m *ClusterManager) String() string { return proto.CompactTextString(m) }
func (*ClusterManager) ProtoMessage()    {}
func (*ClusterManager) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1197defdf9c5e6a, []int{2}
}

func (m *ClusterManager) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterManager.Unmarshal(m, b)
}
func (m *ClusterManager) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterManager.Marshal(b, m, deterministic)
}
func (m *ClusterManager) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterManager.Merge(m, src)
}
func (m *ClusterManager) XXX_Size() int {
	return xxx_messageInfo_ClusterManager.Size(m)
}
func (m *ClusterManager) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterManager.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterManager proto.InternalMessageInfo

func (m *ClusterManager) GetLocalClusterName() string {
	if m != nil {
		return m.LocalClusterName
	}
	return ""
}

func (m *ClusterManager) GetOutlierDetection() *ClusterManager_OutlierDetection {
	if m != nil {
		return m.OutlierDetection
	}
	return nil
}

func (m *ClusterManager) GetUpstreamBindConfig() *core.BindConfig {
	if m != nil {
		return m.UpstreamBindConfig
	}
	return nil
}

func (m *ClusterManager) GetLoadStatsConfig() *core.ApiConfigSource {
	if m != nil {
		return m.LoadStatsConfig
	}
	return nil
}

type ClusterManager_OutlierDetection struct {
	// Specifies the path to the outlier event log.
	EventLogPath         string   `protobuf:"bytes,1,opt,name=event_log_path,json=eventLogPath,proto3" json:"event_log_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterManager_OutlierDetection) Reset()         { *m = ClusterManager_OutlierDetection{} }
func (m *ClusterManager_OutlierDetection) String() string { return proto.CompactTextString(m) }
func (*ClusterManager_OutlierDetection) ProtoMessage()    {}
func (*ClusterManager_OutlierDetection) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1197defdf9c5e6a, []int{2, 0}
}

func (m *ClusterManager_OutlierDetection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterManager_OutlierDetection.Unmarshal(m, b)
}
func (m *ClusterManager_OutlierDetection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterManager_OutlierDetection.Marshal(b, m, deterministic)
}
func (m *ClusterManager_OutlierDetection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterManager_OutlierDetection.Merge(m, src)
}
func (m *ClusterManager_OutlierDetection) XXX_Size() int {
	return xxx_messageInfo_ClusterManager_OutlierDetection.Size(m)
}
func (m *ClusterManager_OutlierDetection) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterManager_OutlierDetection.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterManager_OutlierDetection proto.InternalMessageInfo

func (m *ClusterManager_OutlierDetection) GetEventLogPath() string {
	if m != nil {
		return m.EventLogPath
	}
	return ""
}

// Envoy process watchdog configuration. When configured, this monitors for
// nonresponsive threads and kills the process after the configured thresholds.
type Watchdog struct {
	// The duration after which Envoy counts a nonresponsive thread in the
	// *server.watchdog_miss* statistic. If not specified the default is 200ms.
	MissTimeout *duration.Duration `protobuf:"bytes,1,opt,name=miss_timeout,json=missTimeout,proto3" json:"miss_timeout,omitempty"`
	// The duration after which Envoy counts a nonresponsive thread in the
	// *server.watchdog_mega_miss* statistic. If not specified the default is
	// 1000ms.
	MegamissTimeout *duration.Duration `protobuf:"bytes,2,opt,name=megamiss_timeout,json=megamissTimeout,proto3" json:"megamiss_timeout,omitempty"`
	// If a watched thread has been nonresponsive for this duration, assume a
	// programming error and kill the entire Envoy process. Set to 0 to disable
	// kill behavior. If not specified the default is 0 (disabled).
	KillTimeout *duration.Duration `protobuf:"bytes,3,opt,name=kill_timeout,json=killTimeout,proto3" json:"kill_timeout,omitempty"`
	// If at least two watched threads have been nonresponsive for at least this
	// duration assume a true deadlock and kill the entire Envoy process. Set to 0
	// to disable this behavior. If not specified the default is 0 (disabled).
	MultikillTimeout     *duration.Duration `protobuf:"bytes,4,opt,name=multikill_timeout,json=multikillTimeout,proto3" json:"multikill_timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Watchdog) Reset()         { *m = Watchdog{} }
func (m *Watchdog) String() string { return proto.CompactTextString(m) }
func (*Watchdog) ProtoMessage()    {}
func (*Watchdog) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1197defdf9c5e6a, []int{3}
}

func (m *Watchdog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Watchdog.Unmarshal(m, b)
}
func (m *Watchdog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Watchdog.Marshal(b, m, deterministic)
}
func (m *Watchdog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Watchdog.Merge(m, src)
}
func (m *Watchdog) XXX_Size() int {
	return xxx_messageInfo_Watchdog.Size(m)
}
func (m *Watchdog) XXX_DiscardUnknown() {
	xxx_messageInfo_Watchdog.DiscardUnknown(m)
}

var xxx_messageInfo_Watchdog proto.InternalMessageInfo

func (m *Watchdog) GetMissTimeout() *duration.Duration {
	if m != nil {
		return m.MissTimeout
	}
	return nil
}

func (m *Watchdog) GetMegamissTimeout() *duration.Duration {
	if m != nil {
		return m.MegamissTimeout
	}
	return nil
}

func (m *Watchdog) GetKillTimeout() *duration.Duration {
	if m != nil {
		return m.KillTimeout
	}
	return nil
}

func (m *Watchdog) GetMultikillTimeout() *duration.Duration {
	if m != nil {
		return m.MultikillTimeout
	}
	return nil
}

// Runtime :ref:`configuration overview <config_runtime>`.
type Runtime struct {
	// The implementation assumes that the file system tree is accessed via a
	// symbolic link. An atomic link swap is used when a new tree should be
	// switched to. This parameter specifies the path to the symbolic link. Envoy
	// will watch the location for changes and reload the file system tree when
	// they happen.
	SymlinkRoot string `protobuf:"bytes,1,opt,name=symlink_root,json=symlinkRoot,proto3" json:"symlink_root,omitempty"`
	// Specifies the subdirectory to load within the root directory. This is
	// useful if multiple systems share the same delivery mechanism. Envoy
	// configuration elements can be contained in a dedicated subdirectory.
	Subdirectory string `protobuf:"bytes,2,opt,name=subdirectory,proto3" json:"subdirectory,omitempty"`
	// Specifies an optional subdirectory to load within the root directory. If
	// specified and the directory exists, configuration values within this
	// directory will override those found in the primary subdirectory. This is
	// useful when Envoy is deployed across many different types of servers.
	// Sometimes it is useful to have a per service cluster directory for runtime
	// configuration. See below for exactly how the override directory is used.
	OverrideSubdirectory string   `protobuf:"bytes,3,opt,name=override_subdirectory,json=overrideSubdirectory,proto3" json:"override_subdirectory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Runtime) Reset()         { *m = Runtime{} }
func (m *Runtime) String() string { return proto.CompactTextString(m) }
func (*Runtime) ProtoMessage()    {}
func (*Runtime) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1197defdf9c5e6a, []int{4}
}

func (m *Runtime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Runtime.Unmarshal(m, b)
}
func (m *Runtime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Runtime.Marshal(b, m, deterministic)
}
func (m *Runtime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Runtime.Merge(m, src)
}
func (m *Runtime) XXX_Size() int {
	return xxx_messageInfo_Runtime.Size(m)
}
func (m *Runtime) XXX_DiscardUnknown() {
	xxx_messageInfo_Runtime.DiscardUnknown(m)
}

var xxx_messageInfo_Runtime proto.InternalMessageInfo

func (m *Runtime) GetSymlinkRoot() string {
	if m != nil {
		return m.SymlinkRoot
	}
	return ""
}

func (m *Runtime) GetSubdirectory() string {
	if m != nil {
		return m.Subdirectory
	}
	return ""
}

func (m *Runtime) GetOverrideSubdirectory() string {
	if m != nil {
		return m.OverrideSubdirectory
	}
	return ""
}

func init() {
	proto.RegisterType((*Bootstrap)(nil), "envoy.config.bootstrap.v2.Bootstrap")
	proto.RegisterType((*Bootstrap_StaticResources)(nil), "envoy.config.bootstrap.v2.Bootstrap.StaticResources")
	proto.RegisterType((*Bootstrap_DynamicResources)(nil), "envoy.config.bootstrap.v2.Bootstrap.DynamicResources")
	proto.RegisterType((*Bootstrap_DynamicResources_DeprecatedV1)(nil), "envoy.config.bootstrap.v2.Bootstrap.DynamicResources.DeprecatedV1")
	proto.RegisterType((*Admin)(nil), "envoy.config.bootstrap.v2.Admin")
	proto.RegisterType((*ClusterManager)(nil), "envoy.config.bootstrap.v2.ClusterManager")
	proto.RegisterType((*ClusterManager_OutlierDetection)(nil), "envoy.config.bootstrap.v2.ClusterManager.OutlierDetection")
	proto.RegisterType((*Watchdog)(nil), "envoy.config.bootstrap.v2.Watchdog")
	proto.RegisterType((*Runtime)(nil), "envoy.config.bootstrap.v2.Runtime")
}

func init() {
	proto.RegisterFile("envoy/config/bootstrap/v2/bootstrap.proto", fileDescriptor_f1197defdf9c5e6a)
}

var fileDescriptor_f1197defdf9c5e6a = []byte{
	// 1204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0x5f, 0x6f, 0x1b, 0xc5,
	0x17, 0xfd, 0xf9, 0x4f, 0xeb, 0xf8, 0xda, 0x89, 0x9d, 0x51, 0xda, 0x6e, 0xad, 0x5f, 0x9b, 0xd4,
	0x2d, 0xa8, 0x15, 0xd5, 0x5a, 0x75, 0x0b, 0x94, 0xaa, 0x2a, 0x8a, 0x1b, 0x15, 0x21, 0x95, 0x14,
	0xd6, 0x15, 0x08, 0x5e, 0x56, 0xe3, 0xdd, 0xc9, 0x7a, 0x94, 0xdd, 0x1d, 0x6b, 0x66, 0x76, 0x51,
	0x5e, 0x79, 0xe2, 0x11, 0x21, 0x1e, 0x10, 0x1f, 0x85, 0x27, 0x78, 0xe3, 0x53, 0xc0, 0x33, 0xdf,
	0x02, 0xcd, 0x9f, 0x5d, 0x67, 0x5d, 0x37, 0x09, 0xbc, 0xad, 0xef, 0x9c, 0x73, 0xee, 0x9d, 0x99,
	0x7b, 0xcf, 0x18, 0xee, 0x91, 0x34, 0x67, 0x27, 0xa3, 0x80, 0xa5, 0x47, 0x34, 0x1a, 0xcd, 0x18,
	0x93, 0x42, 0x72, 0xbc, 0x18, 0xe5, 0xe3, 0xe5, 0x0f, 0x77, 0xc1, 0x99, 0x64, 0xe8, 0xba, 0x86,
	0xba, 0x06, 0xea, 0x2e, 0x57, 0xf3, 0xf1, 0x60, 0xd7, 0xa8, 0xe0, 0x05, 0x55, 0xc4, 0x80, 0x71,
	0x32, 0xc2, 0x61, 0xc8, 0x89, 0x10, 0x86, 0x3b, 0xf8, 0xff, 0x9b, 0x80, 0x19, 0x16, 0x64, 0xed,
	0x2a, 0xce, 0xe4, 0x7c, 0x14, 0x10, 0x2e, 0xed, 0xea, 0x3b, 0x6f, 0x72, 0x4d, 0x0d, 0xbe, 0x60,
	0x19, 0x0f, 0x0a, 0x91, 0xab, 0x55, 0x58, 0x28, 0xd6, 0xc6, 0xe3, 0x32, 0x7e, 0xab, 0xb2, 0x73,
	0xc9, 0x71, 0x40, 0x14, 0x40, 0x7f, 0x58, 0xc8, 0xed, 0x0a, 0x24, 0x21, 0x92, 0xd3, 0x40, 0x28,
	0x90, 0x90, 0x58, 0x16, 0x3a, 0xf7, 0x2b, 0x20, 0x96, 0x13, 0x1e, 0x33, 0x1c, 0x8e, 0xf2, 0x31,
	0x8e, 0x17, 0x73, 0x5c, 0x06, 0xd6, 0x4a, 0x72, 0x2c, 0x49, 0x4c, 0x13, 0x2a, 0x95, 0x28, 0x8f,
	0x0b, 0xc9, 0x9b, 0x11, 0x63, 0x51, 0x4c, 0x46, 0xfa, 0xd7, 0x2c, 0x3b, 0x1a, 0x85, 0x19, 0xc7,
	0x92, 0xb2, 0xd4, 0xae, 0x5f, 0xcb, 0x71, 0x4c, 0x43, 0x2c, 0xc9, 0xa8, 0xf8, 0xb0, 0x0b, 0x3b,
	0x11, 0x8b, 0x98, 0xfe, 0x1c, 0xa9, 0x2f, 0x13, 0x1d, 0xfe, 0xb2, 0x09, 0xed, 0x49, 0x71, 0x5d,
	0xe8, 0x3d, 0x68, 0xa6, 0x2c, 0x24, 0x4e, 0x6d, 0xaf, 0x76, 0xb7, 0x33, 0xbe, 0xe6, 0x9a, 0x5b,
	0xc5, 0x0b, 0xea, 0xe6, 0x63, 0x57, 0x9d, 0xae, 0x7b, 0xc8, 0x42, 0xe2, 0x69, 0x10, 0xf2, 0xa1,
	0xaf, 0xf6, 0x4a, 0x03, 0x9f, 0x13, 0x73, 0xda, 0xc2, 0xa9, 0x6b, 0xe2, 0x23, 0xf7, 0xad, 0xed,
	0xe0, 0x96, 0xc9, 0xdc, 0xa9, 0x26, 0x7b, 0x05, 0xd7, 0xeb, 0x89, 0x6a, 0x00, 0xcd, 0x60, 0x3b,
	0x3c, 0x49, 0x71, 0x52, 0xc9, 0xd0, 0xd0, 0x19, 0xde, 0xbf, 0x50, 0x86, 0x03, 0xc3, 0x5e, 0xa6,
	0xe8, 0x87, 0x2b, 0x11, 0xe4, 0x41, 0x2f, 0x88, 0x33, 0x21, 0x09, 0xf7, 0x13, 0x9c, 0xe2, 0x88,
	0x70, 0xa7, 0xa9, 0x33, 0xdc, 0x3b, 0x23, 0xc3, 0x73, 0xc3, 0xf8, 0xcc, 0x10, 0xbc, 0xad, 0xa0,
	0xf2, 0x1b, 0xed, 0x03, 0xcc, 0x43, 0xe1, 0x1b, 0xa6, 0xb3, 0xa5, 0xe5, 0x86, 0x6b, 0xce, 0x72,
	0x7f, 0x41, 0x9f, 0x6b, 0xcc, 0x54, 0x17, 0xe3, 0xb5, 0xe7, 0xa1, 0x30, 0x01, 0x74, 0x03, 0xe0,
	0x28, 0xc6, 0x91, 0xf0, 0x17, 0x58, 0xce, 0x9d, 0x4b, 0x7b, 0xb5, 0xbb, 0x6d, 0xaf, 0xad, 0x23,
	0x9f, 0x63, 0x39, 0x47, 0xcf, 0xa1, 0xa3, 0xdb, 0xcc, 0x17, 0x34, 0x3d, 0x16, 0xce, 0xe5, 0xbd,
	0xc6, 0xa9, 0x14, 0xb6, 0x62, 0xdb, 0x92, 0x2a, 0x9b, 0x3a, 0x69, 0x31, 0xa5, 0xe9, 0xb1, 0x07,
	0xa2, 0xf8, 0x14, 0xe8, 0x13, 0xe8, 0x1a, 0x11, 0x5b, 0xe8, 0xa6, 0x2e, 0xf4, 0xce, 0xd9, 0x2a,
	0xa6, 0x3e, 0xcf, 0xa4, 0xb7, 0xc5, 0x7e, 0x01, 0x3b, 0x46, 0xe8, 0x28, 0xce, 0xc4, 0xdc, 0xa7,
	0xa9, 0x24, 0x3c, 0xc7, 0xb1, 0xd3, 0xd2, 0x82, 0xd7, 0x5d, 0xd3, 0xb1, 0x6e, 0xd1, 0xb1, 0xee,
	0x81, 0xed, 0xd8, 0x49, 0xf3, 0xe7, 0xbf, 0x76, 0x6b, 0x1e, 0xd2, 0xe4, 0x17, 0x8a, 0xfb, 0xa9,
	0xa5, 0xa2, 0x8f, 0x61, 0xe3, 0x5b, 0x2c, 0x83, 0x79, 0xc8, 0x22, 0x67, 0x43, 0xcb, 0xdc, 0x3e,
	0xe3, 0x3e, 0xbe, 0xb2, 0x50, 0xaf, 0x24, 0xa1, 0xc7, 0xd0, 0x52, 0xd3, 0x4a, 0xd3, 0xc8, 0x69,
	0x6b, 0xfe, 0xcd, 0x2a, 0xdf, 0x8c, 0x72, 0x3e, 0x76, 0x5f, 0x1b, 0x94, 0x57, 0xc0, 0x91, 0x0f,
	0x48, 0x8d, 0x9e, 0xaf, 0x67, 0xcf, 0x17, 0x84, 0xe7, 0x34, 0x20, 0x0e, 0x68, 0x91, 0x07, 0x55,
	0x91, 0x72, 0x44, 0x95, 0x90, 0x87, 0x25, 0x79, 0xa9, 0x7e, 0x4c, 0x0d, 0xc5, 0x9e, 0x54, 0x9f,
	0xaf, 0xc4, 0xd1, 0x53, 0x68, 0xf1, 0x2c, 0x95, 0x34, 0x21, 0x4e, 0xa7, 0xd2, 0x1b, 0xeb, 0xb6,
	0xe6, 0x19, 0xa4, 0x57, 0x50, 0xd0, 0x07, 0x70, 0x09, 0x87, 0x09, 0x4d, 0x9d, 0xae, 0xe6, 0xee,
	0x9d, 0xc1, 0xdd, 0x57, 0x38, 0xcf, 0xc0, 0xd1, 0xd7, 0xd0, 0x2f, 0xec, 0xa6, 0xec, 0xf4, 0x9e,
	0x96, 0x70, 0xab, 0x12, 0xa5, 0x29, 0x59, 0x97, 0x72, 0x5f, 0xd9, 0x40, 0xd1, 0xee, 0x3d, 0x56,
	0x0d, 0x0c, 0x7e, 0xaf, 0x41, 0x6f, 0x65, 0x98, 0xd1, 0x13, 0x68, 0xc7, 0x54, 0x48, 0x92, 0x12,
	0x2e, 0x9c, 0x9a, 0xee, 0xcf, 0xab, 0xd5, 0x11, 0x78, 0x69, 0x97, 0x27, 0xcd, 0x3f, 0xfe, 0xdc,
	0xfd, 0x9f, 0xb7, 0x84, 0xa3, 0x0f, 0x61, 0xc3, 0x4e, 0x94, 0x32, 0x14, 0x45, 0xbd, 0x52, 0xa5,
	0xda, 0xf9, 0xb3, 0xcc, 0x12, 0x8c, 0x3e, 0x82, 0x96, 0x20, 0x01, 0x27, 0x52, 0xd9, 0x44, 0x43,
	0xf7, 0x5e, 0x85, 0xa7, 0x5e, 0x0f, 0x77, 0xaa, 0x11, 0x96, 0x5b, 0xe0, 0x07, 0x3f, 0x36, 0xa0,
	0xbf, 0x6a, 0x17, 0xe8, 0x19, 0x40, 0xbc, 0x1c, 0x64, 0x63, 0x8a, 0xbb, 0x6b, 0x06, 0xb9, 0x3a,
	0xc5, 0x71, 0x39, 0xc5, 0xcf, 0x00, 0x82, 0x25, 0xbf, 0x7e, 0x41, 0x7e, 0x50, 0xf2, 0xf7, 0x01,
	0xf0, 0x92, 0xdf, 0xb8, 0xb8, 0x91, 0xe0, 0x52, 0xe2, 0x18, 0x36, 0x43, 0xb2, 0xe0, 0x24, 0xc0,
	0x92, 0x84, 0x7e, 0xfe, 0xc0, 0xba, 0xdb, 0xe4, 0x3f, 0xf9, 0xa7, 0x7b, 0x50, 0x4a, 0x7d, 0xf9,
	0x60, 0x52, 0x77, 0x6a, 0x5e, 0x37, 0x3c, 0x15, 0x19, 0x1c, 0x42, 0xf7, 0x34, 0x42, 0xed, 0x5f,
	0xfc, 0xfb, 0xf3, 0x13, 0x45, 0xf1, 0xc3, 0xef, 0x6b, 0x70, 0x49, 0x37, 0x31, 0x7a, 0x17, 0x7a,
	0x38, 0x08, 0x88, 0x10, 0x7e, 0xcc, 0x22, 0x63, 0x8a, 0x35, 0x6d, 0x8a, 0x9b, 0x26, 0xfc, 0x92,
	0x45, 0xda, 0x18, 0x6f, 0x41, 0x77, 0xc1, 0xd9, 0x11, 0x8d, 0x89, 0x01, 0xd5, 0x35, 0xa8, 0x63,
	0x63, 0x1a, 0xf2, 0x08, 0x5a, 0xf6, 0xff, 0x87, 0x3d, 0xd1, 0xc1, 0xba, 0x13, 0x35, 0x08, 0xaf,
	0x80, 0x0e, 0xbf, 0x6b, 0xc0, 0x56, 0xd5, 0xf6, 0xd1, 0x7d, 0x40, 0x31, 0x0b, 0x70, 0xec, 0x17,
	0x0f, 0x48, 0x8a, 0x13, 0x62, 0xcb, 0xea, 0xeb, 0x15, 0x4b, 0x38, 0xc4, 0x09, 0x41, 0x11, 0x6c,
	0xb3, 0x4c, 0xc6, 0x94, 0x70, 0x3f, 0x24, 0x92, 0x04, 0xca, 0x00, 0x6d, 0x4b, 0x3c, 0xb9, 0xf0,
	0x53, 0xe3, 0xbe, 0x32, 0x12, 0x07, 0x85, 0x82, 0xd7, 0x67, 0x2b, 0x11, 0xf4, 0x0a, 0x76, 0xb2,
	0x85, 0x90, 0x9c, 0xe0, 0xc4, 0x9f, 0xd1, 0x34, 0xac, 0xb6, 0xcf, 0x8d, 0x35, 0x9b, 0x9d, 0xd0,
	0x34, 0xb4, 0x6e, 0x85, 0x0a, 0xea, 0x32, 0x86, 0x0e, 0x61, 0x5b, 0xbb, 0x46, 0xe5, 0xb1, 0x68,
	0x5e, 0xb8, 0x19, 0x7b, 0x8a, 0x7c, 0xea, 0xed, 0x18, 0x3c, 0x86, 0xfe, 0xea, 0x36, 0xd0, 0x1d,
	0xd8, 0x22, 0x39, 0x49, 0xe5, 0xea, 0xf5, 0x76, 0x75, 0xd4, 0xde, 0xee, 0xf0, 0xa7, 0x3a, 0x6c,
	0x14, 0x5e, 0x8f, 0x9e, 0x42, 0x37, 0xa1, 0x42, 0xf8, 0xca, 0x15, 0x59, 0x26, 0x6d, 0x7b, 0xbd,
	0xfd, 0xb5, 0xf1, 0x3a, 0x0a, 0xfe, 0xda, 0xa0, 0xd1, 0x01, 0xf4, 0x13, 0x12, 0xe1, 0x8a, 0x42,
	0xfd, 0x3c, 0x85, 0x5e, 0x41, 0x29, 0x54, 0x9e, 0x42, 0xf7, 0x98, 0xc6, 0x71, 0xa9, 0xd0, 0x38,
	0xb7, 0x06, 0x05, 0x2f, 0xd8, 0x2f, 0x60, 0x3b, 0xc9, 0x62, 0x49, 0x2b, 0x12, 0xcd, 0xf3, 0x24,
	0xfa, 0x25, 0xc7, 0xea, 0x0c, 0x7f, 0xa8, 0x41, 0xcb, 0xbe, 0x13, 0xe8, 0x3e, 0x74, 0xc5, 0x49,
	0x12, 0xd3, 0xf4, 0xd8, 0xe7, 0x8c, 0x99, 0x53, 0x69, 0x4f, 0xda, 0xbf, 0xfe, 0xfd, 0x5b, 0xa3,
	0xc9, 0xeb, 0x7b, 0x35, 0xaf, 0x63, 0x97, 0x3d, 0xc6, 0x24, 0x1a, 0x42, 0x57, 0x64, 0xb3, 0x90,
	0x72, 0x12, 0x48, 0xc6, 0x4f, 0xec, 0xb8, 0x54, 0x62, 0xe8, 0x21, 0x5c, 0x51, 0x86, 0xcf, 0x69,
	0x48, 0xfc, 0x0a, 0xb8, 0xa1, 0xc1, 0x3b, 0xc5, 0xe2, 0xf4, 0xd4, 0xda, 0xa4, 0xf9, 0x4d, 0x3d,
	0x1f, 0xcf, 0x2e, 0xeb, 0xea, 0x1f, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0x70, 0x9f, 0xeb, 0x38,
	0x49, 0x0c, 0x00, 0x00,
}
