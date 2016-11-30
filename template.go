package main

// Template is the output template
var outputTemplate = `# generated by beanform from {{.Source}}
	
provider "aws" {
}

resource "aws_elastic_beanstalk_application" "app" {
  name        = "{{.Application}}"
  description = "{{.Description}}"
}

resource "aws_elastic_beanstalk_environment" "{{.Environment}}" {
  name                = "{{.Environment}}"
  application         = "${aws_elastic_beanstalk_application.app.name}"
  solution_stack_name = "{{.SolutionStack}}"

	setting {
		namespace = "aws:autoscaling:trigger"
		name      = "EvaluationPeriods"
		value     = "{{.EvaluationPeriods}}"  
	}

	setting {
		namespace = "aws:elb:listener:80"
		name      = "ListenerProtocol"
		value     = "{{.ListenerProtocol}}"  
	}

	setting {
		namespace = "aws:elb:loadbalancer"
		name      = "LoadBalancerPortProtocol"
		value     = "{{.LoadBalancerPortProtocol}}"  
	}

	setting {
		namespace = "aws:autoscaling:asg"
		name      = "Availability Zones"
		value     = "{{.AvailabilityZones}}"  
	}

	setting {
		namespace = "aws:autoscaling:updatepolicy:rollingupdate"
		name      = "RollingUpdateEnabled"
		value     = "{{.RollingUpdateEnabled}}"  
	}

	setting {
		namespace = "aws:cloudformation:template:parameter"
		name      = "AppSource"
		value     = "{{.AppSource}}"  
	}

	setting {
		namespace = "aws:ec2:vpc"
		name      = "AssociatePublicIpAddress"
		value     = "{{.AssociatePublicIPAddress}}"  
	}

	setting {
		namespace = "aws:elasticbeanstalk:command"
		name      = "IgnoreHealthCheck"
		value     = "{{.IgnoreHealthCheck}}"  
	}

	setting {
		namespace = "aws:elb:loadbalancer"
		name      = "CrossZone"
		value     = "{{.CrossZone}}"  
	}

	setting {
		namespace = "aws:autoscaling:launchconfiguration"
		name      = "RootVolumeType"
		value     = "{{.RootVolumeType}}"  
	}

	setting {
		namespace = "aws:elb:listener:80"
		name      = "ListenerEnabled"
		value     = "{{.ListenerEnabled}}"  
	}

	setting {
		namespace = "aws:autoscaling:asg"
		name      = "Custom Availability Zones"
		value     = "{{.CustomAvailabilityZones}}"  
	}

	setting {
		namespace = "aws:elasticbeanstalk:command"
		name      = "BatchSizeType"
		value     = "{{.BatchSizeType}}"  
	}

	setting {
		namespace = "aws:elb:loadbalancer"
		name      = "LoadBalancerSSLPortProtocol"
		value     = "{{.LoadBalancerSSLPortProtocol}}"  
	}

	setting {
		namespace = "aws:elasticbeanstalk:sns:topics"
		name      = "Notification Topic ARN"
		value     = "{{.NotificationTopicARN}}"  
	}

	setting {
		namespace = "aws:autoscaling:trigger"
		name      = "LowerThreshold"
		value     = "{{.LowerThreshold}}"  
	}

	setting {
		namespace = "aws:elb:listener:80"
		name      = "SSLCertificateId"
		value     = "{{.SSLCertificateIDListener}}"  
	}

	setting {
		namespace = "aws:autoscaling:launchconfiguration"
		name      = "BlockDeviceMappings"
		value     = "{{.BlockDeviceMappings}}"  
	}

	setting {
		namespace = "aws:ec2:vpc"
		name      = "ELBScheme"
		value     = "{{.ELBScheme}}"  
	}

	setting {
		namespace = "aws:elb:healthcheck"
		name      = "UnhealthyThreshold"
		value     = "{{.UnhealthyThreshold}}"  
	}

	setting {
		namespace = "aws:elb:listener:80"
		name      = "InstanceProtocol"
		value     = "{{.InstanceProtocol}}"  
	}

	setting {
		namespace = "aws:elasticbeanstalk:command"
		name      = "BatchSize"
		value     = "{{.BatchSize}}"  
	}

	setting {
		namespace = "aws:autoscaling:updatepolicy:rollingupdate"
		name      = "MinInstancesInService"
		value     = "{{.MinInstancesInService}}"  
	}

	setting {
		namespace = "aws:elasticbeanstalk:control"
		name      = "LaunchType"
		value     = "{{.LaunchType}}"  
	}

	setting {
		namespace = "aws:elb:healthcheck"
		name      = "Interval"
		value     = "{{.Interval}}"  
	}

	setting {
		namespace = "aws:autoscaling:launchconfiguration"
		name      = "IamInstanceProfile"
		value     = "{{.IamInstanceProfile}}"  
	}

	setting {
		namespace = "aws:autoscaling:trigger"
		name      = "Period"
		value     = "{{.Period}}"  
	}

	setting {
		namespace = "aws:autoscaling:trigger"
		name      = "UpperBreachScaleIncrement"
		value     = "{{.UpperBreachScaleIncrement}}"  
	}

	setting {
		namespace = "aws:elasticbeanstalk:hostmanager"
		name      = "LogPublicationControl"
		value     = "{{.LogPublicationControl}}"  
	}

	setting {
		namespace = "aws:elasticbeanstalk:environment"
		name      = "LoadBalancerType"
		value     = "{{.LoadBalancerType}}"  
	}

	setting {
		namespace = "aws:cloudformation:template:parameter"
		name      = "InstancePort"
		value     = "{{.InstancePort}}"  
	}

	setting {
		namespace = "aws:autoscaling:trigger"
		name      = "Unit"
		value     = "{{.Unit}}"  
	}

	setting {
		namespace = "aws:elasticbeanstalk:container:dotnet:apppool"
		name      = "Enable 32-bit Applications"
		value     = "{{.Enable32bitApplications}}"  
	}

	setting {
		namespace = "aws:autoscaling:updatepolicy:rollingupdate"
		name      = "RollingUpdateType"
		value     = "{{.RollingUpdateType}}"  
	}

	setting {
		namespace = "aws:elasticbeanstalk:application"
		name      = "Application Healthcheck URL"
		value     = "{{.ApplicationHealthcheckURL}}"  
	}

	setting {
		namespace = "aws:elb:policies"
		name      = "ConnectionSettingIdleTimeout"
		value     = "{{.ConnectionSettingIdleTimeout}}"  
	}

	setting {
		namespace = "aws:elb:listener:80"
		name      = "PolicyNames"
		value     = "{{.PolicyNames}}"  
	}

	setting {
		namespace = "aws:autoscaling:launchconfiguration"
		name      = "RootVolumeSize"
		value     = "{{.RootVolumeSize}}"  
	}

	setting {
		namespace = "aws:elasticbeanstalk:environment"
		name      = "EnvironmentType"
		value     = "{{.EnvironmentType}}"  
	}

	setting {
		namespace = "aws:elb:policies"
		name      = "ConnectionDrainingTimeout"
		value     = "{{.ConnectionDrainingTimeout}}"  
	}

	setting {
		namespace = "aws:elasticbeanstalk:sns:topics"
		name      = "Notification Endpoint"
		value     = "{{.NotificationEndpoint}}"  
	}

	setting {
		namespace = "aws:ec2:vpc"
		name      = "ELBSubnets"
		value     = "{{.ELBSubnets}}"  
	}

	setting {
		namespace = "aws:elasticbeanstalk:environment"
		name      = "ExternalExtensionsS3Key"
		value     = "{{.ExternalExtensionsS3Key}}"  
	}

	setting {
		namespace = "aws:cloudformation:template:parameter"
		name      = "EnvironmentVariables"
		value     = "{{.EnvironmentVariables}}"  
	}

	setting {
		namespace = "aws:autoscaling:launchconfiguration"
		name      = "RootVolumeIOPS"
		value     = "{{.RootVolumeIOPS}}"  
	}

	setting {
		namespace = "aws:autoscaling:trigger"
		name      = "MeasureName"
		value     = "{{.MeasureName}}"  
	}

	setting {
		namespace = "aws:autoscaling:updatepolicy:rollingupdate"
		name      = "Timeout"
		value     = "{{.RollingUpdateTimeout}}"  
	}

	setting {
		namespace = "aws:elasticbeanstalk:healthreporting:system"
		name      = "HealthCheckSuccessThreshold"
		value     = "{{.HealthCheckSuccessThreshold}}"  
	}

	setting {
		namespace = "aws:autoscaling:trigger"
		name      = "Statistic"
		value     = "{{.Statistic}}"  
	}

	setting {
		namespace = "aws:autoscaling:launchconfiguration"
		name      = "MonitoringInterval"
		value     = "{{.MonitoringInterval}}"  
	}

	setting {
		namespace = "aws:elasticbeanstalk:command"
		name      = "DeploymentPolicy"
		value     = "{{.DeploymentPolicy}}"  
	}

	setting {
		namespace = "aws:autoscaling:launchconfiguration"
		name      = "InstanceType"
		value     = "{{.InstanceType}}"  
	}

	setting {
		namespace = "aws:elasticbeanstalk:sns:topics"
		name      = "Notification Topic Name"
		value     = "{{.NotificationTopicName}}"  
	}

	setting {
		namespace = "aws:elasticbeanstalk:environment"
		name      = "ExternalExtensionsS3Bucket"
		value     = "{{.ExternalExtensionsS3Bucket}}"  
	}

	setting {
		namespace = "aws:autoscaling:asg"
		name      = "Cooldown"
		value     = "{{.Cooldown}}"  
	}

	setting {
		namespace = "aws:autoscaling:launchconfiguration"
		name      = "SecurityGroups"
		value     = "{{.SecurityGroups}}"  
	}

	setting {
		namespace = "aws:elasticbeanstalk:environment"
		name      = "ServiceRole"
		value     = "{{.ServiceRole}}"  
	}

	setting {
		namespace = "aws:elasticbeanstalk:sns:topics"
		name      = "Notification Protocol"
		value     = "{{.NotificationProtocol}}"  
	}

	setting {
		namespace = "aws:autoscaling:trigger"
		name      = "LowerBreachScaleIncrement"
		value     = "{{.LowerBreachScaleIncrement}}"  
	}

	setting {
		namespace = "aws:autoscaling:launchconfiguration"
		name      = "ImageId"
		value     = "{{.ImageID}}"  
	}

	setting {
		namespace = "aws:autoscaling:trigger"
		name      = "UpperThreshold"
		value     = "{{.UpperThreshold}}"  
	}

	setting {
		namespace = "aws:elb:loadbalancer"
		name      = "SSLCertificateId"
		value     = "{{.SSLCertificateID}}"  
	}

	setting {
		namespace = "aws:elasticbeanstalk:monitoring"
		name      = "Automatically Terminate Unhealthy Instances"
		value     = "{{.AutomaticallyTerminateUnhealthyInstances}}"  
	}

	setting {
		namespace = "aws:autoscaling:asg"
		name      = "MinSize"
		value     = "{{.MinSize}}"  
	}

	setting {
		namespace = "aws:autoscaling:updatepolicy:rollingupdate"
		name      = "PauseTime"
		value     = "{{.PauseTime}}"  
	}

	setting {
		namespace = "aws:elb:listener:80"
		name      = "InstancePort"
		value     = "{{.InstancePort}}"  
	}

	setting {
		namespace = "aws:ec2:vpc"
		name      = "VPCId"
		value     = "{{.VPCId}}"  
	}

	setting {
		namespace = "aws:elb:policies"
		name      = "ConnectionDrainingEnabled"
		value     = "{{.ConnectionDrainingEnabled}}"  
	}

	setting {
		namespace = "aws:autoscaling:trigger"
		name      = "BreachDuration"
		value     = "{{.BreachDuration}}"  
	}

	setting {
		namespace = "aws:elb:healthcheck"
		name      = "Target"
		value     = "{{.Target}}"  
	}

	setting {
		namespace = "aws:elasticbeanstalk:control"
		name      = "LaunchTimeout"
		value     = "{{.LaunchTimeout}}"  
	}

	setting {
		namespace = "aws:elb:loadbalancer"
		name      = "SecurityGroups"
		value     = "{{.SecurityGroupsELB}}"  
	}

	setting {
		namespace = "aws:autoscaling:launchconfiguration"
		name      = "EC2KeyName"
		value     = "{{.EC2KeyName}}"  
	}

	setting {
		namespace = "aws:elb:healthcheck"
		name      = "Timeout"
		value     = "{{.HealthCheckTimeout}}"  
	}

	setting {
		namespace = "aws:elasticbeanstalk:control"
		name      = "DefaultSSHPort"
		value     = "{{.DefaultSSHPort}}"  
	}

	setting {
		namespace = "aws:autoscaling:updatepolicy:rollingupdate"
		name      = "MaxBatchSize"
		value     = "{{.MaxBatchSize}}"  
	}

	setting {
		namespace = "aws:autoscaling:launchconfiguration"
		name      = "SSHSourceRestriction"
		value     = "{{.SSHSourceRestriction}}"  
	}

	setting {
		namespace = "aws:elb:healthcheck"
		name      = "HealthyThreshold"
		value     = "{{.HealthyThreshold}}"  
	}

	setting {
		namespace = "aws:elasticbeanstalk:container:dotnet:apppool"
		name      = "Target Runtime"
		value     = "{{.TargetRuntime}}"  
	}

	setting {
		namespace = "aws:autoscaling:asg"
		name      = "MaxSize"
		value     = "{{.MaxSize}}"  
	}

	setting {
		namespace = "aws:elb:loadbalancer"
		name      = "LoadBalancerHTTPSPort"
		value     = "{{.LoadBalancerHTTPSPort}}"  
	}

	setting {
		namespace = "aws:elb:loadbalancer"
		name      = "LoadBalancerHTTPPort"
		value     = "{{.LoadBalancerHTTPPort}}"  
	}

	setting {
		namespace = "aws:ec2:vpc"
		name      = "Subnets"
		value     = "{{.Subnets}}"  
	}

	setting {
		namespace = "aws:elasticbeanstalk:command"
		name      = "Timeout"
		value     = "{{.CommandTimeout}}"  
	}

	setting {
		namespace = "aws:elasticbeanstalk:control"
		name      = "RollbackLaunchOnFailure"
		value     = "{{.RollbackLaunchOnFailure}}"  
	}


}
`