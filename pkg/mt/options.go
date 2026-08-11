package mt

import "time"

type Option func(client *MTClient)

func WithApiKey(key string) Option {
	return func(client *MTClient) {
		client.http.SetHeader("X-Api-Key", key)
	}
}

func WithBaseUrl(baseUrl string) Option {
	return func(client *MTClient) {
		client.http.SetBaseURL(baseUrl)
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(client *MTClient) {
		client.http.SetTimeout(timeout)
	}
}

func WithRetryCount(retryCount int) Option {
	return func(client *MTClient) {
		client.http.SetRetryCount(retryCount)
	}
}
