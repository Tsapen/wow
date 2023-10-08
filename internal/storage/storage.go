package storage

import (
	"context"
	"crypto/rand"
	"math/big"
	"sync"
)

type storage struct {
	mux     *sync.RWMutex
	content []string
}

func New() *storage {
	content := []string{
		"And again, strong drinks are not for the belly, but for the washing of your bodies.",
		"And barley for all useful animals, and for mild drinks, as also other grain.",
		"And it is pleasing unto me that they should not be used, only in times of winter, or of cold, or famine.",
		"And shall find wisdom and great treasures of knowledge, even hidden treasures.",
		"And it is not pleasing in my sight, saith the Lord.",
		"And it is pleasing unto me that they should not be used, only in times of winter, or of cold, or famine.",
		"And whoso forbiddeth to abstain from meats, that man should not eat the same, is not ordained of God.",
		"And it is not pleasing in my sight, saith the Lord.",
		"Tobacco is not for the body, neither for the belly, and is not good for man.",
		"In consequence of evils and designs which do and will exist in the hearts of conspiring men in the last days.",
	}

	return &storage{
		content: content,
		mux:     &sync.RWMutex{},
	}
}

func (s *storage) Quote(context.Context) (string, error) {
	s.mux.RLock()
	defer s.mux.RUnlock()

	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(s.content))))
	if err != nil {
		return "", err
	}

	return s.content[int(nBig.Int64())], nil
}
