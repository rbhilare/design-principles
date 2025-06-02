#!/bin/bash

sudo apt update
sudo apt install -y datadog-agent
systemctl enable datadog-agent