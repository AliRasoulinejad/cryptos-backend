package middlewares

func skipper(skipURLs []string) map[string]struct{} {
	sURLs := make(map[string]struct{})
	for _, u := range skipURLs {
		sURLs[u] = struct{}{}
	}

	return sURLs
}
