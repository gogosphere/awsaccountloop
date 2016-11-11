package awsaccountloop

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/iam"
)

var accounts = []string{""}
var regions = []string{"us-east-1", "us-east-2", "us-west-1", "us-west-2", "ap-south-1", "ap-northeast-2",
	"ap-southeast-1", "ap-southeast-2", "ap-northeast-1", "ec2.eu-central-1.amazonaws.com", "ec2.eu-west-1.amazonaws.com", "sa-east-1"}
var credentialsfile = os.Getenv("HOME") + "/.aws/credentials"

type RawCreds struct {
	Creds    map[string]*ec2.EC2
	Sessions map[string]*session.Session
	Iams     map[string]*iam.IAM
}

type RCInterface interface {
	CAssign()
	IAssign()
	SAssign()
}

/*
 */

func New() *RawCreds {
	rc := &RawCreds{}
	rc.CAssign()
	rc.IAssign()
	rc.SAssign()
	return rc
}

func (r *RawCreds) CAssign() {
	Accountcreds := make(map[string]*ec2.EC2)

	for _, v := range accounts {

		credentialObject := credentials.NewSharedCredentials(credentialsfile, v)
		svc := ec2.New(session.New(aws.NewConfig().WithRegion("us-east-1").WithMaxRetries(2).WithCredentials(credentialObject)))
		Accountcreds[v] = svc
	}

	r.Creds = Accountcreds
}

func (r *RawCreds) IAssign() {
	Accountcreds := make(map[string]*iam.IAM)

	for _, v := range accounts {

		credentialObject := credentials.NewSharedCredentials(credentialsfile, v)
		svc := iam.New(session.New(aws.NewConfig().WithRegion("us-east-1").WithMaxRetries(2).WithCredentials(credentialObject)))
		Accountcreds[v] = svc
	}

	r.Iams = Accountcreds
}

func (r *RawCreds) SAssign() {
	Sessioncreds := make(map[string]*session.Session)

	for _, v := range accounts {

		credentialObject := credentials.NewSharedCredentials(credentialsfile, v)
		sess := session.New(aws.NewConfig().WithRegion("us-east-1").WithMaxRetries(2).WithCredentials(credentialObject))
		Sessioncreds[v] = sess
	}

	r.Sessions = Sessioncreds
}
