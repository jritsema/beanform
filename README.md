### beanform

Reverse engineer an AWS Elasticbeanstalk app into Terraform source code

This is a CLI tool (work in progress) that generates AWS Elasticbeanstalk terraform .tf source files based on .tfstate files outputted by `terraform import`.


#### usage

Import state from an existing beanstalk app environment.

```
$ terraform import aws_elastic_beanstalk_environment.my-app e-jmunxte7ww

aws_elastic_beanstalk_environment.my-app: Importing from ID "e-jmunxte7ww"...
aws_elastic_beanstalk_environment.my-app: Import complete!
  Imported aws_elastic_beanstalk_environment (ID: e-jmunxte7ww)
aws_elastic_beanstalk_environment.my-app: Refreshing state... (ID: e-jmunxte7ww)

Import success! The resources imported are shown above. These are
now in your Terraform state. Import does not currently generate
configuration, so you must do this next. If you do not create configuration
for the above resources, then the next `terraform plan` will mark
them for destruction.
```

Then generate a main.tf source file

```
$ beanform

main.tf
```

or 

```
$ beanform -t foo.tfstate

main.tf
```
