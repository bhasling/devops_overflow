########
# create_bucket.sh
#
# This uses cloudformation to create an AWS S3 bucket to hold the
# static html files and the issues and responses persisted by the app.
#
# You must remove all data in the S3 bucket before deleting the
# cloud formation stack that will remove the S3 bucket.
#     aws s3 rm s3://devops-overflow-bucket --recursive
########
HtmlBucket='devops-overflow-bucket'

aws cloudformation create-stack --stack-name devops-overflow-bucket-stack --template-body file://create_bucket.yaml --parameters ParameterKey=HtmlBucket,ParameterValue=${HtmlBucket}
