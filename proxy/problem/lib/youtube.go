// This is library code, you cannot change it
package lib

type YoutubeService struct {
}

func NewYoutubeService() *YoutubeService {
	return &YoutubeService{}
}

func (y *YoutubeService) ListTrendVideos() []string {
	// some heavy request, that should be cached

	return []string{}
}
