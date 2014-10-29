# Goploy

Goploy is a silly utility to build AWS Elastic Beanstalk Push URLs. 

## Getting Started

  $ go get github.com/ingenieux/goploy/goploy

## Using Goploy

Oh, glad you asked. Set those variables:

  * AWS_ACCESS_KEY_ID
  * AWS_SECRET_ACCESS_KEY
  * AWS_DEFAULT_REGION
  * AWSEB_APPLICATION_NAME: AWS Elastic Beanstalk Application Name
  * AWSEB_ENVIRONMENT_NAME: AWS Elastic Beanstalk Environment Name (optional)

Then, given a commitId in your repository (hint: git ```rev-parse head```). This should suffice:

  $ COMMIT_ID=$(git rev-parse head)
  $ URL=`goploy $COMMIT_ID` git push $URL master

Then you're set!
