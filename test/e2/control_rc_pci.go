// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"testing"
	"time"

	e2tapi "github.com/onosproject/onos-api/go/onos/e2t/e2"
	ransimtypes "github.com/onosproject/onos-api/go/onos/ransim/types"

	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/indication"

	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	"google.golang.org/protobuf/proto"

	e2sm_rc_pre_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre/v1/e2sm-rc-pre-ies"

	"github.com/stretchr/testify/assert"

	e2client "github.com/onosproject/onos-ric-sdk-go/pkg/e2"

	"github.com/onosproject/onos-e2t/test/utils"
)

const (
	ranParameterValue int32 = 200
	ranParameterName        = "pci"
	ranParameterID    int32 = 1
	defaultPlmnID           = 314628
	priority                = 10
)

// TestControl tests E2 control procedure using ransim and SDK
func (s *TestSuite) TestControl(t *testing.T) {
	sim := utils.CreateRanSimulatorWithNameOrDie(t, "control")
	assert.NotNil(t, sim)
	ch := make(chan indication.Indication)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	clientConfig := e2client.Config{
		AppID: "control-pci-test",
		E2TService: e2client.ServiceConfig{
			Host: utils.E2TServiceHost,
			Port: utils.E2TServicePort,
		},
		SubscriptionService: e2client.ServiceConfig{
			Host: utils.SubscriptionServiceHost,
			Port: utils.SubscriptionServicePort,
		},
	}
	nodeClient := utils.GetRansimNodeClient(t, sim)
	assert.NotNil(t, nodeClient)
	cellClient := utils.GetRansimCellClient(t, sim)
	assert.NotNil(t, cellClient)
	client, err := e2client.NewClient(clientConfig)
	assert.NoError(t, err)

	nodeIDs, err := utils.GetNodeIDs()
	assert.NoError(t, err)
	testNodeID := nodeIDs[0]

	// Subscription
	eventTriggerBytes, err := utils.CreateRcEventTrigger()
	assert.NoError(t, err)

	subRequest := utils.Subscription{
		NodeID:               testNodeID,
		EncodingType:         subapi.Encoding_ENCODING_PROTO,
		ActionType:           subapi.ActionType_ACTION_TYPE_REPORT,
		EventTrigger:         eventTriggerBytes,
		ServiceModelID:       utils.RcServiceModelID,
		ActionID:             100,
		SubSequentActionType: subapi.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
		TimeToWait:           subapi.TimeToWait_TIME_TO_WAIT_ZERO,
	}

	subReq, err := subRequest.Create()
	assert.NoError(t, err)

	sub, err := client.Subscribe(ctx, subReq, ch)
	assert.NoError(t, err)
	indMessage := <-ch
	header := indMessage.Payload.Header
	ricIndicationHeader := e2sm_rc_pre_ies.E2SmRcPreIndicationHeader{}
	err = proto.Unmarshal(header, &ricIndicationHeader)
	assert.NoError(t, err)
	testEci := ricIndicationHeader.GetIndicationHeaderFormat1().GetCgi().GetEUtraCgi().GetEUtracellIdentity().Value.Value
	plmnID := ransimtypes.NewUint24(defaultPlmnID)

	rcControlHeader := utils.RcControlHeader{
		Priority: priority,
		CellID:   testEci,
		PlmnID:   plmnID.ToBytes(),
	}

	rcControlMessage := utils.RcControlMessage{
		RanParameterName:  ranParameterName,
		RanParameterID:    ranParameterID,
		RanParameterValue: ranParameterValue,
	}

	controlMessageBytes, err := rcControlMessage.CreateRcControlMessage()
	assert.NoError(t, err)
	controlHeaderBytes, err := rcControlHeader.CreateRcControlHeader()
	assert.NoError(t, err)

	controlRequest := utils.Control{
		NodeID:            testNodeID,
		EncodingType:      e2tapi.EncodingType_PROTO,
		ServiceModelID:    utils.RcServiceModelID,
		ControlAckRequest: e2tapi.ControlAckRequest_ACK,
		ControlMessage:    controlMessageBytes,
		ControlHeader:     controlHeaderBytes,
	}

	request, err := controlRequest.Create()
	assert.NoError(t, err)
	response, err := client.Control(ctx, request)
	assert.NoError(t, err)
	if response == nil {
		t.Fail()
	}

	ack := response.GetControlAcknowledge()
	failure := response.GetControlFailure()
	if ack != nil {
		controlOutcome := &e2sm_rc_pre_ies.E2SmRcPreControlOutcome{}
		err = proto.Unmarshal(ack.GetControlOutcome(), controlOutcome)
		assert.NoError(t, err)

		outcomeRanParameterID := controlOutcome.
			GetControlOutcomeFormat1().
			GetOutcomeElementList()[0].
			RanParameterId.Value
		outcomeRanParameterValue := controlOutcome.
			GetControlOutcomeFormat1().
			GetOutcomeElementList()[0].
			RanParameterValue.
			GetValueInt()

		assert.Equal(t, ranParameterID, outcomeRanParameterID)
		assert.Equal(t, ranParameterValue, outcomeRanParameterValue)
	}
	if failure != nil {
		t.Fail()
	}

	err = sub.Close()
	assert.NoError(t, err)
	err = sim.Uninstall()
	assert.NoError(t, err)

}