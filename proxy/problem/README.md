# Proxy Problem

Basically, the problem is that every time we need to get trending videos from YouTube we make a request to the Youtube server.
It's not efficient and also it may cause problems and lead to a ban for your Youtube account because it's done too much requests. So probably it'd be better to cache these videos at least for an hour instead of requesting them every time we call the *ListTrendVideos()* method. 

Keep in mind that we cannot change the Youtube library code.
