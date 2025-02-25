rm -rf ollama && wget http://cloud.loongnix.cn/releases/loongarch64abi1/ollama/ollama/0.4.0/ollama && chmod 755 ollama
docker build --build-arg https_proxy=$https_proxy -t cr.loongnix.cn/ollama/ollama:0.4.0 -f Dockerfile.loong64 .
docker push cr.loongnix.cn/ollama/ollama:0.4.0
