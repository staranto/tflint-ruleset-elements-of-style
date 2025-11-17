resource "aws_instance" "my_instance" {
  ami           = "ami-12345678"
  instance_type = "t2.micro"
}

# tflint-ignore: eos_shout
resource "aws_instance" "MY_INSTANCE" {
  ami           = "ami-12345678"
  instance_type = "t2.micro"
}

# tflint-ignore: eos_length
resource "aws_instance" "very_long_instance_name_that_exceeds_limit" {
  ami           = "ami-12345678"
  instance_type = "t2.micro"
}

# tflint-ignore: eos_type_echo
resource "aws_s3_bucket" "my_bucket" {
  bucket = "my-bucket"
}

resource "aws_s3_bucket" "not_echoed" {
  bucket = "not-echoed"
}

data "aws_ami" "ubuntu" {
  most_recent = true
  owners      = ["amazon"]
}

data "aws_ami" "UBUNTU" {
  most_recent = true
  owners      = ["amazon"]
}

variable "instance_count" {
  description = "Number of instances"
  type        = number
  default     = 1
}

output "instance_id" {
  value = aws_instance.my_instance.id
}

output "INSTANCE_ID" {
  value = aws_instance.MY_INSTANCE.id
}

locals {
  common_tags = {
    Environment = "dev"
  }
}

locals {
  COMMON_TAGS = {
    Environment = "dev"
  }
}

check "health_check" {
  data "http" "example" {
    url = "https://httpbin.org/uuid"
  }

  assert {
    condition     = data.http.example.status_code == 200
    error_message = "HTTP request failed"
  }
}

check "HEALTH_CHECK" {
  data "http" "EXAMPLE" {
    url = "https://httpbin.org/uuid"
  }

  assert {
    condition     = data.http.EXAMPLE.status_code == 200
    error_message = "HTTP request failed"
  }
}
