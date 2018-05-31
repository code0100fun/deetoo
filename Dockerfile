FROM golang:1.10

RUN useradd -m -u 1000 deetoo
RUN curl -L -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 && chmod +x /usr/local/bin/dep
RUN apt-get update && apt-get install -y --no-install-recommends \
  automake \
  bison \
  build-essential \
  bzip2 \
  ca-certificates \
  clang \
  cpio \
  curl \
  debhelper \
  file \
  g++-multilib \
  gcc-multilib \
  genisoimage \
  git \
  gobject-introspection \
  gzip \
  intltool \
  libgirepository1.0-dev \
  libgsf-1-dev \
  libssl-dev \
  libtool \
  libxml2-dev \
  llvm-dev \
  make \
  mingw-w64 \
  patch \
  rpm \
  sed \
  uuid-dev \
  valac \
  wget \
  xz-utils

# install ruby
RUN mkdir -p /opt/ruby-2.5.1/ && \
  curl -s https://s3-external-1.amazonaws.com/heroku-buildpack-ruby/heroku-16/ruby-2.5.1.tgz | tar xzC /opt/ruby-2.5.1/

ENV PATH /opt/ruby-2.5.1/bin:$PATH
RUN mkdir -p /home/deetoo && chown deetoo:deetoo -R /home/deetoo

USER 1000
