package main

import "github.com/arthurshafikov/patterns/proxy/problem/lib"

func main() {
	youtubeService := lib.NewYoutubeService()

	firstFunctionSomewhereInProgram(youtubeService)
	secondFunctionSomewhereInProgram(youtubeService)
}

func firstFunctionSomewhereInProgram(youtubeService *lib.YoutubeService) {
	youtubeService.ListTrendVideos()
	// first request to YouTube
}

func secondFunctionSomewhereInProgram(youtubeService *lib.YoutubeService) {
	youtubeService.ListTrendVideos()
	// second request to YouTube
}
