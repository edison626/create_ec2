package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"), // 替换为您的AWS区域
	})
	if err != nil {
		fmt.Println("创建会话失败:", err)
		return
	}

	svc := ec2.New(sess)

	runResult, err := svc.RunInstances(&ec2.RunInstancesInput{
		ImageId:      aws.String("ami-01da42fa32830f2d0"),
		InstanceType: aws.String("t3.small"),
		KeyName:      aws.String("ec2-user"),
		MinCount:     aws.Int64(1),
		MaxCount:     aws.Int64(1), // 只创建一台实例
		SecurityGroupIds: []*string{
			aws.String("sg-033a6552e3ffe1a48"),
		},
		SubnetId: aws.String("vpc-0cadb665c480c21d1"), // 替换为您的子网ID
		BlockDeviceMappings: []*ec2.BlockDeviceMapping{
			{
				DeviceName: aws.String("/dev/sdh"),
				Ebs: &ec2.EbsBlockDevice{
					VolumeSize: aws.Int64(100), // 100GB存储
					VolumeType: aws.String("gp2"),
				},
			},
		},
		TagSpecifications: []*ec2.TagSpecification{
			{
				ResourceType: aws.String("instance"),
				Tags: []*ec2.Tag{
					{
						Key:   aws.String("Name"),
						Value: aws.String("MySingleInstance"), // 指定实例名称
					},
				},
			},
		},
	})

	if err != nil {
		fmt.Println("无法创建实例:", err)
		return
	}

	fmt.Println("已成功创建实例:", runResult.Instances)
}
