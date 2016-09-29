# Example

```
package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/gogosphere/awsaccountloop"
)

func main() {

	a := awsaccountloop.New()

	for accountname, ec2object := range a.Accountcreds {
		a := &ec2.DescribeInstancesInput{}

		ec2objectassigned, _ := ec2object.DescribeInstances(a)

		for _, vv := range ec2objectassigned.Reservations {
			fmt.Println(accountname, *vv.Instances[0].InstanceId)
		}
	}
}

```
