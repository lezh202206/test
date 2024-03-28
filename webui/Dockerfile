FROM ubuntu:latest
LABEL maintainer="lezh<1157640697@qq.com>"
# 相关依赖
RUN apt-get update && \
    apt-get install -y git software-properties-common && \
    add-apt-repository ppa:deadsnakes/ppa && \
    apt-get update && \
    apt-get install -y ffmpeg libsm6 libxext6 python3.10 python3-pip python3.10-venv

RUN update-alternatives --install /usr/bin/python3 python3 /usr/bin/python3.10 1

# sd-webui 启动不允许 管理员
RUN useradd -m sd

# 拷贝远程 && 设置工作目录 && 赋权
RUN git clone https://github.com/AUTOMATIC1111/stable-diffusion-webui /home/sd/stable-diffusion-webui
WORKDIR /home/sd/stable-diffusion-webui
RUN chmod +x ./webui.sh && chown -R sd:sd /home/sd
USER sd

# 放行 7860
EXPOSE 7860

CMD ["./webui.sh", "--listen", "--skip-torch-cuda-test", "--precision", "full", "--no-half", "--allow-code", "--enable-insecure-extension-access", "--api", "--theme", "dark"]