package server

import (
	"encoding/json"

	"github.com/red-hat-storage/ocs-operator/v4/services/provider/common"
	pb "github.com/red-hat-storage/ocs-operator/v4/services/provider/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func mockOnboardConsumer(mockError common.MockError) (*pb.OnboardConsumerResponse, error) { //nolint:deadcode,unused
	switch mockError {
	case common.OnboardInternalError:
		return nil, status.Errorf(codes.Internal, "mock error message")
	case common.OnboardInvalidToken:
		return nil, status.Errorf(codes.Unauthenticated, "mock error message")
	case common.OnboardInvalidArg:
		return nil, status.Errorf(codes.InvalidArgument, "mock error message")
	}

	return &pb.OnboardConsumerResponse{
		StorageConsumerUUID: common.MockConsumerID,
		GrantedCapacity:     common.MockGrantedCapacity,
	}, nil
}

func mockGetStorageConfig(mockError common.MockError) (*pb.StorageConfigResponse, error) { //nolint:deadcode,unused
	switch mockError {
	case common.StorageConfigInternalError:
		return nil, status.Errorf(codes.Internal, "mock error message")
	case common.StorageConfigInvalidUID:
		return nil, status.Errorf(codes.Unauthenticated, "mock error message")
	case common.StorageConfigConsumerNotReady:
		return nil, status.Errorf(codes.Unavailable, "mock error message")
	}

	monSecretData, _ := json.Marshal(common.MockMonSecretData)
	monConfigMapData, _ := json.Marshal(common.MockMonConfigMapData)
	return &pb.StorageConfigResponse{
		ExternalResource: []*pb.ExternalResource{
			{
				Name: "rook-ceph-mon",
				Kind: "Secret",
				Data: monSecretData,
			},
			{
				Name: "rook-ceph-mon-endpoints",
				Kind: "ConfigMap",
				Data: monConfigMapData,
			},
		},
	}, nil
}

func mockUpdateCapacity(mockError common.MockError) (*pb.UpdateCapacityResponse, error) { //nolint:deadcode,unused
	switch mockError {
	case common.UpdateInternalError:
		return nil, status.Errorf(codes.Internal, "mock error message")
	case common.UpdateInvalidArg:
		return nil, status.Errorf(codes.InvalidArgument, "mock error message")
	case common.UpdateInvalidUID:
		return nil, status.Errorf(codes.Unauthenticated, "mock error message")
	case common.UpdateConsumerNotFound:
		return nil, status.Errorf(codes.NotFound, "mock error message")
	}
	return &pb.UpdateCapacityResponse{
		GrantedCapacity: common.MockGrantedCapacity,
	}, nil
}

func mockOffboardConsumer(mockError common.MockError) (*pb.OffboardConsumerResponse, error) { //nolint:deadcode,unused
	switch mockError {
	case common.OffboardInternalError:
		return nil, status.Errorf(codes.Internal, "mock error message")
	case common.OffboardInvalidUID:
		return nil, status.Errorf(codes.Unauthenticated, "mock error message")
	case common.OffBoardConsumerNotFound:
		return nil, status.Errorf(codes.NotFound, "mock error message")
	}
	return &pb.OffboardConsumerResponse{}, nil
}
