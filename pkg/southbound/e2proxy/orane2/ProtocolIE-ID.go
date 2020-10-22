// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package orane2

//#include "ProtocolIE-ID.h"
import "C"
import (
	"fmt"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2proxy/e2ctypes"
)

func protocolIeIDToC(pcIeID e2ctypes.ProtocolIE_IDT) (C.ProtocolIE_ID_t, error) {
	switch pcIeID {
	case e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_Cause:
		return C.ProtocolIE_ID_id_Cause, nil
	case e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_CriticalityDiagnostics:
		return C.ProtocolIE_ID_id_CriticalityDiagnostics, nil
	case e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_GlobalE2node_ID:
		return C.ProtocolIE_ID_id_GlobalE2node_ID, nil
	case e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_GlobalRIC_ID:
		return C.ProtocolIE_ID_id_GlobalRIC_ID, nil
	case e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_RANfunctionID:
		return C.ProtocolIE_ID_id_RANfunctionID, nil
	case e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_RANfunctionID_Item:
		return C.ProtocolIE_ID_id_RANfunctionID_Item, nil
	case e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_RANfunctionIEcause_Item:
		return C.ProtocolIE_ID_id_RANfunctionIEcause_Item, nil
	case e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_RANfunction_Item:
		return C.ProtocolIE_ID_id_RANfunction_Item, nil
	case e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_RANfunctionsAccepted:
		return C.ProtocolIE_ID_id_RANfunctionsAccepted, nil
	case e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_RANfunctionsAdded:
		return C.ProtocolIE_ID_id_RANfunctionsAdded, nil
	case e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_RANfunctionsDeleted:
		return C.ProtocolIE_ID_id_RANfunctionsDeleted, nil
	case e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_RANfunctionsModified:
		return C.ProtocolIE_ID_id_RANfunctionsModified, nil
	case e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_RANfunctionsRejected:
		return C.ProtocolIE_ID_id_RANfunctionsRejected, nil
	case e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_RICaction_Admitted_Item:
		return C.ProtocolIE_ID_id_RICaction_Admitted_Item, nil
	case e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_RICactionID:
		return C.ProtocolIE_ID_id_RICactionID, nil
	default:
		return protocolIeIDToC2(pcIeID)
	}
}

// had to split up in to 2 parts because of
// cyclomatic complexity 35 of func `protocolIeIDToC` is high (> 30) (gocyclo)
func protocolIeIDToC2(pcIeID e2ctypes.ProtocolIE_IDT) (C.ProtocolIE_ID_t, error) {
	switch pcIeID {
	case e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_RICaction_NotAdmitted_Item:
		return C.ProtocolIE_ID_id_RICaction_NotAdmitted_Item, nil
	case e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_RICactions_Admitted:
		return C.ProtocolIE_ID_id_RICactions_Admitted, nil
	case e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_RICactions_NotAdmitted:
		return C.ProtocolIE_ID_id_RICactions_NotAdmitted, nil
	case e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_RICaction_ToBeSetup_Item:
		return C.ProtocolIE_ID_id_RICaction_ToBeSetup_Item, nil
	case e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_RICcallProcessID:
		return C.ProtocolIE_ID_id_RICcallProcessID, nil
	case e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_RICcontrolAckRequest:
		return C.ProtocolIE_ID_id_RICcontrolAckRequest, nil
	case e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_RICcontrolHeader:
		return C.ProtocolIE_ID_id_RICcontrolHeader, nil
	case e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_RICcontrolMessage:
		return C.ProtocolIE_ID_id_RICcontrolMessage, nil
	case e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_RICcontrolStatus:
		return C.ProtocolIE_ID_id_RICcontrolStatus, nil
	case e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_RICindicationHeader:
		return C.ProtocolIE_ID_id_RICindicationHeader, nil
	case e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_RICindicationMessage:
		return C.ProtocolIE_ID_id_RICindicationMessage, nil
	case e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_RICindicationSN:
		return C.ProtocolIE_ID_id_RICindicationSN, nil
	case e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_RICindicationType:
		return C.ProtocolIE_ID_id_RICindicationType, nil
	case e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_RICrequestID:
		return C.ProtocolIE_ID_id_RICrequestID, nil
	case e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_RICsubscriptionDetails:
		return C.ProtocolIE_ID_id_RICsubscriptionDetails, nil
	case e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_TimeToWait:
		return C.ProtocolIE_ID_id_TimeToWait, nil
	case e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_RICcontrolOutcome:
		return C.ProtocolIE_ID_id_RICcontrolOutcome, nil
	case e2ctypes.ProtocolIE_IDT_ProtocolIE_ID_id_Dummy:
		fallthrough
	default:
		return C.ProtocolIE_ID_t(-1), fmt.Errorf("unexpected value for ProtocolIE_IDT %v", pcIeID)
	}
}