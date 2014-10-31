# Goploy

Goploy is a silly utility to build AWS Elastic Beanstalk Push URLs and optionally push it.

## Getting Started

  $ go get github.com/ingenieux/goploy/goploy

## Using Goploy

Oh, glad you asked. Set those variables:

  * AWS_ACCESS_KEY_ID
  * AWS_SECRET_ACCESS_KEY
  * AWS_DEFAULT_REGION: Optional

Then simply run:

  $ goploy push myapplication [myenvname]

Then you're set!

Run `goploy -h` to get its usage options.

## Support

No support is provided. I [do accept donations](http://beanstalker.ingenieux.com.br/donate.html), though :]

## Related

  * [Beanstalker](http://beanstalker.ingenieux.com.br/beanstalker-maven-plugin/)
  * [AWSEB Deployment Plugin](https://wiki.jenkins-ci.org/display/JENKINS/AWSEB+Deployment+Plugin)