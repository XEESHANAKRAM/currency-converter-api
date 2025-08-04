#!/bin/bash
# user-data.sh - Script to set up the server

set -e

# Update system
yum update -y

# Install Docker
yum install -y docker
systemctl start docker
systemctl enable docker
usermod -a -G docker ec2-user

# Install Docker Compose
curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose

# Install Git
yum install -y git

# Install Go
wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> /etc/profile
export PATH=$PATH:/usr/local/go/bin

# Install CloudWatch agent
wget https://s3.amazonaws.com/amazoncloudwatch-agent/amazon_linux/amd64/latest/amazon-cloudwatch-agent.rpm
rpm -U ./amazon-cloudwatch-agent.rpm

# Create application directory
mkdir -p /opt/currency-converter
cd /opt/currency-converter

# Clone application code
git clone ${github_repo} .

# Create environment file
cat << EOF > .env
PORT=8080
EXCHANGE_API_KEY=${exchange_api_key}
ENVIRONMENT=production
EOF

# Build and run application with Docker
docker build -t currency-converter .
docker run -d \
  --name currency-converter \
  -p 80:8080 \
  --env-file .env \
  --restart unless-stopped \
  currency-converter

# Configure CloudWatch agent
cat << EOF > /opt/aws/amazon-cloudwatch-agent/etc/amazon-cloudwatch-agent.json
{
  "logs": {
    "logs_collected": {
      "files": {
        "collect_list": [
          {
            "file_path": "/var/log/messages",
            "log_group_name": "currency-converter-system",
            "log_stream_name": "{instance_id}/system"
          }
          