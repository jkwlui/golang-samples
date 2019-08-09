// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package subscription

// [START pubsub_create_pull_subscription]
import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"
)

func create(client *pubsub.Client, subName string, topic *pubsub.Topic) error {
	ctx := context.Background()
	sub, err := client.CreateSubscription(ctx, subName, pubsub.SubscriptionConfig{
		Topic:       topic,
		AckDeadline: 20 * time.Second,
	})
	if err != nil {
		return fmt.Errorf("CreateSubscription: %v", err)
	}
	fmt.Printf("Created subscription: %v\n", sub)
	return nil
}

// [END pubsub_create_pull_subscription]