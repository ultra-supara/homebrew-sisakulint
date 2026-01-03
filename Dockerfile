FROM python:3.17-alpine

RUN pip3 install --upgrade pip && \
  rm -r /root/.cache

RUN apk --update add git curl && \
  rm -rf /var/lib/apt/lists/* && \
  rm /var/cache/apk/*

# install reviewdog
ENV REVIEWDOG_VERSION=v0.17.0
RUN wget -O - -q https://raw.githubusercontent.com/reviewdog/reviewdog/master/install.sh | sh -s -- -b /usr/local/bin/ ${REVIEWDOG_VERSION}

# install sisakulint
ENV SISAKULINT_VERSION=0.1.2
ENV OSTYPE=linux-gnu

# sisakulintのダウンロードと解凍
RUN wget https://github.com/sisaku-security/sisakulint/releases/download/v${SISAKULINT_VERSION}/sisakulint_${SISAKULINT_VERSION}_linux_amd64.tar.gz -O sisakulint.tar.gz
RUN tar -xzf sisakulint.tar.gz -C /usr/local/bin

# Copy entrypoint script and set permission
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# 作業ディレクトリの設定（必要に応じて）
WORKDIR /usr/local/bin

ENTRYPOINT ["/entrypoint.sh"]
