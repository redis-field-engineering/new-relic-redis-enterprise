# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.box      = "ubuntu/bionic64"
  config.vm.hostname = "newrelic"
  config.vm.provider "virtualbox" do |v|
  config.vm.boot_timeout = 600
    v.memory = 2048
    v.cpus = 2
  end

  config.vm.provision "ansible" do |ansible|
    ansible.playbook           = "ansible/playbook.yml"
  end
end
