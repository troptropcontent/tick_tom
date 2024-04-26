FROM mcr.microsoft.com/devcontainers/go:1-1.22-bookworm

# Install dependencies
RUN sudo apt update \ 
    && sudo apt install gh -y

# Install sass from github
ENV SASS_VERSION 1.75.0
ADD https://github.com/sass/dart-sass/releases/download/${SASS_VERSION}/dart-sass-${SASS_VERSION}-linux-arm64.tar.gz /opt/
RUN tar -C /opt/ -xzvf /opt/dart-sass-${SASS_VERSION}-linux-arm64.tar.gz

# Install tailwind
RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-arm64 
RUN chmod +x tailwindcss-linux-arm64
RUN mv tailwindcss-linux-arm64 /tailwindcss

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# [Optional] Uncomment the next lines to use go get to install anything else you need
# USER vscode
# RUN go get -x <your-dependency-or-tool>
# USER root

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1
