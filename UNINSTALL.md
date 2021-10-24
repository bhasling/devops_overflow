# Uninstalling
To uninstall from the AWS cloud do the following.

1. Use AWS Console to Open CloudFront services. Disable the distribution with the origin 'devops-overflow-bucket.s3.amazonaws.com'. It takes about 5 minutes for the distribution to be 
disabled. Be patient.

2. After the distribution is disabled you can delete the distribution.

3. You must also remove the 'Origin access identity' using the Security Tab in the CloudFront AWS Services with the AWS console.

4. Using the AWS Console CloudFormation delete the 'devops-overflow-cdn-stack'. This will remove everything except the S3 bucket.

5. Remove all the files in the S3 bucket using the command line:

    aws s3 rm s3://devops-overflow-bucket --recursive

6. Using the AWS Console CloudFormation delete the 'devops-overflow-bucket-stack'. This will remove the S3 bucket and the cloud formation stack.