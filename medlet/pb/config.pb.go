// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: config.proto

/*
Package medletpb is a generated protocol buffer package.

It is generated from these files:
	config.proto

It has these top-level messages:
	Config
	NetworkConfig
	ChainConfig
	RPCConfig
	AppConfig
	PprofConfig
	MiscConfig
*/
package medletpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Med global configurations.
type Config struct {
	// Network config.
	Network *NetworkConfig `protobuf:"bytes,1,opt,name=network" json:"network,omitempty"`
	// Chain config.
	Chain *ChainConfig `protobuf:"bytes,2,opt,name=chain" json:"chain,omitempty"`
	// RPC config.
	Rpc *RPCConfig `protobuf:"bytes,3,opt,name=rpc" json:"rpc,omitempty"`
	// Misc config.
	Misc *MiscConfig `protobuf:"bytes,100,opt,name=misc" json:"misc,omitempty"`
	// App Config.
	App *AppConfig `protobuf:"bytes,101,opt,name=app" json:"app,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

func (m *Config) GetNetwork() *NetworkConfig {
	if m != nil {
		return m.Network
	}
	return nil
}

func (m *Config) GetChain() *ChainConfig {
	if m != nil {
		return m.Chain
	}
	return nil
}

func (m *Config) GetRpc() *RPCConfig {
	if m != nil {
		return m.Rpc
	}
	return nil
}

func (m *Config) GetMisc() *MiscConfig {
	if m != nil {
		return m.Misc
	}
	return nil
}

func (m *Config) GetApp() *AppConfig {
	if m != nil {
		return m.App
	}
	return nil
}

type NetworkConfig struct {
	// Med seed node address.
	Seed []string `protobuf:"bytes,1,rep,name=seed" json:"seed,omitempty"`
	// Listen addresses.
	Listen []string `protobuf:"bytes,2,rep,name=listen" json:"listen,omitempty"`
	// Network node privateKey address. If nil, generate a new node.
	PrivateKey string `protobuf:"bytes,3,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	// Network ID
	NetworkId uint32 `protobuf:"varint,4,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
}

func (m *NetworkConfig) Reset()                    { *m = NetworkConfig{} }
func (m *NetworkConfig) String() string            { return proto.CompactTextString(m) }
func (*NetworkConfig) ProtoMessage()               {}
func (*NetworkConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{1} }

func (m *NetworkConfig) GetSeed() []string {
	if m != nil {
		return m.Seed
	}
	return nil
}

func (m *NetworkConfig) GetListen() []string {
	if m != nil {
		return m.Listen
	}
	return nil
}

func (m *NetworkConfig) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

func (m *NetworkConfig) GetNetworkId() uint32 {
	if m != nil {
		return m.NetworkId
	}
	return 0
}

type ChainConfig struct {
	// ChainID.
	ChainId uint32 `protobuf:"varint,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// genesis conf file path
	Genesis string `protobuf:"bytes,2,opt,name=genesis,proto3" json:"genesis,omitempty"`
	// Data dir.
	Datadir string `protobuf:"bytes,11,opt,name=datadir,proto3" json:"datadir,omitempty"`
	// Key dir.
	Keydir string `protobuf:"bytes,12,opt,name=keydir,proto3" json:"keydir,omitempty"`
	// start mine at launch
	StartMine bool `protobuf:"varint,20,opt,name=start_mine,json=startMine,proto3" json:"start_mine,omitempty"`
	// Coinbase.
	Coinbase string `protobuf:"bytes,21,opt,name=coinbase,proto3" json:"coinbase,omitempty"`
	// Miner.
	Miner string `protobuf:"bytes,22,opt,name=miner,proto3" json:"miner,omitempty"`
	// Passphrase.
	Passphrase string `protobuf:"bytes,23,opt,name=passphrase,proto3" json:"passphrase,omitempty"`
	// Supported signature cipher list. ["ECC_SECP256K1"]
	SignatureCiphers []string `protobuf:"bytes,24,rep,name=signature_ciphers,json=signatureCiphers" json:"signature_ciphers,omitempty"`
}

func (m *ChainConfig) Reset()                    { *m = ChainConfig{} }
func (m *ChainConfig) String() string            { return proto.CompactTextString(m) }
func (*ChainConfig) ProtoMessage()               {}
func (*ChainConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{2} }

func (m *ChainConfig) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *ChainConfig) GetGenesis() string {
	if m != nil {
		return m.Genesis
	}
	return ""
}

func (m *ChainConfig) GetDatadir() string {
	if m != nil {
		return m.Datadir
	}
	return ""
}

func (m *ChainConfig) GetKeydir() string {
	if m != nil {
		return m.Keydir
	}
	return ""
}

func (m *ChainConfig) GetStartMine() bool {
	if m != nil {
		return m.StartMine
	}
	return false
}

func (m *ChainConfig) GetCoinbase() string {
	if m != nil {
		return m.Coinbase
	}
	return ""
}

func (m *ChainConfig) GetMiner() string {
	if m != nil {
		return m.Miner
	}
	return ""
}

func (m *ChainConfig) GetPassphrase() string {
	if m != nil {
		return m.Passphrase
	}
	return ""
}

func (m *ChainConfig) GetSignatureCiphers() []string {
	if m != nil {
		return m.SignatureCiphers
	}
	return nil
}

type RPCConfig struct {
	// RPC listen addresses.
	RpcListen []string `protobuf:"bytes,1,rep,name=rpc_listen,json=rpcListen" json:"rpc_listen,omitempty"`
	// HTTP listen addresses.
	HttpListen []string `protobuf:"bytes,2,rep,name=http_listen,json=httpListen" json:"http_listen,omitempty"`
	// Enabled HTTP modules.["api", "admin"]
	HttpModule []string `protobuf:"bytes,3,rep,name=http_module,json=httpModule" json:"http_module,omitempty"`
	// Connection limit.
	ConnectionLimits int32 `protobuf:"varint,4,opt,name=connection_limits,json=connectionLimits,proto3" json:"connection_limits,omitempty"`
}

func (m *RPCConfig) Reset()                    { *m = RPCConfig{} }
func (m *RPCConfig) String() string            { return proto.CompactTextString(m) }
func (*RPCConfig) ProtoMessage()               {}
func (*RPCConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{3} }

func (m *RPCConfig) GetRpcListen() []string {
	if m != nil {
		return m.RpcListen
	}
	return nil
}

func (m *RPCConfig) GetHttpListen() []string {
	if m != nil {
		return m.HttpListen
	}
	return nil
}

func (m *RPCConfig) GetHttpModule() []string {
	if m != nil {
		return m.HttpModule
	}
	return nil
}

func (m *RPCConfig) GetConnectionLimits() int32 {
	if m != nil {
		return m.ConnectionLimits
	}
	return 0
}

type AppConfig struct {
	// log level
	LogLevel string `protobuf:"bytes,1,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
	// log file path
	LogFile string `protobuf:"bytes,2,opt,name=log_file,json=logFile,proto3" json:"log_file,omitempty"`
	// log file age, unit is s.
	LogAge uint32 `protobuf:"varint,3,opt,name=log_age,json=logAge,proto3" json:"log_age,omitempty"`
	// pprof config
	Pprof *PprofConfig `protobuf:"bytes,4,opt,name=pprof" json:"pprof,omitempty"`
	// App version
	Version string `protobuf:"bytes,100,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *AppConfig) Reset()                    { *m = AppConfig{} }
func (m *AppConfig) String() string            { return proto.CompactTextString(m) }
func (*AppConfig) ProtoMessage()               {}
func (*AppConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{4} }

func (m *AppConfig) GetLogLevel() string {
	if m != nil {
		return m.LogLevel
	}
	return ""
}

func (m *AppConfig) GetLogFile() string {
	if m != nil {
		return m.LogFile
	}
	return ""
}

func (m *AppConfig) GetLogAge() uint32 {
	if m != nil {
		return m.LogAge
	}
	return 0
}

func (m *AppConfig) GetPprof() *PprofConfig {
	if m != nil {
		return m.Pprof
	}
	return nil
}

func (m *AppConfig) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type PprofConfig struct {
	// pprof listen address, if not configured, the function closes.
	HttpListen string `protobuf:"bytes,1,opt,name=http_listen,json=httpListen,proto3" json:"http_listen,omitempty"`
	// cpu profiling file, if not configured, the profiling not start
	Cpuprofile string `protobuf:"bytes,2,opt,name=cpuprofile,proto3" json:"cpuprofile,omitempty"`
	// memory profiling file, if not configured, the profiling not start
	Memprofile string `protobuf:"bytes,3,opt,name=memprofile,proto3" json:"memprofile,omitempty"`
}

func (m *PprofConfig) Reset()                    { *m = PprofConfig{} }
func (m *PprofConfig) String() string            { return proto.CompactTextString(m) }
func (*PprofConfig) ProtoMessage()               {}
func (*PprofConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{5} }

func (m *PprofConfig) GetHttpListen() string {
	if m != nil {
		return m.HttpListen
	}
	return ""
}

func (m *PprofConfig) GetCpuprofile() string {
	if m != nil {
		return m.Cpuprofile
	}
	return ""
}

func (m *PprofConfig) GetMemprofile() string {
	if m != nil {
		return m.Memprofile
	}
	return ""
}

type MiscConfig struct {
	// Default encryption ciper when create new keystore file.
	DefaultKeystoreFileCiper string `protobuf:"bytes,1,opt,name=default_keystore_file_ciper,json=defaultKeystoreFileCiper,proto3" json:"default_keystore_file_ciper,omitempty"`
}

func (m *MiscConfig) Reset()                    { *m = MiscConfig{} }
func (m *MiscConfig) String() string            { return proto.CompactTextString(m) }
func (*MiscConfig) ProtoMessage()               {}
func (*MiscConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{6} }

func (m *MiscConfig) GetDefaultKeystoreFileCiper() string {
	if m != nil {
		return m.DefaultKeystoreFileCiper
	}
	return ""
}

func init() {
	proto.RegisterType((*Config)(nil), "medletpb.Config")
	proto.RegisterType((*NetworkConfig)(nil), "medletpb.NetworkConfig")
	proto.RegisterType((*ChainConfig)(nil), "medletpb.ChainConfig")
	proto.RegisterType((*RPCConfig)(nil), "medletpb.RPCConfig")
	proto.RegisterType((*AppConfig)(nil), "medletpb.AppConfig")
	proto.RegisterType((*PprofConfig)(nil), "medletpb.PprofConfig")
	proto.RegisterType((*MiscConfig)(nil), "medletpb.MiscConfig")
}

func init() { proto.RegisterFile("config.proto", fileDescriptorConfig) }

var fileDescriptorConfig = []byte{
	// 607 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x94, 0xcf, 0x6a, 0xdb, 0x40,
	0x10, 0xc6, 0x51, 0x1c, 0x3b, 0xd6, 0x38, 0x81, 0x74, 0x9b, 0x3f, 0xdb, 0x86, 0xa6, 0xc6, 0x50,
	0x30, 0x04, 0x02, 0x6d, 0xcf, 0x3d, 0x04, 0x43, 0x21, 0x24, 0x29, 0x41, 0x2f, 0x20, 0x14, 0x69,
	0x2c, 0x2f, 0x91, 0xb4, 0xcb, 0xee, 0x3a, 0x25, 0xf4, 0x2d, 0x7a, 0xeb, 0xb9, 0x2f, 0xd6, 0x47,
	0x29, 0x33, 0x5a, 0x59, 0x4e, 0x6e, 0x9a, 0xef, 0xfb, 0x8d, 0x34, 0xfe, 0x66, 0xd7, 0xb0, 0x9f,
	0xeb, 0x66, 0xa9, 0xca, 0x4b, 0x63, 0xb5, 0xd7, 0x62, 0x5c, 0x63, 0x51, 0xa1, 0x37, 0x0f, 0xb3,
	0x7f, 0x11, 0x8c, 0x16, 0x6c, 0x89, 0xcf, 0xb0, 0xd7, 0xa0, 0xff, 0xa9, 0xed, 0xa3, 0x8c, 0xa6,
	0xd1, 0x7c, 0xf2, 0xe5, 0xf4, 0xb2, 0xc3, 0x2e, 0x7f, 0xb4, 0x46, 0x4b, 0x26, 0x1d, 0x27, 0x2e,
	0x60, 0x98, 0xaf, 0x32, 0xd5, 0xc8, 0x1d, 0x6e, 0x38, 0xee, 0x1b, 0x16, 0x24, 0x07, 0xbc, 0x65,
	0xc4, 0x27, 0x18, 0x58, 0x93, 0xcb, 0x01, 0xa3, 0x6f, 0x7b, 0x34, 0xb9, 0x5f, 0x04, 0x90, 0x7c,
	0x31, 0x87, 0xdd, 0x5a, 0xb9, 0x5c, 0x16, 0xcc, 0x1d, 0xf5, 0xdc, 0x9d, 0x72, 0x79, 0x00, 0x99,
	0xa0, 0x17, 0x66, 0xc6, 0x48, 0x7c, 0xfd, 0xc2, 0x2b, 0x63, 0xba, 0x17, 0x66, 0xc6, 0xcc, 0x7e,
	0xc1, 0xc1, 0x8b, 0xf1, 0x85, 0x80, 0x5d, 0x87, 0x58, 0xc8, 0x68, 0x3a, 0x98, 0xc7, 0x09, 0x3f,
	0x8b, 0x13, 0x18, 0x55, 0xca, 0x79, 0xa4, 0x9f, 0x42, 0x6a, 0xa8, 0xc4, 0x47, 0x98, 0x18, 0xab,
	0x9e, 0x32, 0x8f, 0xe9, 0x23, 0x3e, 0xf3, 0xf0, 0x71, 0x02, 0x41, 0xba, 0xc1, 0x67, 0xf1, 0x01,
	0x20, 0xa4, 0x91, 0xaa, 0x42, 0xee, 0x4e, 0xa3, 0xf9, 0x41, 0x12, 0x07, 0xe5, 0xba, 0x98, 0xfd,
	0xde, 0x81, 0xc9, 0x56, 0x16, 0xe2, 0x1d, 0x8c, 0x39, 0x0d, 0x82, 0x23, 0x86, 0xf7, 0xb8, 0xbe,
	0x2e, 0x84, 0x84, 0xbd, 0x12, 0x1b, 0x74, 0xca, 0x71, 0x9c, 0x71, 0xd2, 0x95, 0xe4, 0x14, 0x99,
	0xcf, 0x0a, 0x65, 0xe5, 0xa4, 0x75, 0x42, 0x49, 0x63, 0x3f, 0xe2, 0x33, 0x19, 0xfb, 0x6c, 0x84,
	0x8a, 0xa6, 0x72, 0x3e, 0xb3, 0x3e, 0xad, 0x55, 0x83, 0xf2, 0x68, 0x1a, 0xcd, 0xc7, 0x49, 0xcc,
	0xca, 0x9d, 0x6a, 0x50, 0xbc, 0x87, 0x71, 0xae, 0x55, 0xf3, 0x90, 0x39, 0x94, 0xc7, 0xdc, 0xb8,
	0xa9, 0xc5, 0x11, 0x0c, 0xa9, 0xc9, 0xca, 0x13, 0x36, 0xda, 0x42, 0x9c, 0x03, 0x98, 0xcc, 0x39,
	0xb3, 0xb2, 0xd4, 0x73, 0x1a, 0x62, 0xd8, 0x28, 0xe2, 0x02, 0xde, 0x38, 0x55, 0x36, 0x99, 0x5f,
	0x5b, 0x4c, 0x73, 0x65, 0x56, 0x68, 0x9d, 0x94, 0x1c, 0xe5, 0xe1, 0xc6, 0x58, 0xb4, 0xfa, 0xec,
	0x4f, 0x04, 0xf1, 0x66, 0xeb, 0x34, 0xab, 0x35, 0x79, 0x1a, 0xe2, 0x6f, 0x97, 0x12, 0x5b, 0x93,
	0xdf, 0x6e, 0x36, 0xb0, 0xf2, 0xde, 0xa4, 0x2f, 0xd6, 0x03, 0x24, 0xbd, 0x02, 0x6a, 0x5d, 0xac,
	0x2b, 0x94, 0x83, 0x1e, 0xb8, 0x63, 0x85, 0x66, 0xcb, 0x75, 0xd3, 0x60, 0xee, 0x95, 0x6e, 0xd2,
	0x4a, 0xd5, 0xca, 0x3b, 0xde, 0xd4, 0x30, 0x39, 0xec, 0x8d, 0x5b, 0xd6, 0x67, 0x7f, 0x23, 0x88,
	0x37, 0x07, 0x48, 0x9c, 0x41, 0x5c, 0xe9, 0x32, 0xad, 0xf0, 0x09, 0x2b, 0xde, 0x57, 0x9c, 0x8c,
	0x2b, 0x5d, 0xde, 0x52, 0x4d, 0xbb, 0x24, 0x73, 0xa9, 0x2a, 0xec, 0x36, 0x56, 0xe9, 0xf2, 0xbb,
	0xaa, 0x50, 0x9c, 0x02, 0x3d, 0xa6, 0x59, 0x89, 0x7c, 0x64, 0x0e, 0x92, 0x51, 0xa5, 0xcb, 0xab,
	0x92, 0x66, 0x19, 0x1a, 0x63, 0xf5, 0x92, 0xbf, 0xff, 0xe2, 0xc6, 0xdc, 0x93, 0xdc, 0xdd, 0x18,
	0x66, 0x68, 0xef, 0x4f, 0x68, 0x9d, 0xd2, 0x0d, 0xdf, 0x86, 0x38, 0xe9, 0xca, 0x59, 0x03, 0x93,
	0x2d, 0xfe, 0x75, 0x46, 0xed, 0xa0, 0xdb, 0x19, 0x9d, 0x03, 0xe4, 0x66, 0x4d, 0x1d, 0xfd, 0xb0,
	0x5b, 0x0a, 0xf9, 0x35, 0xd6, 0x9d, 0x1f, 0x4e, 0x79, 0xaf, 0xcc, 0x6e, 0x00, 0xfa, 0xeb, 0x27,
	0xbe, 0xc1, 0x59, 0x81, 0xcb, 0x6c, 0x5d, 0x79, 0xba, 0x14, 0xce, 0x6b, 0x8b, 0x9c, 0x02, 0x2d,
	0x1e, 0x6d, 0xf8, 0xbc, 0x0c, 0xc8, 0x4d, 0x20, 0x28, 0x97, 0x05, 0xf9, 0x0f, 0x23, 0xfe, 0x13,
	0xfa, 0xfa, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x88, 0xcf, 0x3a, 0x94, 0x04, 0x00, 0x00,
}