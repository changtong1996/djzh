package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)


type MsgCreateArticle struct {
	Creator      sdk.AccAddress `json:"creator"`           // address of the article creator
	A_text       string         `json:"a_text"`              // 
	A_title      string         `json:"a_title"`
	Tag          string         `json:"tag"`
	Article_id   string         `json:"article_id"`
	Tid			 string         `json:"tid"`
	Uid          string         `json:"uid"`
	A_timestamp  string         `json:"a_timestamp"`
	Reward       sdk.Coins      `json:"reward"`            // reward of the article
}

func NewMsgCreateArticle(creator sdk.AccAddress, a_text string, a_title string, tag string, article_id string,
 tid string, uid string, a_timestamp string, reward sdk.Coins) MsgCreateArticle {
	return MsgCreateArticle{
		Creator:      creator,
		A_text:       a_text,
		A_title:      a_title,
		Tag:          tag,
		Article_id:   article_id,
		Tid:          tid,
 		Uid:          uid,
		A_timestamp:  a_timestamp,
		Reward:       reward,
	}
}

const CreateArticleConst = "CreateArticle"


//nolint
func (msg MsgCreateArticle) Route() string { return RouterKey }
func (msg MsgCreateArticle) Type()  string { return CreateArticleConst }

func (msg MsgCreateArticle) GetSigners() []sdk.AccAddress{
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateArticle) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgCreateArticle) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	if msg.Article_id == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "articleid can't be empty")
	}
	return nil
}






type MsgCreateComment struct {
	Creator      sdk.AccAddress  `json:"creator"`           // address of the article creator
	Comment_id   string          `json:"comment_id"`                     // id of the comment
	Article_id   string          `json:"article_id"`                     // id of the article
	Tid          string          `json:"tid"`                            // id of the transaction
	Uid          string          `json:"uid"`							 // id of the user
	C_timestamp  string          `json:"c_timestamp"`                    // timestamp of the comment 
	C_text       string          `json:"c_text"`                         // context of the comment
	Reward       sdk.Coins       `json:"reward"`
}

func NewMsgCreateComment(creator sdk.AccAddress, comment_id string, article_id string, tid string, uid string,
 c_timestamp string, c_text string, reward sdk.Coins) MsgCreateComment {
	return MsgCreateComment{
		Creator:      creator,
 		Comment_id:   comment_id,
		Article_id:   article_id,
		Tid:          tid,
 		Uid:          uid,
		C_timestamp:  c_timestamp,
		C_text:       c_text,
	    Reward:       reward,
	}
}

const CreateCommentConst = "CreateComment"


//nolint
func (msg MsgCreateComment) Route() string { return RouterKey }
func (msg MsgCreateComment) Type()  string { return CreateCommentConst }
func (msg MsgCreateComment) GetSigners() []sdk.AccAddress{
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateComment) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgCreateComment) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	if msg.Comment_id == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "solutionScavengerHash can't be empty")
	}
	return nil
}









type MsgCreateReturnVisit struct {
	Creator          sdk.AccAddress   `json:"creator"`           // address of the article creator
	Return_visit_id  string           `json:"return_visit_id"`
	Article_id       string           `json:"article_id"`
	Tid              string           `json:"tid"`
	Uid              string           `json:"uid"`
	Rv_timestamp     string           `json:"rv_timestamp"`
	Rv_text			 string           `json:"rv_text"`
	Flag             string           `json:"flag"` 
	Reward           sdk.Coins        `json:"reward"`            // reward of the article

}

func NewMsgCreateReturnVisit(creator sdk.AccAddress, return_visit_id string, article_id string, tid string, uid string,
 rv_timestamp string, rv_text string, flag string, reward sdk.Coins) MsgCreateReturnVisit {
	return MsgCreateReturnVisit{
		Creator:             creator,
		Return_visit_id:     return_visit_id,
		Article_id:          article_id,
		Tid:                 tid,
 		Uid:                 uid,
		Rv_timestamp:        rv_timestamp,
		Rv_text:             rv_text,
		Flag:                flag,
		Reward:              reward,
	}
}

const CreateReturnVisitConst = "CreateReturnVisit"


//nolint
func (msg MsgCreateReturnVisit) Route() string { return RouterKey }
func (msg MsgCreateReturnVisit) Type()  string { return CreateReturnVisitConst }
func (msg MsgCreateReturnVisit) GetSigners() []sdk.AccAddress{
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateReturnVisit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgCreateReturnVisit) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	if msg.Return_visit_id == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "rv can't be empty")
	}
	return nil
}








type MsgCreateAVote struct {
	Creator          sdk.AccAddress   `json:"creator"`           // address of the article creator
	Article_id       string           `json:"article_id"`        // id of the article
	VoteUP           int              `json:"voteUP"`
	VoteDOWN         int              `json:"voteDOWN"`
	Num              int              `json:"num"`
}

func NewMsgCreateAVote(creator sdk.AccAddress, article_id string, voteUP int, voteDOWN int, num int) MsgCreateAVote {
	return MsgCreateAVote{
		Creator:             creator,
		Article_id:          article_id,
		VoteUP:              voteUP,    
	    VoteDOWN:            voteDOWN,
	    Num:                 num,
	}
}

const CreateAVoteConst = "CreateAVote"


//nolint
func (msg MsgCreateAVote) Route() string { return RouterKey }
func (msg MsgCreateAVote) Type()  string { return CreateAVoteConst }
func (msg MsgCreateAVote) GetSigners() []sdk.AccAddress{
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateAVote) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgCreateAVote) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	if msg.Article_id == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "solutionScavengerHash can't be empty")
	}
	return nil
}







type MsgCreateCVote struct {
	Creator          sdk.AccAddress   `json:"creator"`           // address of the article creator
	Comment_id       string           `json:"comment_id"`                     // id of the article
	VoteUP           int              `json:"voteUP"`
	VoteDOWN         int              `json:"voteDOWN"`
	Num              int              `json:"num"`

}

func NewMsgCreateCVote(creator sdk.AccAddress, comment_id string, voteUP int, voteDOWN int, num int) MsgCreateCVote {
	return MsgCreateCVote{
		Creator:             creator,
		Comment_id:          comment_id,
		VoteUP:              voteUP,    
	    VoteDOWN:            voteDOWN,
	    Num:                 num,
	}
}

const CreateCVoteConst = "CreateCVote"


//nolint
func (msg MsgCreateCVote) Route() string { return RouterKey }
func (msg MsgCreateCVote) Type()  string { return CreateCVoteConst }
func (msg MsgCreateCVote) GetSigners() []sdk.AccAddress{
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateCVote) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgCreateCVote) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	if msg.Comment_id == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "comment_id can't be empty")
	}
	return nil
}






type MsgSendStake struct {
	FromAddr          sdk.AccAddress           `json:"fromaddr"`           
	ToAddr            sdk.AccAddress           `json:"toaddr"`
	Signer1           sdk.AccAddress           `json:"signer1"`
	Signer2           sdk.AccAddress           `json:"signer2"`                    
	Amount    	      sdk.Coins                `json:"amount"`
}

func NewMsgSendStake(fromaddr sdk.AccAddress, toaddr sdk.AccAddress, signer1 sdk.AccAddress, signer2 sdk.AccAddress, amount sdk.Coins) MsgSendStake {
	return MsgSendStake{
		FromAddr:            fromaddr,
		ToAddr:              toaddr,
		Signer1:             signer1,
		Signer2:             signer2,
		Amount:              amount,
	}
}

const SendStakeConst = "SendStake"


//nolint
func (msg MsgSendStake) Route() string { return RouterKey }
func (msg MsgSendStake) Type()  string { return SendStakeConst }
func (msg MsgSendStake) GetSigners() []sdk.AccAddress{
	return []sdk.AccAddress{sdk.AccAddress(msg.FromAddr), sdk.AccAddress(msg.Signer1), sdk.AccAddress(msg.Signer2)}
}

func (msg MsgSendStake) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgSendStake) ValidateBasic() error {
	if msg.FromAddr.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "FromAddr can't be empty")
	}
	if msg.ToAddr.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "ToAddr can't be empty")
	}
	if !msg.Amount.IsAllPositive() {
		return sdkerrors.ErrInsufficientFunds
	}
	return nil
}






type MsgSendToken struct {        
	ToAddr            sdk.AccAddress   `json:"toaddr"`                    
	Amount    	      sdk.Coins        `json:"amount"`
	Percentage        string           `json:"percentage"`
}

func NewMsgSendToken(toaddr sdk.AccAddress, amount sdk.Coins, percentage string) MsgSendToken {
	return MsgSendToken{
		ToAddr:              toaddr,
		Amount:              amount,
		Percentage:          percentage,
	}
}

const SendTokenConst = "SendToken"


//nolint
func (msg MsgSendToken) Route() string { return RouterKey }
func (msg MsgSendToken) Type()  string { return SendTokenConst }
func (msg MsgSendToken) GetSigners() []sdk.AccAddress{
	return []sdk.AccAddress{sdk.AccAddress(msg.ToAddr)}
}

func (msg MsgSendToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgSendToken) ValidateBasic() error {
	if msg.ToAddr.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "ToAddr can't be empty")
	}
	if !msg.Amount.IsAllPositive() {
		return sdkerrors.ErrInsufficientFunds
	}
	return nil
}


// TODO: Describe your actions, these will implment the interface of `sdk.Msg`
/*
verify interface at compile time
var _ sdk.Msg = &Msg<Action>{}

Msg<Action> - struct for unjailing jailed validator
type Msg<Action> struct {
	ValidatorAddr sdk.ValAddress `json:"address" yaml:"address"` // address of the validator operator
}

NewMsg<Action> creates a new Msg<Action> instance
func NewMsg<Action>(validatorAddr sdk.ValAddress) Msg<Action> {
	return Msg<Action>{
		ValidatorAddr: validatorAddr,
	}
}

const <action>Const = "<action>"

// nolint
func (msg Msg<Action>) Route() string { return RouterKey }
func (msg Msg<Action>) Type() string  { return <action>Const }
func (msg Msg<Action>) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.ValidatorAddr)}
}

GetSignBytes gets the bytes for the message signer to sign on
func (msg Msg<Action>) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

ValidateBasic validity check for the AnteHandler
func (msg Msg<Action>) ValidateBasic() error {
	if msg.ValidatorAddr.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing validator address"
	}
	return nil
}
*/
