import requests as rq
import sys
from pathlib import Path
import os

no_gender = 'TCCM/cheng/030100.cha,\
TCCM/cheng/030300.cha,\
TCCM/cheng/030400.cha,\
TCCM/cheng/030700.cha,\
TCCM/cheng/030800.cha,\
TCCM/chou/020400.cha,\
TCCM/chou/020700.cha,\
TCCM/chou/020800.cha,\
TCCM/chou/021100.cha,\
TCCM/chou/030000.cha,\
TCCM/chw/030900.cha,\
TCCM/chw/031100.cha,\
TCCM/chw/040000.cha,\
TCCM/jc/020900.cha,\
TCCM/jc/030100.cha,\
TCCM/wang/020600.cha,\
TCCM/wang/020800.cha,\
TCCM/wang/021100.cha,\
TCCM/wang/030100.cha,\
TCCM/wang/030300.cha,\
TCCM/wu/010700.cha,\
TCCM/wu/010800.cha,\
TCCM/wuys/020700.cha,\
TCCM/wuys/020800.cha,\
TCCM/wuys/030200.cha,\
TCCM/wuys/030300.cha,\
TCCM/wuys/031000.cha,\
TCCM/xu/010700.cha,\
TCCM/xu/011100.cha,\
TCCM/xu/020000.cha,\
TCCM/xu/020100.cha,\
TCCM/xu/020200.cha,\
TCCM/yang/011000.cha,\
TCCM/yang/020200.cha,\
TCCM/yang/020800.cha,\
TCCM/yang/020900.cha,'.split(',')


'''
上傳步驟
1. 把 router 打開
2. 把 data.db 刪掉
3. 把 cha_store 資料夾清空
4. 跑腳本，跑法：python3 corpus_upload.py <根資料夾位置>
'''

def upload(src, dst):
    url = f'https://kideval.hdfs.ntnu.edu.tw/api/upload?file={dst}'
    # url = f'http://noj.tw:8777/api/upload?file={dst}'

    with rq.Session() as sess:
        filehandle = open(src)
        r = sess.post(url, files={'file': filehandle})
        print(r.text)


root = ''


def dfs(src):
    for path in Path(src).iterdir():
        name = str(path).split(root)[1][1:]
        if path.is_dir():
            dfs(str(path))
        elif name not in no_gender and '.cha' in name:
            upload(str(path), name)
            print(name)


if __name__ == "__main__":
    src = sys.argv[1] if len(sys.argv) >= 2 else None
    root = src

    if src is not None:
        dfs(src)
