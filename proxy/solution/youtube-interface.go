package main

type YoutubeServiceInterface interface {
	ListTrendVideos() []string
}
