package controllers

import (
	"context"
	"sync"

	"cloud.google.com/go/firestore"
	"github.com/sirupsen/logrus"
)

var mu sync.Mutex
var logger = logrus.New()

var ctx context.Context
var client *firestore.Client

// Initializing firestore connection
func init() {
	ctx = context.Background()
	c, err := firestore.NewClient(ctx, "spry-blade-403304")
	if err != nil {
		logger.Errorf("Failed to create Firestore client: %v", err)
	}
	client = c
}
