package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type EC2Config struct {
	ImageId         string
	InstanceType    string
	KeyName         string
	SecurityGroupID string
	SubnetID        string
	TagValue        string
	VolumeSize      int64
}

// 创建一个EC2Config切片包含不同的实例配置
func configEC2Instances(batch string) []EC2Config {
	return []EC2Config{
		{
			ImageId:      "ami-01da42fa32830f2d0",
			InstanceType: "t3.medium",
			TagValue:     batch + "vn-prod-web-proxy01",
			VolumeSize:   100,
		},
		{
			ImageId:      "ami-01da42fa32830f2d0",
			InstanceType: "t3.small",
			TagValue:     batch + "prod-web-proxy02",
			VolumeSize:   100,
		},
		{
			ImageId:      "ami-01da42fa32830f2d0",
			InstanceType: "t3.small",
			TagValue:     batch + "vn-prod-callback",
			VolumeSize:   100,
		},
		{
			ImageId:      "ami-01da42fa32830f2d0",
			InstanceType: "t3.small",
			TagValue:     batch + "vn-prod-houtai",
			VolumeSize:   100,
		},
		{
			ImageId:      "ami-0e8849aa060c28662",
			InstanceType: "c5.xlarge",
			TagValue:     batch + "vn-prod-mongodb-01",
			VolumeSize:   500,
		},
		{
			ImageId:      "ami-0e8849aa060c28662",
			InstanceType: "c5.xlarge",
			TagValue:     batch + "vn-prod-mongodb-02",
			VolumeSize:   500,
		},
		{
			ImageId:      "ami-0e8849aa060c28662",
			InstanceType: "c5.xlarge",
			TagValue:     batch + "vn-prod-mongodb-03",
			VolumeSize:   500,
		},
		{
			ImageId:      "ami-0e8849aa060c28662",
			InstanceType: "t3.xlarge",
			TagValue:     batch + "vn-prod-cgcron-clinet-task01",
			VolumeSize:   300,
		},
		{
			ImageId:      "ami-0e8849aa060c28662",
			InstanceType: "c5.xlarge",
			TagValue:     batch + "vn-prod-cgcron-clinet-task02",
			VolumeSize:   300,
		},
		{
			ImageId:      "ami-0e8849aa060c28662",
			InstanceType: "c5.xlarge",
			TagValue:     batch + "prod-redis-etcd01",
			VolumeSize:   150,
		},
		{
			ImageId:      "ami-0e8849aa060c28662",
			InstanceType: "c5.xlarge",
			TagValue:     batch + "prod-redis-etcd02",
			VolumeSize:   150,
		},
		{
			ImageId:      "ami-0e8849aa060c28662",
			InstanceType: "c5.xlarge",
			TagValue:     batch + "prod-redis-etcd03",
			VolumeSize:   150,
		},
		{
			ImageId:      "ami-0e8849aa060c28662",
			InstanceType: "c5.2xlarge",
			TagValue:     batch + "starrocks-be-01",
			VolumeSize:   300,
		},
		{
			ImageId:      "ami-0e8849aa060c28662",
			InstanceType: "c5.2xlarge",
			TagValue:     batch + "starrocks-be-02",
			VolumeSize:   300,
		},
		{
			ImageId:      "ami-0e8849aa060c28662",
			InstanceType: "c5.2xlarge",
			TagValue:     batch + "starrocks-be-03",
			VolumeSize:   300,
		},
		{
			ImageId:      "ami-0e8849aa060c28662",
			InstanceType: "c5.xlarge",
			TagValue:     batch + "starrocks-fe-01",
			VolumeSize:   100,
		},
		{
			ImageId:      "ami-0e8849aa060c28662",
			InstanceType: "c5.xlarge",
			TagValue:     batch + "starrocks-fe-02",
			VolumeSize:   100,
		},
		{
			ImageId:      "ami-0e8849aa060c28662",
			InstanceType: "c5.xlarge",
			TagValue:     batch + "starrocks-fe-03",
			VolumeSize:   100,
		},
		{
			ImageId:      "ami-0e8849aa060c28662",
			InstanceType: "c5.xlarge",
			TagValue:     batch + "wukong-yq-a06",
			VolumeSize:   200,
		},
		{
			ImageId:      "ami-0e8849aa060c28662",
			InstanceType: "c5.xlarge",
			TagValue:     batch + "wukong-yq-a07",
			VolumeSize:   200,
		},
		{
			ImageId:      "ami-0e8849aa060c28662",
			InstanceType: "c5.xlarge",
			TagValue:     batch + "wukong-yq-a08",
			VolumeSize:   200,
		},
		{
			ImageId:      "ami-0e8849aa060c28662",
			InstanceType: "c5.2xlarge",
			TagValue:     batch + "yq-prod-cloudcanal-01",
			VolumeSize:   500,
		},
		{
			ImageId:      "ami-0e8849aa060c28662",
			InstanceType: "c5.2xlarge",
			TagValue:     batch + "yq-prod-cloudcanal-02",
			VolumeSize:   500,
		},
		{
			ImageId:      "ami-0e8849aa060c28662",
			InstanceType: "c5.2xlarge",
			TagValue:     batch + "yq-prod-cloudcanal-03",
			VolumeSize:   500,
		},
		{
			ImageId:      "ami-0e8849aa060c28662",
			InstanceType: "c5.xlarge",
			TagValue:     batch + "yq-rocketmq-prod1",
			VolumeSize:   100,
		},
		{
			ImageId:      "ami-0e8849aa060c28662",
			InstanceType: "c5.xlarge",
			TagValue:     batch + "yq-rocketmq-prod2",
			VolumeSize:   100,
		},
		{
			ImageId:      "ami-0e8849aa060c28662",
			InstanceType: "c5.xlarge",
			TagValue:     batch + "vn-prod-zinc-beanstalkd",
			VolumeSize:   100,
		},
		{
			ImageId:      "ami-0e8849aa060c28662",
			InstanceType: "c5.xlarge",
			TagValue:     batch + "vn-prod-rmqtt01",
			VolumeSize:   100,
		},
		{
			ImageId:      "ami-0e8849aa060c28662",
			InstanceType: "c5.xlarge",
			TagValue:     batch + "vn-prod-rmqtt02",
			VolumeSize:   100,
		},
	}
}

func main() {

	//创建AWS会话和EC2服务客户端代码
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-east-1"), // 替换为您的AWS区域
	})
	if err != nil {
		fmt.Println("创建会话失败:", err)
		return
	}

	svc := ec2.New(sess)

	//服务器名 - 前缀
	batch := "b3-"

	//实例内容
	configs := configEC2Instances(batch)

	// ... 创建实例的代码
	for _, config := range configs {
		runResult, err := svc.RunInstances(&ec2.RunInstancesInput{
			ImageId:      aws.String(config.ImageId),
			InstanceType: aws.String(config.InstanceType),
			KeyName:      aws.String("ec2-user"),
			MinCount:     aws.Int64(1),
			MaxCount:     aws.Int64(1), // 只创建一台实例
			SecurityGroupIds: []*string{
				aws.String("sg-033a6552e3ffe1a48"),
			},
			SubnetId: aws.String("subnet-0a7e140afbc1f8f9b"), // 替换为您的子网ID
			BlockDeviceMappings: []*ec2.BlockDeviceMapping{
				{
					DeviceName: aws.String("/dev/sdh"),
					Ebs: &ec2.EbsBlockDevice{
						VolumeSize: aws.Int64(config.VolumeSize), // 100GB存储
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
							Value: aws.String(config.TagValue), // 指定实例名称
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

		instanceId := runResult.Instances[0].InstanceId

		// 等待实例变为running状态
		fmt.Println("等待实例启动...")
		for {
			descInstances, err := svc.DescribeInstances(&ec2.DescribeInstancesInput{
				InstanceIds: []*string{instanceId},
			})
			if err != nil {
				fmt.Println("无法获取实例状态:", err)
				return
			}

			state := descInstances.Reservations[0].Instances[0].State.Name
			if *state == "running" {
				break
			}

			time.Sleep(10 * time.Second)
		}
		fmt.Println("实例已启动,正在分配弹性IP...")

		// 申请弹性IP
		allocRes, err := svc.AllocateAddress(&ec2.AllocateAddressInput{
			Domain: aws.String("vpc-0cadb665c480c21d1"), // VPC网络
		})
		if err != nil {
			fmt.Println("无法分配弹性IP:", err)
			return
		}

		// 关联弹性IP到实例
		_, err = svc.AssociateAddress(&ec2.AssociateAddressInput{
			InstanceId:   instanceId,
			AllocationId: allocRes.AllocationId,
		})
		if err != nil {
			fmt.Println("无法关联弹性IP:", err)
			return
		}

		fmt.Println("弹性IP已成功关联到实例:", *instanceId)
	}
}
