# -*- mode: ruby -*-
# vi: set ft=ruby :

# CHECK OS
module OS
  def OS.windows?
      (/cygwin|mswin|mingw|bccwin|wince|emx/ =~ RUBY_PLATFORM) != nil
  end

  def OS.mac?
      (/darwin/ =~ RUBY_PLATFORM) != nil
  end

  def OS.unix?
      !OS.windows?
  end

  def OS.linux?
      OS.unix? and not OS.mac?
  end
end

if OS.windows?
  unless Vagrant.has_plugin?('vagrant-vbguest')
    raise <<~ERROR
      "vagrant-vbguest" is not installed!
      Please execute the below command to install plugins.
      - $ vagrant plugin install vagrant-vbguest
    ERROR
  end
end

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.
Vagrant.configure("2") do |config|
  # The most common configuration options are documented and commented below.
  # For a complete reference, please see the online documentation at
  # https://docs.vagrantup.com.

  # Every Vagrant development environment requires a box. You can search for
  # boxes at https://vagrantcloud.com/search.
  config.vm.box = "ubuntu/xenial64"
  # NOTE: ubuntu/xenial64 のイメージはログが出る
  config.vm.provider "virtualbox" do |vb|
    vb.customize [ "modifyvm", :id, "--uartmode1", "disconnected" ]
  end

  # Disable automatic box update checking. If you disable this, then
  # boxes will only be checked for updates when the user runs
  # `vagrant box outdated`. This is not recommended.
  # config.vm.box_check_update = false

  # Create a forwarded port mapping which allows access to a specific port
  # within the machine from a port on the host machine. In the example below,
  # accessing "localhost:8080" will access port 80 on the guest machine.
  # NOTE: This will enable public access to the opened port
  # config.vm.network "forwarded_port", guest: 80, host: 8080

  # Create a forwarded port mapping which allows access to a specific port
  # within the machine from a port on the host machine and only allow access
  # via 127.0.0.1 to disable public access
  # config.vm.network "forwarded_port", guest: 80, host: 8080, host_ip: "127.0.0.1"



  # Create a private network, which allows host-only access to the machine
  # using a specific IP.
  config.vm.network "private_network", ip: "192.168.33.11"

  # Create a public network, which generally matched to bridged network.
  # Bridged networks make the machine appear as another physical device on
  # your network.
  # config.vm.network "public_network"

  # Share an additional folder to the guest VM. The first argument is
  # the path on the host to the actual folder. The second argument is
  # the path on the guest to mount the folder. And the optional third
  # argument is a set of non-required options.
  # config.vm.synced_folder "../data", "/vagrant_data"
  if OS.windows?
    config.vm.synced_folder "./", "/go/src/github.com/Mushus/cms", type: "virtualbox", mount_options: ['dmode=777','fmode=777']
  else
    config.vm.synced_folder "./", "/go/src/github.com/Mushus/cms", type: "nfs"
  end

  # Provider-specific configuration so you can fine-tune various
  # backing providers for Vagrant. These expose provider-specific options.
  # Example for VirtualBox:
  #
  # config.vm.provider "virtualbox" do |vb|
  #   # Display the VirtualBox GUI when booting the machine
  #   vb.gui = true
  #
  #   # Customize the amount of memory on the VM:
  #   vb.memory = "1024"
  # end
  #
  # View the documentation for the provider you are using for more
  # information on available options.

  # Enable provisioning with a shell script. Additional provisioners such as
  # Puppet, Chef, Ansible, Salt, and Docker are also available. Please see the
  # documentation for more information about their specific syntax and use.
  # config.vm.provision "shell", inline: <<-SHELL
  #   apt-get update
  #   apt-get install -y apache2
  # SHELL
  config.vm.provision "shell", inline: <<~SHELL
    GO_VERSION=1.9.2
    GO_OS=linux
    GO_ARCH=amd64

    apt-get update
    apt-get install git -y

    # install golang
    wget https://redirector.gvt1.com/edgedl/go/go$GO_VERSION.$GO_OS-$GO_ARCH.tar.gz -q -O- | tar -C /usr/local -xz

    cat <<'EOS' > /etc/profile.d/gopath.sh
    export GOROOT=/usr/local/go
    export GOPATH=/go
    export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
    EOS
    source /etc/profile.d/gopath.sh

    mkdir -p /go/pkg /go/bin /go/src
    chown -R vagrant /go/pkg /go/bin
    chmod -R 777 /go/pkg /go/bin

    chown vagrant /go/src /go/src/github.com
    chmod 777 /go/src /go/src/github.com
  SHELL

  config.vm.provision "shell", run: "always", inline: <<~SHELL
  SHELL
end
