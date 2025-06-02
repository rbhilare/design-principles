# Violating Open Closed Principle
# Creating WebServer without Open Closed Principle

resource "aws_instance" "webserver" {
  ami           = "ami-rohit00000"
  instance_type = "t3.micro"

  user_data = <<-EOF
              #!/bin/bash
              apt update
              apt install -y nginx
              apt install -y datadog-agent
              systemctl start nginx
              systemctl start datadog-agent
              EOF

  tags = {
    Name = "WebServer"
  }
}
