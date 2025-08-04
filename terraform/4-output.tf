# outputs.tf - Output values

output "public_ip" {
  description = "Public IP address of the web server"
  value = aws_eip.web.public_ip
}

output "public_dns" {
  description = "Public DNS name of the web server"
  value = aws_instance.web.public_dns
}

output "application_url" {
  description = "URL to access the application"
  value = "http://${aws_eip.web.public_ip}"
}

output "ssh_command" {
  description = "SSH command to connect to the instance"
  value = "ssh -i ~/.ssh/id_rsa ec2-user@${aws_eip.web.public_ip}"
}

output "instance_id" {
  description = "ID of the EC2 instance"
  value = aws_instance.web.id
}