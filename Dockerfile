# 表示依赖 alpine 最新版
FROM golang:alpine

# 为镜像设置必要的环境变量
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io"

# 移动到工作目录： /build
WORKDIR /build
COPY . .

# 将代码编译成二进制可执行文件 app
RUN go build -o app .

# 移动到用于存放生成的二进制文件的 /dist 目录
WORKDIR /dist

# 设置时区为上海
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

# 设置编码
ENV LANG C.UTF-8

# 将二进制文件从 /build 目录复制到这里
RUN cp /build/app .

# RUN mkdir templates .
RUN cp -r /build/templates ./templates
RUN cp -r /build/application.yml ./application.yml


EXPOSE 8092

# 运行golang程序的命令
CMD  ["/dist/app"]