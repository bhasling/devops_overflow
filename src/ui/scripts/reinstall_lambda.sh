HtmlBucket='devops-overflow-bucket'
LambdaFunctionName='devops-overflow-lambda'
Region='us-east-1'

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
aws lambda update-function-code --function-name $LambdaFunctionName --s3-bucket $HtmlBucket --s3-key lambda_function.zip

