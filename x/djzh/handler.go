package djzh

import (
	"fmt"

	"github.com/changtong1996/djzh/x/djzh/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

)

// NewHandler creates an sdk.Handler for all the djzh type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgCreateArticle:
			return handleMsgCreateArticle(ctx, k, msg)
	
		case MsgCreateComment:
			return handleMsgCreateComment(ctx, k, msg)
	
		case MsgCreateReturnVisit:
			return handleMsgCreateReturnVisit(ctx, k, msg)
		
		case MsgCreateAVote:
			return handleMsgCreateAVote(ctx, k, msg)

		case MsgCreateCVote:
			return handleMsgCreateCVote(ctx, k, msg)		
			
		case MsgSendStake:
			return handleMsgSendStake(ctx, k, msg)	
			
		case MsgSendToken:
			return handleMsgSendToken(ctx, k, msg)	

		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName,  msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// handleMsgCreateArticle creates a new article and moves the reward into escrow
func handleMsgCreateArticle(ctx sdk.Context, k Keeper, msg MsgCreateArticle) (*sdk.Result, error) {
	var article = types.Article{
		Creator:      msg.Creator,
		A_text:       msg.A_text,
		A_title:      msg.A_title,
		Tag:          msg.Tag,
		Article_id:   msg.Article_id,
		Tid:          msg.Tid,
 		Uid:          msg.Uid,
		A_timestamp:  msg.A_timestamp,
		Reward:       msg.Reward,
	}
	_, err := k.GetArticle(ctx, article.Article_id)
	if err == nil {
		return nil, sdkerrors.Wrap(err, "Article with that hash already exists")
	}

	_, err1 := k.CoinKeeper.AddCoins(ctx, article.Creator, article.Reward)
	if err1 !=  nil {
		return nil, err1
	}

	k.SetArticle(ctx, article)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeCreateArticle),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator.String()),
			sdk.NewAttribute(types.AttributeText, msg.A_text),
			sdk.NewAttribute(types.AttributeA_title, msg.A_title),
			sdk.NewAttribute(types.AttributeTag, msg.Tag),
			sdk.NewAttribute(types.AttributeArticle_id, msg.Article_id),
			sdk.NewAttribute(types.AttributeTid, msg.Tid),
			sdk.NewAttribute(types.AttributeUid, msg.Uid),
			sdk.NewAttribute(types.AttributeA_timestamp, msg.A_timestamp),
			sdk.NewAttribute(types.AttributeReward, msg.Reward.String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}



func handleMsgCreateComment(ctx sdk.Context, k Keeper, msg MsgCreateComment) (*sdk.Result, error) {
	var comment = types.Comment{
		Creator:     msg.Creator,
		Comment_id:  msg.Comment_id,
		Article_id:  msg.Article_id,
		Tid:         msg.Tid,
 		Uid:         msg.Uid,
		C_timestamp: msg.C_timestamp,
		C_text:      msg.C_text,
		Reward:      msg.Reward,
	}
	_, err := k.GetComment(ctx, comment.Comment_id)
	if err == nil {
		return nil, sdkerrors.Wrap(err, "Comment with that id already exists")
	}

	_, err1 := k.CoinKeeper.AddCoins(ctx, comment.Creator, comment.Reward)
	if err1 !=  nil {
		return nil, err1
	}
	k.SetComment(ctx, comment)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeCreateComment),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator.String()),
			sdk.NewAttribute(types.AttributeReward, msg.Reward.String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}



func handleMsgCreateReturnVisit(ctx sdk.Context, k Keeper, msg MsgCreateReturnVisit) (*sdk.Result, error) {
	var rv = types.ReturnVisit{
		Creator:             msg.Creator,
		Return_visit_id:     msg.Return_visit_id,
		Article_id:          msg.Article_id,
		Tid:                 msg.Tid,
 		Uid:                 msg.Uid,
		Rv_timestamp:        msg.Rv_timestamp,
		Rv_text:             msg.Rv_text,
		Flag:                msg.Flag,
		Reward:              msg.Reward,
	}
	_, err := k.GetReturnVisit(ctx, rv.Return_visit_id)
	if err == nil {
		return nil, sdkerrors.Wrap(err, "Rv with that id already exists")
	}
	_, err1 := k.CoinKeeper.AddCoins(ctx, rv.Creator, rv.Reward)
	if err1 !=  nil {
		return nil, err1
	}

	k.SetReturnVisit(ctx, rv)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeCreateReturnVisit),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator.String()),
			sdk.NewAttribute(types.AttributeReward, msg.Reward.String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}



func handleMsgCreateAVote(ctx sdk.Context, k Keeper, msg MsgCreateAVote) (*sdk.Result, error) {
	var aVote = types.ArticleVote{
		Creator:             msg.Creator,
		Article_id:          msg.Article_id,
		VoteUP:              msg.VoteUP,    
	    VoteDOWN:            msg.VoteDOWN,
	    Num:                 msg.Num,
	}
	_, err := k.GetAVote(ctx, aVote.Article_id)
	if err == nil {
		return nil, sdkerrors.Wrap(err, "Rv with that id already exists")
	}

	
	k.SetAVote(ctx, aVote)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeCreateAVote),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator.String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}





func handleMsgCreateCVote(ctx sdk.Context, k Keeper, msg MsgCreateCVote) (*sdk.Result, error) {
	var cVote = types.CommentVote{
		Creator:             msg.Creator,
		Comment_id:          msg.Comment_id,
		VoteUP:              msg.VoteUP,    
	    VoteDOWN:            msg.VoteDOWN,
	    Num:                 msg.Num,
	}
	_, err := k.GetCVote(ctx, cVote.Comment_id)
	if err == nil {
		return nil, sdkerrors.Wrap(err, "Rv with that id already exists")
	}
	

	k.SetCVote(ctx, cVote)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeCreateCVote),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator.String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}





// Send the stake from the admin account to user account
func handleMsgSendStake(ctx sdk.Context, k Keeper, msg MsgSendStake) (*sdk.Result, error) {
	err := k.CoinKeeper.SendCoins(ctx, msg.FromAddr, msg.ToAddr, msg.Amount)
	if err != nil {
		return nil, err
	}
	return &sdk.Result{}, nil
}

// Send the award token according to user's percentage of stake 
func handleMsgSendToken(ctx sdk.Context, k Keeper, msg MsgSendToken) (*sdk.Result, error) {
	var sendtoken = types.SendToken{
		ToAddr:           msg.ToAddr,
		Amount:           msg.Amount,
		Percentage:       msg.Percentage,
	}
	_, err := k.CoinKeeper.AddCoins(ctx, msg.ToAddr, msg.Amount)
	if err !=  nil {
		return nil, err
	}

	k.SetSendToken(ctx, sendtoken)
	return &sdk.Result{}, nil
}



