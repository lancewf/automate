// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/config/deployment/config_request.proto

package deployment // import "github.com/chef/automate/api/config/deployment"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import shared "github.com/chef/automate/api/config/shared"
import _ "github.com/chef/automate/components/automate-grpc/protoc-gen-a2-config/api/a2conf"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ConfigRequest struct {
	V1                   *ConfigRequest_V1 `protobuf:"bytes,3,opt,name=v1,proto3" json:"v1,omitempty" toml:"v1,omitempty" mapstructure:"v1,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte            `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32             `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest) Reset()         { *m = ConfigRequest{} }
func (m *ConfigRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest) ProtoMessage()    {}
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_3799ee30fbfa1658, []int{0}
}
func (m *ConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest.Unmarshal(m, b)
}
func (m *ConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest.Merge(dst, src)
}
func (m *ConfigRequest) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest.Size(m)
}
func (m *ConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest proto.InternalMessageInfo

func (m *ConfigRequest) GetV1() *ConfigRequest_V1 {
	if m != nil {
		return m.V1
	}
	return nil
}

type ConfigRequest_V1 struct {
	Sys                  *ConfigRequest_V1_System  `protobuf:"bytes,1,opt,name=sys,proto3" json:"sys,omitempty" toml:"sys,omitempty" mapstructure:"sys,omitempty"`
	Svc                  *ConfigRequest_V1_Service `protobuf:"bytes,2,opt,name=svc,proto3" json:"svc,omitempty" toml:"svc,omitempty" mapstructure:"svc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                    `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                     `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1) Reset()         { *m = ConfigRequest_V1{} }
func (m *ConfigRequest_V1) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1) ProtoMessage()    {}
func (*ConfigRequest_V1) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_3799ee30fbfa1658, []int{0, 0}
}
func (m *ConfigRequest_V1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1.Unmarshal(m, b)
}
func (m *ConfigRequest_V1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest_V1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1.Merge(dst, src)
}
func (m *ConfigRequest_V1) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1.Size(m)
}
func (m *ConfigRequest_V1) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1 proto.InternalMessageInfo

func (m *ConfigRequest_V1) GetSys() *ConfigRequest_V1_System {
	if m != nil {
		return m.Sys
	}
	return nil
}

func (m *ConfigRequest_V1) GetSvc() *ConfigRequest_V1_Service {
	if m != nil {
		return m.Svc
	}
	return nil
}

type ConfigRequest_V1_System struct {
	Mlsa                 *shared.Mlsa                        `protobuf:"bytes,1,opt,name=mlsa,proto3" json:"mlsa,omitempty" toml:"mlsa,omitempty" mapstructure:"mlsa,omitempty"`
	Service              *ConfigRequest_V1_System_Service    `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty" toml:"service,omitempty" mapstructure:"service,omitempty"`
	Log                  *ConfigRequest_V1_System_Log        `protobuf:"bytes,3,opt,name=log,proto3" json:"log,omitempty" toml:"log,omitempty" mapstructure:"log,omitempty"`
	GatherLogs           *ConfigRequest_V1_System_GatherLogs `protobuf:"bytes,4,opt,name=gather_logs,json=gatherLogs,proto3" json:"gather_logs,omitempty" toml:"gather_logs,omitempty" mapstructure:"gather_logs,omitempty"`
	Proxy                *ConfigRequest_V1_System_Proxy      `protobuf:"bytes,5,opt,name=proxy,proto3" json:"proxy,omitempty" toml:"proxy,omitempty" mapstructure:"proxy,omitempty"`
	Backup               *ConfigRequest_V1_System_Backup     `protobuf:"bytes,6,opt,name=backup,proto3" json:"backup,omitempty" toml:"backup,omitempty" mapstructure:"backup,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                               `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System) Reset()         { *m = ConfigRequest_V1_System{} }
func (m *ConfigRequest_V1_System) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System) ProtoMessage()    {}
func (*ConfigRequest_V1_System) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_3799ee30fbfa1658, []int{0, 0, 0}
}
func (m *ConfigRequest_V1_System) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest_V1_System) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System.Merge(dst, src)
}
func (m *ConfigRequest_V1_System) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System.Size(m)
}
func (m *ConfigRequest_V1_System) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System proto.InternalMessageInfo

func (m *ConfigRequest_V1_System) GetMlsa() *shared.Mlsa {
	if m != nil {
		return m.Mlsa
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetService() *ConfigRequest_V1_System_Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetLog() *ConfigRequest_V1_System_Log {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetGatherLogs() *ConfigRequest_V1_System_GatherLogs {
	if m != nil {
		return m.GatherLogs
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetProxy() *ConfigRequest_V1_System_Proxy {
	if m != nil {
		return m.Proxy
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetBackup() *ConfigRequest_V1_System_Backup {
	if m != nil {
		return m.Backup
	}
	return nil
}

type ConfigRequest_V1_System_Service struct {
	ListenAddress        *wrappers.StringValue `protobuf:"bytes,1,opt,name=listen_address,json=listenAddress,proto3" json:"listen_address,omitempty" toml:"listen_address,omitempty" mapstructure:"listen_address,omitempty"`
	Port                 *wrappers.Int32Value  `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty" toml:"port,omitempty" mapstructure:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Service) Reset()         { *m = ConfigRequest_V1_System_Service{} }
func (m *ConfigRequest_V1_System_Service) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Service) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_3799ee30fbfa1658, []int{0, 0, 0, 0}
}
func (m *ConfigRequest_V1_System_Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Service.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Service.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest_V1_System_Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Service.Merge(dst, src)
}
func (m *ConfigRequest_V1_System_Service) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Service.Size(m)
}
func (m *ConfigRequest_V1_System_Service) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Service.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Service proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Service) GetListenAddress() *wrappers.StringValue {
	if m != nil {
		return m.ListenAddress
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetPort() *wrappers.Int32Value {
	if m != nil {
		return m.Port
	}
	return nil
}

type ConfigRequest_V1_System_Log struct {
	Level                *wrappers.StringValue `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty" toml:"level,omitempty" mapstructure:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Log) Reset()         { *m = ConfigRequest_V1_System_Log{} }
func (m *ConfigRequest_V1_System_Log) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Log) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_3799ee30fbfa1658, []int{0, 0, 0, 1}
}
func (m *ConfigRequest_V1_System_Log) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Log.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Log.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest_V1_System_Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Log.Merge(dst, src)
}
func (m *ConfigRequest_V1_System_Log) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Log.Size(m)
}
func (m *ConfigRequest_V1_System_Log) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Log.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Log proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Log) GetLevel() *wrappers.StringValue {
	if m != nil {
		return m.Level
	}
	return nil
}

type ConfigRequest_V1_System_GatherLogs struct {
	StagingDir           *wrappers.StringValue `protobuf:"bytes,1,opt,name=staging_dir,json=stagingDir,proto3" json:"staging_dir,omitempty" toml:"staging_dir,omitempty" mapstructure:"staging_dir,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_GatherLogs) Reset()         { *m = ConfigRequest_V1_System_GatherLogs{} }
func (m *ConfigRequest_V1_System_GatherLogs) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_GatherLogs) ProtoMessage()    {}
func (*ConfigRequest_V1_System_GatherLogs) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_3799ee30fbfa1658, []int{0, 0, 0, 2}
}
func (m *ConfigRequest_V1_System_GatherLogs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_GatherLogs.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_GatherLogs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_GatherLogs.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest_V1_System_GatherLogs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_GatherLogs.Merge(dst, src)
}
func (m *ConfigRequest_V1_System_GatherLogs) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_GatherLogs.Size(m)
}
func (m *ConfigRequest_V1_System_GatherLogs) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_GatherLogs.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_GatherLogs proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_GatherLogs) GetStagingDir() *wrappers.StringValue {
	if m != nil {
		return m.StagingDir
	}
	return nil
}

type ConfigRequest_V1_System_Proxy struct {
	ConnectionString     *wrappers.StringValue `protobuf:"bytes,1,opt,name=connection_string,json=connectionString,proto3" json:"connection_string,omitempty" toml:"connection_string,omitempty" mapstructure:"connection_string,omitempty"`
	NoProxyString        *wrappers.StringValue `protobuf:"bytes,2,opt,name=no_proxy_string,json=noProxyString,proto3" json:"no_proxy_string,omitempty" toml:"no_proxy_string,omitempty" mapstructure:"no_proxy_string,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Proxy) Reset()         { *m = ConfigRequest_V1_System_Proxy{} }
func (m *ConfigRequest_V1_System_Proxy) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Proxy) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Proxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_3799ee30fbfa1658, []int{0, 0, 0, 3}
}
func (m *ConfigRequest_V1_System_Proxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Proxy.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Proxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Proxy.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest_V1_System_Proxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Proxy.Merge(dst, src)
}
func (m *ConfigRequest_V1_System_Proxy) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Proxy.Size(m)
}
func (m *ConfigRequest_V1_System_Proxy) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Proxy.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Proxy proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Proxy) GetConnectionString() *wrappers.StringValue {
	if m != nil {
		return m.ConnectionString
	}
	return nil
}

func (m *ConfigRequest_V1_System_Proxy) GetNoProxyString() *wrappers.StringValue {
	if m != nil {
		return m.NoProxyString
	}
	return nil
}

type ConfigRequest_V1_System_Backup struct {
	Filesystem           *ConfigRequest_V1_System_Backup_Filesystem `protobuf:"bytes,1,opt,name=filesystem,proto3" json:"filesystem,omitempty" toml:"filesystem,omitempty" mapstructure:"filesystem,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                                     `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                                      `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Backup) Reset()         { *m = ConfigRequest_V1_System_Backup{} }
func (m *ConfigRequest_V1_System_Backup) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Backup) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Backup) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_3799ee30fbfa1658, []int{0, 0, 0, 4}
}
func (m *ConfigRequest_V1_System_Backup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Backup.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Backup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Backup.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest_V1_System_Backup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Backup.Merge(dst, src)
}
func (m *ConfigRequest_V1_System_Backup) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Backup.Size(m)
}
func (m *ConfigRequest_V1_System_Backup) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Backup.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Backup proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Backup) GetFilesystem() *ConfigRequest_V1_System_Backup_Filesystem {
	if m != nil {
		return m.Filesystem
	}
	return nil
}

type ConfigRequest_V1_System_Backup_Filesystem struct {
	Path                 *wrappers.StringValue `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty" toml:"path,omitempty" mapstructure:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Backup_Filesystem) Reset() {
	*m = ConfigRequest_V1_System_Backup_Filesystem{}
}
func (m *ConfigRequest_V1_System_Backup_Filesystem) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Backup_Filesystem) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Backup_Filesystem) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_3799ee30fbfa1658, []int{0, 0, 0, 4, 0}
}
func (m *ConfigRequest_V1_System_Backup_Filesystem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Backup_Filesystem.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Backup_Filesystem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Backup_Filesystem.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest_V1_System_Backup_Filesystem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Backup_Filesystem.Merge(dst, src)
}
func (m *ConfigRequest_V1_System_Backup_Filesystem) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Backup_Filesystem.Size(m)
}
func (m *ConfigRequest_V1_System_Backup_Filesystem) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Backup_Filesystem.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Backup_Filesystem proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Backup_Filesystem) GetPath() *wrappers.StringValue {
	if m != nil {
		return m.Path
	}
	return nil
}

type ConfigRequest_V1_Service struct {
	Name                            *wrappers.StringValue       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	Origin                          *wrappers.StringValue       `protobuf:"bytes,2,opt,name=origin,proto3" json:"origin,omitempty" toml:"origin,omitempty" mapstructure:"origin,omitempty"`
	Channel                         *wrappers.StringValue       `protobuf:"bytes,3,opt,name=channel,proto3" json:"channel,omitempty" toml:"channel,omitempty" mapstructure:"channel,omitempty"`
	UpgradeStrategy                 *wrappers.StringValue       `protobuf:"bytes,4,opt,name=upgrade_strategy,json=upgradeStrategy,proto3" json:"upgrade_strategy,omitempty" toml:"upgrade_strategy,omitempty" mapstructure:"upgrade_strategy,omitempty"`
	DeploymentType                  *wrappers.StringValue       `protobuf:"bytes,5,opt,name=deployment_type,json=deploymentType,proto3" json:"deployment_type,omitempty" toml:"deployment_type,omitempty" mapstructure:"deployment_type,omitempty"`
	OverrideOrigin                  *wrappers.StringValue       `protobuf:"bytes,6,opt,name=override_origin,json=overrideOrigin,proto3" json:"override_origin,omitempty" toml:"override_origin,omitempty" mapstructure:"override_origin,omitempty"`
	HartifactsPath                  *wrappers.StringValue       `protobuf:"bytes,7,opt,name=hartifacts_path,json=hartifactsPath,proto3" json:"hartifacts_path,omitempty" toml:"hartifacts_path,omitempty" mapstructure:"hartifacts_path,omitempty"`
	AdminUser                       *ConfigRequest_V1_AdminUser `protobuf:"bytes,8,opt,name=admin_user,json=adminUser,proto3" json:"admin_user,omitempty" toml:"admin_user,omitempty" mapstructure:"admin_user,omitempty"`
	ManifestCacheExpiry             *wrappers.StringValue       `protobuf:"bytes,9,opt,name=manifest_cache_expiry,json=manifestCacheExpiry,proto3" json:"manifest_cache_expiry,omitempty" toml:"manifest_cache_expiry,omitempty" mapstructure:"manifest_cache_expiry,omitempty"`
	ManifestDirectory               *wrappers.StringValue       `protobuf:"bytes,10,opt,name=manifest_directory,json=manifestDirectory,proto3" json:"manifest_directory,omitempty" toml:"manifest_directory,omitempty" mapstructure:"manifest_directory,omitempty"`
	EnableChefServer                *wrappers.BoolValue         `protobuf:"bytes,11,opt,name=enable_chef_server,json=enableChefServer,proto3" json:"enable_chef_server,omitempty" toml:"enable_chef_server,omitempty" mapstructure:"enable_chef_server,omitempty"`
	EnableDeploymentOrderStressMode *wrappers.BoolValue         `protobuf:"bytes,12,opt,name=enable_deployment_order_stress_mode,json=enableDeploymentOrderStressMode,proto3" json:"enable_deployment_order_stress_mode,omitempty" toml:"enable_deployment_order_stress_mode,omitempty" mapstructure:"enable_deployment_order_stress_mode,omitempty"`
	FeatureFlagS3Backups            *wrappers.BoolValue         `protobuf:"bytes,512,opt,name=feature_flag_s3_backups,json=featureFlagS3Backups,proto3" json:"feature_flag_s3_backups,omitempty" toml:"feature_flag_s3_backups,omitempty" mapstructure:"feature_flag_s3_backups,omitempty"`
	EnableWorkflow                  *wrappers.BoolValue         `protobuf:"bytes,13,opt,name=enable_workflow,json=enableWorkflow,proto3" json:"enable_workflow,omitempty" toml:"enable_workflow,omitempty" mapstructure:"enable_workflow,omitempty"`
	PackageCleanupMode              *wrappers.StringValue       `protobuf:"bytes,14,opt,name=package_cleanup_mode,json=packageCleanupMode,proto3" json:"package_cleanup_mode,omitempty" toml:"package_cleanup_mode,omitempty" mapstructure:"package_cleanup_mode,omitempty"`
	EnableDevMonitoring             *wrappers.BoolValue         `protobuf:"bytes,15,opt,name=enable_dev_monitoring,json=enableDevMonitoring,proto3" json:"enable_dev_monitoring,omitempty" toml:"enable_dev_monitoring,omitempty" mapstructure:"enable_dev_monitoring,omitempty"`
	XXX_NoUnkeyedLiteral            struct{}                    `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized                []byte                      `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache                   int32                       `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_Service) Reset()         { *m = ConfigRequest_V1_Service{} }
func (m *ConfigRequest_V1_Service) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_Service) ProtoMessage()    {}
func (*ConfigRequest_V1_Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_3799ee30fbfa1658, []int{0, 0, 1}
}
func (m *ConfigRequest_V1_Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_Service.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_Service.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest_V1_Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_Service.Merge(dst, src)
}
func (m *ConfigRequest_V1_Service) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_Service.Size(m)
}
func (m *ConfigRequest_V1_Service) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_Service.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_Service proto.InternalMessageInfo

func (m *ConfigRequest_V1_Service) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ConfigRequest_V1_Service) GetOrigin() *wrappers.StringValue {
	if m != nil {
		return m.Origin
	}
	return nil
}

func (m *ConfigRequest_V1_Service) GetChannel() *wrappers.StringValue {
	if m != nil {
		return m.Channel
	}
	return nil
}

func (m *ConfigRequest_V1_Service) GetUpgradeStrategy() *wrappers.StringValue {
	if m != nil {
		return m.UpgradeStrategy
	}
	return nil
}

func (m *ConfigRequest_V1_Service) GetDeploymentType() *wrappers.StringValue {
	if m != nil {
		return m.DeploymentType
	}
	return nil
}

func (m *ConfigRequest_V1_Service) GetOverrideOrigin() *wrappers.StringValue {
	if m != nil {
		return m.OverrideOrigin
	}
	return nil
}

func (m *ConfigRequest_V1_Service) GetHartifactsPath() *wrappers.StringValue {
	if m != nil {
		return m.HartifactsPath
	}
	return nil
}

func (m *ConfigRequest_V1_Service) GetAdminUser() *ConfigRequest_V1_AdminUser {
	if m != nil {
		return m.AdminUser
	}
	return nil
}

func (m *ConfigRequest_V1_Service) GetManifestCacheExpiry() *wrappers.StringValue {
	if m != nil {
		return m.ManifestCacheExpiry
	}
	return nil
}

func (m *ConfigRequest_V1_Service) GetManifestDirectory() *wrappers.StringValue {
	if m != nil {
		return m.ManifestDirectory
	}
	return nil
}

func (m *ConfigRequest_V1_Service) GetEnableChefServer() *wrappers.BoolValue {
	if m != nil {
		return m.EnableChefServer
	}
	return nil
}

func (m *ConfigRequest_V1_Service) GetEnableDeploymentOrderStressMode() *wrappers.BoolValue {
	if m != nil {
		return m.EnableDeploymentOrderStressMode
	}
	return nil
}

func (m *ConfigRequest_V1_Service) GetFeatureFlagS3Backups() *wrappers.BoolValue {
	if m != nil {
		return m.FeatureFlagS3Backups
	}
	return nil
}

func (m *ConfigRequest_V1_Service) GetEnableWorkflow() *wrappers.BoolValue {
	if m != nil {
		return m.EnableWorkflow
	}
	return nil
}

func (m *ConfigRequest_V1_Service) GetPackageCleanupMode() *wrappers.StringValue {
	if m != nil {
		return m.PackageCleanupMode
	}
	return nil
}

func (m *ConfigRequest_V1_Service) GetEnableDevMonitoring() *wrappers.BoolValue {
	if m != nil {
		return m.EnableDevMonitoring
	}
	return nil
}

type ConfigRequest_V1_AdminUser struct {
	// use 'name' instead of 'email' for new code. It hasn't been reserved
	// because we'll support using 'email' from installations.
	Email                *wrappers.StringValue `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty" toml:"email,omitempty" mapstructure:"email,omitempty"` // Deprecated: Do not use.
	Username             *wrappers.StringValue `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty" toml:"username,omitempty" mapstructure:"username,omitempty"`
	Name                 *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	Password             *wrappers.StringValue `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty" toml:"password,omitempty" mapstructure:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_AdminUser) Reset()         { *m = ConfigRequest_V1_AdminUser{} }
func (m *ConfigRequest_V1_AdminUser) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_AdminUser) ProtoMessage()    {}
func (*ConfigRequest_V1_AdminUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_3799ee30fbfa1658, []int{0, 0, 2}
}
func (m *ConfigRequest_V1_AdminUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_AdminUser.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_AdminUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_AdminUser.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest_V1_AdminUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_AdminUser.Merge(dst, src)
}
func (m *ConfigRequest_V1_AdminUser) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_AdminUser.Size(m)
}
func (m *ConfigRequest_V1_AdminUser) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_AdminUser.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_AdminUser proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *ConfigRequest_V1_AdminUser) GetEmail() *wrappers.StringValue {
	if m != nil {
		return m.Email
	}
	return nil
}

func (m *ConfigRequest_V1_AdminUser) GetUsername() *wrappers.StringValue {
	if m != nil {
		return m.Username
	}
	return nil
}

func (m *ConfigRequest_V1_AdminUser) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ConfigRequest_V1_AdminUser) GetPassword() *wrappers.StringValue {
	if m != nil {
		return m.Password
	}
	return nil
}

func init() {
	proto.RegisterType((*ConfigRequest)(nil), "chef.automate.domain.deployment.ConfigRequest")
	proto.RegisterType((*ConfigRequest_V1)(nil), "chef.automate.domain.deployment.ConfigRequest.V1")
	proto.RegisterType((*ConfigRequest_V1_System)(nil), "chef.automate.domain.deployment.ConfigRequest.V1.System")
	proto.RegisterType((*ConfigRequest_V1_System_Service)(nil), "chef.automate.domain.deployment.ConfigRequest.V1.System.Service")
	proto.RegisterType((*ConfigRequest_V1_System_Log)(nil), "chef.automate.domain.deployment.ConfigRequest.V1.System.Log")
	proto.RegisterType((*ConfigRequest_V1_System_GatherLogs)(nil), "chef.automate.domain.deployment.ConfigRequest.V1.System.GatherLogs")
	proto.RegisterType((*ConfigRequest_V1_System_Proxy)(nil), "chef.automate.domain.deployment.ConfigRequest.V1.System.Proxy")
	proto.RegisterType((*ConfigRequest_V1_System_Backup)(nil), "chef.automate.domain.deployment.ConfigRequest.V1.System.Backup")
	proto.RegisterType((*ConfigRequest_V1_System_Backup_Filesystem)(nil), "chef.automate.domain.deployment.ConfigRequest.V1.System.Backup.Filesystem")
	proto.RegisterType((*ConfigRequest_V1_Service)(nil), "chef.automate.domain.deployment.ConfigRequest.V1.Service")
	proto.RegisterType((*ConfigRequest_V1_AdminUser)(nil), "chef.automate.domain.deployment.ConfigRequest.V1.AdminUser")
}

func init() {
	proto.RegisterFile("api/config/deployment/config_request.proto", fileDescriptor_config_request_3799ee30fbfa1658)
}

var fileDescriptor_config_request_3799ee30fbfa1658 = []byte{
	// 1093 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x97, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x80, 0xa1, 0x1f, 0xcb, 0xf6, 0xb8, 0xb1, 0x9c, 0xcd, 0x4f, 0x09, 0xb6, 0x48, 0x82, 0xf6,
	0x52, 0x14, 0x30, 0x15, 0xdb, 0x41, 0x90, 0xf4, 0x27, 0x6d, 0x2c, 0x3b, 0x6e, 0xfc, 0x0f, 0x2a,
	0x75, 0x00, 0x5f, 0x88, 0x15, 0x39, 0xa2, 0x58, 0x93, 0xbb, 0xec, 0xee, 0x4a, 0x8e, 0x6e, 0x7d,
	0x8b, 0xf6, 0x09, 0x7a, 0x6d, 0x0b, 0xf4, 0x09, 0xf2, 0x36, 0x3d, 0xfb, 0x05, 0x0a, 0x2e, 0x97,
	0x94, 0xdb, 0x14, 0x36, 0xa3, 0xde, 0x24, 0x6a, 0xbe, 0x8f, 0xb3, 0x33, 0xc3, 0xa1, 0x0d, 0x9f,
	0xd3, 0x34, 0xea, 0xf8, 0x9c, 0x0d, 0xa2, 0xb0, 0x13, 0x60, 0x1a, 0xf3, 0x49, 0x82, 0x4c, 0x99,
	0x2b, 0x9e, 0xc0, 0x1f, 0x47, 0x28, 0x95, 0x93, 0x0a, 0xae, 0x38, 0xb9, 0xef, 0x0f, 0x71, 0xe0,
	0xd0, 0x91, 0xe2, 0x09, 0x55, 0xe8, 0x04, 0x3c, 0xa1, 0x11, 0x73, 0xa6, 0x94, 0x7d, 0xef, 0x92,
	0x4c, 0x0e, 0xa9, 0xc0, 0xa0, 0x13, 0xc6, 0xbc, 0x4f, 0xe3, 0x5c, 0x60, 0xdf, 0x0b, 0x39, 0x0f,
	0x63, 0xec, 0xe8, 0x6f, 0xfd, 0xd1, 0xa0, 0x73, 0x2e, 0x68, 0x9a, 0xa2, 0x90, 0xe6, 0xf7, 0x5d,
	0x9f, 0x27, 0x29, 0x67, 0xc8, 0x94, 0xec, 0x14, 0xb7, 0x59, 0x0d, 0x45, 0xea, 0xe7, 0x84, 0xbf,
	0x1a, 0x22, 0x5b, 0xa5, 0xeb, 0xab, 0xe6, 0x16, 0xd9, 0xdd, 0xe8, 0x7a, 0xf6, 0xa5, 0x43, 0x19,
	0xe3, 0x8a, 0xaa, 0x88, 0x33, 0xe3, 0xfa, 0xe4, 0xd7, 0xbb, 0x70, 0xa3, 0xab, 0xe3, 0xdc, 0xfc,
	0x10, 0xe4, 0x39, 0xd4, 0xc7, 0x6b, 0x56, 0xe3, 0x41, 0xed, 0xb3, 0xa5, 0xf5, 0x35, 0xe7, 0x9a,
	0xb3, 0x38, 0xff, 0x60, 0x9d, 0x93, 0x35, 0xb7, 0x3e, 0x5e, 0xb3, 0xff, 0xb8, 0x03, 0xf5, 0x93,
	0x35, 0xb2, 0x0b, 0x0d, 0x39, 0x91, 0x56, 0x4d, 0xab, 0x9e, 0xbc, 0xb7, 0xca, 0xe9, 0x4d, 0xa4,
	0xc2, 0xc4, 0xcd, 0x24, 0x64, 0x0f, 0x1a, 0x72, 0xec, 0x5b, 0x75, 0xed, 0x7a, 0x3a, 0x83, 0x0b,
	0xc5, 0x38, 0xf2, 0xd1, 0xcd, 0x2c, 0xf6, 0xc5, 0x02, 0xb4, 0x72, 0x39, 0x79, 0x04, 0xcd, 0x24,
	0x96, 0xd4, 0x24, 0xf9, 0xe0, 0x5f, 0xe2, 0x88, 0x0d, 0x04, 0x75, 0xf2, 0x3a, 0x3a, 0x07, 0xb1,
	0xa4, 0xae, 0x8e, 0x26, 0xa7, 0x30, 0x2f, 0x73, 0xa1, 0xc9, 0xe8, 0xdb, 0x59, 0x4f, 0x57, 0x26,
	0x56, 0x08, 0xc9, 0x21, 0x34, 0x62, 0x1e, 0x9a, 0x06, 0x7c, 0x35, 0xb3, 0x77, 0x9f, 0x87, 0x6e,
	0x26, 0x22, 0x01, 0x2c, 0x85, 0x54, 0x0d, 0x51, 0x78, 0x31, 0x0f, 0xa5, 0xd5, 0xd4, 0xde, 0xee,
	0xcc, 0xde, 0x1d, 0xed, 0xda, 0xe7, 0xa1, 0x74, 0x21, 0x2c, 0x3f, 0x93, 0x57, 0x30, 0x97, 0x0a,
	0xfe, 0x66, 0x62, 0xcd, 0x69, 0xff, 0xb3, 0x99, 0xfd, 0xc7, 0x99, 0xc5, 0xcd, 0x65, 0xe4, 0x35,
	0xb4, 0xfa, 0xd4, 0x3f, 0x1b, 0xa5, 0x56, 0x4b, 0x6b, 0xbf, 0x99, 0x59, 0xbb, 0xa9, 0x35, 0xae,
	0xd1, 0xd9, 0x3f, 0xd7, 0x60, 0xde, 0x54, 0x9e, 0x74, 0x61, 0x39, 0x8e, 0xa4, 0x42, 0xe6, 0xd1,
	0x20, 0x10, 0x28, 0x8b, 0x89, 0xfd, 0xd8, 0xc9, 0x9f, 0x43, 0xa7, 0x78, 0x0e, 0x9d, 0x9e, 0x12,
	0x11, 0x0b, 0x4f, 0x68, 0x3c, 0x42, 0xf7, 0x46, 0xce, 0x3c, 0xcf, 0x11, 0xb2, 0x03, 0xcd, 0x94,
	0x0b, 0x65, 0xc6, 0xe1, 0xa3, 0x77, 0xd0, 0x97, 0x4c, 0x6d, 0xac, 0x6b, 0x72, 0xf3, 0xee, 0xdb,
	0x0b, 0x8b, 0x94, 0x03, 0xb4, 0xf2, 0xfb, 0x91, 0xdd, 0xcc, 0x1e, 0x5e, 0x57, 0x0b, 0xec, 0xa7,
	0xd0, 0xd8, 0xe7, 0x21, 0x59, 0x87, 0xb9, 0x18, 0xc7, 0x18, 0x57, 0xca, 0x25, 0x0f, 0xb5, 0xf7,
	0x00, 0xa6, 0xdd, 0x21, 0x5f, 0xc3, 0x92, 0x54, 0x34, 0x8c, 0x58, 0xe8, 0x05, 0x91, 0xa8, 0xe4,
	0x01, 0x03, 0x6c, 0x45, 0xc2, 0xfe, 0xa5, 0x06, 0x73, 0xba, 0x17, 0xe4, 0x25, 0xdc, 0xf4, 0x39,
	0x63, 0xe8, 0x67, 0x7b, 0xc3, 0x93, 0x3a, 0xbe, 0x92, 0x6e, 0x65, 0x8a, 0xe5, 0x97, 0xc9, 0x16,
	0xb4, 0x19, 0xf7, 0x74, 0x6f, 0x0b, 0x51, 0xbd, 0x4a, 0xad, 0x19, 0xd7, 0xa9, 0xe4, 0xd7, 0xec,
	0x3f, 0x6b, 0xd0, 0xca, 0xfb, 0x49, 0x7e, 0x00, 0x18, 0x44, 0x31, 0x4a, 0xdd, 0x64, 0x93, 0xd4,
	0xee, 0xff, 0x1c, 0x12, 0xe7, 0x45, 0x69, 0x74, 0x2f, 0xd9, 0xed, 0x67, 0x00, 0xd3, 0x5f, 0xc8,
	0x43, 0x68, 0xa6, 0x54, 0x0d, 0x2b, 0x15, 0x42, 0x47, 0xda, 0xbf, 0x2d, 0x4e, 0x67, 0xee, 0x21,
	0x34, 0x19, 0x4d, 0xb0, 0x1a, 0x9d, 0x45, 0x92, 0x47, 0xd0, 0xe2, 0x22, 0x0a, 0x23, 0x56, 0xa9,
	0x62, 0x26, 0x96, 0x3c, 0x86, 0x79, 0x7f, 0x48, 0x19, 0xc3, 0xd8, 0x2c, 0x94, 0xab, 0xb1, 0x22,
	0x98, 0xec, 0xc0, 0xca, 0x28, 0x0d, 0x05, 0x0d, 0x30, 0xeb, 0x13, 0x55, 0x18, 0x4e, 0xcc, 0xe6,
	0xb8, 0x5a, 0xd0, 0x36, 0x54, 0xcf, 0x40, 0x64, 0x1b, 0xda, 0xd3, 0xc2, 0x7b, 0x6a, 0x92, 0xa2,
	0xd9, 0x10, 0x57, 0x7b, 0x96, 0xa7, 0xd0, 0xab, 0x49, 0x8a, 0x99, 0x86, 0x8f, 0x51, 0x88, 0x28,
	0x40, 0xcf, 0x94, 0xa1, 0x55, 0x45, 0x53, 0x40, 0x47, 0x79, 0x39, 0xb6, 0xa1, 0x3d, 0xa4, 0x42,
	0x45, 0x03, 0xea, 0x2b, 0xe9, 0xe9, 0xfe, 0xcd, 0x57, 0xd1, 0x4c, 0xa1, 0x63, 0xaa, 0x86, 0xe4,
	0x14, 0x80, 0x06, 0x49, 0xc4, 0xbc, 0x91, 0x44, 0x61, 0x2d, 0x68, 0xc3, 0x97, 0xef, 0x3f, 0x75,
	0xcf, 0x33, 0xc7, 0xf7, 0x12, 0x85, 0xbb, 0x48, 0x8b, 0x8f, 0xe4, 0x18, 0xee, 0x24, 0x94, 0x45,
	0x03, 0x94, 0xca, 0xf3, 0xa9, 0x3f, 0x44, 0x0f, 0xdf, 0xa4, 0x91, 0x98, 0x58, 0x8b, 0x15, 0x12,
	0xbd, 0x55, 0xa0, 0xdd, 0x8c, 0xdc, 0xd6, 0x20, 0xd9, 0x03, 0x52, 0x1a, 0x83, 0x48, 0xa0, 0xaf,
	0xb8, 0x98, 0x58, 0x50, 0x41, 0x77, 0xb3, 0xe0, 0xb6, 0x0a, 0x8c, 0x7c, 0x07, 0x04, 0x19, 0xed,
	0xc7, 0xe8, 0x65, 0xc7, 0xf5, 0xb2, 0x25, 0x86, 0xc2, 0x5a, 0xd2, 0x32, 0xfb, 0x1d, 0xd9, 0x26,
	0xe7, 0xb1, 0xd9, 0x05, 0x39, 0xd5, 0x1d, 0xe2, 0xa0, 0xa7, 0x19, 0x32, 0x84, 0x4f, 0x8d, 0xe9,
	0xd2, 0x80, 0x70, 0x11, 0xa0, 0xc8, 0x46, 0x0e, 0xa5, 0xf4, 0x12, 0x1e, 0xa0, 0xf5, 0xc1, 0xb5,
	0xea, 0xfb, 0xb9, 0x66, 0xab, 0xb4, 0x1c, 0x65, 0x92, 0x9e, 0x76, 0x1c, 0xf0, 0x00, 0x89, 0x0b,
	0x1f, 0x0e, 0x90, 0xaa, 0x91, 0x40, 0x6f, 0x10, 0xd3, 0xd0, 0x93, 0x1b, 0x5e, 0xfe, 0x1a, 0x90,
	0xd6, 0x4f, 0xcd, 0x6b, 0xf5, 0xb7, 0x0d, 0xfb, 0x22, 0xa6, 0x61, 0x6f, 0x23, 0x5f, 0x11, 0x92,
	0x74, 0xa1, 0x6d, 0xb2, 0x3f, 0xe7, 0xe2, 0x6c, 0x10, 0xf3, 0x73, 0xeb, 0xc6, 0xb5, 0xaa, 0xe5,
	0x1c, 0x79, 0x6d, 0x08, 0x72, 0x08, 0xb7, 0x53, 0xea, 0x9f, 0xd1, 0x10, 0x3d, 0x3f, 0x46, 0xca,
	0x46, 0x69, 0x7e, 0xe6, 0xe5, 0x0a, 0xbd, 0x21, 0x86, 0xec, 0xe6, 0xa0, 0x3e, 0xe8, 0x21, 0xdc,
	0x29, 0x4b, 0x3a, 0xf6, 0x12, 0xce, 0x22, 0xc5, 0xf5, 0x92, 0x6d, 0x5f, 0x9b, 0xda, 0xad, 0xa2,
	0x88, 0xe3, 0x83, 0x12, 0xb3, 0xff, 0xaa, 0xc1, 0x62, 0x39, 0xa4, 0xe4, 0x31, 0xcc, 0x61, 0x42,
	0xa3, 0x4a, 0xaf, 0xa4, 0xcd, 0xba, 0x55, 0x73, 0xf3, 0x70, 0xf2, 0x04, 0x16, 0xb2, 0xe7, 0x44,
	0xef, 0xbb, 0x2a, 0xbb, 0xab, 0x8c, 0x2e, 0xb7, 0x64, 0xb3, 0xf2, 0x96, 0x7c, 0x02, 0x0b, 0x29,
	0x95, 0xf2, 0x9c, 0x8b, 0xa0, 0xd2, 0xc2, 0x2b, 0xa3, 0xbf, 0xb0, 0xde, 0x5e, 0x58, 0xb7, 0x81,
	0x4c, 0x67, 0x71, 0xd5, 0xbc, 0xa0, 0x77, 0x9b, 0x0b, 0xb5, 0x95, 0xc6, 0xe6, 0xc3, 0x53, 0x27,
	0x8c, 0xd4, 0x70, 0xd4, 0x77, 0x7c, 0x9e, 0x74, 0xb2, 0xe1, 0x2f, 0xff, 0xf6, 0xee, 0xfc, 0xe7,
	0x3f, 0x07, 0xfd, 0x96, 0xbe, 0xe3, 0xc6, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd2, 0xdd, 0x05,
	0x25, 0x3c, 0x0c, 0x00, 0x00,
}
