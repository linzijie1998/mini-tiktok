// Code generated by hertz generator.

package favorite

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/jinzhu/copier"
	favorite "github.com/linzijie1998/mini-tiktok/cmd/api/biz/model/douyin/favorite"
	"github.com/linzijie1998/mini-tiktok/cmd/api/rpc"
	rpcfavorite "github.com/linzijie1998/mini-tiktok/kitex_gen/douyin/favorite"
)

// FavoriteAction .
// @router /douyin/favorite/action/ [POST]
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req favorite.ActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(favorite.ActionResponse)
	rpcResp := new(rpcfavorite.ActionResponse)

	rpcResp, err = rpc.FavoriteAction(ctx, &rpcfavorite.ActionRequest{
		Token:      req.Token,
		VideoId:    req.VideoID,
		ActionType: req.ActionType,
	})

	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}

	if err = copier.Copy(resp, rpcResp); err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// FavoriteList .
// @router /douyin/favorite/list/ [GET]
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req favorite.ListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(favorite.ListResponse)
	rpcResp := new(rpcfavorite.ListResponse)

	rpcResp, err = rpc.FavoriteList(ctx, &rpcfavorite.ListRequest{
		UserId: req.UserID,
		Token:  req.Token,
	})

	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}

	if err = copier.Copy(resp, rpcResp); err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}

	for i, video := range rpcResp.VideoList {
		if err = copier.Copy(resp.VideoList[i], video); err != nil {
			c.String(consts.StatusInternalServerError, err.Error())
			return
		}
		resp.VideoList[i].ID = video.Id
		resp.VideoList[i].PlayURL = video.PlayUrl
		resp.VideoList[i].CoverURL = video.CoverUrl
		resp.VideoList[i].Author.ID = video.Author.Id
	}

	c.JSON(consts.StatusOK, resp)
}
