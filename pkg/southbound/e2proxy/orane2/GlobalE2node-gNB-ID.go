// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "GlobalE2node-gNB-ID.h"
import "C"
import (
	"fmt"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
)

func newGlobalE2nodegNBID(gnbID *e2ctypes.GlobalE2NodeGNB_ID) (*C.GlobalE2node_gNB_ID_t, error) {

	globalgNBID, err := newGlobalgNBID(gnbID.GlobalGNB_ID)
	if err != nil {
		return nil, err
	}

	globalgNBIDC := C.GlobalE2node_gNB_ID_t{
		global_gNB_ID: *globalgNBID,
		gNB_CU_UP_ID:  nil,
		gNB_DU_ID:     nil,
	}

	return &globalgNBIDC, nil
}

func decodeGlobalE2nodegNBID(gNBC *C.GlobalE2node_gNB_ID_t) (*e2ctypes.GlobalE2NodeGNB_ID, error) {
	result := new(e2ctypes.GlobalE2NodeGNB_ID)
	var err error
	result.GlobalGNB_ID, err = decodeGlobalGnbID(&gNBC.global_gNB_ID)
	if err != nil {
		return nil, fmt.Errorf("error decodeGlobalE2nodegNBID() %v", err)
	}

	return result, nil
}