# ec2.tf - EC2 instance configuration

# Key Pair for SSH access
resource "aws_key_pair" "main" {
  key_name = "currency-converter-key"
  public_key = file(var.public_key_path)
}

# User data script for Docker installation and app deployment
locals {
  user_data = base64encode(templatefile("${path.module}/user-data.sh", {
    exchange_api_key = var.exchange_api_key
    github_repo = var.github_repo
  }))
}

# EC2 Instance
resource "aws_instance" "web" {
  ami = data.aws_ami.amazon_linux.id
  instance_type = var.instance_type
  key_name = aws_key_pair.main.key_name
  vpc_security_group_ids = [aws_security_group.web.id]
  subnet_id = aws_subnet.public.id
  iam_instance_profile = aws_iam_instance_profile.ec2_profile.name

  user_data_base64 = local.user_data

  root_block_device {
    volume_type = "gp3"
    volume_size = 20
    encrypted = true
  }

  tags = {
    Name = "currency-converter-web"
    Environment = var.environment
  }
}

# Elastic IP
resource "aws_eip" "web" {
  domain = "vpc"
  instance = aws_instance.web.id

  tags = {
    Name = "currency-converter-eip"
  }
}

# Application Load Balancer (optional for production)
resource "aws_lb" "main" {
  name = "currency-converter-alb"
  internal = false
  load_balancer_type = "application"
  security_groups = [aws_security_group.web.id]
  subnets = [aws_subnet.public.id]

  enable_deletion_protection = false

  tags = {
    Name = "currency-converter-alb"
  }
}