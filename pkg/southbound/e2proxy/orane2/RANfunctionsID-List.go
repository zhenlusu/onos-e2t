// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

//#cgo CFLAGS: -I. -D_DEFAULT_SOURCE -DASN_DISABLE_OER_SUPPORT
//#cgo LDFLAGS: -lm
//#include <stdio.h>
//#include <stdlib.h>
//#include <assert.h>
//#include "ProtocolIE-Field.h"
import "C"

//func newRANfunctionsID_List(rfIDs *e2ctypes.RANfunctionsID_ListT) (*C.RANfunctionsID_List_t, error) {
//	rfIDIEsC := C.RANfunctionsID_List_t{}
//	for _, ie := range rfIDs.GetList() {
//		ieC, err := newRANfunctionID_ItemIEs(ie)
//		if err != nil {
//			return nil, err
//		}
//		if _, err = C.asn_sequence_add(unsafe.Pointer(&rfIDIEsC), unsafe.Pointer(ieC)); err != nil {
//			return nil, err
//		}
//	}
//
//	return &rfIDIEsC, nil
//}