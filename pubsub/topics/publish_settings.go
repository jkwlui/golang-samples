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

package topics

// [START pubsub_publisher_batch_settings]

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"
)

func publishWithSettings(client *pubsub.Client, topic string, msg []byte) error {
	ctx := context.Background()
	t := client.Topic(topic)
	t.PublishSettings.ByteThreshold = 5000
	t.PublishSettings.CountThreshold = 10
	t.PublishSettings.DelayThreshold = 100 * time.Millisecond

	result := t.Publish(ctx, &pubsub.Message{Data: msg})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("Get: %v", err)
	}
	fmt.Printf("Published a message; msg ID: %v\n", id)
	return nil
}

// [END pubsub_publisher_batch_settings]