package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	storetypes "cosmossdk.io/core/store"
	"github.com/0xlb/rpschain/x/rps/types"
	"github.com/cosmos/cosmos-sdk/codec"
)

type Keeper struct {
	cdc          codec.BinaryCodec
	addressCodec address.Codec

	// authority is the address capable of executing a MsgUpdateParams and other authority-gated message.
	// typically, this should be the x/gov module account.
	authority string

	// state management
	Schema     collections.Schema
	Params     collections.Item[types.Params]
	GameNumber collections.Sequence
	Games      collections.Map[uint64, types.Game]
	// ActiveGamesQueue stores the game expiration and the game id as key
	ActiveGamesQueue collections.KeySet[collections.Pair[uint64, uint64]]
}

// NewKeeper creates a new Keeper instance
func NewKeeper(cdc codec.BinaryCodec, addressCodec address.Codec, storeService storetypes.KVStoreService, authority string) Keeper {
	if _, err := addressCodec.StringToBytes(authority); err != nil {
		panic(fmt.Errorf("invalid authority address: %w", err))
	}

	sb := collections.NewSchemaBuilder(storeService)
	k := Keeper{
		cdc:          cdc,
		addressCodec: addressCodec,
		authority:    authority,
		Params:       collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		GameNumber:   collections.NewSequence(sb, types.GameNumberKey, "game_number"),
		Games:        collections.NewMap(sb, types.GamesKey, "games", collections.Uint64Key, codec.CollValue[types.Game](cdc)),
		ActiveGamesQueue: collections.NewKeySet(
			sb,
			types.ActiveGamesQueueKey,
			"active_games_queue",
			collections.PairKeyCodec(collections.Uint64Key, collections.Uint64Key),
		),
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	k.Schema = schema

	return k
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// NextGameNumber returns and increments the global game number counter.
func (k Keeper) NextGameNumber(ctx context.Context) uint64 {
	n, err := k.GameNumber.Next(ctx)
	if err != nil {
		panic(err)
	}
	// sequence starts at 0 but we want first game to start at 1s
	return n + 1
}
