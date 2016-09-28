package awsaccountloop

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var (
	credentialsPath = os.Getenv("HOME") + "/.aws/credentials"
	accountcreds    = make(map[*ec2.EC2]string)
	r               = readFile(credentialsPath)
	p               = pullAccounts(r)
)

// AWSAccount holding the account number and the svc token here
type AWSAccount struct {
	accountcreds    map[string]*ec2.EC2
	accountnames    []string
	credentialsPath string
}

func awsaccountloop() {
	var awsloop = &AWSAccount{
		credentialsPath: credentialsPath,
		accountnames:    p,
		accountcreds:    assignToken(),
	}
	fmt.Println(awsloop.accountcreds)
}

func readFile(cf string) []string {
	var f []string
	file, err := os.Open(cf)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f = append(f, scanner.Text())
	}
	return f
}
func pullAccounts(pa []string) []string {
	r := regexp.MustCompile("[a-z]+-[a-z]+_.*admin")
	var a []string
	for _, v := range pa {
		match := r.FindString(v)
		if len(match) != 0 {
			a = append(a, match)
		}
	}
	return a
}
func assignToken() map[string]*ec2.EC2 {
	accountcreds := make(map[string]*ec2.EC2)
	for _, v := range p {
		credentialObject := credentials.NewSharedCredentials(credentialsPath, v)
		svc := ec2.New(session.New(aws.NewConfig().WithRegion("us-east-1").
			WithMaxRetries(2).WithCredentials(credentialObject)))
		accountcreds[v] = svc
	}
	return accountcreds
}

func er(e error) {
	if e != nil {
		log.Printf("Error: %v\n", e)
	}
}
