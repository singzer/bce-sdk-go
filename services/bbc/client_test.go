package bbc

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"testing"

	"github.com/baidubce/bce-sdk-go/services/bbc/api"
	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	BBC_CLIENT              *Client
	BBC_TestCdsId           string
	BBC_TestBccId           string
	BBC_TestSecurityGroupId string
	BBC_TestImageId         string
	BBC_TestSnapshotId      string
	BBC_TestAspId           string
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK       string
	SK       string
	Endpoint string
}

func init() {
	var err error
	ak:= os.Getenv("BAIDUCLOUD_SECRET_ID")
	sk:= os.Getenv("BAIDUCLOUD_SECRET_KEY")
	ep:= os.Getenv("BAIDUCLOUD_ENDPOINT")
	BBC_CLIENT, err = NewClient(ak,sk,ep)
	if err!=nil{
		log.Fatal("ak/sk or endpoint error")
		os.Exit(1)
	}
	log.SetLogLevel(log.WARN)
	//log.SetLogLevel(log.DEBUG)
}

// ExpectEqual is the helper function for test each case
func ExpectEqual(alert func(format string, args ...interface{}),
	expected interface{}, actual interface{}) bool {
	expectedValue, actualValue := reflect.ValueOf(expected), reflect.ValueOf(actual)
	equal := false
	switch {
	case expected == nil && actual == nil:
		return true
	case expected != nil && actual == nil:
		equal = expectedValue.IsNil()
	case expected == nil && actual != nil:
		equal = actualValue.IsNil()
	default:
		if actualType := reflect.TypeOf(actual); actualType != nil {
			if expectedValue.IsValid() && expectedValue.Type().ConvertibleTo(actualType) {
				equal = reflect.DeepEqual(expectedValue.Convert(actualType).Interface(), actual)
			}
		}
	}
	if !equal {
		_, file, line, _ := runtime.Caller(1)
		alert("%s:%d: missmatch, expect %v but %v", file, line, expected, actual)
		return false
	}
	return true
}

//func TestCreateInstance(t *testing.T) {
//	createInstanceArgs := &api.CreateInstanceArgs{
//		ImageId: "m-DpgNg8lO",
//		Billing: api.Billing{
//			PaymentTiming: api.PaymentTimingPostPaid,
//		},
//		InstanceType:        api.InstanceTypeN1,
//		CpuCount:            1,
//		MemoryCapacityInGB:  1,
//		RootDiskSizeInGb:    40,
//		RootDiskStorageType: api.StorageTypeCloudHP1,
//		CreateCdsList: []api.CreateCdsModel{
//			{
//				StorageType: api.StorageTypeSSD,
//				CdsSizeInGB: 0,
//			},
//		},
//		AdminPass: "123qaz!@#",
//		Name:      "sdkTest",
//	}
//	createResult, err := BBC_CLIENT.CreateInstance(createInstanceArgs)
//	ExpectEqual(t.Errorf, err, nil)
//	BCC_TestBccId = createResult.InstanceIds[0]
//}
//
//func TestCreateSecurityGroup(t *testing.T) {
//	args := &api.CreateSecurityGroupArgs{
//		Name: "testSecurityGroup",
//		Rules: []api.SecurityGroupRuleModel{
//			{
//				Remark:        "备注",
//				Protocol:      "tcp",
//				PortRange:     "1-65535",
//				Direction:     "ingress",
//				SourceIp:      "",
//				SourceGroupId: "",
//			},
//		},
//	}
//	result, err := BBC_CLIENT.CreateSecurityGroup(args)
//	ExpectEqual(t.Errorf, err, nil)
//	BCC_TestSecurityGroupId = result.SecurityGroupId
//}
//
//func TestCreateImage(t *testing.T) {
//	args := &api.CreateImageArgs{
//		ImageName:  "testImageName",
//		InstanceId: BCC_TestBccId,
//	}
//	result, err := BBC_CLIENT.CreateImage(args)
//	ExpectEqual(t.Errorf, err, nil)
//	BCC_TestImageId = result.ImageId
//}

func TestListInstances(t *testing.T) {
	listArgs := &api.ListInstanceArgs{}
	_, err := BBC_CLIENT.ListInstances(listArgs)
	ExpectEqual(t.Errorf, err, nil)
}

//func TestGetInstanceDetail(t *testing.T) {
//	_, err := BBC_CLIENT.GetInstanceDetail(BCC_TestBccId)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestResizeInstance(t *testing.T) {
//	resizeArgs := &api.ResizeInstanceArgs{
//		CpuCount:           2,
//		MemoryCapacityInGB: 4,
//	}
//	err := BBC_CLIENT.ResizeInstance(BCC_TestBccId, resizeArgs)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestStopInstance(t *testing.T) {
//	err := BBC_CLIENT.StopInstance(BCC_TestBccId, true)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestStartInstance(t *testing.T) {
//	err := BBC_CLIENT.StartInstance(BCC_TestBccId)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestRebootInstance(t *testing.T) {
//	err := BBC_CLIENT.RebootInstance(BCC_TestBccId, true)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestRebuildInstance(t *testing.T) {
//	rebuildArgs := &api.RebuildInstanceArgs{
//		ImageId:   "m-DpgNg8lO",
//		AdminPass: "123qaz!@#",
//	}
//	err := BBC_CLIENT.RebuildInstance(BCC_TestBccId, rebuildArgs)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestChangeInstancePass(t *testing.T) {
//	changeArgs := &api.ChangeInstancePassArgs{
//		AdminPass: "321zaq#@!",
//	}
//	err := BBC_CLIENT.ChangeInstancePass(BCC_TestBccId, changeArgs)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestModifyInstanceAttribute(t *testing.T) {
//	modifyArgs := &api.ModifyInstanceAttributeArgs{
//		Name: "test-modify",
//	}
//	err := BBC_CLIENT.ModifyInstanceAttribute(BCC_TestBccId, modifyArgs)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestModifyInstanceDesc(t *testing.T) {
//	modifyArgs := &api.ModifyInstanceDescArgs{
//		Description: "test modify",
//	}
//	err := BBC_CLIENT.ModifyInstanceDesc(BCC_TestBccId, modifyArgs)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestGetInstanceVNC(t *testing.T) {
//	_, err := BBC_CLIENT.GetInstanceVNC(BCC_TestBccId)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestBindSecurityGroup(t *testing.T) {
//	err := BBC_CLIENT.BindSecurityGroup(BCC_TestBccId, BCC_TestSecurityGroupId)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestUnBindSecurityGroup(t *testing.T) {
//	err := BBC_CLIENT.UnBindSecurityGroup(BCC_TestBccId, BCC_TestSecurityGroupId)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestCreateCDSVolume(t *testing.T) {
//	args := &api.CreateCDSVolumeArgs{
//		PurchaseCount: 1,
//		CdsSizeInGB:   40,
//		StorageType:   api.StorageTypeSSD,
//		Billing: &api.Billing{
//			PaymentTiming: api.PaymentTimingPostPaid,
//		},
//	}
//
//	result, err := BBC_CLIENT.CreateCDSVolume(args)
//	ExpectEqual(t.Errorf, err, nil)
//	BCC_TestCdsId = result.VolumeIds[0]
//}
//
//func TestCreateSnapshot(t *testing.T) {
//	args := &api.CreateSnapshotArgs{
//		VolumeId:     BCC_TestCdsId,
//		SnapshotName: "testSnapshotName",
//	}
//	result, err := BBC_CLIENT.CreateSnapshot(args)
//	ExpectEqual(t.Errorf, err, nil)
//
//	BCC_TestSnapshotId = result.SnapshotId
//}
//
//func TestListSnapshot(t *testing.T) {
//	args := &api.ListSnapshotArgs{}
//	_, err := BBC_CLIENT.ListSnapshot(args)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestGetSnapshotDetail(t *testing.T) {
//	_, err := BBC_CLIENT.GetSnapshotDetail(BCC_TestSnapshotId)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestDeleteSnapshot(t *testing.T) {
//	err := BBC_CLIENT.DeleteSnapshot(BCC_TestSnapshotId)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestCreateAutoSnapshotPolicy(t *testing.T) {
//	args := &api.CreateASPArgs{
//		Name:           "testAspName",
//		TimePoints:     []string{"20"},
//		RepeatWeekdays: []string{"1", "5"},
//		RetentionDays:  "7",
//	}
//	result, err := BBC_CLIENT.CreateAutoSnapshotPolicy(args)
//	ExpectEqual(t.Errorf, err, nil)
//	BCC_TestAspId = result.AspId
//}
//
//func TestAttachAutoSnapshotPolicy(t *testing.T) {
//	args := &api.AttachASPArgs{
//		VolumeIds: []string{BCC_TestCdsId},
//	}
//	err := BBC_CLIENT.AttachAutoSnapshotPolicy(BCC_TestAspId, args)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestDetachAutoSnapshotPolicy(t *testing.T) {
//	args := &api.DetachASPArgs{
//		VolumeIds: []string{BCC_TestCdsId},
//	}
//	err := BBC_CLIENT.DetachAutoSnapshotPolicy(BCC_TestAspId, args)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestListAutoSnapshotPolicy(t *testing.T) {
//	args := &api.ListASPArgs{}
//	_, err := BBC_CLIENT.ListAutoSnapshotPolicy(args)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestGetAutoSnapshotPolicy(t *testing.T) {
//	_, err := BBC_CLIENT.GetAutoSnapshotPolicy(BCC_TestAspId)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestDeleteAutoSnapshotPolicy(t *testing.T) {
//	err := BBC_CLIENT.DeleteAutoSnapshotPolicy(BCC_TestAspId)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestListCDSVolume(t *testing.T) {
//	queryArgs := &api.ListCDSVolumeArgs{}
//	_, err := BBC_CLIENT.ListCDSVolume(queryArgs)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestGetCDSVolumeDetail(t *testing.T) {
//	_, err := BBC_CLIENT.GetCDSVolumeDetail(BCC_TestCdsId)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestAttachCDSVolume(t *testing.T) {
//	args := &api.AttachVolumeArgs{
//		InstanceId: BCC_TestBccId,
//	}
//
//	_, err := BBC_CLIENT.AttachCDSVolume(BCC_TestCdsId, args)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestDetachCDSVolume(t *testing.T) {
//	args := &api.DetachVolumeArgs{
//		InstanceId: BCC_TestBccId,
//	}
//
//	err := BBC_CLIENT.DetachCDSVolume(BCC_TestCdsId, args)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestResizeCDSVolume(t *testing.T) {
//	args := &api.ResizeCSDVolumeArgs{
//		NewCdsSizeInGB: 100,
//	}
//
//	err := BBC_CLIENT.ResizeCDSVolume(BCC_TestCdsId, args)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestPurchaseReservedCDSVolume(t *testing.T) {
//	args := &api.PurchaseReservedCSDVolumeArgs{
//		Billing: &api.Billing{
//			Reservation: &api.Reservation{
//				ReservationLength:   1,
//				ReservationTimeUnit: "Month",
//			},
//		},
//	}
//
//	err := BBC_CLIENT.PurchaseReservedCDSVolume(BCC_TestCdsId, args)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestRenameCDSVolume(t *testing.T) {
//	args := &api.RenameCSDVolumeArgs{
//		Name: "testRenamedName",
//	}
//
//	err := BBC_CLIENT.RenameCDSVolume(BCC_TestCdsId, args)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestModifyCDSVolume(t *testing.T) {
//	args := &api.ModifyCSDVolumeArgs{
//		CdsName: "modifiedName",
//		Desc:    "modifiedDesc",
//	}
//
//	err := BBC_CLIENT.ModifyCDSVolume(BCC_TestCdsId, args)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestModifyChargeTypeCDSVolume(t *testing.T) {
//	args := &api.ModifyChargeTypeCSDVolumeArgs{
//		Billing: &api.Billing{
//			Reservation: &api.Reservation{
//				ReservationLength: 1,
//			},
//		},
//	}
//
//	err := BBC_CLIENT.ModifyChargeTypeCDSVolume(BCC_TestCdsId, args)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestDeleteCDSVolumeNew(t *testing.T) {
//	args := &api.DeleteCDSVolumeArgs{
//		AutoSnapshot: "on",
//	}
//
//	err := BBC_CLIENT.DeleteCDSVolumeNew(BCC_TestCdsId, args)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestDeleteCDSVolume(t *testing.T) {
//	err := BBC_CLIENT.DeleteCDSVolume(BCC_TestCdsId)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestListSecurityGroup(t *testing.T) {
//	queryArgs := &api.ListSecurityGroupArgs{}
//	_, err := BBC_CLIENT.ListSecurityGroup(queryArgs)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestAuthorizeSecurityGroupRule(t *testing.T) {
//	args := &api.AuthorizeSecurityGroupArgs{
//		Rule: &api.SecurityGroupRuleModel{
//			Remark:        "备注",
//			Protocol:      "udp",
//			PortRange:     "1-65535",
//			Direction:     "ingress",
//			SourceIp:      "",
//			SourceGroupId: "",
//		},
//	}
//	err := BBC_CLIENT.AuthorizeSecurityGroupRule(BCC_TestSecurityGroupId, args)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestRevokeSecurityGroupRule(t *testing.T) {
//	args := &api.RevokeSecurityGroupArgs{
//		Rule: &api.SecurityGroupRuleModel{
//			Remark:        "备注",
//			Protocol:      "udp",
//			PortRange:     "1-65535",
//			Direction:     "ingress",
//			SourceIp:      "",
//			SourceGroupId: "",
//		},
//	}
//	err := BBC_CLIENT.RevokeSecurityGroupRule(BCC_TestSecurityGroupId, args)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestDeleteSecurityGroupRule(t *testing.T) {
//	err := BBC_CLIENT.DeleteSecurityGroup(BCC_TestSecurityGroupId)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestListImage(t *testing.T) {
//	queryArgs := &api.ListImageArgs{}
//	_, err := BBC_CLIENT.ListImage(queryArgs)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestGetImageDetail(t *testing.T) {
//	_, err := BBC_CLIENT.GetImageDetail(BCC_TestImageId)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestRemoteCopyImage(t *testing.T) {
//	args := &api.RemoteCopyImageArgs{
//		Name:       "testRemoteCopy",
//		DestRegion: []string{"bj"},
//	}
//	err := BBC_CLIENT.RemoteCopyImage(BCC_TestImageId, args)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestCancelRemoteCopyImage(t *testing.T) {
//	err := BBC_CLIENT.CancelRemoteCopyImage(BCC_TestImageId)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestShareImage(t *testing.T) {
//	args := &api.SharedUser{
//		AccountId: "id",
//	}
//	err := BBC_CLIENT.ShareImage(BCC_TestImageId, args)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestUnShareImage(t *testing.T) {
//	args := &api.SharedUser{
//		AccountId: "id",
//	}
//	err := BBC_CLIENT.UnShareImage(BCC_TestImageId, args)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestGetImageSharedUser(t *testing.T) {
//	_, err := BBC_CLIENT.GetImageSharedUser(BCC_TestImageId)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestGetImageOS(t *testing.T) {
//	args := &api.GetImageOsArgs{}
//	_, err := BBC_CLIENT.GetImageOS(args)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestDeleteImage(t *testing.T) {
//	err := BBC_CLIENT.DeleteImage(BCC_TestImageId)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestDeleteInstance(t *testing.T) {
//	err := BBC_CLIENT.DeleteInstance(BCC_TestBccId)
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestListSpec(t *testing.T) {
//	_, err := BBC_CLIENT.ListSpec()
//	ExpectEqual(t.Errorf, err, nil)
//}
//
//func TestListZone(t *testing.T) {
//	_, err := BBC_CLIENT.ListZone()
//	ExpectEqual(t.Errorf, err, nil)
//}
//
