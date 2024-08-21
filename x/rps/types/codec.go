package types

import (
	types "github.com/cosmos/cosmos-sdk/codec/types"
)

// RegisterInterfaces registers the interfaces types with the interface registry.
func RegisterInterfaces(registry types.InterfaceRegistry) {
	// registry.RegisterImplementations((*transaction.Msg)(nil),
	//	&MsgUpdateParams{},
	//	&MsgIncrementCounter{},
	// )
	// msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
