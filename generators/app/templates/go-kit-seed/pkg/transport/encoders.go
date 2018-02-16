package transport

import (
	"context"
	"fmt"
	"reflect"

	"github.com/golang/protobuf/ptypes"

	"github.com/golang/protobuf/ptypes/any"

	pb "dev.adeoweb.biz/pas/be-buckets/pkg/pb/buckets"

	"dev.adeoweb.biz/pas/be-buckets/pkg/models"
	"dev.adeoweb.biz/pas/be-buckets/pkg/service"
)

func encodeValidationError(err service.ValidationErrorDetail) (*any.Any, error) {
	e := &pb.ValidationError{
		Message: err.Message,
	}

	return ptypes.MarshalAny(e)
}

func encodeAccessDeniedError(err service.AccessDeniedErrorDetail) (*any.Any, error) {
	e := &pb.AccessDenied{
		Message: err.Message,
	}

	return ptypes.MarshalAny(e)
}

func encodeJSONError(errorDetail service.JSONErrorDetail) (*any.Any, error) {
	jsonData, err := errorDetail.Value()
	if err != nil {
		return nil, err
	}

	e := &pb.JSONError{
		Data: jsonData,
	}

	return ptypes.MarshalAny(e)

}

func encodeErrorStatus(errorStatus service.ErrorStatus) (*pb.ErrorStatus, error) {
	parsed := &pb.ErrorStatus{
		Message: errorStatus.Message(),
	}

	for _, detail := range errorStatus.Details() {
		if detail.Type() == service.ValidationError {
			e := detail.(service.ValidationErrorDetail)
			errorDetail, err := encodeValidationError(e)
			if err != nil {
				return &pb.ErrorStatus{}, err
			}
			parsed.Errors = append(parsed.Errors, errorDetail)
		}
		if detail.Type() == service.JSONError {
			e := detail.(service.JSONErrorDetail)
			errorDetail, err := encodeJSONError(e)
			if err != nil {
				return &pb.ErrorStatus{}, err
			}
			parsed.Errors = append(parsed.Errors, errorDetail)
		}

		if detail.Type() == service.AccessDenied {
			e := detail.(service.AccessDeniedErrorDetail)
			errorDetail, err := encodeAccessDeniedError(e)
			if err != nil {
				return &pb.ErrorStatus{}, err
			}
			parsed.Errors = append(parsed.Errors, errorDetail)
		}
	}

	return parsed, nil
}

func encodeBuckets(buckets []models.Bucket) []*pb.Bucket {
	response := []*pb.Bucket{}
	for _, bucket := range buckets {
		accessKeys := []*pb.AccessKey{}
		for _, key := range bucket.AccessKeys {
			accessKey := &pb.AccessKey{
				UserID:                string(key.UserID),
				CipherText:            key.CipherText,
				EncryptionPublicKeyID: string(key.EncryptionPublicKeyID),
			}
			accessKeys = append(accessKeys, accessKey)
		}

		b := &pb.Bucket{
			ID:         string(bucket.ID),
			AccessKeys: accessKeys,
			Meta: &pb.Meta{
				Signature: &pb.Signature{
					HashType: pb.HashType(bucket.Meta.Signature.HashType),
					Body:     bucket.Meta.Signature.Body,
					VerificationPublicKeyID: string(bucket.Meta.Signature.VerificationPublicKeyID),
				},
			},
		}

		response = append(response, b)
	}

	return response
}

func encodeCreateRequest(ctx context.Context, req interface{}) (res interface{}, err error) {
	request, ok := req.(service.CreateRequest)
	if !ok {
		err = fmt.Errorf("incorect interface{} expected [service.CreateRequest] got [%s]", reflect.TypeOf(request))
		return
	}

	pbBuckets := encodeBuckets(request.Buckets)

	return &pb.CreateRequest{
		Buckets: pbBuckets,
	}, nil
}

func encodeCreateResponse(ctx context.Context, req interface{}) (res interface{}, err error) {
	request, ok := req.(service.CreateResponse)
	response := &pb.CreateResponse{}
	if !ok {
		err = fmt.Errorf("incorect interface{} expected [service.CreateResponse] got [%s]", reflect.TypeOf(request))
		return
	}

	response.Buckets = encodeBuckets(request.Buckets)
	if request.ErrorStatus != nil {
		response.ErrorStatus, err = encodeErrorStatus(request.ErrorStatus)
		if err != nil {
			return nil, err
		}
	}
	return response, nil
}

func encodeGetByIDsRequest(ctx context.Context, req interface{}) (res interface{}, err error) {
	request, ok := req.(service.GetByIDsRequest)
	if !ok {
		err = fmt.Errorf("incorect interface{} expected [service.GetByIDsRequest] got [%s]", reflect.TypeOf(request))
		return
	}

	return &pb.GetByIDsRequest{
		IDs: request.IDs,
	}, nil
}

func encodeGetByIDsResponse(ctx context.Context, req interface{}) (res interface{}, err error) {
	request, ok := req.(service.GetByIDsResponse)
	response := &pb.GetByIDsResponse{}

	if !ok {
		err = fmt.Errorf("incorect interface{} expected [service.GetByIDsResponse] got [%s]", reflect.TypeOf(request))
		return
	}

	pbBuckets := encodeBuckets(request.Buckets)
	response.Buckets = pbBuckets
	if request.ErrorStatus != nil {
		response.ErrorStatus, err = encodeErrorStatus(request.ErrorStatus)
		if err != nil {
			return nil, err
		}
	}

	return response, nil
}
