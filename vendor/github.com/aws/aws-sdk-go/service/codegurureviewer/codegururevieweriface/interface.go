// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package codegururevieweriface provides an interface to enable mocking the Amazon CodeGuru Reviewer service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package codegururevieweriface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/codegurureviewer"
)

// CodeGuruReviewerAPI provides an interface to enable mocking the
// codegurureviewer.CodeGuruReviewer service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon CodeGuru Reviewer.
//    func myFunc(svc codegururevieweriface.CodeGuruReviewerAPI) bool {
//        // Make svc.AssociateRepository request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := codegurureviewer.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCodeGuruReviewerClient struct {
//        codegururevieweriface.CodeGuruReviewerAPI
//    }
//    func (m *mockCodeGuruReviewerClient) AssociateRepository(input *codegurureviewer.AssociateRepositoryInput) (*codegurureviewer.AssociateRepositoryOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCodeGuruReviewerClient{}
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
type CodeGuruReviewerAPI interface {
	AssociateRepository(*codegurureviewer.AssociateRepositoryInput) (*codegurureviewer.AssociateRepositoryOutput, error)
	AssociateRepositoryWithContext(aws.Context, *codegurureviewer.AssociateRepositoryInput, ...request.Option) (*codegurureviewer.AssociateRepositoryOutput, error)
	AssociateRepositoryRequest(*codegurureviewer.AssociateRepositoryInput) (*request.Request, *codegurureviewer.AssociateRepositoryOutput)

	DescribeRepositoryAssociation(*codegurureviewer.DescribeRepositoryAssociationInput) (*codegurureviewer.DescribeRepositoryAssociationOutput, error)
	DescribeRepositoryAssociationWithContext(aws.Context, *codegurureviewer.DescribeRepositoryAssociationInput, ...request.Option) (*codegurureviewer.DescribeRepositoryAssociationOutput, error)
	DescribeRepositoryAssociationRequest(*codegurureviewer.DescribeRepositoryAssociationInput) (*request.Request, *codegurureviewer.DescribeRepositoryAssociationOutput)

	DisassociateRepository(*codegurureviewer.DisassociateRepositoryInput) (*codegurureviewer.DisassociateRepositoryOutput, error)
	DisassociateRepositoryWithContext(aws.Context, *codegurureviewer.DisassociateRepositoryInput, ...request.Option) (*codegurureviewer.DisassociateRepositoryOutput, error)
	DisassociateRepositoryRequest(*codegurureviewer.DisassociateRepositoryInput) (*request.Request, *codegurureviewer.DisassociateRepositoryOutput)

	ListRepositoryAssociations(*codegurureviewer.ListRepositoryAssociationsInput) (*codegurureviewer.ListRepositoryAssociationsOutput, error)
	ListRepositoryAssociationsWithContext(aws.Context, *codegurureviewer.ListRepositoryAssociationsInput, ...request.Option) (*codegurureviewer.ListRepositoryAssociationsOutput, error)
	ListRepositoryAssociationsRequest(*codegurureviewer.ListRepositoryAssociationsInput) (*request.Request, *codegurureviewer.ListRepositoryAssociationsOutput)

	ListRepositoryAssociationsPages(*codegurureviewer.ListRepositoryAssociationsInput, func(*codegurureviewer.ListRepositoryAssociationsOutput, bool) bool) error
	ListRepositoryAssociationsPagesWithContext(aws.Context, *codegurureviewer.ListRepositoryAssociationsInput, func(*codegurureviewer.ListRepositoryAssociationsOutput, bool) bool, ...request.Option) error
}

var _ CodeGuruReviewerAPI = (*codegurureviewer.CodeGuruReviewer)(nil)