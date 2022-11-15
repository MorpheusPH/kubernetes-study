package utils

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

func PatchStatus(ctx context.Context, obj client.Object, kubeClient client.Client) error {
	namespaceName := client.ObjectKeyFromObject(obj)
	current := obj.DeepCopyObject().(client.Object)
	if err := kubeClient.Get(ctx, namespaceName, current); err != nil {
		return err
	}

	return kubeClient.Status().Patch(ctx, obj, client.MergeFrom(current))
}
