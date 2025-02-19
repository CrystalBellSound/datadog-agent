// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2024-present Datadog, Inc.

package aggregator

import (
	"fmt"
	"time"

	agentmodel "github.com/DataDog/agent-payload/v5/sbom"

	"github.com/DataDog/datadog-agent/test/fakeintake/api"

	"google.golang.org/protobuf/proto"
)

// SBOMPayload is a payload type for the sbom check
type SBOMPayload struct {
	*agentmodel.SBOMEntity
	collectedTime time.Time
}

func (p *SBOMPayload) name() string {
	return p.Id
}

// GetTags return the tags from a payload
func (p *SBOMPayload) GetTags() []string {
	return p.DdTags
}

// GetCollectedTime returns the time that the payload was received by the fake intake
func (p *SBOMPayload) GetCollectedTime() time.Time {
	return p.collectedTime
}

// ParseSbomPayload parses an api.Payload into a list of SbomPayload
func ParseSbomPayload(payload api.Payload) ([]*SBOMPayload, error) {
	inflated, err := inflate(payload.Data, payload.Encoding)
	if err != nil {
		return nil, fmt.Errorf("could not inflate payload: %w", err)
	}

	msg := agentmodel.SBOMPayload{}
	if err := proto.Unmarshal(inflated, &msg); err != nil {
		return nil, err
	}

	payloads := make([]*SBOMPayload, len(msg.Entities))
	for i, sbomEntity := range msg.Entities {
		payloads[i] = &SBOMPayload{SBOMEntity: sbomEntity, collectedTime: payload.Timestamp}
	}
	return payloads, nil
}

// SBOMAggregator is an Aggregator for SbomPayload
type SBOMAggregator struct {
	Aggregator[*SBOMPayload]
}

// NewSBOMAggregator returns a new SbomAggregator
func NewSBOMAggregator() SBOMAggregator {
	return SBOMAggregator{
		Aggregator: newAggregator(ParseSbomPayload),
	}
}
