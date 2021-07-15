/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mockautoscaling

import (
	"sync"

	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/autoscaling/autoscalingiface"
)

type MockAutoscaling struct {
	// Mock out interface
	autoscalingiface.AutoScalingAPI

	mutex             sync.Mutex
	Groups            map[string]*autoscaling.Group
	WarmPoolInstances map[string][]*autoscaling.Instance
	LifecycleHooks    map[string]*autoscaling.LifecycleHook
}

var _ autoscalingiface.AutoScalingAPI = &MockAutoscaling{}
