package v1alpha2

import (
	"context"

	pb "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

func (c *service) Version(
	ctx context.Context, req *pb.VersionRequest,
) (*pb.VersionResponse, error) {
	return nil, nil
}
