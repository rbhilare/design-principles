# With Open Closed Principle
# Creating WebServer with Open Closed Principle and extending it

module "base_webserver" {
  source        = "./modules/webserver"
  ami_id        = "ami-rohit00000"
  instance_type = "t3.micro"
  user_data     = file("scripts/install.sh")
  tags = {
    Name = "WebServer"
  }
}