// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package subscriptiondelete

import (
	"github.com/wangxn2015/onos-lib-go/pkg/errors"

	v2 "github.com/wangxn2015/onos-e2t/api/e2ap/v2"
	e2appducontents "github.com/wangxn2015/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
)

// GetRequesterID gets requester ID
func GetRequesterID(request *e2appducontents.RicsubscriptionDeleteRequest) (*int32, error) {
	var res int32 = -1
	for _, v := range request.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDRicrequestID) {
			res = v.GetValue().GetRicrequestId().GetRicRequestorId()
			break
		}
	}

	if res == -1 {
		return nil, errors.NewNotFound("RicRequestID was not found")
	}

	return &res, nil
}

// GetRanFunctionID gets ran function ID
func GetRanFunctionID(request *e2appducontents.RicsubscriptionDeleteRequest) (*int32, error) {
	var res int32 = -1
	for _, v := range request.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDRanfunctionID) {
			res = v.GetValue().GetRanfunctionId().GetValue()
			break
		}
	}

	if res == -1 {
		return nil, errors.NewNotFound("RanFunctionID was not found")
	}

	return &res, nil
}

// GetRicInstanceID gets ric instance ID
func GetRicInstanceID(request *e2appducontents.RicsubscriptionDeleteRequest) (*int32, error) {
	var res int32 = -1
	for _, v := range request.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDRicrequestID) {
			res = v.GetValue().GetRicrequestId().GetRicInstanceId()
			break
		}
	}

	if res == -1 {
		return nil, errors.NewNotFound("RicInstanceID was not found")
	}

	return &res, nil
}
