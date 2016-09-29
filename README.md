#  What and why
I have many accounts to deal with and  its a pain either checking each one, or putting some sort of a loop  in my code.  This makes doing an all account query to the AWS API very easy.
# Example 
You can put the code in a main.go file, and ```go get -d``` to pull the library down and  you should see a list of all instances in all your accounts.  (accounts are defined by [whatever] in your ~/.aws/credentials file.
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
		parameters := &ec2.DescribeInstancesInput{}

		ec2objectassigned, _ := ec2object.DescribeInstances(parameters)

		for _, vv := range ec2objectassigned.Reservations {
			fmt.Println(accountname, *vv.Instances[0].InstanceId)
		}
	}
}

```
