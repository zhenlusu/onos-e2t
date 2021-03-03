// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"testing"
	"time"

	subtaskapi "github.com/onosproject/onos-api/go/onos/e2sub/task"

	subapi "github.com/onosproject/onos-api/go/onos/e2sub/subscription"
	"github.com/onosproject/onos-e2t/test/utils"
	e2client "github.com/onosproject/onos-ric-sdk-go/pkg/e2"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/indication"
	"github.com/stretchr/testify/assert"
)

// TestInvalidActionID tests invalid action ID (i.e. INSERT action) for kpm service model that
// supports just REPORT action
func (s *TestSuite) TestInvalidActionID(t *testing.T) {
	sim := utils.CreateRanSimulatorWithNameOrDie(t, "ran-simulator")

	clientConfig := e2client.Config{
		AppID: "invalid-action-id",
		SubscriptionService: e2client.ServiceConfig{
			Host: utils.SubscriptionServiceHost,
			Port: utils.SubscriptionServicePort,
		},
	}
	client, err := e2client.NewClient(clientConfig)
	assert.NoError(t, err)

	ch := make(chan indication.Indication)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nodeIDs, err := utils.GetNodeIDs()
	assert.NoError(t, err)

	eventTriggerBytes, err := utils.CreateKpmEventTrigger(12)
	assert.NoError(t, err)

	subRequest := utils.Subscription{
		NodeID:               nodeIDs[0],
		EncodingType:         subapi.Encoding_ENCODING_PROTO,
		ActionType:           subapi.ActionType_ACTION_TYPE_INSERT,
		EventTrigger:         eventTriggerBytes,
		ServiceModelID:       utils.KpmServiceModelID,
		ActionID:             100,
		SubSequentActionType: subapi.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
		TimeToWait:           subapi.TimeToWait_TIME_TO_WAIT_ZERO,
	}

	subReq, err := subRequest.Create()
	assert.NoError(t, err)

	sub, err := client.Subscribe(ctx, subReq, ch)
	assert.NoError(t, err)

	select {
	case err = <-sub.Err():
		t.Log(err.Error())
		assert.Equal(t, err.Error(), subtaskapi.Cause_CAUSE_RIC_ACTION_NOT_SUPPORTED.String())
	case <-time.After(10 * time.Second):
		t.Fatal("test is failed because of timeout")

	}

	err = sub.Close()
	assert.NoError(t, err)
	err = sim.Uninstall()
	assert.NoError(t, err)

}