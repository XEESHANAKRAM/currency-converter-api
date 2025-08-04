# variables.tf - Input variables

variable "aws_region" {
  description = "AWS region for resources"
  type = string
  default = "us-east-1"
}

variable "instance_type" {
  description = "EC2 instance type"
  type = string
  default = "t3.micro"
}

variable "environment" {
  description = "Environment name"
  type = string
  default = "production"
}

variable "public_key_path" {
  description = "Path to public SSH key"
  type = string
  default = "~/.ssh/id_rsa.pub"
}

variable "allowed_cidr" {
  description = "CIDR block allowed for SSH access"
  type = string
  default = "0.0.0.0/0"
}

variable "exchange_api_key" {
  description = "API key for exchange rate service"
  type = string
  default = ""
}

variable "github_repo" {
  description = "GitHub repository URL"
  type = string
  default = "https://github.com/xeeshanakram/currency-converter.git"
}

variable "aws_access_key" {
  description = "AWS access key"
  type = string
  default = ""
}

variable "aws_secret_key" {
  description = "AWS secret key"
  type = string
  default =""
}