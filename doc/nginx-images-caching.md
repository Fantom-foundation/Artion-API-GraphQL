Following Nginx caching proxy configuration is needed:

```
http {
	# Artion images cache
	proxy_cache_path
		/srv/nginx-cache # directory where to store the cache
		levels=2:2 # 2-letters/2-letters dir hierarchy
		keys_zone=artimgs:100m # cache name "artimgs", 100 MB keys database
		min_free=1000m # keep at least 1 GB fee at the disk
		inactive=2y; # cache invalidation after 2 years (cannot be infinite)
	proxy_cache_key "$scheme$request_method$host$request_uri";
}

server {
	# images (resized)
	location /images/ {
		proxy_cache artimgs; # name of cache defined above
		add_header X-Cache-Status $upstream_cache_status; # add reponse header with HIT/MISS status
		proxy_ignore_headers Set-Cookie Expires Cache-Control; # ignore expiration defined by this headers
		proxy_hide_header Set-Cookie; # skip cookies
		proxy_cache_valid 200 302 10y; # OK status keep in cache for 10 years
		proxy_cache_valid any 1m; # error status keep in cache only 1 minute
		proxy_pass http://localhost:16761;
	}
}

```

