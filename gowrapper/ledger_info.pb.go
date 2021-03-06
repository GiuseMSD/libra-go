// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ledger_info.proto

package gowrapper

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

/// Even though we don't always need all hashes, we pass them in and return them
/// always so that we keep them in sync on the client and don't make the client
/// worry about which one(s) to pass in which cases
///
/// This structure serves a dual purpose.
///
/// First, if this structure is signed by 2f+1 validators it signifies the state
/// of the ledger at version `version` -- it contains the transaction
/// accumulator at that version which commits to all historical transactions.
/// This structure may be expanded to include other information that is derived
/// from that accumulator (e.g. the current time according to the time contract)
/// to reduce the number of proofs a client must get.
///
/// Second, the structure contains a `consensus_data_hash` value. This is the
/// hash of an internal data structure that represents a block that is voted on
/// by consensus.
///
/// Combining these two concepts when the consensus algorithm votes on a block B
/// it votes for a LedgerInfo with the `version` being the latest version that
/// will be committed if B gets 2f+1 votes. It sets `consensus_data_hash` to
/// represent B so that if those 2f+1 votes are gathered, the block is valid to
/// commit
type LedgerInfo struct {
	// Current latest version of the system
	Version uint64 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Root hash of transaction accumulator at this version
	TransactionAccumulatorHash []byte `protobuf:"bytes,2,opt,name=transaction_accumulator_hash,json=transactionAccumulatorHash,proto3" json:"transaction_accumulator_hash,omitempty"`
	// Hash of consensus-specific data that is opaque to all parts of the system
	// other than consensus.  This is needed to verify signatures because
	// consensus signing includes this hash
	ConsensusDataHash []byte `protobuf:"bytes,3,opt,name=consensus_data_hash,json=consensusDataHash,proto3" json:"consensus_data_hash,omitempty"`
	// The block id of the last committed block corresponding to this ledger info.
	// This field is not particularly interesting to the clients, but can be used
	// by the validators for synchronization.
	ConsensusBlockId []byte `protobuf:"bytes,4,opt,name=consensus_block_id,json=consensusBlockId,proto3" json:"consensus_block_id,omitempty"`
	// Epoch number corresponds to the set of validators that are active for this
	// ledger info. The main motivation for keeping the epoch number in the
	// LedgerInfo is to ensure that the client has enough information to verify
	// that the signatures for this info are coming from the validators that
	// indeed form a quorum. Without epoch number a potential attack could reuse
	// the signatures from the validators in one epoch in order to sign the wrong
	// info belonging to another epoch, in which these validators do not form a
	// quorum. The very first epoch number is 0.
	EpochNum uint64 `protobuf:"varint,5,opt,name=epoch_num,json=epochNum,proto3" json:"epoch_num,omitempty"`
	// Timestamp that represents the microseconds since the epoch (unix time) that is
	// generated by the proposer of the block.  This is strictly increasing with every block.
	// If a client reads a timestamp > the one they specified for transaction expiration time,
	// they can be certain that their transaction will never be included in a block in the future
	// (assuming that their transaction has not yet been included)
	TimestampUsecs       uint64   `protobuf:"varint,6,opt,name=timestamp_usecs,json=timestampUsecs,proto3" json:"timestamp_usecs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LedgerInfo) Reset()         { *m = LedgerInfo{} }
func (m *LedgerInfo) String() string { return proto.CompactTextString(m) }
func (*LedgerInfo) ProtoMessage()    {}
func (*LedgerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0a2c689d95021ed, []int{0}
}

func (m *LedgerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LedgerInfo.Unmarshal(m, b)
}
func (m *LedgerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LedgerInfo.Marshal(b, m, deterministic)
}
func (m *LedgerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LedgerInfo.Merge(m, src)
}
func (m *LedgerInfo) XXX_Size() int {
	return xxx_messageInfo_LedgerInfo.Size(m)
}
func (m *LedgerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LedgerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LedgerInfo proto.InternalMessageInfo

func (m *LedgerInfo) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *LedgerInfo) GetTransactionAccumulatorHash() []byte {
	if m != nil {
		return m.TransactionAccumulatorHash
	}
	return nil
}

func (m *LedgerInfo) GetConsensusDataHash() []byte {
	if m != nil {
		return m.ConsensusDataHash
	}
	return nil
}

func (m *LedgerInfo) GetConsensusBlockId() []byte {
	if m != nil {
		return m.ConsensusBlockId
	}
	return nil
}

func (m *LedgerInfo) GetEpochNum() uint64 {
	if m != nil {
		return m.EpochNum
	}
	return 0
}

func (m *LedgerInfo) GetTimestampUsecs() uint64 {
	if m != nil {
		return m.TimestampUsecs
	}
	return 0
}

/// The validator node returns this structure which includes signatures
/// from each validator to confirm the state.  The client needs to only pass
/// back the LedgerInfo element since the validator node doesn't need to know
/// the signatures again when the client performs a query, those are only there
/// for the client to be able to verify the state
type LedgerInfoWithSignatures struct {
	// Signatures of the root node from each validator
	Signatures           []*ValidatorSignature `protobuf:"bytes,1,rep,name=signatures,proto3" json:"signatures,omitempty"`
	LedgerInfo           *LedgerInfo           `protobuf:"bytes,2,opt,name=ledger_info,json=ledgerInfo,proto3" json:"ledger_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *LedgerInfoWithSignatures) Reset()         { *m = LedgerInfoWithSignatures{} }
func (m *LedgerInfoWithSignatures) String() string { return proto.CompactTextString(m) }
func (*LedgerInfoWithSignatures) ProtoMessage()    {}
func (*LedgerInfoWithSignatures) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0a2c689d95021ed, []int{1}
}

func (m *LedgerInfoWithSignatures) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LedgerInfoWithSignatures.Unmarshal(m, b)
}
func (m *LedgerInfoWithSignatures) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LedgerInfoWithSignatures.Marshal(b, m, deterministic)
}
func (m *LedgerInfoWithSignatures) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LedgerInfoWithSignatures.Merge(m, src)
}
func (m *LedgerInfoWithSignatures) XXX_Size() int {
	return xxx_messageInfo_LedgerInfoWithSignatures.Size(m)
}
func (m *LedgerInfoWithSignatures) XXX_DiscardUnknown() {
	xxx_messageInfo_LedgerInfoWithSignatures.DiscardUnknown(m)
}

var xxx_messageInfo_LedgerInfoWithSignatures proto.InternalMessageInfo

func (m *LedgerInfoWithSignatures) GetSignatures() []*ValidatorSignature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func (m *LedgerInfoWithSignatures) GetLedgerInfo() *LedgerInfo {
	if m != nil {
		return m.LedgerInfo
	}
	return nil
}

type ValidatorSignature struct {
	// The account address of the validator, which can be used for retrieving its
	// public key during the given epoch.
	ValidatorId          []byte   `protobuf:"bytes,1,opt,name=validator_id,json=validatorId,proto3" json:"validator_id,omitempty"`
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidatorSignature) Reset()         { *m = ValidatorSignature{} }
func (m *ValidatorSignature) String() string { return proto.CompactTextString(m) }
func (*ValidatorSignature) ProtoMessage()    {}
func (*ValidatorSignature) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0a2c689d95021ed, []int{2}
}

func (m *ValidatorSignature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidatorSignature.Unmarshal(m, b)
}
func (m *ValidatorSignature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidatorSignature.Marshal(b, m, deterministic)
}
func (m *ValidatorSignature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorSignature.Merge(m, src)
}
func (m *ValidatorSignature) XXX_Size() int {
	return xxx_messageInfo_ValidatorSignature.Size(m)
}
func (m *ValidatorSignature) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorSignature.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorSignature proto.InternalMessageInfo

func (m *ValidatorSignature) GetValidatorId() []byte {
	if m != nil {
		return m.ValidatorId
	}
	return nil
}

func (m *ValidatorSignature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*LedgerInfo)(nil), "types.LedgerInfo")
	proto.RegisterType((*LedgerInfoWithSignatures)(nil), "types.LedgerInfoWithSignatures")
	proto.RegisterType((*ValidatorSignature)(nil), "types.ValidatorSignature")
}

func init() { proto.RegisterFile("ledger_info.proto", fileDescriptor_d0a2c689d95021ed) }

var fileDescriptor_d0a2c689d95021ed = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x4f, 0x8f, 0xd3, 0x30,
	0x10, 0xc5, 0x95, 0xfe, 0x83, 0x4e, 0x2a, 0xa0, 0xe6, 0x62, 0xa0, 0x87, 0xd2, 0x0b, 0x3d, 0x54,
	0xa9, 0x54, 0x4e, 0xdc, 0xa0, 0xe2, 0x40, 0x25, 0xc4, 0x21, 0xa8, 0x20, 0x71, 0x89, 0x26, 0x8e,
	0x9b, 0x58, 0x24, 0x76, 0xe4, 0x3f, 0x45, 0x7c, 0x84, 0x95, 0xf6, 0x43, 0xaf, 0xe2, 0x6d, 0x9d,
	0x4a, 0x7b, 0xcb, 0xbc, 0xf7, 0x9b, 0x8c, 0xfd, 0xc6, 0x30, 0xaf, 0x79, 0x51, 0x72, 0x9d, 0x09,
	0x79, 0x52, 0x49, 0xab, 0x95, 0x55, 0x64, 0x6c, 0xff, 0xb7, 0xdc, 0xac, 0xee, 0x07, 0x00, 0xdf,
	0xbd, 0x79, 0x90, 0x27, 0x45, 0x28, 0x3c, 0x3b, 0x73, 0x6d, 0x84, 0x92, 0x34, 0x5a, 0x46, 0xeb,
	0x51, 0x7a, 0x2d, 0xc9, 0x67, 0x58, 0x58, 0x8d, 0xd2, 0x20, 0xb3, 0x42, 0xc9, 0x0c, 0x19, 0x73,
	0x8d, 0xab, 0xd1, 0x2a, 0x9d, 0x55, 0x68, 0x2a, 0x3a, 0x58, 0x46, 0xeb, 0x59, 0xfa, 0xf6, 0x86,
	0xf9, 0xd2, 0x23, 0xdf, 0xd0, 0x54, 0x24, 0x81, 0xd7, 0x4c, 0x49, 0xc3, 0xa5, 0x71, 0x26, 0x2b,
	0xd0, 0xe2, 0x63, 0xe3, 0xd0, 0x37, 0xce, 0x83, 0xf5, 0x15, 0x2d, 0x7a, 0x7e, 0x03, 0xa4, 0xe7,
	0xf3, 0x5a, 0xb1, 0xbf, 0x99, 0x28, 0xe8, 0xc8, 0xe3, 0xaf, 0x82, 0xb3, 0xef, 0x8c, 0x43, 0x41,
	0xde, 0xc1, 0x94, 0xb7, 0x8a, 0x55, 0x99, 0x74, 0x0d, 0x1d, 0xfb, 0xb3, 0x3f, 0xf7, 0xc2, 0x0f,
	0xd7, 0x90, 0x0f, 0xf0, 0xd2, 0x8a, 0x86, 0x1b, 0x8b, 0x4d, 0x9b, 0x39, 0xc3, 0x99, 0xa1, 0x13,
	0x8f, 0xbc, 0x08, 0xf2, 0xb1, 0x53, 0x57, 0x77, 0x11, 0xd0, 0x3e, 0x8e, 0xdf, 0xc2, 0x56, 0x3f,
	0x45, 0x29, 0xd1, 0x3a, 0xcd, 0x0d, 0xf9, 0x04, 0x60, 0x42, 0x45, 0xa3, 0xe5, 0x70, 0x1d, 0xef,
	0xde, 0x24, 0x3e, 0xc7, 0xe4, 0x17, 0xd6, 0xa2, 0xe8, 0xae, 0x1a, 0xf8, 0xf4, 0x06, 0x26, 0x3b,
	0x88, 0x6f, 0x56, 0xe0, 0xc3, 0x8a, 0x77, 0xf3, 0x4b, 0x6f, 0x3f, 0x30, 0x85, 0x3a, 0x7c, 0xaf,
	0x8e, 0x40, 0x9e, 0xfe, 0x95, 0xbc, 0x87, 0xd9, 0xf9, 0xaa, 0x76, 0x79, 0x44, 0x3e, 0x8f, 0x38,
	0x68, 0x87, 0x82, 0x2c, 0x60, 0x1a, 0x46, 0x5f, 0xf6, 0xd2, 0x0b, 0xfb, 0xe4, 0xcf, 0xa6, 0x14,
	0xb6, 0x72, 0x79, 0xc2, 0x54, 0xb3, 0x65, 0xaa, 0xe0, 0x0d, 0x9e, 0xb9, 0x16, 0x6c, 0x5b, 0x8b,
	0x5c, 0x23, 0xab, 0x05, 0x97, 0x76, 0x5b, 0xaa, 0x7f, 0x1a, 0xdb, 0x96, 0xeb, 0x7c, 0xe2, 0xdf,
	0xcb, 0xc7, 0x87, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1a, 0xac, 0xc7, 0x44, 0x44, 0x02, 0x00, 0x00,
}
