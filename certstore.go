package goproxy

import (
	"crypto/tls"
	"github.com/muesli/cache2go"
	"time"
)

type DefaultCertStore struct {
	cache *cache2go.CacheTable
}

func newCache() *cache2go.CacheTable {
	cache := cache2go.Cache("cacheDefaultCertStore")
	return cache
}
func (s *DefaultCertStore) Fetch(hostname string, gen func() (*tls.Certificate, error)) (*tls.Certificate, error) {
	res, ok := s.cache.Value(hostname)
	if ok != nil {
		cert, err := gen()
		s.cache.Add(hostname, time.Hour, cert)
		return cert, err
	} else {
		return res.Data().(*tls.Certificate), nil
	}
}
func NewDefaultCertStore() *DefaultCertStore {
	return &DefaultCertStore{
		cache: newCache(),
	}
}
