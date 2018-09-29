// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package redshiftiface provides an interface to enable mocking the Amazon Redshift service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package redshiftiface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
)

// RedshiftAPI provides an interface to enable mocking the
// redshift.Redshift service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Redshift.
//    func myFunc(svc redshiftiface.RedshiftAPI) bool {
//        // Make svc.AcceptReservedNodeExchange request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := redshift.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockRedshiftClient struct {
//        redshiftiface.RedshiftAPI
//    }
//    func (m *mockRedshiftClient) AcceptReservedNodeExchange(input *redshift.AcceptReservedNodeExchangeInput) (*redshift.AcceptReservedNodeExchangeOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockRedshiftClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type RedshiftAPI interface {
	AcceptReservedNodeExchangeRequest(*redshift.AcceptReservedNodeExchangeInput) redshift.AcceptReservedNodeExchangeRequest

	AuthorizeClusterSecurityGroupIngressRequest(*redshift.AuthorizeClusterSecurityGroupIngressInput) redshift.AuthorizeClusterSecurityGroupIngressRequest

	AuthorizeSnapshotAccessRequest(*redshift.AuthorizeSnapshotAccessInput) redshift.AuthorizeSnapshotAccessRequest

	CopyClusterSnapshotRequest(*redshift.CopyClusterSnapshotInput) redshift.CopyClusterSnapshotRequest

	CreateClusterRequest(*redshift.CreateClusterInput) redshift.CreateClusterRequest

	CreateClusterParameterGroupRequest(*redshift.CreateClusterParameterGroupInput) redshift.CreateClusterParameterGroupRequest

	CreateClusterSecurityGroupRequest(*redshift.CreateClusterSecurityGroupInput) redshift.CreateClusterSecurityGroupRequest

	CreateClusterSnapshotRequest(*redshift.CreateClusterSnapshotInput) redshift.CreateClusterSnapshotRequest

	CreateClusterSubnetGroupRequest(*redshift.CreateClusterSubnetGroupInput) redshift.CreateClusterSubnetGroupRequest

	CreateEventSubscriptionRequest(*redshift.CreateEventSubscriptionInput) redshift.CreateEventSubscriptionRequest

	CreateHsmClientCertificateRequest(*redshift.CreateHsmClientCertificateInput) redshift.CreateHsmClientCertificateRequest

	CreateHsmConfigurationRequest(*redshift.CreateHsmConfigurationInput) redshift.CreateHsmConfigurationRequest

	CreateSnapshotCopyGrantRequest(*redshift.CreateSnapshotCopyGrantInput) redshift.CreateSnapshotCopyGrantRequest

	CreateTagsRequest(*redshift.CreateTagsInput) redshift.CreateTagsRequest

	DeleteClusterRequest(*redshift.DeleteClusterInput) redshift.DeleteClusterRequest

	DeleteClusterParameterGroupRequest(*redshift.DeleteClusterParameterGroupInput) redshift.DeleteClusterParameterGroupRequest

	DeleteClusterSecurityGroupRequest(*redshift.DeleteClusterSecurityGroupInput) redshift.DeleteClusterSecurityGroupRequest

	DeleteClusterSnapshotRequest(*redshift.DeleteClusterSnapshotInput) redshift.DeleteClusterSnapshotRequest

	DeleteClusterSubnetGroupRequest(*redshift.DeleteClusterSubnetGroupInput) redshift.DeleteClusterSubnetGroupRequest

	DeleteEventSubscriptionRequest(*redshift.DeleteEventSubscriptionInput) redshift.DeleteEventSubscriptionRequest

	DeleteHsmClientCertificateRequest(*redshift.DeleteHsmClientCertificateInput) redshift.DeleteHsmClientCertificateRequest

	DeleteHsmConfigurationRequest(*redshift.DeleteHsmConfigurationInput) redshift.DeleteHsmConfigurationRequest

	DeleteSnapshotCopyGrantRequest(*redshift.DeleteSnapshotCopyGrantInput) redshift.DeleteSnapshotCopyGrantRequest

	DeleteTagsRequest(*redshift.DeleteTagsInput) redshift.DeleteTagsRequest

	DescribeClusterDbRevisionsRequest(*redshift.DescribeClusterDbRevisionsInput) redshift.DescribeClusterDbRevisionsRequest

	DescribeClusterParameterGroupsRequest(*redshift.DescribeClusterParameterGroupsInput) redshift.DescribeClusterParameterGroupsRequest

	DescribeClusterParametersRequest(*redshift.DescribeClusterParametersInput) redshift.DescribeClusterParametersRequest

	DescribeClusterSecurityGroupsRequest(*redshift.DescribeClusterSecurityGroupsInput) redshift.DescribeClusterSecurityGroupsRequest

	DescribeClusterSnapshotsRequest(*redshift.DescribeClusterSnapshotsInput) redshift.DescribeClusterSnapshotsRequest

	DescribeClusterSubnetGroupsRequest(*redshift.DescribeClusterSubnetGroupsInput) redshift.DescribeClusterSubnetGroupsRequest

	DescribeClusterTracksRequest(*redshift.DescribeClusterTracksInput) redshift.DescribeClusterTracksRequest

	DescribeClusterVersionsRequest(*redshift.DescribeClusterVersionsInput) redshift.DescribeClusterVersionsRequest

	DescribeClustersRequest(*redshift.DescribeClustersInput) redshift.DescribeClustersRequest

	DescribeDefaultClusterParametersRequest(*redshift.DescribeDefaultClusterParametersInput) redshift.DescribeDefaultClusterParametersRequest

	DescribeEventCategoriesRequest(*redshift.DescribeEventCategoriesInput) redshift.DescribeEventCategoriesRequest

	DescribeEventSubscriptionsRequest(*redshift.DescribeEventSubscriptionsInput) redshift.DescribeEventSubscriptionsRequest

	DescribeEventsRequest(*redshift.DescribeEventsInput) redshift.DescribeEventsRequest

	DescribeHsmClientCertificatesRequest(*redshift.DescribeHsmClientCertificatesInput) redshift.DescribeHsmClientCertificatesRequest

	DescribeHsmConfigurationsRequest(*redshift.DescribeHsmConfigurationsInput) redshift.DescribeHsmConfigurationsRequest

	DescribeLoggingStatusRequest(*redshift.DescribeLoggingStatusInput) redshift.DescribeLoggingStatusRequest

	DescribeOrderableClusterOptionsRequest(*redshift.DescribeOrderableClusterOptionsInput) redshift.DescribeOrderableClusterOptionsRequest

	DescribeReservedNodeOfferingsRequest(*redshift.DescribeReservedNodeOfferingsInput) redshift.DescribeReservedNodeOfferingsRequest

	DescribeReservedNodesRequest(*redshift.DescribeReservedNodesInput) redshift.DescribeReservedNodesRequest

	DescribeResizeRequest(*redshift.DescribeResizeInput) redshift.DescribeResizeRequest

	DescribeSnapshotCopyGrantsRequest(*redshift.DescribeSnapshotCopyGrantsInput) redshift.DescribeSnapshotCopyGrantsRequest

	DescribeTableRestoreStatusRequest(*redshift.DescribeTableRestoreStatusInput) redshift.DescribeTableRestoreStatusRequest

	DescribeTagsRequest(*redshift.DescribeTagsInput) redshift.DescribeTagsRequest

	DisableLoggingRequest(*redshift.DisableLoggingInput) redshift.DisableLoggingRequest

	DisableSnapshotCopyRequest(*redshift.DisableSnapshotCopyInput) redshift.DisableSnapshotCopyRequest

	EnableLoggingRequest(*redshift.EnableLoggingInput) redshift.EnableLoggingRequest

	EnableSnapshotCopyRequest(*redshift.EnableSnapshotCopyInput) redshift.EnableSnapshotCopyRequest

	GetClusterCredentialsRequest(*redshift.GetClusterCredentialsInput) redshift.GetClusterCredentialsRequest

	GetReservedNodeExchangeOfferingsRequest(*redshift.GetReservedNodeExchangeOfferingsInput) redshift.GetReservedNodeExchangeOfferingsRequest

	ModifyClusterRequest(*redshift.ModifyClusterInput) redshift.ModifyClusterRequest

	ModifyClusterDbRevisionRequest(*redshift.ModifyClusterDbRevisionInput) redshift.ModifyClusterDbRevisionRequest

	ModifyClusterIamRolesRequest(*redshift.ModifyClusterIamRolesInput) redshift.ModifyClusterIamRolesRequest

	ModifyClusterParameterGroupRequest(*redshift.ModifyClusterParameterGroupInput) redshift.ModifyClusterParameterGroupRequest

	ModifyClusterSubnetGroupRequest(*redshift.ModifyClusterSubnetGroupInput) redshift.ModifyClusterSubnetGroupRequest

	ModifyEventSubscriptionRequest(*redshift.ModifyEventSubscriptionInput) redshift.ModifyEventSubscriptionRequest

	ModifySnapshotCopyRetentionPeriodRequest(*redshift.ModifySnapshotCopyRetentionPeriodInput) redshift.ModifySnapshotCopyRetentionPeriodRequest

	PurchaseReservedNodeOfferingRequest(*redshift.PurchaseReservedNodeOfferingInput) redshift.PurchaseReservedNodeOfferingRequest

	RebootClusterRequest(*redshift.RebootClusterInput) redshift.RebootClusterRequest

	ResetClusterParameterGroupRequest(*redshift.ResetClusterParameterGroupInput) redshift.ResetClusterParameterGroupRequest

	ResizeClusterRequest(*redshift.ResizeClusterInput) redshift.ResizeClusterRequest

	RestoreFromClusterSnapshotRequest(*redshift.RestoreFromClusterSnapshotInput) redshift.RestoreFromClusterSnapshotRequest

	RestoreTableFromClusterSnapshotRequest(*redshift.RestoreTableFromClusterSnapshotInput) redshift.RestoreTableFromClusterSnapshotRequest

	RevokeClusterSecurityGroupIngressRequest(*redshift.RevokeClusterSecurityGroupIngressInput) redshift.RevokeClusterSecurityGroupIngressRequest

	RevokeSnapshotAccessRequest(*redshift.RevokeSnapshotAccessInput) redshift.RevokeSnapshotAccessRequest

	RotateEncryptionKeyRequest(*redshift.RotateEncryptionKeyInput) redshift.RotateEncryptionKeyRequest

	WaitUntilClusterAvailable(*redshift.DescribeClustersInput) error
	WaitUntilClusterAvailableWithContext(aws.Context, *redshift.DescribeClustersInput, ...aws.WaiterOption) error

	WaitUntilClusterDeleted(*redshift.DescribeClustersInput) error
	WaitUntilClusterDeletedWithContext(aws.Context, *redshift.DescribeClustersInput, ...aws.WaiterOption) error

	WaitUntilClusterRestored(*redshift.DescribeClustersInput) error
	WaitUntilClusterRestoredWithContext(aws.Context, *redshift.DescribeClustersInput, ...aws.WaiterOption) error

	WaitUntilSnapshotAvailable(*redshift.DescribeClusterSnapshotsInput) error
	WaitUntilSnapshotAvailableWithContext(aws.Context, *redshift.DescribeClusterSnapshotsInput, ...aws.WaiterOption) error
}

var _ RedshiftAPI = (*redshift.Redshift)(nil)
