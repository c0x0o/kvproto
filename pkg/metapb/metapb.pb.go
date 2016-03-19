// Code generated by protoc-gen-go.
// source: metapb.proto
// DO NOT EDIT!

/*
Package metapb is a generated protocol buffer package.

It is generated from these files:
	metapb.proto

It has these top-level messages:
	Cluster
	Node
	Store
	Peer
	Region
*/
package metapb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Cluster struct {
	ClusterId *uint64 `protobuf:"varint,1,opt,name=cluster_id" json:"cluster_id,omitempty"`
	// max peer number for a region.
	// pd will do the auto-balance if region peer number mismatches.
	MaxPeerNumber    *uint32 `protobuf:"varint,2,opt,name=max_peer_number" json:"max_peer_number,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Cluster) Reset()                    { *m = Cluster{} }
func (m *Cluster) String() string            { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()               {}
func (*Cluster) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Cluster) GetClusterId() uint64 {
	if m != nil && m.ClusterId != nil {
		return *m.ClusterId
	}
	return 0
}

func (m *Cluster) GetMaxPeerNumber() uint32 {
	if m != nil && m.MaxPeerNumber != nil {
		return *m.MaxPeerNumber
	}
	return 0
}

type Node struct {
	NodeId           *uint64 `protobuf:"varint,1,opt,name=node_id" json:"node_id,omitempty"`
	Address          *string `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Node) GetNodeId() uint64 {
	if m != nil && m.NodeId != nil {
		return *m.NodeId
	}
	return 0
}

func (m *Node) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

type Store struct {
	StoreId          *uint64 `protobuf:"varint,1,opt,name=store_id" json:"store_id,omitempty"`
	NodeId           *uint64 `protobuf:"varint,2,opt,name=node_id" json:"node_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Store) Reset()                    { *m = Store{} }
func (m *Store) String() string            { return proto.CompactTextString(m) }
func (*Store) ProtoMessage()               {}
func (*Store) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Store) GetStoreId() uint64 {
	if m != nil && m.StoreId != nil {
		return *m.StoreId
	}
	return 0
}

func (m *Store) GetNodeId() uint64 {
	if m != nil && m.NodeId != nil {
		return *m.NodeId
	}
	return 0
}

type Peer struct {
	PeerId           *uint64 `protobuf:"varint,1,opt,name=peer_id" json:"peer_id,omitempty"`
	NodeId           *uint64 `protobuf:"varint,2,opt,name=node_id" json:"node_id,omitempty"`
	StoreId          *uint64 `protobuf:"varint,3,opt,name=store_id" json:"store_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Peer) Reset()                    { *m = Peer{} }
func (m *Peer) String() string            { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()               {}
func (*Peer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Peer) GetPeerId() uint64 {
	if m != nil && m.PeerId != nil {
		return *m.PeerId
	}
	return 0
}

func (m *Peer) GetNodeId() uint64 {
	if m != nil && m.NodeId != nil {
		return *m.NodeId
	}
	return 0
}

func (m *Peer) GetStoreId() uint64 {
	if m != nil && m.StoreId != nil {
		return *m.StoreId
	}
	return 0
}

type Region struct {
	RegionId *uint64 `protobuf:"varint,1,opt,name=region_id" json:"region_id,omitempty"`
	// Region key range [start_key, end_key).
	StartKey []byte `protobuf:"bytes,2,opt,name=start_key" json:"start_key,omitempty"`
	EndKey   []byte `protobuf:"bytes,3,opt,name=end_key" json:"end_key,omitempty"`
	// Conf change version, auto increment when add or remove peer
	PeerConfVer *uint64 `protobuf:"varint,4,opt,name=peer_conf_ver" json:"peer_conf_ver,omitempty"`
	// Region version, auto increment when split or merge
	Version          *uint64 `protobuf:"varint,5,opt,name=version" json:"version,omitempty"`
	Peers            []*Peer `protobuf:"bytes,6,rep,name=peers" json:"peers,omitempty"`
	MaxPeerId        *uint64 `protobuf:"varint,7,opt,name=max_peer_id" json:"max_peer_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Region) Reset()                    { *m = Region{} }
func (m *Region) String() string            { return proto.CompactTextString(m) }
func (*Region) ProtoMessage()               {}
func (*Region) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Region) GetRegionId() uint64 {
	if m != nil && m.RegionId != nil {
		return *m.RegionId
	}
	return 0
}

func (m *Region) GetStartKey() []byte {
	if m != nil {
		return m.StartKey
	}
	return nil
}

func (m *Region) GetEndKey() []byte {
	if m != nil {
		return m.EndKey
	}
	return nil
}

func (m *Region) GetPeerConfVer() uint64 {
	if m != nil && m.PeerConfVer != nil {
		return *m.PeerConfVer
	}
	return 0
}

func (m *Region) GetVersion() uint64 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *Region) GetPeers() []*Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

func (m *Region) GetMaxPeerId() uint64 {
	if m != nil && m.MaxPeerId != nil {
		return *m.MaxPeerId
	}
	return 0
}

func init() {
	proto.RegisterType((*Cluster)(nil), "metapb.Cluster")
	proto.RegisterType((*Node)(nil), "metapb.Node")
	proto.RegisterType((*Store)(nil), "metapb.Store")
	proto.RegisterType((*Peer)(nil), "metapb.Peer")
	proto.RegisterType((*Region)(nil), "metapb.Region")
}

var fileDescriptor0 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x50, 0xbd, 0x4e, 0xc3, 0x30,
	0x10, 0x56, 0x88, 0x93, 0xd0, 0x6b, 0xaa, 0x82, 0x11, 0x22, 0x12, 0x4b, 0x95, 0x29, 0x62, 0xe8,
	0xc0, 0xc0, 0xc0, 0xca, 0x8e, 0x10, 0x3c, 0x80, 0xe5, 0xd6, 0x07, 0xaa, 0x20, 0x76, 0x64, 0xbb,
	0x08, 0x1e, 0x85, 0xb7, 0xe5, 0x7c, 0xa5, 0x54, 0x19, 0x3a, 0xe5, 0xbb, 0xef, 0xbe, 0x9f, 0xf8,
	0xa0, 0xee, 0x31, 0xea, 0x61, 0xb5, 0x1c, 0xbc, 0x8b, 0x4e, 0x96, 0xbb, 0xa9, 0xbd, 0x83, 0xea,
	0xe1, 0x63, 0x1b, 0x22, 0x7a, 0x29, 0x01, 0xd6, 0x3b, 0xa8, 0x36, 0xa6, 0xc9, 0x16, 0x59, 0x27,
	0xe4, 0x15, 0xcc, 0x7b, 0xfd, 0xa5, 0x06, 0x24, 0xd2, 0x6e, 0xfb, 0x15, 0xfa, 0xe6, 0x84, 0x16,
	0xb3, 0xb6, 0x03, 0xf1, 0xe8, 0x0c, 0xca, 0x39, 0x54, 0x96, 0xbe, 0x07, 0x07, 0x11, 0xda, 0x18,
	0x8f, 0x21, 0xb0, 0x72, 0xd2, 0xde, 0x40, 0xf1, 0x12, 0x9d, 0x47, 0x79, 0x06, 0xa7, 0x21, 0x81,
	0x91, 0x76, 0x6f, 0x4e, 0x5a, 0xd1, 0xde, 0x83, 0x78, 0xa2, 0xaa, 0xb4, 0xe0, 0xca, 0xa3, 0xca,
	0x51, 0x58, 0xce, 0xde, 0x9f, 0x0c, 0xca, 0x67, 0x7c, 0xdb, 0x38, 0x2b, 0xcf, 0x61, 0xe2, 0x19,
	0x1d, 0x02, 0x88, 0x0a, 0x51, 0xfb, 0xa8, 0xde, 0xf1, 0x9b, 0x23, 0xea, 0x94, 0x89, 0xd6, 0x30,
	0x91, 0x33, 0x71, 0x09, 0x33, 0x6e, 0x5d, 0x3b, 0xfb, 0xaa, 0x3e, 0xe9, 0xa9, 0x62, 0xdf, 0x4d,
	0x43, 0xa0, 0xb8, 0xa6, 0x60, 0xe2, 0x1a, 0x8a, 0xa4, 0x0b, 0x4d, 0xb9, 0xc8, 0xbb, 0xe9, 0x6d,
	0xbd, 0xfc, 0xbb, 0x2c, 0xff, 0xfa, 0x05, 0x4c, 0xff, 0x2f, 0x46, 0xed, 0x55, 0x72, 0xfc, 0x06,
	0x00, 0x00, 0xff, 0xff, 0xf1, 0x2d, 0xc7, 0x9b, 0x7c, 0x01, 0x00, 0x00,
}
