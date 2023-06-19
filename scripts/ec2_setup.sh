ssh -i ~/.ssh/laptop.pem ec2-user@ec2-13-49-73-201.eu-north-1.compute.amazonaws.com

sudo yum install docker -y
sudo usermod -a -G docker ec2-user
id ec2-user
newgrp docker
sudo systemctl enable docker.service
sudo systemctl start docker.service



### Setup docker login
aws ecr get-login-password --region eu-north-1 | sudo docker login --username AWS --password-stdin 530826328503.dkr.ecr.eu-north-1.amazonaws.com

