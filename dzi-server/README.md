#### 导出依赖
- pip freeze > requirements.txt
- conda env export > environment.yml
#### 打包 exe
- Pyinstaller -F -w -i favicon.ico -n dzi-server main.py 