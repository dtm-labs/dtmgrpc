/*
 * Copyright (c) 2021 yedf. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package dtmgrpc

import (
	context "context"

	"github.com/dtm-labs/dtmcli"
	"github.com/dtm-labs/dtmcli/dtmimp"
	"github.com/dtm-labs/dtmgrpc/dtmgimp"
	"github.com/dtm-labs/dtmdriver"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// DtmError2GrpcError translate dtm error to grpc error
func DtmError2GrpcError(res interface{}) error {
	e, ok := res.(error)
	if ok && e == dtmimp.ErrFailure {
		return status.New(codes.Aborted, dtmcli.ResultFailure).Err()
	} else if ok && e == dtmimp.ErrOngoing {
		return status.New(codes.FailedPrecondition, dtmcli.ResultOngoing).Err()
	}
	return e
}

// MustGenGid must gen a gid from grpcServer
func MustGenGid(grpcServer string) string {
	dc := dtmgimp.MustGetDtmClient(grpcServer)
	r, err := dc.NewGid(context.Background(), &emptypb.Empty{})
	dtmimp.E2P(err)
	return r.Gid
}

// UseDriver use the specified driver to handle grpc urls
func UseDriver(driverName string) error {
	return dtmdriver.Use(driverName)
}

// AddUnaryInterceptor adds grpc.UnaryClientInterceptor
func AddUnaryInterceptor(interceptor grpc.UnaryClientInterceptor) {
	dtmgimp.ClientInterceptors = append(dtmgimp.ClientInterceptors, interceptor)
}
