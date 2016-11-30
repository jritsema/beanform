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

	//parse .tfstate file
	tfstateData, err := ioutil.ReadFile(tfStateFile)
	if err != nil {
		log.Fatal(err)
	}

	var tfState TfState
	err = json.Unmarshal([]byte(tfstateData), &tfState)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	//then copy the settings and output a main.tf
	beanEnv := PrimaryBeanstalkEnvironment{}
	for k := range tfState.Modules[0].Resources {
		beanEnv = tfState.Modules[0].Resources[k].Primary
		break
	}

	//map tfState into BeanstalkHCL
	beanstalkHCL := mapData(beanEnv)

	//marshal beanstalkHCL
	outFile, _ := os.Create("main.tf")
	const padding = 2
	w := tabwriter.NewWriter(outFile, 0, 0, padding, ' ', tabwriter.DiscardEmptyColumns)

	//create a formatted template
	tmpl, err := template.New("out").Parse(outputTemplate)

	//execute the template with the data
	err = tmpl.Execute(w, beanstalkHCL)
	if err != nil {
		log.Fatal(err)
	}
	w.Flush()

	fmt.Println("main.tf")
}

func mapData(beanEnv PrimaryBeanstalkEnvironment) BeanstalkHCL {
	return BeanstalkHCL{
		Source:                      tfStateFile,
		Application:                 beanEnv.Attributes["application"],
		Environment:                 beanEnv.Attributes["name"],
		Description:                 beanEnv.Attributes["description"],
		SolutionStack:               beanEnv.Attributes["solution_stack_name"],
		EvaluationPeriods:           beanEnv.Attributes["all_settings.1036703857.value"],
		ListenerProtocol:            beanEnv.Attributes["all_settings.1093784765.value"],
		LoadBalancerPortProtocol:    beanEnv.Attributes["all_settings.1105127146.value"],
		AvailabilityZones:           beanEnv.Attributes["all_settings.1225551841.value"],
		RollingUpdateEnabled:        beanEnv.Attributes["all_settings.1311926541.value"],
		AppSource:                   beanEnv.Attributes["all_settings.1374365602.value"],
		AssociatePublicIPAddress:    beanEnv.Attributes["all_settings.1429395430.value"],
		IgnoreHealthCheck:           beanEnv.Attributes["all_settings.1487137936.value"],
		CrossZone:                   beanEnv.Attributes["all_settings.1497664210.value"],
		RootVolumeType:              beanEnv.Attributes["all_settings.1499800684.value"],
		ListenerEnabled:             beanEnv.Attributes["all_settings.1539976252.value"],
		CustomAvailabilityZones:     beanEnv.Attributes["all_settings.1584820072.value"],
		BatchSizeType:               beanEnv.Attributes["all_settings.1610162273.value"],
		LoadBalancerSSLPortProtocol: beanEnv.Attributes["all_settings.1709622552.value"],
		NotificationTopicARN:        beanEnv.Attributes["all_settings.1818591148.value"],
		LowerThreshold:              beanEnv.Attributes["all_settings.1830179792.value"],
		SSLCertificateIDListener:    beanEnv.Attributes["all_settings.1835511052.value"],
		BlockDeviceMappings:         beanEnv.Attributes["all_settings.1839941942.value"],
		ELBScheme:                   beanEnv.Attributes["all_settings.1886990778.value"],
		UnhealthyThreshold:          beanEnv.Attributes["all_settings.1932615024.value"],
		InstanceProtocol:            beanEnv.Attributes["all_settings.2145095834.value"],
		BatchSize:                   beanEnv.Attributes["all_settings.2175068570.value"],
		MinInstancesInService:       beanEnv.Attributes["all_settings.2207438838.value"],
		LaunchType:                  beanEnv.Attributes["all_settings.2210680385.value"],
		Interval:                    beanEnv.Attributes["all_settings.2394718604.value"],
		IamInstanceProfile:          beanEnv.Attributes["all_settings.2395212047.value"],
		Period:                      beanEnv.Attributes["all_settings.2402433859.value"],
		UpperBreachScaleIncrement:  beanEnv.Attributes["all_settings.240742627.value"],
		LogPublicationControl:      beanEnv.Attributes["all_settings.2492456712.value"],
		LoadBalancerType:           beanEnv.Attributes["all_settings.2542474489.value"],
		InstancePortCloudFormation: beanEnv.Attributes["all_settings.2578083670.value"],
		Unit: beanEnv.Attributes["all_settings.2740395520.value"],
		Enable32bitApplications:                  beanEnv.Attributes["all_settings.2810849049.value"],
		RollingUpdateType:                        beanEnv.Attributes["all_settings.2824564098.value"],
		ApplicationHealthcheckURL:                beanEnv.Attributes["all_settings.28260639.value"],
		ConnectionSettingIdleTimeout:             beanEnv.Attributes["all_settings.2827474164.value"],
		PolicyNames:                              beanEnv.Attributes["all_settings.285213017.value"],
		RootVolumeSize:                           beanEnv.Attributes["all_settings.2883801498.value"],
		EnvironmentType:                          beanEnv.Attributes["all_settings.2929485696.value"],
		ConnectionDrainingTimeout:                beanEnv.Attributes["all_settings.304444642.value"],
		NotificationEndpoint:                     beanEnv.Attributes["all_settings.306296839.value"],
		ELBSubnets:                               beanEnv.Attributes["all_settings.312884338.value"],
		ExternalExtensionsS3Key:                  beanEnv.Attributes["all_settings.3243179558.value"],
		EnvironmentVariables:                     beanEnv.Attributes["all_settings.325585145.value"],
		RootVolumeIOPS:                           beanEnv.Attributes["all_settings.3263411188.value"],
		MeasureName:                              beanEnv.Attributes["all_settings.3341012745.value"],
		RollingUpdateTimeout:                     beanEnv.Attributes["all_settings.3354438189.value"],
		HealthCheckSuccessThreshold:              beanEnv.Attributes["all_settings.3396910454.value"],
		Statistic:                                beanEnv.Attributes["all_settings.3402994671.value"],
		MonitoringInterval:                       beanEnv.Attributes["all_settings.3438018982.value"],
		DeploymentPolicy:                         beanEnv.Attributes["all_settings.3553402416.value"],
		InstanceType:                             beanEnv.Attributes["all_settings.3558272587.value"],
		NotificationTopicName:                    beanEnv.Attributes["all_settings.3584475532.value"],
		ExternalExtensionsS3Bucket:               beanEnv.Attributes["all_settings.3615826456.value"],
		Cooldown:                                 beanEnv.Attributes["all_settings.3648591417.value"],
		SecurityGroups:                           beanEnv.Attributes["all_settings.3665611509.value"],
		ServiceRole:                              beanEnv.Attributes["all_settings.37040285.value"],
		NotificationProtocol:                     beanEnv.Attributes["all_settings.3705073300.value"],
		LowerBreachScaleIncrement:                beanEnv.Attributes["all_settings.3708017666.value"],
		ImageID:                                  beanEnv.Attributes["all_settings.3755611818.value"],
		UpperThreshold:                           beanEnv.Attributes["all_settings.3896318158.value"],
		SSLCertificateID:                         beanEnv.Attributes["all_settings.3961867433.value"],
		AutomaticallyTerminateUnhealthyInstances: beanEnv.Attributes["all_settings.3981805791.value"],
		MinSize:      beanEnv.Attributes["all_settings.399529023.value"],
		PauseTime:    beanEnv.Attributes["all_settings.4020160483.value"],
		InstancePort: beanEnv.Attributes["all_settings.4020527343.value"],
		VPCId:        beanEnv.Attributes["all_settings.4058884305.value"],
		ConnectionDrainingEnabled: beanEnv.Attributes["all_settings.4065641051.value"],
		BreachDuration:            beanEnv.Attributes["all_settings.4100247781.value"],
		Target:                    beanEnv.Attributes["all_settings.4107770397.value"],
		LaunchTimeout:             beanEnv.Attributes["all_settings.4167503353.value"],
		SecurityGroupsELB:         beanEnv.Attributes["all_settings.4232919657.value"],
		EC2KeyName:                beanEnv.Attributes["all_settings.4236355112.value"],
		HealthCheckTimeout:        beanEnv.Attributes["all_settings.4267362780.value"],
		DefaultSSHPort:            beanEnv.Attributes["all_settings.433110370.value"],
		MaxBatchSize:              beanEnv.Attributes["all_settings.493815232.value"],
		SSHSourceRestriction:      beanEnv.Attributes["all_settings.502734328.value"],
		HealthyThreshold:          beanEnv.Attributes["all_settings.59521051.value"],
		TargetRuntime:             beanEnv.Attributes["all_settings.642110545.value"],
		MaxSize:                   beanEnv.Attributes["all_settings.693650251.value"],
		LoadBalancerHTTPSPort:     beanEnv.Attributes["all_settings.731293825.value"],
		LoadBalancerHTTPPort:      beanEnv.Attributes["all_settings.784312882.value"],
		Subnets:                   beanEnv.Attributes["all_settings.814243164.value"],
		CommandTimeout:            beanEnv.Attributes["all_settings.881749867.value"],
		RollbackLaunchOnFailure:   beanEnv.Attributes["all_settings.940069638.value"],
	}

}
