package djzh

import (
	"github.com/changtong1996/djzh/x/djzh/internal/keeper"
	"github.com/changtong1996/djzh/x/djzh/internal/types"
)

const (
	// TODO: define constants that you would like exposed from the internal package

	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	DefaultParamspace = types.DefaultParamspace
	//QueryParams       = types.QueryParams
	QuerierRoute      = types.QuerierRoute
)

var (
	// functions aliases
	NewKeeper           = keeper.NewKeeper
	NewQuerier          = keeper.NewQuerier
	RegisterCodec       = types.RegisterCodec
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis
	// TODO: Fill out function aliases

	// variable aliases
	ModuleCdc     = types.ModuleCdc
	// TODO: Fill out variable aliases
	NewMsgSendStake        = types.NewMsgSendStake
	NewMsgSendToken        = types.NewMsgSendToken
)

type (
	Keeper              = keeper.Keeper
	GenesisState        = types.GenesisState
	Params              = types.Params
/*	QueryArticleNames = types.QueryArticleNames*/
	// TODO: Fill out module types
/*	QueryVoteNum      = types.QueryVoteNum*/
	MsgCreateArticle    = types.MsgCreateArticle
	MsgCreateComment    = types.MsgCreateComment
	MsgCreateReturnVisit= types.MsgCreateReturnVisit
	MsgCreateAVote      = types.MsgCreateAVote
	MsgCreateCVote      = types.MsgCreateCVote
	MsgSendStake        = types.MsgSendStake
	MsgSendToken        = types.MsgSendToken
)
