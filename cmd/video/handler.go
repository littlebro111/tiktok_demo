package main

import (
	"context"
	"github.com/Pinklr/tiktok_demo/cmd/video/pack"
	"github.com/Pinklr/tiktok_demo/cmd/video/service"
	video "github.com/Pinklr/tiktok_demo/kitex_gen/video"
	"github.com/Pinklr/tiktok_demo/pkg/errno"
	"log"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	resp = new(video.FeedResponse)
	var latestTime int64 = 0
	if req != nil {
		latestTime = req.LatestTime
	}
	if latestTime < 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return
	}
	videos, nextTime, err := service.Feed(ctx, latestTime)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Videos = videos
	resp.NextTime = nextTime
	return
}

// VideoAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) VideoAction(ctx context.Context, req *video.VideoActionRequest) (resp *video.VideoActionResponse, err error) {
	//上传视频
	resp = new(video.VideoActionResponse)

	v := req.Video
	log.Println(v)
	if v.Author.Id <= 0 || len(v.PlayUrl) == 0 || len(v.Title) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return
	}
	err = service.UploadVideo(ctx, v.Author.Id, v.PlayUrl, v.CoverUrl, v.Title)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// List implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) List(ctx context.Context, req *video.ListRequest) (resp *video.ListResponse, err error) {
	// TODO: Your code here...
	return
}

// MGetVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) MGetVideo(ctx context.Context, req *video.MGetVideoRequest) (resp *video.MGetVideoResponse, err error) {
	// TODO: Your code here...
	return
}
