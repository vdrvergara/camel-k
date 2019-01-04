/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package integration

import (
	"context"

	"github.com/apache/camel-k/pkg/apis/camel/v1alpha1"
	"github.com/apache/camel-k/pkg/client"
)

// Action --
type Action interface {
	client.Injectable

	// a user friendly name for the action
	Name() string

	// returns true if the action can handle the integration
	CanHandle(integration *v1alpha1.Integration) bool

	// executes the handling function
	Handle(ctx context.Context, integration *v1alpha1.Integration) error
}

type baseAction struct {
	client client.Client
}

func (action *baseAction) InjectClient(client client.Client) {
	action.client = client
}