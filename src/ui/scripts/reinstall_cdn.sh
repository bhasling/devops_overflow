
########
# reinstall_devops_cdn.sh
#
# This copies the latest UI code to the S3 bucket to re-install the latest UI for lambda.
# 
########

HtmlBucket='devops-overflow-bucket'
Region='us-east-1'

# Copy the static HTML of the UI to the S3 Bucket for use by CloudFront
yarn run build
yarn run export
aws s3 cp ../out/ s3://$HtmlBucket/ --recursive

