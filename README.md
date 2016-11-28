### beanform

A CLI tool that generates AWS Elasticbeanstalk terraform .tf source files based on .tfstate files outputted by `terraform import` for true reverse engineering.

#### usage

import state from an existing beanstalk app

```
$ terraform import aws_elastic_beanstalk_environment.development e-jmunxte7ww

aws_elastic_beanstalk_environment.development: Importing from ID "e-jmunxte7ww"...
aws_elastic_beanstalk_environment.development: Import complete!
  Imported aws_elastic_beanstalk_environment (ID: e-jmunxte7ww)
aws_elastic_beanstalk_environment.development: Refreshing state... (ID: e-jmunxte7ww)

Import success! The resources imported are shown above. These are
now in your Terraform state. Import does not currently generate
configuration, so you must do this next. If you do not create configuration
for the above resources, then the next `terraform plan` will mark
them for destruction.
```

then generate a main.tf source file

```
$ beanform terraform.tfstate

main.tf
```
