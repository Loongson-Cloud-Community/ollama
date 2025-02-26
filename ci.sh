rm -rf ollama*
wget http://cloud.loongnix.cn/releases/loongarch64/ollama/ollama/0.5.12/ollama && chmod 755 ollama
docker build --build-arg https_proxy=$https_proxy -t lcr.loongnix.cn/ollama/ollama:0.5.12 -f Dockerfile.loong64 .
docker push lcr.loongnix.cn/ollama/ollama:0.5.12
