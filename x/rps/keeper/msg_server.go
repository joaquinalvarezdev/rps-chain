package keeper

import (
	"context"
	"fmt"
	"strings"

	"github.com/0xlb/rpschain/x/rps/rules"
	"github.com/0xlb/rpschain/x/rps/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	k Keeper
}

var _ types.MsgServer = msgServer{}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{k: keeper}
}

func (ms msgServer) CreateGame(ctx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	newGame := types.Game{
		GameNumber: ms.k.NextGameNumber(ctx),
		PlayerA:    msg.Creator,
		PlayerB:    msg.Oponent,
		Rounds:     msg.Rounds,
		Status:     rules.StatusWaiting,
		Score:      []uint64{0, 0},
	}

	if err := newGame.Validate(); err != nil {
		return nil, err
	}

	if err := ms.k.Games.Set(ctx, newGame.GameNumber, newGame); err != nil {
		return nil, err
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.EventManager().EmitTypedEvent(&types.EventCreateGame{
		GameNumber: newGame.GameNumber,
		PlayerA:    newGame.PlayerA,
		PlayerB:    newGame.PlayerB,
	})

	return &types.MsgCreateGameResponse{}, nil
}

func (ms msgServer) MakeMove(ctx context.Context, msg *types.MsgMakeMove) (*types.MsgMakeMoveResponse, error) {

	if ok := rules.IsValidMove(msg.Move); !ok {
		return nil, types.ErrInvalidMove
	}

	game, err := ms.k.Games.Get(ctx, msg.GameNumber)
	if err != nil {
		return nil, err
	}

	if game.Status != rules.StatusInProgress && game.Status != rules.StatusWaiting {
		return nil, types.ErrInvalidGameStatus
	}

	var player rules.Player
	switch msg.Player {
	case game.PlayerA:
		player = rules.PlayerA
		game.PlayerAMoves = append(game.PlayerAMoves, msg.Move)
	case game.PlayerB:
		player = rules.PlayerB
		game.PlayerBMoves = append(game.PlayerBMoves, msg.Move)
	}

	if player == rules.InvalidPlayer {
		return nil, types.ErrInvalidPlayer
	}

	playerAMovesCount, playerBMovesCount := len(game.PlayerAMoves), len(game.PlayerBMoves)

	if ok := rules.CanMakeMove(player, playerAMovesCount, playerBMovesCount); !ok {
		return nil, types.ErrPlayerCantMakeMove
	}

	if playerAMovesCount == playerBMovesCount {
		playerAResult := rules.DetermineRoundWinner(rules.Choice(game.PlayerAMoves[playerAMovesCount-1]), rules.Choice(game.PlayerBMoves[playerBMovesCount-1]))
		if playerAResult == rules.Win {
			game.AddWintoPlayerA()
		} else if playerAResult == rules.Loss {
			game.AddWintoPlayerB()
		}
	}

	game.Status = rules.GameResultByMajority(game.GetPlayerAScore(), game.GetPlayerBScore(), game.Rounds)

	if err := game.Validate(); err != nil {
		return nil, err
	}

	if err := ms.k.Games.Set(ctx, game.GameNumber, game); err != nil {
		return nil, err
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.EventManager().EmitTypedEvent(&types.EventMakeMove{
		GameNumber: msg.GameNumber,
		Player:     msg.Player,
		Move:       msg.Move,
	})

	if game.Ended() {
		sdkCtx.EventManager().EmitTypedEvent(&types.EventEndGame{
			GameNumber: msg.GameNumber,
			Status:     game.Status,
		})

	}

	return &types.MsgMakeMoveResponse{}, nil
}

func (ms msgServer) UpdateParams(ctx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if _, err := ms.k.addressCodec.StringToBytes(msg.Authority); err != nil {
		return nil, fmt.Errorf("invalid authority address: %w", err)
	}

	if authority := ms.k.GetAuthority(); !strings.EqualFold(msg.Authority, authority) {
		return nil, fmt.Errorf("unauthorized, authority does not match the module's authority: got %s, want %s", msg.Authority, authority)
	}

	if err := msg.Params.Validate(); err != nil {
		return nil, err
	}

	if err := ms.k.Params.Set(ctx, msg.Params); err != nil {
		return nil, err
	}

	return &types.MsgUpdateParamsResponse{}, nil
}
