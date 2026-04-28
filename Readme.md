（1）在loongarch平台上关闭GGML_CPU_ALL_VARIANTS
该变量的作用常用于编译包含多种指令集优化(如x86的SSE/AVX)的““胖二进制或通用库，以确保生成的库能在不同型号的CPU上运行。
（2）loongarch平台上不使用文件ggml-cpu/arch/loongarch/quants.c
该文件是针对每个架构编写的加速函数，虽然在文件ml/backend/ggml/ggml/src/ggml-cpu/CMakeLists.txt中调用了这个文件，但是目前官方代码中还没有这个文件，直接构建官方代码报错
