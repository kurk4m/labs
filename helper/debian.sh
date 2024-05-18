wget https://go.dev/dl/go1.22.3.linux-amd64.tar.gz
tar -zxvf go1.22.3.linux-amd64.tar.gz
mv go /usr/local
export PATH=$PATH:/usr/local/go/bin

sudo apt update
sudo apt install openssh-server
sudo systemctl start ssh
sudo systemctl enable ssh
sudo systemctl status ssh

sudo apt-get install git -y


git config --global user.name "name"
git config --global user.email "email"
git config --global --list