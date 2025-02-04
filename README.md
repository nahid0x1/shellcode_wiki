# Shellcode Wiki

## Introduction
Shellcode is a small piece of executable code used in exploitation to control a system after successfully exploiting a vulnerability. It is typically written in assembly language and injected into a process's memory to execute arbitrary instructions.

- You can also optionally pack the output executable with UPX which will reduce the binary size from ~10MB to ~3MB. To do this, install [UPX](https://github.com/upx/upx/releases/) and run

```bash
upx.exe --ultra-brute payload.exe
```
