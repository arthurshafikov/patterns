# Proxy Solution

This is one of many situations where we can use the **Proxy** pattern. *YoutubeServiceProxy* is kind of a middleware between our application and the library code. Thanks to this class we can easily perform operations either before or after a request to the Youtube server. In this particular case I used this middleware for caching *trendingVideos* list to optimize our application.
