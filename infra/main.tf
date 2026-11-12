# LABORATÓRIO DE SAST: configuração insegura de propósito. Nunca aplicar.

provider "aws" {
  region = "sa-east-1"
}

# T1: bucket sem bloqueio de acesso público
resource "aws_s3_bucket" "lab" {
  bucket = "lab-sast-exemplo-nao-aplicar"
}

resource "aws_s3_bucket_public_access_block" "lab" {
  bucket                  = aws_s3_bucket.lab.id
  block_public_acls       = false
  block_public_policy     = false
  ignore_public_acls      = false
  restrict_public_buckets = false
}

# T2: SSH aberto para a internet
resource "aws_security_group" "ssh_aberto" {
  name        = "lab-ssh-aberto"
  description = "Laboratorio SAST"

  ingress {
    description = "SSH"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# T3, T4, T5: banco sem criptografia, público e com senha no código
resource "aws_db_instance" "lab" {
  identifier          = "lab-sast"
  engine              = "postgres"
  instance_class      = "db.t3.micro"
  allocated_storage   = 20
  username            = "admin_lab"
  password            = "senha-fixa-123"
  storage_encrypted   = false
  publicly_accessible = true
  skip_final_snapshot = true
}
