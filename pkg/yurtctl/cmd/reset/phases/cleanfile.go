/*
Copyright 2021 The OpenYurt Authors.

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

package phases

import (
	"fmt"
	"os"

	"k8s.io/klog/v2"
	"k8s.io/kubernetes/cmd/kubeadm/app/cmd/phases/workflow"

	"github.com/openyurtio/openyurt/pkg/yurtctl/constants"
	"github.com/openyurtio/openyurt/pkg/yurtctl/util/edgenode"
)

func NewCleanfilePhase() workflow.Phase {
	return workflow.Phase{
		Name:  "Clean up the directories and files related to kubelet and yurthub.",
		Short: "Clean up the directories and files related to kubelet and yurthub.",
		Run:   runCleanfile,
	}
}

func runCleanfile(c workflow.RunData) error {
	for _, comp := range []string{"kubectl", "kubeadm", "kubelet"} {
		target := fmt.Sprintf("/usr/bin/%s", comp)
		if err := os.RemoveAll(target); err != nil {
			klog.Warningf("Clean file %s fail: %v, please clean it manually.", target, err)
		}
	}

	for _, file := range []string{constants.KubeletWorkdir,
		constants.YurttunnelAgentWorkdir,
		constants.YurttunnelServerWorkdir,
		constants.YurtHubWorkdir,
		edgenode.KubeletSvcPath,
		constants.KubeletServiceFilepath,
		constants.KubeCniDir,
		constants.KubeletConfigureDir,
		constants.Sysctl_k8s_config} {
		if err := os.RemoveAll(file); err != nil {
			klog.Warningf("Clean file %s fail: %v, please clean it manually.", file, err)
		}
	}
	return nil
}
