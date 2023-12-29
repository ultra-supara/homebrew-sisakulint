FROM ubuntu:latest

SHELL ["/bin/bash", "-c"]
RUN apt-get update && apt-get install -y curl git vim
RUN /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
RUN echo 'export PATH="/home/linuxbrew/.linuxbrew/bin:$PATH"' >> ~/.bashrc
RUN source ~/.bashrc
RUN /home/linuxbrew/.linuxbrew/bin/brew tap ultra-supara/homebrew-sisakulint

CMD ["/bin/bash"]
