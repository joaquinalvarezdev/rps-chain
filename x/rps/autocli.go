package rps

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	rpsv1 "github.com/0xlb/rpschain/api/rps/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: rpsv1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "GetGame",
					Use:       "game [game_number]",
					Short:     "Get the current game with the provided game number",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "index"},
					},
				},
				{
					RpcMethod: "GetParams",
					Use:       "params",
					Short:     "Get the current module parameters",
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: rpsv1.Msg_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "CreateGame",
					Use:       "create oponent rounds",
					Short:     "Creates a new Rock, Paper & Scissors game for the message sender and the chose opponent",
					Long:      "Creates a new Rock, Paper & Scissors game for the message sender and the chose opponent. Input parameters are the opponent address and the rounds number for the game",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "oponent"},
						{ProtoField: "rounds"},
					},
				},
				{
					RpcMethod: "MakeMove",
					Use:       "make-move game_number move",
					Short:     "Submits a new move for a specific Rock, Paper & Scissors game",
					Long:      "Submits a new move for a specific Rock, Paper & Scissors game. Valid move options are Rock, Paper or Scissors",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "game_number"},
						{ProtoField: "move"},
					},
				},
				{
					RpcMethod: "RevealMove",
					Use:       "reveal-move game_number revealed_move salt",
					Short:     "Reveals a submitted commitment for a specific Rock, Paper & Scissors game",
					Long:      "Reveals a submitted commitment for a specific Rock, Paper & Scissors game. To do this, need to provide the valid move and the salt used for generating the hash. Valid move options 'Rock', 'Paper' or 'Scissors'",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "game_number"},
						{ProtoField: "revealed_move"},
						{ProtoField: "salt"},
					},
				},
			},
		},
	}
}
