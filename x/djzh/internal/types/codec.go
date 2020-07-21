package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	// TODO: Register the modules msgs
	cdc.RegisterConcrete(MsgCreateArticle{}, "djzh/MsgCreateArticle", nil)
	cdc.RegisterConcrete(MsgCreateComment{}, "djzh/MsgCreateComment", nil)
	cdc.RegisterConcrete(MsgCreateReturnVisit{}, "djzh/MsgCreateReturnVisit", nil)
	cdc.RegisterConcrete(MsgCreateAVote{}, "djzh/MsgCreateAVote", nil)
	cdc.RegisterConcrete(MsgCreateCVote{}, "djzh/MsgCreateCVote", nil)
	cdc.RegisterConcrete(MsgSendStake{}, "djzh/MsgSendStake", nil)
	cdc.RegisterConcrete(MsgSendToken{}, "djzh/MsgSendToken", nil)

}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
