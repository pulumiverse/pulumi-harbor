FROM  --platform=linux/amd64 pulumi/pulumi:3.185.0


# create a directory for pulumictl and download the binary to it and set to PATH
RUN mkdir -p /root/pulumictl && cd /root/pulumictl/
RUN wget https://github.com/pulumi/pulumictl/releases/download/v0.0.42/pulumictl-v0.0.42-linux-amd64.tar.gz -O /root/pulumictl/pulumictl-v0.0.42-linux-amd64.tar.gz
RUN tar -xvf /root/pulumictl/pulumictl-v0.0.42-linux-amd64.tar.gz -C /root/pulumictl/
ENV PATH="//root/pulumictl/:${PATH}"

RUN apt update
RUN apt install sudo -y

RUN type -p curl >/dev/null || (sudo apt update && sudo apt install curl -y)
RUN curl -fsSL https://cli.github.com/packages/githubcli-archive-keyring.gpg | sudo dd of=/usr/share/keyrings/githubcli-archive-keyring.gpg
RUN sudo chmod go+r /usr/share/keyrings/githubcli-archive-keyring.gpg
RUN echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/githubcli-archive-keyring.gpg] https://cli.github.com/packages stable main" | sudo tee /etc/apt/sources.list.d/github-cli.list > /dev/null
RUN sudo apt update
RUN sudo apt install gh -y
RUN sudo apt install vim -y

RUN go install github.com/pulumi/upgrade-provider@main
RUN pip install setuptools
