package main

import "github.com/arthurshafikov/patterns/proxy/solution/lib"

func main() {
	youtubeService := lib.NewYoutubeService()
	youtubeServiceProxy := NewYoutubeServiceProxy(youtubeService)

	firstFunctionSomewhereInProgram(youtubeServiceProxy)
	secondFunctionSomewhereInProgram(youtubeServiceProxy)
}

func firstFunctionSomewhereInProgram(youtubeService YoutubeServiceInterface) {
	youtubeService.ListTrendVideos()
	// first request to YouTube
}

func secondFunctionSomewhereInProgram(youtubeService YoutubeServiceInterface) {
	youtubeService.ListTrendVideos()
	// second request to YouTube
}
