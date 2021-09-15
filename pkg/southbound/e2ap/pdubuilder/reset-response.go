// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0
package pdubuilder

import (
	"fmt"
	"github.com/onosproject/onos-e2t/api/e2ap/v1beta2"
	e2ap_commondatatypes "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-commondatatypes"
	e2ap_constants "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-constants"
	e2ap_ies "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-ies"
	e2appducontents "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-contents"
	e2appdudescriptions "github.com/onosproject/onos-e2t/api/e2ap/v1beta2/e2ap-pdu-descriptions"
	"github.com/onosproject/onos-e2t/pkg/southbound/e2ap/types"
)

func CreateResetResponseE2apPdu(failureProcCode v1beta2.ProcedureCodeT, failureCrit e2ap_commondatatypes.Criticality,
	failureTrigMsg e2ap_commondatatypes.TriggeringMessage, reqID *types.RicRequest,
	critDiags []*types.CritDiag) (*e2appdudescriptions.E2ApPdu, error) {

	if reqID == nil && critDiags == nil {
		return nil, fmt.Errorf("no input parameters were passed - you should have at least one")
	}

	criticalityDiagnostics := e2appducontents.ResetResponseIes_ResetResponseIes2{
		Id:          int32(v1beta2.ProtocolIeIDCriticalityDiagnostics),
		Criticality: int32(e2ap_commondatatypes.Criticality_CRITICALITY_IGNORE),
		Value: &e2ap_ies.CriticalityDiagnostics{
			ProcedureCode: &e2ap_commondatatypes.ProcedureCode{
				Value: int32(failureProcCode), // range of Integer from e2ap-v01.00.asn1:1206, value were taken from line 1236 (same file)
			},
			TriggeringMessage:    failureTrigMsg,
			ProcedureCriticality: failureCrit, // from e2ap-v01.00.asn1:153
			RicRequestorId: &e2ap_ies.RicrequestId{
				RicRequestorId: int32(reqID.RequestorID),
				RicInstanceId:  int32(reqID.InstanceID),
			},
			IEsCriticalityDiagnostics: &e2ap_ies.CriticalityDiagnosticsIeList{
				Value: make([]*e2ap_ies.CriticalityDiagnosticsIeItem, 0),
			},
		},
		Presence: int32(e2ap_commondatatypes.Presence_PRESENCE_OPTIONAL),
	}

	for _, critDiag := range critDiags {
		criticDiagnostics := e2ap_ies.CriticalityDiagnosticsIeItem{
			IEcriticality: critDiag.IECriticality,
			IEId: &e2ap_commondatatypes.ProtocolIeId{
				Value: int32(critDiag.IEId), // value were taken from e2ap-v01.00.asn1:1278
			},
			TypeOfError: critDiag.TypeOfError,
		}
		criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value = append(criticalityDiagnostics.Value.IEsCriticalityDiagnostics.Value, &criticDiagnostics)
	}

	e2apPdu := e2appdudescriptions.E2ApPdu{
		E2ApPdu: &e2appdudescriptions.E2ApPdu_SuccessfulOutcome{
			SuccessfulOutcome: &e2appdudescriptions.SuccessfulOutcome{
				ProcedureCode: &e2appdudescriptions.E2ApElementaryProcedures{
					Reset_: &e2appdudescriptions.Reset{
						SuccessfulOutcome: &e2appducontents.ResetResponse{
							ProtocolIes: &e2appducontents.ResetResponseIes{
								ResetResponseIes2: &criticalityDiagnostics, //Criticality Diagnostics
							},
						},
						ProcedureCode: &e2ap_constants.IdReset{
							Value: int32(v1beta2.ProcedureCodeIDReset),
						},
						Criticality: &e2ap_commondatatypes.CriticalityReject{
							Criticality: e2ap_commondatatypes.Criticality_CRITICALITY_REJECT,
						},
					},
				},
			},
		},
	}
	if err := e2apPdu.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2ApPDU %s", err.Error())
	}
	return &e2apPdu, nil
}