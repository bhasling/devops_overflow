
########
# install_devops_cdn.sh
#
# This script uses cloud formation to create the AWS objects to deploy DevOps Overflow to AWS.
# This creates a CloudFront distribution as a Content Delivery Network to view NextJS static HTML
# and also an API Gateway viewing a lambda function containing the GO web server.
# This script copies the static HTML UI to an S3 bucket and copies the GO program to a .zip file
# into the S3 bucket to be installed.
# 
########

HtmlBucket='devops-overflow-bucket'
Region='us-east-1'

# Copy the static HTML of the UI to the S3 Bucket for use by CloudFront
aws s3 cp ../out/ s3://$HtmlBucket/ --recursive

# Copy the GO server source code to a build folder and build it and create zip file
rm -rf ../unixlambdabuild
mkdir ../unixlambdabuild
cp -r ../../server/* ../unixlambdabuild
pushd ../unixlambdabuild

# Write GO server configuration file with configuration information so this goes into build
echo "S3BucketName: $HtmlBucket" > config.yaml
echo "Region: $Region" >> config.yaml
echo "StaticFolder: '.'" >> config.yaml

# Build the GO main executable
rm -f main
go build .

# Create a zip file with the GO executable
rm -f lambda_function.zip
zip lambda_function.zip main config.yaml
cd ../out
zip -ru ../unixlambdabuild/lambda_function.zip .
cd ../unixlambdabuild

# Copy the zip file to the S3 bucket for use by cloudformation to install
aws s3 cp lambda_function.zip s3://$HtmlBucket/lambda_function.zip

popd

# Run the CloudFormation stack to create the cloudfront distribution object
aws cloudformation create-stack --stack-name devops-overflow-cdn-stack --template-body file://install_devops_cdn.yaml --capabilities CAPABILITY_IAM  --parameters ParameterKey=HtmlBucket,ParameterValue=${HtmlBucket}
