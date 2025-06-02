resource "aws_instance" "web" {
  ami           = var.ami_id
  instance_type = var.instance_type
  user_data     = var.user_data

  tags = var.tags
}