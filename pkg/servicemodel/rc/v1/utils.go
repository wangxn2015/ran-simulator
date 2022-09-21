// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	ransimtypes "github.com/onosproject/onos-api/go/onos/ransim/types"
	"github.com/wangxn2015/onos-e2-sm/servicemodels/e2sm_rc/pdubuilder"
	e2smrc "github.com/wangxn2015/onos-e2-sm/servicemodels/e2sm_rc/servicemodel"
	e2smcommonies "github.com/wangxn2015/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-common-ies"
	e2smrcies "github.com/wangxn2015/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-rc-ies"
	v2 "github.com/wangxn2015/onos-e2t/api/e2ap/v2"
	e2appducontents "github.com/wangxn2015/onos-e2t/api/e2ap/v2/e2ap-pdu-contents"
	e2aptypes "github.com/wangxn2015/onos-e2t/pkg/southbound/e2ap/types"
	"github.com/wangxn2015/onos-lib-go/api/asn1/v1/asn1"
	"github.com/wangxn2015/onos-lib-go/pkg/errors"
	"github.com/wangxn2015/ran-simulator/pkg/utils"
	indicationutils "github.com/wangxn2015/ran-simulator/pkg/utils/e2ap/indication"
	subutils "github.com/wangxn2015/ran-simulator/pkg/utils/e2ap/subscription"
	"github.com/wangxn2015/ran-simulator/pkg/utils/e2sm/rc/v1/indication/headers/format1"
	"github.com/wangxn2015/ran-simulator/pkg/utils/e2sm/rc/v1/indication/messages/format3"
	"google.golang.org/protobuf/proto"
)

func getActionDefinitionMap(actionList []*e2appducontents.RicactionToBeSetupItemIes, ricActionsAccepted []*e2aptypes.RicActionID) (map[*e2aptypes.RicActionID]*e2smrcies.E2SmRcActionDefinition, error) {
	actionDefinitionsMap := make(map[*e2aptypes.RicActionID]*e2smrcies.E2SmRcActionDefinition)
	for _, action := range actionList {
		for _, actionID := range ricActionsAccepted {
			if action.GetValue().GetRicactionToBeSetupItem().GetRicActionId().GetValue() == int32(*actionID) {
				actionDefinitionBytes := action.GetValue().GetRicactionToBeSetupItem().GetRicActionDefinition().GetValue()
				var rcServiceModel e2smrc.RCServiceModel

				actionDefinitionProtoBytes, err := rcServiceModel.ActionDefinitionASN1toProto(actionDefinitionBytes)
				if err != nil {
					return nil, err
				}

				actionDefinition := &e2smrcies.E2SmRcActionDefinition{}
				err = proto.Unmarshal(actionDefinitionProtoBytes, actionDefinition)
				if err != nil {
					return nil, err
				}
				actionDefinitionsMap[actionID] = actionDefinition
			}
		}
	}
	return actionDefinitionsMap, nil
}

func getEventTrigger(request *e2appducontents.RicsubscriptionRequest) (*e2smrcies.E2SmRcEventTrigger, error) {
	var eventTriggerAsnBytes []byte
	for _, v := range request.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDRicsubscriptionDetails) {
			eventTriggerAsnBytes = v.GetValue().GetRicsubscriptionDetails().GetRicEventTriggerDefinition().GetValue()
			break
		}
	}

	var rcServiceModel e2smrc.RCServiceModel
	eventTriggerProtoBytes, err := rcServiceModel.EventTriggerDefinitionASN1toProto(eventTriggerAsnBytes)
	if err != nil {
		return nil, err
	}
	eventTriggerDefinition := &e2smrcies.E2SmRcEventTrigger{}
	err = proto.Unmarshal(eventTriggerProtoBytes, eventTriggerDefinition)
	if err != nil {
		return nil, err
	}

	return eventTriggerDefinition, nil
}

func getControlMessage(request *e2appducontents.RiccontrolRequest) (*e2smrcies.E2SmRcControlMessage, error) {
	var rcServiceModel e2smrc.RCServiceModel
	var controlMessageAsnBytes []byte
	for _, v := range request.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDRiccontrolMessage) {
			controlMessageAsnBytes = v.GetValue().GetRiccontrolMessage().GetValue()
			break
		}
	}
	controlMessageProtoBytes, err := rcServiceModel.ControlMessageASN1toProto(controlMessageAsnBytes)
	if err != nil {
		return nil, err
	}
	controlMessage := &e2smrcies.E2SmRcControlMessage{}
	err = proto.Unmarshal(controlMessageProtoBytes, controlMessage)

	if err != nil {
		return nil, err
	}
	return controlMessage, nil
}

func getControlHeader(request *e2appducontents.RiccontrolRequest) (*e2smrcies.E2SmRcControlHeader, error) {
	var rcServiceModel e2smrc.RCServiceModel
	var controlHeaderAsnBytes []byte
	for _, v := range request.GetProtocolIes() {
		if v.Id == int32(v2.ProtocolIeIDRiccontrolHeader) {
			controlHeaderAsnBytes = v.GetValue().GetRiccontrolHeader().GetValue()
			break
		}
	}
	controlHeaderProtoBytes, err := rcServiceModel.ControlHeaderASN1toProto(controlHeaderAsnBytes)
	if err != nil {
		return nil, err
	}
	controlHeader := &e2smrcies.E2SmRcControlHeader{}
	err = proto.Unmarshal(controlHeaderProtoBytes, controlHeader)
	if err != nil {
		return nil, err
	}

	return controlHeader, nil
}

func createRANParametersInsertStyle3List() ([]*e2smrcies.InsertIndicationRanparameterItem, error) {
	// RAN Parameters for Insert Style 3
	insertRANParametersStyle3List := make([]*e2smrcies.InsertIndicationRanparameterItem, 0)
	ranParameter1, err := pdubuilder.CreateInsertIndicationRanparameterItem(1, "Target Primary Cell ID")
	if err != nil {
		return nil, err
	}
	insertRANParametersStyle3List = append(insertRANParametersStyle3List, ranParameter1)

	ranParameter2, err := pdubuilder.CreateInsertIndicationRanparameterItem(2, "Target Cell")
	if err != nil {
		return nil, err
	}
	insertRANParametersStyle3List = append(insertRANParametersStyle3List, ranParameter2)

	ranParameter3, err := pdubuilder.CreateInsertIndicationRanparameterItem(3, "NR Cell")
	if err != nil {
		return nil, err
	}
	insertRANParametersStyle3List = append(insertRANParametersStyle3List, ranParameter3)

	ranParameter4, err := pdubuilder.CreateInsertIndicationRanparameterItem(4, "NR CGI")
	if err != nil {
		return nil, err
	}
	insertRANParametersStyle3List = append(insertRANParametersStyle3List, ranParameter4)

	ranParameter5, err := pdubuilder.CreateInsertIndicationRanparameterItem(7, "List of PDU sessions for handover")
	if err != nil {
		return nil, err
	}
	insertRANParametersStyle3List = append(insertRANParametersStyle3List, ranParameter5)

	ranParameter6, err := pdubuilder.CreateInsertIndicationRanparameterItem(13, "List of DRBs for handover")
	if err != nil {
		return nil, err
	}
	insertRANParametersStyle3List = append(insertRANParametersStyle3List, ranParameter6)
	return insertRANParametersStyle3List, nil

}

func createRANParametersReportStyle3List() ([]*e2smrcies.ReportRanparameterItem, error) {
	// RAN Parameters for Report Style 3
	reportParametersStyle3List := make([]*e2smrcies.ReportRanparameterItem, 0)

	return reportParametersStyle3List, nil
}

func createRANParametersReportStyle2List() ([]*e2smrcies.ReportRanparameterItem, error) {
	// RAN Parameters for Report Style 2
	reportParametersStyle2List := make([]*e2smrcies.ReportRanparameterItem, 0)
	ranParameter1, err := pdubuilder.CreateReportRanparameterItem(1, "Current UE ID")
	if err != nil {
		return nil, err
	}

	reportParametersStyle2List = append(reportParametersStyle2List, ranParameter1)

	ranParameter2, err := pdubuilder.CreateReportRanparameterItem(21001, "S-NSSAI")
	if err != nil {
		return nil, err
	}
	reportParametersStyle2List = append(reportParametersStyle2List, ranParameter2)

	ranParameter3, err := pdubuilder.CreateReportRanparameterItem(21002, "SST")
	if err != nil {
		return nil, err
	}
	reportParametersStyle2List = append(reportParametersStyle2List, ranParameter3)

	ranParameter4, err := pdubuilder.CreateReportRanparameterItem(21003, "SD")
	if err != nil {
		return nil, err
	}
	reportParametersStyle2List = append(reportParametersStyle2List, ranParameter4)

	ranParameter5, err := pdubuilder.CreateReportRanparameterItem(27108, "Best Neighboring Cell")
	if err != nil {
		return nil, err
	}
	reportParametersStyle2List = append(reportParametersStyle2List, ranParameter5)

	ranParameter6, err := pdubuilder.CreateReportRanparameterItem(21528, "List of Neighbor cells")
	if err != nil {
		return nil, err
	}
	reportParametersStyle2List = append(reportParametersStyle2List, ranParameter6)

	ranParameter7, err := pdubuilder.CreateReportRanparameterItem(10102, "Cell Results")
	if err != nil {
		return nil, err
	}
	reportParametersStyle2List = append(reportParametersStyle2List, ranParameter7)

	ranParameter8, err := pdubuilder.CreateReportRanparameterItem(10103, "SSB Results")
	if err != nil {
		return nil, err
	}
	reportParametersStyle2List = append(reportParametersStyle2List, ranParameter8)

	ranParameter9, err := pdubuilder.CreateReportRanparameterItem(12501, "RSRP")
	if err != nil {
		return nil, err
	}
	reportParametersStyle2List = append(reportParametersStyle2List, ranParameter9)

	ranParameter10, err := pdubuilder.CreateReportRanparameterItem(12502, "RSRQ")
	if err != nil {
		return nil, err
	}
	reportParametersStyle2List = append(reportParametersStyle2List, ranParameter10)

	ranParameter11, err := pdubuilder.CreateReportRanparameterItem(12503, "SINR")
	if err != nil {
		return nil, err
	}
	reportParametersStyle2List = append(reportParametersStyle2List, ranParameter11)
	return reportParametersStyle2List, nil
}

func (c *Client) getCellPCI(ctx context.Context, ncgi ransimtypes.NCGI) (int32, error) {
	cell, err := c.ServiceModel.CellStore.Get(ctx, ncgi)
	if err != nil {
		return 0, err
	}

	return int32(cell.PCI), nil
}

func (c *Client) getARFCN(ctx context.Context, ncgi ransimtypes.NCGI) (int32, error) {
	cell, err := c.ServiceModel.CellStore.Get(ctx, ncgi)
	if err != nil {
		return 0, err
	}

	return int32(cell.Earfcn), nil
}

func (c *Client) createRICIndicationFormat3(ctx context.Context, cells []ransimtypes.NCGI, subscription *subutils.Subscription, e2NodeInfoChangeID int32) (*e2appducontents.Ricindication, error) {
	headerFormat1 := format1.NewIndicationHeader(format1.WithEventConditionID(e2NodeInfoChangeID))
	indicationHeaderAsn1Bytes, err := headerFormat1.ToAsn1Bytes()
	if err != nil {
		return nil, err
	}
	plmnID := c.getPlmnID()

	cellInfoList := make([]*e2smrcies.E2SmRcIndicationMessageFormat3Item, 0)
	for _, ncgi := range cells {
		cell, err := c.ServiceModel.CellStore.Get(ctx, ncgi)
		if err != nil {
			return nil, err
		}

		nci := ransimtypes.GetNCI(ncgi)
		nrCGI, err := pdubuilder.CreateNrCgi(plmnID.Value().ToBytes(), &asn1.BitString{
			Value: utils.Uint64ToBitString(uint64(nci), 36),
			Len:   36,
		})
		if err != nil {
			return nil, err
		}

		cgi, err := pdubuilder.CreateCgiNRCgi(nrCGI)
		if err != nil {
			return nil, err
		}

		pci, err := c.getCellPCI(ctx, ncgi)
		if err != nil {
			return nil, err
		}
		cellPCI, err := pdubuilder.CreateServingCellPciNR(pci)
		if err != nil {
			return nil, err
		}

		earfcn, err := c.getARFCN(ctx, ncgi)
		if err != nil {
			return nil, err
		}
		cellArfcn, err := pdubuilder.CreateServingCellArfcnNR(earfcn)
		if err != nil {
			return nil, err
		}

		neighborCellList := make([]*e2smrcies.NeighborCellItem, 0)
		for _, neighborNCGI := range cell.Neighbors {
			neighborNci := ransimtypes.GetNCI(neighborNCGI)
			neighborNrCGI, err := pdubuilder.CreateNrCgi(plmnID.Value().ToBytes(), &asn1.BitString{
				Value: utils.Uint64ToBitString(uint64(neighborNci), 36),
				Len:   36,
			})
			if err != nil {
				return nil, err
			}

			neighborPci, err := c.getCellPCI(ctx, neighborNCGI)
			if err != nil {
				return nil, err
			}

			neighborEarfcn, err := c.getARFCN(ctx, neighborNCGI)
			if err != nil {
				return nil, err
			}

			neighborNrArfcn, err := pdubuilder.CreateNrArfcn(neighborEarfcn)
			if err != nil {
				return nil, err
			}

			nrFrequencyBandList := &e2smcommonies.NrfrequencyBandList{
				Value: make([]*e2smcommonies.NrfrequencyBandItem, 0),
			}

			supportedSulbandList := make([]*e2smcommonies.SupportedSulfreqBandItem, 0)
			frequencyBandItem, err := pdubuilder.CreateNrfrequencyBandItem(1, &e2smcommonies.SupportedSulbandList{
				Value: supportedSulbandList,
			})
			if err != nil {
				return nil, err
			}

			nrFrequencyBandList.Value = append(nrFrequencyBandList.Value, frequencyBandItem)

			nrFrequencyInfo, err := pdubuilder.CreateNrfrequencyInfo(neighborNrArfcn, nrFrequencyBandList)
			if err != nil {
				return nil, err
			}
			nrFrequencyInfo.SetFrequencyShift7P5Khz(pdubuilder.CreateNrfrequencyShift7P5KhzTrue())

			neighborCellItem, err := pdubuilder.CreateNeighborCellItemRanTypeChoiceNr(neighborNrCGI, neighborPci, []byte{0xFF, 0xFF, 0xFF}, pdubuilder.CreateNRModeInfoFDD(),
				nrFrequencyInfo, pdubuilder.CreateX2XNEstablishedTrue(), pdubuilder.CreateHOValidatedTrue(), 1)
			if err != nil {
				return nil, err
			}
			neighborCellList = append(neighborCellList, neighborCellItem)

		}

		neighborRelationTable, err := pdubuilder.CreateNeighborRelationInfo(cellPCI, cellArfcn, &e2smrcies.NeighborCellList{
			Value: neighborCellList,
		})
		if err != nil {
			return nil, err
		}

		item, err := pdubuilder.CreateE2SmRcIndicationMessageFormat3Item(cgi)
		if err != nil {
			return nil, err
		}
		item.SetNeighborRelationTable(neighborRelationTable).SetCellDeleted(false)
		cellInfoList = append(cellInfoList, item)

	}

	messageFormat3 := format3.NewIndicationMessage(format3.WithMessageItems(cellInfoList))
	indicationMessageAsn1Bytes, err := messageFormat3.ToAsn1Bytes()
	if err != nil {
		return nil, err
	}

	// Creates e2 indication
	indication := indicationutils.NewIndication(
		indicationutils.WithRicInstanceID(subscription.GetRicInstanceID()),
		indicationutils.WithRanFuncID(subscription.GetRanFuncID()),
		indicationutils.WithRequestID(subscription.GetReqID()),
		indicationutils.WithIndicationHeader(indicationHeaderAsn1Bytes),
		indicationutils.WithIndicationMessage(indicationMessageAsn1Bytes))

	ricIndication, err := indication.Build()
	if err != nil {
		return nil, err
	}
	return ricIndication, nil
}

func (c *Client) getPlmnID() ransimtypes.Uint24 {
	plmnIDUint24 := ransimtypes.Uint24{}
	plmnIDUint24.Set(uint32(c.ServiceModel.Model.PlmnID))
	return plmnIDUint24
}

// checkAndSetPCI check if the control header and message including the required info for changing the PCI value for a specific cell
func (c *Client) checkAndSetPCI(ctx context.Context, controlHeader *e2smrcies.E2SmRcControlHeader, controlMessage *e2smrcies.E2SmRcControlMessage) error {
	headerFormat1 := controlHeader.GetRicControlHeaderFormats().GetControlHeaderFormat1()
	if headerFormat1 != nil {
		// Process Control Style 9, Control Action ID 1 (i.e. PCI assignment)
		if headerFormat1.GetRicStyleType().Value == controlStyleType9 &&
			headerFormat1.GetRicControlActionId().Value == controlActionID1 {
			messageFormat1 := controlMessage.GetRicControlMessageFormats().GetControlMessageFormat1()

			if messageFormat1 != nil {
				var pciValue int64
				for _, ranParameter := range messageFormat1.GetRanPList() {
					var ncgi ransimtypes.NCGI
					// Extracts NR PCI ran parameter
					ranParameterID := ranParameter.GetRanParameterId().Value
					if ranParameterID == PCIRANParameterID {
						ranParameterValue := ranParameter.GetRanParameterValueType().GetRanPChoiceStructure().GetRanParameterStructure().GetSequenceOfRanParameters()[0].GetRanParameterValueType().GetRanPChoiceElementFalse()
						if ranParameterValue != nil {
							pciValue = ranParameterValue.GetRanParameterValue().GetValueInt()
						} else {
							return errors.NewInvalid("PCI ran parameter is not set")
						}
					}
					// Extracts NCGI ran parameter
					if ranParameterID == NCGIRANParameterID {
						ncgiStruct := ranParameter.GetRanParameterValueType().GetRanPChoiceStructure().GetRanParameterStructure().GetSequenceOfRanParameters()[0].GetRanParameterValueType().GetRanPChoiceStructure()
						if ncgiStruct != nil {
							ncgiFields := ncgiStruct.GetRanParameterStructure().GetSequenceOfRanParameters()
							if len(ncgiFields) == 2 {
								plmnIDField := ncgiFields[0]
								var plmnID ransimtypes.PlmnID
								var nci ransimtypes.NCI
								if plmnIDField != nil {
									plmnIDBitString := plmnIDField.GetRanParameterValueType().GetRanPChoiceElementFalse().GetRanParameterValue().GetValueOctS()
									plmnID = ransimtypes.PlmnID(ransimtypes.Uint24ToUint32(plmnIDBitString))

								} else {
									return errors.NewInvalid("plmn ID ran parameter is not set")
								}
								nciField := ncgiFields[1]
								if nciField != nil {
									nciBitString := nciField.GetRanParameterValueType().GetRanPChoiceElementFalse().GetRanParameterValue().GetValueBitS()
									nci = ransimtypes.NCI(utils.BitStringToUint64(nciBitString.GetValue(), int(nciBitString.GetLen())))
								} else {
									return errors.NewInvalid("NCI ran parameter is not set")
								}
								ncgi = ransimtypes.ToNCGI(plmnID, nci)
								cell, err := c.ServiceModel.CellStore.Get(ctx, ncgi)
								if err != nil {
									return err
								}
								cell.PCI = uint32(pciValue)
								err = c.ServiceModel.CellStore.Update(ctx, cell)
								if err != nil {
									return err
								}
							}
						} else {
							return errors.NewInvalid("NCGI ran parameter is not set")
						}
					}
				}
			}
		}
	}

	return nil

}
