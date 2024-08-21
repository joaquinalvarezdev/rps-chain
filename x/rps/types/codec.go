package types

import (
	"cosmossdk.io/core/registry"
)

// RegisterInterfaces registers the interfaces types with the interface registry.
func RegisterInterfaces(registry registry.InterfaceRegistrar) {
	// registry.RegisterImplementations((*transaction.Msg)(nil),
	//	&MsgUpdateParams{},
	//	&MsgIncrementCounter{},
	// )
	// msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
