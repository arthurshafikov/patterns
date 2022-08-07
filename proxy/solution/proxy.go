package main

import (
	"github.com/arthurshafikov/patterns/proxy/solution/lib"
)

type YoutubeServiceProxy struct {
	youtubeService *lib.YoutubeService

	trendingVideos []string
	resetIsNeeded  bool // Reset when TTL (Time To Lose) expires e.g.
}

func NewYoutubeServiceProxy(youtubeService *lib.YoutubeService) *YoutubeServiceProxy {
	return &YoutubeServiceProxy{
		youtubeService: youtubeService,
	}
}

func (y *YoutubeServiceProxy) ListTrendVideos() []string {
	if y.trendingVideos == nil || y.resetIsNeeded {
		y.trendingVideos = y.youtubeService.ListTrendVideos()
	}

	return y.trendingVideos
}
