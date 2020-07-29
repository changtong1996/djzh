package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)


type Article struct {
	Creator      sdk.AccAddress  `json:"creator"`
	Article_id   string          `json:"article_id"`                     // id of the article
	Uid          string          `json:"uid"`							 // id of the user
	Tid          string          `json:"tid"`                            // id of the transaction
	A_timestamp  string          `json:"a_timestamp"`                    // timestamp of the article 
	A_title      string          `json:"a_title"`                        // title of the article 
	A_text       string          `json:"a_text"`                          // text of the article
	Tag          string          `json:"tag"` 
	Flag         int             `json:"flag"` 
	Reward       sdk.Coins       `json:"reward"`
}  

func (a Article) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Creator: %s
	Article_id: %s
	Uid: %s
	Tid: %s
	A_timestamp: %s
	A_title: %s
	A_text: %s
	Tag: %s
	Flag: %d
	Reward:%s`,
		a.Creator,
		a.Article_id,
		a.Uid,
		a.Tid,
		a.A_timestamp,
		a.A_title,
		a.A_text,
		a.Tag,
		a.Flag,
		a.Reward,
	))
}



type ArticleVote struct {
	Creator      sdk.AccAddress  `json:"creator"`
	Article_id   string          `json:"article_id"`                     // id of the article
	VoteUP       int             `json:"voteUP"`
	VoteDOWN     int             `json:"voteDOWN"`
	Num          int             `json:"num"`
} 

func (av ArticleVote) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Creator: %s
	Article_id: %s
	VoteUP: %d
	VoteDOWN: %d
	Num: %d`,
		av.Creator,
		av.Article_id,
		av.VoteUP,
		av.VoteDOWN,
		av.Num,
	))
}



type Comment struct {
	Creator      sdk.AccAddress  `json:"creator"`
	Comment_id   string          `json:"comment_id"`                     // id of the comment
	Article_id   string          `json:"article_id"`                     // id of the article
	Tid          string          `json:"tid"`                            // id of the transaction
	Uid          string          `json:"uid"`							 // id of the user
	C_timestamp  string          `json:"c_timestamp"`                    // timestamp of the comment 
	C_text       string          `json:"c_text"`                         // context of the comment
	Flag         int             `json:"flag"`
	Reward       sdk.Coins       `json:"reward"`
} 


func (c Comment) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Creator: %s
	Comment_id: %s
	Article_id: %s
	Uid: %s
	Tid: %s
	C_timestamp: %s
	C_text: %s
	Flag: %d
	Reward:%s`,
		c.Creator,
		c.Comment_id,
		c.Article_id,
		c.Uid,
		c.Tid,
		c.C_timestamp,
		c.C_text,
		c.Flag,
		c.Reward,
	))
}





type CommentVote struct {
	Creator      sdk.AccAddress  `json:"creator"`
	Comment_id   string          `json:"comment_id"`                     // id of the comment
	VoteUP       int             `json:"voteUP"`
	VoteDOWN     int             `json:"voteDOWN"`
	Num          int             `json:"num"`
} 


func (cv CommentVote) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Creator: %s
	Comment_id: %s
	VoteUP: %d
	VoteDOWN: %d
	Num: %d`,
		cv.Creator,
		cv.Comment_id,
		cv.VoteUP,
		cv.VoteDOWN,
		cv.Num,
	))
}



type Domain struct {
	Domainn      string             `json:"domain"`                         // 
	Ip           string             `json:"ip"`
	Owner        string             `json:"owner"`
	Suffix       string             `json:"suffix"`
} 


type Equity struct {
	Uid          string             `json:"uid"`
	Balance      string             `json:"balance"`                        // 股权资产数
	Detail       string             `json:"detail"`
}


type EquityTransaction struct {
	Et_id            string            `json:"et_id"`
	Source_id        string            `json:"source_id"`
	Destination_id   string            `json:"destination_id"`
	Tid              string            `json:"tid"`
	Balance          string            `json:"balance"`
	Et_timestamp     int               `json:"et_timestamp"`
	Detail           string            `json:"detail"`
}

type ReturnVisit struct{
	Creator          sdk.AccAddress   `json:"creator"`
	Return_visit_id  string           `json:"return_visit_id"`
	Article_id       string           `json:"article_id"`
	Tid              string           `json:"tid"`
	Uid              string           `json:"uid"`
	Rv_timestamp     string           `json:"rv_timestamp"`
	Rv_text			 string           `json:"rv_text"`
	Flag             string           `json:"flag"` 
	Reward           sdk.Coins        `json:"reward"`
}


type SendToken struct{
	ToAddr          sdk.AccAddress    `json:"roaddr"`
	Amount          sdk.Coins         `json:"amount"`
	Percentage      string            `json:"percentage"`
}

func (rv ReturnVisit) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Creator: %s
	Return_visit_id: %s
	Article_id: %s
	Tid:%s
	Uid:%s
	Rv_timestamp:%s
	Rv_text:%s
	Flag:%s
	Reward:%s`,
		rv.Creator,
		rv.Return_visit_id,
		rv.Article_id,
		rv.Tid,
		rv.Uid,
		rv.Rv_timestamp,
		rv.Rv_text,
		rv.Flag,
		rv.Reward,
	))
}

