rm -rf ollama*
wget http://cloud.loongnix.cn/releases/loongarch64/ollama/ollama/2025-0225/ollama && chmod 755 ollama
docker build --build-arg https_proxy=$https_proxy -t lcr.loongnix.cn/ollama/ollama:2025-0225 -f Dockerfile.loong64 .
docker push lcr.loongnix.cn/ollama/ollama:2025-0225
