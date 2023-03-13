// Code generated by MockGen. DO NOT EDIT.
// Source: x/ccv/types/expected_keepers.go

// Package keeper is a generated GoMock package.
package keeper

import (
	context "context"
	reflect "reflect"
	time "time"

	types "github.com/cosmos/cosmos-sdk/types"
	types0 "github.com/cosmos/cosmos-sdk/x/auth/types"
	types1 "github.com/cosmos/cosmos-sdk/x/capability/types"
	types2 "github.com/cosmos/cosmos-sdk/x/evidence/types"
	types3 "github.com/cosmos/cosmos-sdk/x/slashing/types"
	types4 "github.com/cosmos/cosmos-sdk/x/staking/types"
	types5 "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	types6 "github.com/cosmos/ibc-go/v4/modules/core/03-connection/types"
	types7 "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	exported "github.com/cosmos/ibc-go/v4/modules/core/exported"
	gomock "github.com/golang/mock/gomock"
	types8 "github.com/tendermint/tendermint/abci/types"
stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// MockStakingKeeper is a mock of StakingKeeper interface.
type MockStakingKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockStakingKeeperMockRecorder
}

// MockStakingKeeperMockRecorder is the mock recorder for MockStakingKeeper.
type MockStakingKeeperMockRecorder struct {
	mock *MockStakingKeeper
}

// NewMockStakingKeeper creates a new mock instance.
func NewMockStakingKeeper(ctrl *gomock.Controller) *MockStakingKeeper {
	mock := &MockStakingKeeper{ctrl: ctrl}
	mock.recorder = &MockStakingKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStakingKeeper) EXPECT() *MockStakingKeeperMockRecorder {
	return m.recorder
}

// GetLastTotalPower mocks base method.
func (m *MockStakingKeeper) GetLastTotalPower(ctx types.Context) types.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastTotalPower", ctx)
	ret0, _ := ret[0].(types.Int)
	return ret0
}

// GetLastTotalPower indicates an expected call of GetLastTotalPower.
func (mr *MockStakingKeeperMockRecorder) GetLastTotalPower(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastTotalPower", reflect.TypeOf((*MockStakingKeeper)(nil).GetLastTotalPower), ctx)
}

// GetLastValidatorPower mocks base method.
func (m *MockStakingKeeper) GetLastValidatorPower(ctx types.Context, operator types.ValAddress) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastValidatorPower", ctx, operator)
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetLastValidatorPower indicates an expected call of GetLastValidatorPower.
func (mr *MockStakingKeeperMockRecorder) GetLastValidatorPower(ctx, operator interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastValidatorPower", reflect.TypeOf((*MockStakingKeeper)(nil).GetLastValidatorPower), ctx, operator)
}

// GetValidator mocks base method.
func (m *MockStakingKeeper) GetValidator(ctx types.Context, addr types.ValAddress) (types4.Validator, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidator", ctx, addr)
	ret0, _ := ret[0].(types4.Validator)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetValidator indicates an expected call of GetValidator.
func (mr *MockStakingKeeperMockRecorder) GetValidator(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidator", reflect.TypeOf((*MockStakingKeeper)(nil).GetValidator), ctx, addr)
}

// GetValidatorByConsAddr mocks base method.
func (m *MockStakingKeeper) GetValidatorByConsAddr(ctx types.Context, consAddr types.ConsAddress) (types4.Validator, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidatorByConsAddr", ctx, consAddr)
	ret0, _ := ret[0].(types4.Validator)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetValidatorByConsAddr indicates an expected call of GetValidatorByConsAddr.
func (mr *MockStakingKeeperMockRecorder) GetValidatorByConsAddr(ctx, consAddr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorByConsAddr", reflect.TypeOf((*MockStakingKeeper)(nil).GetValidatorByConsAddr), ctx, consAddr)
}

// GetValidatorUpdates mocks base method.
func (m *MockStakingKeeper) GetValidatorUpdates(ctx types.Context) []types8.ValidatorUpdate {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidatorUpdates", ctx)
	ret0, _ := ret[0].([]types8.ValidatorUpdate)
	return ret0
}

// GetValidatorUpdates indicates an expected call of GetValidatorUpdates.
func (mr *MockStakingKeeperMockRecorder) GetValidatorUpdates(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorUpdates", reflect.TypeOf((*MockStakingKeeper)(nil).GetValidatorUpdates), ctx)
}

// IterateLastValidatorPowers mocks base method.
func (m *MockStakingKeeper) IterateLastValidatorPowers(ctx types.Context, cb func(types.ValAddress, int64) bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IterateLastValidatorPowers", ctx, cb)
}

// IterateLastValidatorPowers indicates an expected call of IterateLastValidatorPowers.
func (mr *MockStakingKeeperMockRecorder) IterateLastValidatorPowers(ctx, cb interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateLastValidatorPowers", reflect.TypeOf((*MockStakingKeeper)(nil).IterateLastValidatorPowers), ctx, cb)
}

// Jail mocks base method.
func (m *MockStakingKeeper) Jail(arg0 types.Context, arg1 types.ConsAddress) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Jail", arg0, arg1)
}

// Jail indicates an expected call of Jail.
func (mr *MockStakingKeeperMockRecorder) Jail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Jail", reflect.TypeOf((*MockStakingKeeper)(nil).Jail), arg0, arg1)
}

// PowerReduction mocks base method.
func (m *MockStakingKeeper) PowerReduction(ctx types.Context) types.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PowerReduction", ctx)
	ret0, _ := ret[0].(types.Int)
	return ret0
}

// PowerReduction indicates an expected call of PowerReduction.
func (mr *MockStakingKeeperMockRecorder) PowerReduction(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PowerReduction", reflect.TypeOf((*MockStakingKeeper)(nil).PowerReduction), ctx)
}

// PutUnbondingOnHold mocks base method.
func (m *MockStakingKeeper) PutUnbondingOnHold(ctx types.Context, id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutUnbondingOnHold", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutUnbondingOnHold indicates an expected call of PutUnbondingOnHold.
func (mr *MockStakingKeeperMockRecorder) PutUnbondingOnHold(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutUnbondingOnHold", reflect.TypeOf((*MockStakingKeeper)(nil).PutUnbondingOnHold), ctx, id)
}

// Slash mocks base method.
func (m *MockStakingKeeper) Slash(arg0 types.Context, arg1 types.ConsAddress, arg2, arg3 int64, arg4 types.Dec, arg5 types4.InfractionType) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Slash", arg0, arg1, arg2, arg3, arg4, arg5)
}

// Slash indicates an expected call of Slash.
func (mr *MockStakingKeeperMockRecorder) Slash(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Slash", reflect.TypeOf((*MockStakingKeeper)(nil).Slash), arg0, arg1, arg2, arg3, arg4, arg5)
}

// UnbondingCanComplete mocks base method.
func (m *MockStakingKeeper) UnbondingCanComplete(ctx types.Context, id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnbondingCanComplete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnbondingCanComplete indicates an expected call of UnbondingCanComplete.
func (mr *MockStakingKeeperMockRecorder) UnbondingCanComplete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnbondingCanComplete", reflect.TypeOf((*MockStakingKeeper)(nil).UnbondingCanComplete), ctx, id)
}

// UnbondingTime mocks base method.
func (m *MockStakingKeeper) UnbondingTime(ctx types.Context) time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnbondingTime", ctx)
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// UnbondingTime indicates an expected call of UnbondingTime.
func (mr *MockStakingKeeperMockRecorder) UnbondingTime(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnbondingTime", reflect.TypeOf((*MockStakingKeeper)(nil).UnbondingTime), ctx)
}

// Unjail mocks base method.
func (m *MockStakingKeeper) Unjail(ctx types.Context, addr types.ConsAddress) {
	m.ctrl.T.Helper()
	_ = m.ctrl.Call(m, "Unjail", ctx, addr)
}
	
// Unjail indicates an expected call of Unjail.
func (mr *MockStakingKeeperMockRecorder) Unjail(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unjail", reflect.TypeOf((*MockStakingKeeper)(nil).Unjail), ctx, addr)
}

func (m *MockStakingKeeper) IterateValidators(ctx types.Context, f func(index int64, validator stakingtypes.ValidatorI) (stop bool)) {
	m.ctrl.T.Helper()
	_ = m.ctrl.Call(m, "IterateValidators", ctx, f)
}

func (mr *MockStakingKeeperMockRecorder) IterateValidators(ctx, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateValidators", reflect.TypeOf((*MockStakingKeeper)(nil).IterateValidators), ctx, f)
}

func (m *MockStakingKeeper) Validator(ctx types.Context, valAddr types.ValAddress) stakingtypes.ValidatorI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validator", ctx, valAddr)
	ret0, _ := ret[0].(stakingtypes.ValidatorI)
	return ret0
}

func (mr *MockStakingKeeperMockRecorder) Validator(ctx, valAddr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validator", reflect.TypeOf((*MockStakingKeeper)(nil).Validator), ctx, valAddr)
}

// IsValidatorJailed mocks base method.
func (m *MockStakingKeeper) IsValidatorJailed(ctx types.Context, addr types.ConsAddress) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValidatorJailed", ctx, addr)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsValidatorJailed indicates an expected call of IsValidatorJailed.
func (mr *MockStakingKeeperMockRecorder) IsValidatorJailed(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValidatorJailed", reflect.TypeOf((*MockStakingKeeper)(nil).IsValidatorJailed), ctx, addr)
}


func (m *MockStakingKeeper) ValidatorByConsAddr(ctx types.Context, consAddr types.ConsAddress) stakingtypes.ValidatorI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatorByConsAddr", ctx, consAddr)
	ret0, _ := ret[0].(stakingtypes.ValidatorI)
	return ret0
}

func (mr *MockStakingKeeperMockRecorder) ValidatorByConsAddr(ctx, consAddr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatorByConsAddr", reflect.TypeOf((*MockStakingKeeper)(nil).ValidatorByConsAddr), ctx, consAddr)
}

// Delegation mocks base method.
func (m *MockStakingKeeper) Delegation(ctx types.Context, addr types.AccAddress, valAddr types.ValAddress) stakingtypes.DelegationI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delegation", ctx, addr, valAddr)
	ret0, _ := ret[0].(stakingtypes.DelegationI)
	return ret0
}

// Delegation indicates an expected call of Delegation.
func (mr *MockStakingKeeperMockRecorder) Delegation(ctx, addr, valAddr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delegation", reflect.TypeOf((*MockStakingKeeper)(nil).Delegation), ctx, addr, valAddr)
}

func (m *MockStakingKeeper) MaxValidators(ctx types.Context) uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxValidators", ctx)
	ret0, _ := ret[0].(uint32)
	return ret0
}

func (mr *MockStakingKeeperMockRecorder) MaxValidators(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxValidators", reflect.TypeOf((*MockStakingKeeper)(nil).MaxValidators), ctx)
}

// MockEvidenceKeeper is a mock of EvidenceKeeper interface.
type MockEvidenceKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockEvidenceKeeperMockRecorder
}

// MockEvidenceKeeperMockRecorder is the mock recorder for MockEvidenceKeeper.
type MockEvidenceKeeperMockRecorder struct {
	mock *MockEvidenceKeeper
}

// NewMockEvidenceKeeper creates a new mock instance.
func NewMockEvidenceKeeper(ctrl *gomock.Controller) *MockEvidenceKeeper {
	mock := &MockEvidenceKeeper{ctrl: ctrl}
	mock.recorder = &MockEvidenceKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEvidenceKeeper) EXPECT() *MockEvidenceKeeperMockRecorder {
	return m.recorder
}

// HandleEquivocationEvidence mocks base method.
func (m *MockEvidenceKeeper) HandleEquivocationEvidence(ctx types.Context, evidence *types2.Equivocation) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleEquivocationEvidence", ctx, evidence)
}

// HandleEquivocationEvidence indicates an expected call of HandleEquivocationEvidence.
func (mr *MockEvidenceKeeperMockRecorder) HandleEquivocationEvidence(ctx, evidence interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleEquivocationEvidence", reflect.TypeOf((*MockEvidenceKeeper)(nil).HandleEquivocationEvidence), ctx, evidence)
}

// MockSlashingKeeper is a mock of SlashingKeeper interface.
type MockSlashingKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockSlashingKeeperMockRecorder
}

// MockSlashingKeeperMockRecorder is the mock recorder for MockSlashingKeeper.
type MockSlashingKeeperMockRecorder struct {
	mock *MockSlashingKeeper
}

// NewMockSlashingKeeper creates a new mock instance.
func NewMockSlashingKeeper(ctrl *gomock.Controller) *MockSlashingKeeper {
	mock := &MockSlashingKeeper{ctrl: ctrl}
	mock.recorder = &MockSlashingKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSlashingKeeper) EXPECT() *MockSlashingKeeperMockRecorder {
	return m.recorder
}

// DowntimeJailDuration mocks base method.
func (m *MockSlashingKeeper) DowntimeJailDuration(arg0 types.Context) time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DowntimeJailDuration", arg0)
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// DowntimeJailDuration indicates an expected call of DowntimeJailDuration.
func (mr *MockSlashingKeeperMockRecorder) DowntimeJailDuration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DowntimeJailDuration", reflect.TypeOf((*MockSlashingKeeper)(nil).DowntimeJailDuration), arg0)
}

// GetValidatorSigningInfo mocks base method.
func (m *MockSlashingKeeper) GetValidatorSigningInfo(ctx types.Context, address types.ConsAddress) (types3.ValidatorSigningInfo, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidatorSigningInfo", ctx, address)
	ret0, _ := ret[0].(types3.ValidatorSigningInfo)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetValidatorSigningInfo indicates an expected call of GetValidatorSigningInfo.
func (mr *MockSlashingKeeperMockRecorder) GetValidatorSigningInfo(ctx, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorSigningInfo", reflect.TypeOf((*MockSlashingKeeper)(nil).GetValidatorSigningInfo), ctx, address)
}

// IsTombstoned mocks base method.
func (m *MockSlashingKeeper) IsTombstoned(arg0 types.Context, arg1 types.ConsAddress) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsTombstoned", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsTombstoned indicates an expected call of IsTombstoned.
func (mr *MockSlashingKeeperMockRecorder) IsTombstoned(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsTombstoned", reflect.TypeOf((*MockSlashingKeeper)(nil).IsTombstoned), arg0, arg1)
}

// JailUntil mocks base method.
func (m *MockSlashingKeeper) JailUntil(arg0 types.Context, arg1 types.ConsAddress, arg2 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "JailUntil", arg0, arg1, arg2)
}

// JailUntil indicates an expected call of JailUntil.
func (mr *MockSlashingKeeperMockRecorder) JailUntil(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JailUntil", reflect.TypeOf((*MockSlashingKeeper)(nil).JailUntil), arg0, arg1, arg2)
}

// SlashFractionDoubleSign mocks base method.
func (m *MockSlashingKeeper) SlashFractionDoubleSign(ctx types.Context) types.Dec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SlashFractionDoubleSign", ctx)
	ret0, _ := ret[0].(types.Dec)
	return ret0
}

// SlashFractionDoubleSign indicates an expected call of SlashFractionDoubleSign.
func (mr *MockSlashingKeeperMockRecorder) SlashFractionDoubleSign(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SlashFractionDoubleSign", reflect.TypeOf((*MockSlashingKeeper)(nil).SlashFractionDoubleSign), ctx)
}

// SlashFractionDowntime mocks base method.
func (m *MockSlashingKeeper) SlashFractionDowntime(arg0 types.Context) types.Dec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SlashFractionDowntime", arg0)
	ret0, _ := ret[0].(types.Dec)
	return ret0
}

// SlashFractionDowntime indicates an expected call of SlashFractionDowntime.
func (mr *MockSlashingKeeperMockRecorder) SlashFractionDowntime(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SlashFractionDowntime", reflect.TypeOf((*MockSlashingKeeper)(nil).SlashFractionDowntime), arg0)
}

// Tombstone mocks base method.
func (m *MockSlashingKeeper) Tombstone(arg0 types.Context, arg1 types.ConsAddress) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Tombstone", arg0, arg1)
}

// Tombstone indicates an expected call of Tombstone.
func (mr *MockSlashingKeeperMockRecorder) Tombstone(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tombstone", reflect.TypeOf((*MockSlashingKeeper)(nil).Tombstone), arg0, arg1)
}

// MockChannelKeeper is a mock of ChannelKeeper interface.
type MockChannelKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockChannelKeeperMockRecorder
}

// MockChannelKeeperMockRecorder is the mock recorder for MockChannelKeeper.
type MockChannelKeeperMockRecorder struct {
	mock *MockChannelKeeper
}

// NewMockChannelKeeper creates a new mock instance.
func NewMockChannelKeeper(ctrl *gomock.Controller) *MockChannelKeeper {
	mock := &MockChannelKeeper{ctrl: ctrl}
	mock.recorder = &MockChannelKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChannelKeeper) EXPECT() *MockChannelKeeperMockRecorder {
	return m.recorder
}

// ChanCloseInit mocks base method.
func (m *MockChannelKeeper) ChanCloseInit(ctx types.Context, portID, channelID string, chanCap *types1.Capability) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChanCloseInit", ctx, portID, channelID, chanCap)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChanCloseInit indicates an expected call of ChanCloseInit.
func (mr *MockChannelKeeperMockRecorder) ChanCloseInit(ctx, portID, channelID, chanCap interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChanCloseInit", reflect.TypeOf((*MockChannelKeeper)(nil).ChanCloseInit), ctx, portID, channelID, chanCap)
}

// GetChannel mocks base method.
func (m *MockChannelKeeper) GetChannel(ctx types.Context, srcPort, srcChan string) (types7.Channel, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChannel", ctx, srcPort, srcChan)
	ret0, _ := ret[0].(types7.Channel)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetChannel indicates an expected call of GetChannel.
func (mr *MockChannelKeeperMockRecorder) GetChannel(ctx, srcPort, srcChan interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChannel", reflect.TypeOf((*MockChannelKeeper)(nil).GetChannel), ctx, srcPort, srcChan)
}

// GetNextSequenceSend mocks base method.
func (m *MockChannelKeeper) GetNextSequenceSend(ctx types.Context, portID, channelID string) (uint64, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextSequenceSend", ctx, portID, channelID)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetNextSequenceSend indicates an expected call of GetNextSequenceSend.
func (mr *MockChannelKeeperMockRecorder) GetNextSequenceSend(ctx, portID, channelID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextSequenceSend", reflect.TypeOf((*MockChannelKeeper)(nil).GetNextSequenceSend), ctx, portID, channelID)
}

// SendPacket mocks base method.
func (m *MockChannelKeeper) SendPacket(ctx types.Context, channelCap *types1.Capability, packet exported.PacketI) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendPacket", ctx, channelCap, packet)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendPacket indicates an expected call of SendPacket.
func (mr *MockChannelKeeperMockRecorder) SendPacket(ctx, channelCap, packet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendPacket", reflect.TypeOf((*MockChannelKeeper)(nil).SendPacket), ctx, channelCap, packet)
}

// WriteAcknowledgement mocks base method.
func (m *MockChannelKeeper) WriteAcknowledgement(ctx types.Context, chanCap *types1.Capability, packet exported.PacketI, acknowledgement exported.Acknowledgement) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteAcknowledgement", ctx, chanCap, packet, acknowledgement)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteAcknowledgement indicates an expected call of WriteAcknowledgement.
func (mr *MockChannelKeeperMockRecorder) WriteAcknowledgement(ctx, chanCap, packet, acknowledgement interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteAcknowledgement", reflect.TypeOf((*MockChannelKeeper)(nil).WriteAcknowledgement), ctx, chanCap, packet, acknowledgement)
}

// MockPortKeeper is a mock of PortKeeper interface.
type MockPortKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockPortKeeperMockRecorder
}

// MockPortKeeperMockRecorder is the mock recorder for MockPortKeeper.
type MockPortKeeperMockRecorder struct {
	mock *MockPortKeeper
}

// NewMockPortKeeper creates a new mock instance.
func NewMockPortKeeper(ctrl *gomock.Controller) *MockPortKeeper {
	mock := &MockPortKeeper{ctrl: ctrl}
	mock.recorder = &MockPortKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPortKeeper) EXPECT() *MockPortKeeperMockRecorder {
	return m.recorder
}

// BindPort mocks base method.
func (m *MockPortKeeper) BindPort(ctx types.Context, portID string) *types1.Capability {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindPort", ctx, portID)
	ret0, _ := ret[0].(*types1.Capability)
	return ret0
}

// BindPort indicates an expected call of BindPort.
func (mr *MockPortKeeperMockRecorder) BindPort(ctx, portID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindPort", reflect.TypeOf((*MockPortKeeper)(nil).BindPort), ctx, portID)
}

// MockConnectionKeeper is a mock of ConnectionKeeper interface.
type MockConnectionKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionKeeperMockRecorder
}

// MockConnectionKeeperMockRecorder is the mock recorder for MockConnectionKeeper.
type MockConnectionKeeperMockRecorder struct {
	mock *MockConnectionKeeper
}

// NewMockConnectionKeeper creates a new mock instance.
func NewMockConnectionKeeper(ctrl *gomock.Controller) *MockConnectionKeeper {
	mock := &MockConnectionKeeper{ctrl: ctrl}
	mock.recorder = &MockConnectionKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnectionKeeper) EXPECT() *MockConnectionKeeperMockRecorder {
	return m.recorder
}

// GetConnection mocks base method.
func (m *MockConnectionKeeper) GetConnection(ctx types.Context, connectionID string) (types6.ConnectionEnd, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConnection", ctx, connectionID)
	ret0, _ := ret[0].(types6.ConnectionEnd)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetConnection indicates an expected call of GetConnection.
func (mr *MockConnectionKeeperMockRecorder) GetConnection(ctx, connectionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConnection", reflect.TypeOf((*MockConnectionKeeper)(nil).GetConnection), ctx, connectionID)
}

// MockClientKeeper is a mock of ClientKeeper interface.
type MockClientKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockClientKeeperMockRecorder
}

// MockClientKeeperMockRecorder is the mock recorder for MockClientKeeper.
type MockClientKeeperMockRecorder struct {
	mock *MockClientKeeper
}

// NewMockClientKeeper creates a new mock instance.
func NewMockClientKeeper(ctrl *gomock.Controller) *MockClientKeeper {
	mock := &MockClientKeeper{ctrl: ctrl}
	mock.recorder = &MockClientKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientKeeper) EXPECT() *MockClientKeeperMockRecorder {
	return m.recorder
}

// CreateClient mocks base method.
func (m *MockClientKeeper) CreateClient(ctx types.Context, clientState exported.ClientState, consensusState exported.ConsensusState) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateClient", ctx, clientState, consensusState)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateClient indicates an expected call of CreateClient.
func (mr *MockClientKeeperMockRecorder) CreateClient(ctx, clientState, consensusState interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateClient", reflect.TypeOf((*MockClientKeeper)(nil).CreateClient), ctx, clientState, consensusState)
}

// GetClientState mocks base method.
func (m *MockClientKeeper) GetClientState(ctx types.Context, clientID string) (exported.ClientState, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClientState", ctx, clientID)
	ret0, _ := ret[0].(exported.ClientState)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetClientState indicates an expected call of GetClientState.
func (mr *MockClientKeeperMockRecorder) GetClientState(ctx, clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClientState", reflect.TypeOf((*MockClientKeeper)(nil).GetClientState), ctx, clientID)
}

// GetLatestClientConsensusState mocks base method.
func (m *MockClientKeeper) GetLatestClientConsensusState(ctx types.Context, clientID string) (exported.ConsensusState, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestClientConsensusState", ctx, clientID)
	ret0, _ := ret[0].(exported.ConsensusState)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetLatestClientConsensusState indicates an expected call of GetLatestClientConsensusState.
func (mr *MockClientKeeperMockRecorder) GetLatestClientConsensusState(ctx, clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestClientConsensusState", reflect.TypeOf((*MockClientKeeper)(nil).GetLatestClientConsensusState), ctx, clientID)
}

// GetSelfConsensusState mocks base method.
func (m *MockClientKeeper) GetSelfConsensusState(ctx types.Context, height exported.Height) (exported.ConsensusState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSelfConsensusState", ctx, height)
	ret0, _ := ret[0].(exported.ConsensusState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSelfConsensusState indicates an expected call of GetSelfConsensusState.
func (mr *MockClientKeeperMockRecorder) GetSelfConsensusState(ctx, height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSelfConsensusState", reflect.TypeOf((*MockClientKeeper)(nil).GetSelfConsensusState), ctx, height)
}

// MockConsumerHooks is a mock of ConsumerHooks interface.
type MockConsumerHooks struct {
	ctrl     *gomock.Controller
	recorder *MockConsumerHooksMockRecorder
}

// MockConsumerHooksMockRecorder is the mock recorder for MockConsumerHooks.
type MockConsumerHooksMockRecorder struct {
	mock *MockConsumerHooks
}

// NewMockConsumerHooks creates a new mock instance.
func NewMockConsumerHooks(ctrl *gomock.Controller) *MockConsumerHooks {
	mock := &MockConsumerHooks{ctrl: ctrl}
	mock.recorder = &MockConsumerHooksMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConsumerHooks) EXPECT() *MockConsumerHooksMockRecorder {
	return m.recorder
}

// AfterValidatorBonded mocks base method.
func (m *MockConsumerHooks) AfterValidatorBonded(ctx types.Context, consAddr types.ConsAddress, arg2 types.ValAddress) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AfterValidatorBonded", ctx, consAddr, arg2)
}

// AfterValidatorBonded indicates an expected call of AfterValidatorBonded.
func (mr *MockConsumerHooksMockRecorder) AfterValidatorBonded(ctx, consAddr, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterValidatorBonded", reflect.TypeOf((*MockConsumerHooks)(nil).AfterValidatorBonded), ctx, consAddr, arg2)
}

// MockBankKeeper is a mock of BankKeeper interface.
type MockBankKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankKeeperMockRecorder
}

// MockBankKeeperMockRecorder is the mock recorder for MockBankKeeper.
type MockBankKeeperMockRecorder struct {
	mock *MockBankKeeper
}

// NewMockBankKeeper creates a new mock instance.
func NewMockBankKeeper(ctrl *gomock.Controller) *MockBankKeeper {
	mock := &MockBankKeeper{ctrl: ctrl}
	mock.recorder = &MockBankKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankKeeper) EXPECT() *MockBankKeeperMockRecorder {
	return m.recorder
}

// GetAllBalances mocks base method.
func (m *MockBankKeeper) GetAllBalances(ctx types.Context, addr types.AccAddress) types.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBalances", ctx, addr)
	ret0, _ := ret[0].(types.Coins)
	return ret0
}

// GetAllBalances indicates an expected call of GetAllBalances.
func (mr *MockBankKeeperMockRecorder) GetAllBalances(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBalances", reflect.TypeOf((*MockBankKeeper)(nil).GetAllBalances), ctx, addr)
}

// GetBalance mocks base method.
func (m *MockBankKeeper) GetBalance(ctx types.Context, addr types.AccAddress, denom string) types.Coin {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", ctx, addr, denom)
	ret0, _ := ret[0].(types.Coin)
	return ret0
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockBankKeeperMockRecorder) GetBalance(ctx, addr, denom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockBankKeeper)(nil).GetBalance), ctx, addr, denom)
}

// SendCoinsFromModuleToModule mocks base method.
func (m *MockBankKeeper) SendCoinsFromModuleToModule(ctx types.Context, senderModule, recipientModule string, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromModuleToModule", ctx, senderModule, recipientModule, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromModuleToModule indicates an expected call of SendCoinsFromModuleToModule.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromModuleToModule(ctx, senderModule, recipientModule, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromModuleToModule", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromModuleToModule), ctx, senderModule, recipientModule, amt)
}

// MockAccountKeeper is a mock of AccountKeeper interface.
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
}

// MockAccountKeeperMockRecorder is the mock recorder for MockAccountKeeper.
type MockAccountKeeperMockRecorder struct {
	mock *MockAccountKeeper
}

// NewMockAccountKeeper creates a new mock instance.
func NewMockAccountKeeper(ctrl *gomock.Controller) *MockAccountKeeper {
	mock := &MockAccountKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountKeeper) EXPECT() *MockAccountKeeperMockRecorder {
	return m.recorder
}

// GetModuleAccount mocks base method.
func (m *MockAccountKeeper) GetModuleAccount(ctx types.Context, name string) types0.ModuleAccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAccount", ctx, name)
	ret0, _ := ret[0].(types0.ModuleAccountI)
	return ret0
}

// GetModuleAccount indicates an expected call of GetModuleAccount.
func (mr *MockAccountKeeperMockRecorder) GetModuleAccount(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAccount), ctx, name)
}

// MockIBCTransferKeeper is a mock of IBCTransferKeeper interface.
type MockIBCTransferKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockIBCTransferKeeperMockRecorder
}

// MockIBCTransferKeeperMockRecorder is the mock recorder for MockIBCTransferKeeper.
type MockIBCTransferKeeperMockRecorder struct {
	mock *MockIBCTransferKeeper
}

// NewMockIBCTransferKeeper creates a new mock instance.
func NewMockIBCTransferKeeper(ctrl *gomock.Controller) *MockIBCTransferKeeper {
	mock := &MockIBCTransferKeeper{ctrl: ctrl}
	mock.recorder = &MockIBCTransferKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBCTransferKeeper) EXPECT() *MockIBCTransferKeeperMockRecorder {
	return m.recorder
}

// SendTransfer mocks base method.
func (m *MockIBCTransferKeeper) SendTransfer(ctx types.Context, sourcePort, sourceChannel string, token types.Coin, sender types.AccAddress, receiver string, timeoutHeight types5.Height, timeoutTimestamp uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTransfer", ctx, sourcePort, sourceChannel, token, sender, receiver, timeoutHeight, timeoutTimestamp)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendTransfer indicates an expected call of SendTransfer.
func (mr *MockIBCTransferKeeperMockRecorder) SendTransfer(ctx, sourcePort, sourceChannel, token, sender, receiver, timeoutHeight, timeoutTimestamp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTransfer", reflect.TypeOf((*MockIBCTransferKeeper)(nil).SendTransfer), ctx, sourcePort, sourceChannel, token, sender, receiver, timeoutHeight, timeoutTimestamp)
}

// MockIBCCoreKeeper is a mock of IBCCoreKeeper interface.
type MockIBCCoreKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockIBCCoreKeeperMockRecorder
}

// MockIBCCoreKeeperMockRecorder is the mock recorder for MockIBCCoreKeeper.
type MockIBCCoreKeeperMockRecorder struct {
	mock *MockIBCCoreKeeper
}

// NewMockIBCCoreKeeper creates a new mock instance.
func NewMockIBCCoreKeeper(ctrl *gomock.Controller) *MockIBCCoreKeeper {
	mock := &MockIBCCoreKeeper{ctrl: ctrl}
	mock.recorder = &MockIBCCoreKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBCCoreKeeper) EXPECT() *MockIBCCoreKeeperMockRecorder {
	return m.recorder
}

// ChannelOpenInit mocks base method.
func (m *MockIBCCoreKeeper) ChannelOpenInit(goCtx context.Context, msg *types7.MsgChannelOpenInit) (*types7.MsgChannelOpenInitResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChannelOpenInit", goCtx, msg)
	ret0, _ := ret[0].(*types7.MsgChannelOpenInitResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChannelOpenInit indicates an expected call of ChannelOpenInit.
func (mr *MockIBCCoreKeeperMockRecorder) ChannelOpenInit(goCtx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChannelOpenInit", reflect.TypeOf((*MockIBCCoreKeeper)(nil).ChannelOpenInit), goCtx, msg)
}

// MockScopedKeeper is a mock of ScopedKeeper interface.
type MockScopedKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockScopedKeeperMockRecorder
}

// MockScopedKeeperMockRecorder is the mock recorder for MockScopedKeeper.
type MockScopedKeeperMockRecorder struct {
	mock *MockScopedKeeper
}

// NewMockScopedKeeper creates a new mock instance.
func NewMockScopedKeeper(ctrl *gomock.Controller) *MockScopedKeeper {
	mock := &MockScopedKeeper{ctrl: ctrl}
	mock.recorder = &MockScopedKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScopedKeeper) EXPECT() *MockScopedKeeperMockRecorder {
	return m.recorder
}

// AuthenticateCapability mocks base method.
func (m *MockScopedKeeper) AuthenticateCapability(ctx types.Context, cap *types1.Capability, name string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthenticateCapability", ctx, cap, name)
	ret0, _ := ret[0].(bool)
	return ret0
}

// AuthenticateCapability indicates an expected call of AuthenticateCapability.
func (mr *MockScopedKeeperMockRecorder) AuthenticateCapability(ctx, cap, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthenticateCapability", reflect.TypeOf((*MockScopedKeeper)(nil).AuthenticateCapability), ctx, cap, name)
}

// ClaimCapability mocks base method.
func (m *MockScopedKeeper) ClaimCapability(ctx types.Context, cap *types1.Capability, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClaimCapability", ctx, cap, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClaimCapability indicates an expected call of ClaimCapability.
func (mr *MockScopedKeeperMockRecorder) ClaimCapability(ctx, cap, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClaimCapability", reflect.TypeOf((*MockScopedKeeper)(nil).ClaimCapability), ctx, cap, name)
}

// GetCapability mocks base method.
func (m *MockScopedKeeper) GetCapability(ctx types.Context, name string) (*types1.Capability, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCapability", ctx, name)
	ret0, _ := ret[0].(*types1.Capability)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetCapability indicates an expected call of GetCapability.
func (mr *MockScopedKeeperMockRecorder) GetCapability(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCapability", reflect.TypeOf((*MockScopedKeeper)(nil).GetCapability), ctx, name)
}
