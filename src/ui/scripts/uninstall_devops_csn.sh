########
# unnstall_devops_cdn.sh
#
# This deletes the cloudformation stack that removes the cloud front, api gateway, and lambda functions
# of the DevOps StackOverflow application.
#
# You must manually use the AWS Console to first disable the CloudFront distribution and then
# Delete it first before running this script or else the delete stack operations will fail.
#
# You must also manually remove all the files in the devops-overflow-bucket and delete that
# stack as well to completely cleanup.
# 
########
aws cloudformation delete-stack --stack-name devops-overflow-cdn-stack
