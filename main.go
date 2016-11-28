package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var tfStateFile string

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&tfStateFile, "tfstate", "t", "terraform.tfstate", ".tfstate file you want to import")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
}

var rootCmd = &cobra.Command{
	Use:   "beanform",
	Short: "beanform generates terraform .tf source files from existing aws elastic beanstalk environments",
	Long: `$ terraform import aws_elastic_beanstalk_environment.dev e-jmunxte7ww
$ beanform
`,
	Run: beanform,
}

func beanform(cmd *cobra.Command, args []string) {

	//parse terraform.tfstate file
	tfstateData, err := ioutil.ReadFile(tfStateFile)
	if err != nil {
		log.Fatal(err)
	}

	//parse the .tfstate file
	var tfState TfState
	err = json.Unmarshal([]byte(tfstateData), &tfState)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	//then copy the settings and output a main.tf

	//map tfState into BeanstalkHCL
	beanEnv := PrimaryBeanstalkEnvironment{}
	for k := range tfState.Modules[0].Resources {
		beanEnv = tfState.Modules[0].Resources[k].Primary
		break
	}

	beanstalkHCL := BeanstalkHCL{
		Source:                    tfStateFile,
		Application:               beanEnv.Attributes["application"],
		Environment:               beanEnv.Attributes["name"],
		Description:               beanEnv.Attributes["description"],
		SolutionStack:             beanEnv.Attributes["solution_stack_name"],
		VPCId:                     beanEnv.Attributes["all_settings.4058884305.value"],
		Subnets:                   beanEnv.Attributes["all_settings.814243164.value"],
		ELBScheme:                 beanEnv.Attributes["all_settings.1886990778.value"],
		ApplicationHealthCheckURL: beanEnv.Attributes["all_settings.28260639.value"],
		LoadBalancerType:          beanEnv.Attributes["all_settings.2542474489.value"],
		ServiceRole:               beanEnv.Attributes["all_settings.37040285.value"],
		MinSize:                   beanEnv.Attributes["all_settings.399529023.value"],
		MaxSize:                   beanEnv.Attributes["all_settings.693650251.value"],
		InstanceType:              beanEnv.Attributes["all_settings.3558272587.value"],
		SecurityGroups:            beanEnv.Attributes["all_settings.4232919657.value"],
		IamInstanceProfile:        beanEnv.Attributes["all_settings.2395212047.value"],
		EC2KeyName:                beanEnv.Attributes["all_settings.4236355112.value"],
	}

	//marshal beanstalkHCL

	outFile, _ := os.Create("main.tf")

	const padding = 2
	w := tabwriter.NewWriter(outFile, 0, 0, padding, ' ', tabwriter.DiscardEmptyColumns)

	hcl := `# generated from {{.Source}}
	
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
    namespace = "aws:ec2:vpc"
    name      = "VPCId"
    value     = "{{.VPCId}}"
  }
	
  setting {
    namespace = "aws:ec2:vpc"
    name      = "Subnets"
    value     = "{{.Subnets}}"
  }

  setting {
    namespace = "aws:ec2:vpc"
    name      = "ELBScheme"
    value     = "{{.ELBScheme}}"
  }

  setting {
    namespace = "aws:elasticbeanstalk:application"
    name      = "Application Healthcheck URL"
    value     = "{{.ApplicationHealthCheckURL}}"
  }

  setting {
    namespace = "aws:elasticbeanstalk:environment"
    name      = "LoadBalancerType"
    value     = "{{.LoadBalancerType}}"
  }

  setting {
    namespace = "aws:elasticbeanstalk:environment"
    name      = "ServiceRole"
    value     = "{{.ServiceRole}}"
  }

  setting {
    namespace = "aws:autoscaling:asg"
    name      = "MinSize"
    value     = {{.MinSize}}
  }

  setting {
    namespace = "aws:autoscaling:asg"
    name      = "MaxSize"
    value     = {{.MaxSize}}
  }

  setting {
    namespace = "aws:autoscaling:launchconfiguration"
    name      = "InstanceType"
    value     = "{{.InstanceType}}"
  }

  setting {
    namespace = "aws:autoscaling:launchconfiguration"
    name      = "SecurityGroups"
    value     = "{{.SecurityGroups}}"
  }

  setting {
    namespace = "aws:autoscaling:launchconfiguration"
    name      = "IamInstanceProfile"
    value     = "{{.IamInstanceProfile}}"
  }

  setting {
    namespace = "aws:autoscaling:launchconfiguration"
    name      = "EC2KeyName"
    value     = "{{.EC2KeyName}}"
  }

}	
`

	//create a formatted template
	tmpl, err := template.New("resource").Parse(hcl)

	//execute the template with the data
	err = tmpl.Execute(w, beanstalkHCL)
	if err != nil {
		log.Fatal(err)
	}
	w.Flush()
}
