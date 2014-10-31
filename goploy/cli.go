package main

const usage = `goploy.

Builds a git endpoint for AWS Elastic Beanstalk Deployment and optionally 
pushes to it.

(Required) Environment Variables:
  AWS_ACCESS_KEY_ID      AWS Access Key
  AWS_SECRET_ACCESS_KEY  AWS Secret Access Key
  AWS_DEFAULT_REGION     AWS Region to Use

Usage:
  goploy [options] genurl APPLICATION [ENVIRONMENT]
  goploy [options] push APPLICATION [ENVIRONMENT]
  goploy -h | --help

Options:
  -h, --help               Show this screen.
  -c, --commitId COMMITID  Commit Id
  -d, --directory DIR      Directory where .git is. [default: .]
  -b, --branch BRANCH      Branch to use. [default: HEAD]
`
