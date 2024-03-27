#!/usr/bin/env python3
# -*- coding:utf-8 -*-

"""
    Loki 4.0 Template For Python3

    [URL] https://api.droidtown.co/Loki/BulkAPI/

    Request:
        {
            "username": "your_username",
            "input_list": ["your_input_1", "your_input_2"],
            "loki_key": "your_loki_key",
            "filter_list": ["intent_filter_list"] # optional
        }

    Response:
        {
            "status": True,
            "msg": "Success!",
            "version": "v223",
            "word_count_balance": 2000,
            "result_list": [
                {
                    "status": True,
                    "msg": "Success!",
                    "results": [
                        {
                            "intent": "intentName",
                            "pattern": "matchPattern",
                            "utterance": "matchUtterance",
                            "argument": ["arg1", "arg2", ... "argN"]
                        },
                        ...
                    ]
                },
                {
                    "status": False,
                    "msg": "No matching Intent."
                }
            ]
        }
"""

from copy import deepcopy
from glob import glob
from importlib import import_module
from pathlib import Path
from requests import post
from requests import codes
from pprint import pprint
import json
import math
import os
import re
import sys
import pandas as pd
import numpy as np

BASE_PATH = os.path.dirname(os.path.abspath(__file__))

lokiIntentDICT = {}
for modulePath in glob("{}/intent/Loki_*.py".format(BASE_PATH)):
    moduleNameSTR = Path(modulePath).stem[5:]
    modulePathSTR = modulePath.replace(BASE_PATH, "").replace(".py", "").replace("/", ".").replace("\\", ".")[1:]
    globals()[moduleNameSTR] = import_module(modulePathSTR)
    lokiIntentDICT[moduleNameSTR] = globals()[moduleNameSTR]

LOKI_URL = "https://api.droidtown.co/Loki/BulkAPI/"
try:
    accountInfo = json.load(open(os.path.join(BASE_PATH, "account.info"), encoding="utf-8"))
    USERNAME = accountInfo["username"]
    LOKI_KEY = accountInfo["loki_key"]
except Exception as e:
    print("[ERROR] AccountInfo => {}".format(str(e)))
    USERNAME = ""
    LOKI_KEY = ""

# 意圖過濾器說明
# INTENT_FILTER = []        => 比對全部的意圖 (預設)
# INTENT_FILTER = [intentN] => 僅比對 INTENT_FILTER 內的意圖
INTENT_FILTER = []
INPUT_LIMIT = 20

class LokiResult():
    status = False
    message = ""
    version = ""
    balance = -1
    lokiResultLIST = []

    def __init__(self, inputLIST, filterLIST):
        self.status = False
        self.message = ""
        self.version = ""
        self.balance = -1
        self.lokiResultLIST = []
        # filterLIST 空的就採用預設的 INTENT_FILTER
        if filterLIST == []:
            filterLIST = INTENT_FILTER

        try:
            result = post(LOKI_URL, json={
                "username": USERNAME,
                "input_list": inputLIST,
                "loki_key": LOKI_KEY,
                "filter_list": filterLIST
            })

            if result.status_code == codes.ok:
                result = result.json()
                self.status = result["status"]
                self.message = result["msg"]
                if result["status"]:
                    self.version = result["version"]
                    if "word_count_balance" in result:
                        self.balance = result["word_count_balance"]
                    self.lokiResultLIST = result["result_list"]
            else:
                self.message = "{} Connection failed.".format(result.status_code)
        except Exception as e:
            self.message = str(e)

    def getStatus(self):
        return self.status

    def getMessage(self):
        return self.message

    def getVersion(self):
        return self.version

    def getBalance(self):
        return self.balance

    def getLokiStatus(self, index):
        rst = False
        if index < len(self.lokiResultLIST):
            rst = self.lokiResultLIST[index]["status"]
        return rst

    def getLokiMessage(self, index):
        rst = ""
        if index < len(self.lokiResultLIST):
            rst = self.lokiResultLIST[index]["msg"]
        return rst

    def getLokiLen(self, index):
        rst = 0
        if index < len(self.lokiResultLIST):
            if self.lokiResultLIST[index]["status"]:
                rst = len(self.lokiResultLIST[index]["results"])
        return rst

    def getLokiResult(self, index, resultIndex):
        lokiResultDICT = None
        if resultIndex < self.getLokiLen(index):
            lokiResultDICT = self.lokiResultLIST[index]["results"][resultIndex]
        return lokiResultDICT

    def getIntent(self, index, resultIndex):
        rst = ""
        lokiResultDICT = self.getLokiResult(index, resultIndex)
        if lokiResultDICT:
            rst = lokiResultDICT["intent"]
        return rst

    def getPattern(self, index, resultIndex):
        rst = ""
        lokiResultDICT = self.getLokiResult(index, resultIndex)
        if lokiResultDICT:
            rst = lokiResultDICT["pattern"]
        return rst

    def getUtterance(self, index, resultIndex):
        rst = ""
        lokiResultDICT = self.getLokiResult(index, resultIndex)
        if lokiResultDICT:
            rst = lokiResultDICT["utterance"]
        return rst

    def getArgs(self, index, resultIndex):
        rst = []
        lokiResultDICT = self.getLokiResult(index, resultIndex)
        if lokiResultDICT:
            rst = lokiResultDICT["argument"]
        return rst

def runLoki(inputLIST, filterLIST=[], refDICT={}):
    resultDICT = deepcopy(refDICT)
    lokiRst = LokiResult(inputLIST, filterLIST)
    if lokiRst.getStatus():
        for index, key in enumerate(inputLIST):
            lokiResultDICT = {k: [] for k in refDICT}
            for resultIndex in range(0, lokiRst.getLokiLen(index)):
                if lokiRst.getIntent(index, resultIndex) in lokiIntentDICT:
                    lokiResultDICT = lokiIntentDICT[lokiRst.getIntent(index, resultIndex)].getResult(
                        key, lokiRst.getUtterance(index, resultIndex), lokiRst.getArgs(index, resultIndex),
                        lokiResultDICT, refDICT, pattern=lokiRst.getPattern(index, resultIndex))

            # save lokiResultDICT to resultDICT
            for k in lokiResultDICT:
                if k not in resultDICT:
                    resultDICT[k] = []
                if type(resultDICT[k]) != list:
                    resultDICT[k] = [resultDICT[k]] if resultDICT[k] else []
                if type(lokiResultDICT[k]) == list:
                    resultDICT[k].extend(lokiResultDICT[k])
                else:
                    resultDICT[k].append(lokiResultDICT[k])
    else:
        resultDICT["msg"] = lokiRst.getMessage()
    return resultDICT

def execLoki(content, filterLIST=[], splitLIST=[], refDICT={}):
    """
    input
        content       STR / STR[]    要執行 loki 分析的內容 (可以是字串或字串列表)
        filterLIST    STR[]          指定要比對的意圖 (空列表代表不指定)
        splitLIST     STR[]          指定要斷句的符號 (空列表代表不指定)
                                     * 如果一句 content 內包含同一意圖的多個 utterance，請使用 splitLIST 切割 content
        refDICT       DICT           參考內容

    output
        resultDICT    DICT           合併 runLoki() 的結果

    e.g.
        splitLIST = ["！", "，", "。", "？", "!", ",", "\n", "；", "\u3000", ";"]
        resultDICT = execLoki("今天天氣如何？後天氣象如何？")                      # output => ["今天天氣"]
        resultDICT = execLoki("今天天氣如何？後天氣象如何？", splitLIST=splitLIST) # output => ["今天天氣", "後天氣象"]
        resultDICT = execLoki(["今天天氣如何？", "後天氣象如何？"])                # output => ["今天天氣", "後天氣象"]
    """
    resultDICT = deepcopy(refDICT)
    if resultDICT is None:
        resultDICT = {}

    contentLIST = []
    if type(content) == str:
        contentLIST = [content]
    if type(content) == list:
        contentLIST = content

    if contentLIST:
        if splitLIST:
            # 依 splitLIST 做分句切割
            splitPAT = re.compile("[{}]".format("".join(splitLIST)))
            inputLIST = []
            for c in contentLIST:
                tmpLIST = splitPAT.split(c)
                inputLIST.extend(tmpLIST)
            # 去除空字串
            while "" in inputLIST:
                inputLIST.remove("")
        else:
            # 不做分句切割處理
            inputLIST = contentLIST

        # 依 INPUT_LIMIT 限制批次處理
        for i in range(0, math.ceil(len(inputLIST) / INPUT_LIMIT)):
            resultDICT = runLoki(inputLIST[i*INPUT_LIMIT:(i+1)*INPUT_LIMIT], filterLIST=filterLIST, refDICT=resultDICT)
            if "msg" in resultDICT:
                break

    return resultDICT

def loadFile(filename):
    try:
        with open(filename, 'r', encoding='utf-8') as f:
            return f.readlines()
    except UnicodeDecodeError:
        with open(filename, 'r', encoding='big5') as f:
            return f.readlines()

def generateReportCsv(outputData,outputFilePath):
    columnsList = ["index","sentence","量詞-個","量詞-特定","X 的","X 的 Y","方位","體貌標記","結果補語","趨向補語","情態補語","可能補語","數量補語","動前介詞","動後介詞","把字句","被字句","存現句","連動/兼語","帶連詞複句","緊縮複句","感知/心理狀態主謂句"]
    defaultRaws = [
        [np.nan,np.nan,"名詞短語",np.nan,np.nan,np.nan,np.nan,"動詞短語",np.nan,np.nan,np.nan,np.nan,np.nan,"介詞短語",np.nan,"句子",np.nan,np.nan,np.nan,np.nan,np.nan,np.nan],
        [np.nan,np.nan,"NP1","NP2","NP3","NP4","NP5","VP1","VP2","VP3","VP4","VP5","VP6","PP1","PP2","S1","S2","S3","S4","S5","S6","S7"],
        [np.nan,np.nan,"量詞-個","量詞-特定","X 的","X 的 Y","方位","體貌標記","結果補語","趨向補語","情態補語","可能補語","數量補語","動前介詞","動後介詞","把字句","被字句","存現句","連動/兼語","帶連詞複句","緊縮複句","感知/心理狀態主謂句"]
    ]
    metricsScore = [np.nan,"項目得分"]+[s[0] for s in outputData["Scores"][:20]]
    sentenceScore = [np.nan,"句法類型得分"]+[s[1] for s in outputData["Scores"][:20]]
    totalScore = [['總分','項目得分','句法類型得分'],['名詞短語類',*outputData["Scores"][20]],['動詞短語類',*outputData["Scores"][21]],['介詞短語類',*outputData["Scores"][22]],['句子類',*outputData["Scores"][23]],['語法複雜度總分',*outputData["Scores"][24]]]
    totalScore = [[np.nan,*t] for t in totalScore]
    reportData = defaultRaws + outputData["Sentences"] +[[]] + [metricsScore,sentenceScore] + totalScore
    report_df = pd.DataFrame(reportData, columns=columnsList)
    report_df = report_df.replace("-", np.NaN)
    report_df.to_csv(outputFilePath,header=0,index=0 ,encoding='utf-8-sig')

if __name__ == "__main__":
    
    # 測試所有意圖
    #testIntent()

    # 測試其它句子
    filterLIST = []
    splitLIST = ["！", "，", "。", "？", "!", ",", "\n", "；", "\u3000", ";"]
    # 設定參考資料
    # refDICT = {"key":[]} #refDICT 指定 key 需為字串 str，其值為一空列表 list
    refDICT = {
        '量-個': [],
        '量-特': [],
        'X的': [],
        'X的Y': [],
        '方位': [],
        '體貌': [],
        '結果補語': [],
        '趨向補語': [],
        '情態補語':[],
        '可能補語':[],
        '數量補語':[],
        '動前介詞':[],
        '動後介詞': [],
        '把字句': [],
        '被字句': [],
        '存現句': [],
        '連謂/兼語': [],
        '帶連詞複句': [],
        '緊縮複句': [],
        '感知/心理狀態': []
    }    

    outputData = {
        "Scores":[],
        "Sentences":[]
    }

    # inputSTR = "有一個小朋友，早上起床，覺得肚子很餓，他下床，走到房間外面，他打開冰箱，拿出一瓶牛奶，把牛奶倒進杯子，拿起來喝，喝完了。講完了"
    # 使用示例
    count = 1
    sentencesLIST = []
    itemScoreDICT ={}
    inputFilePath , outputFilePath= sys.argv[1] , sys.argv[2]
    f = loadFile(inputFilePath)

    for line in f:
        inputSTR = re.sub(r'[^\u4e00-\u9fa5]', '', line)
        resultDICT = execLoki(content=inputSTR, splitLIST=splitLIST, refDICT=refDICT,filterLIST=filterLIST) 
        sentencesLIST =[str(count),inputSTR]
        for key,value in resultDICT.items():
            if key not in itemScoreDICT:
                itemScoreDICT[key] = 0
            if isinstance(value, list) and value:
                sentencesLIST.append(str(value[0]))
                itemScoreDICT[key] += value[0]
            elif isinstance(value, list) and not value:
                sentencesLIST.append("-")
        outputData["Sentences"].append(sentencesLIST)
        count = count+1

    itemScoreSum = 0
    categoryScoreSum = 0
    
    for key,value in itemScoreDICT.items():
        itemScoreSum = itemScoreSum + value
    
    categoryScoreDICT = deepcopy(itemScoreDICT)
    for key,value in categoryScoreDICT.items():
        if categoryScoreDICT[key]-2 >= 0:
            categoryScoreDICT[key] = 2
        else:
            pass
        categoryScoreSum = categoryScoreSum + categoryScoreDICT[key]
    
    scoreLIST = []
    npLIST =[0,0]
    vpLIST =[0,0]
    ppLIST =[0,0]
    sLIST =[0,0]
    sumLIST = [itemScoreSum,categoryScoreSum]
    for key,value in itemScoreDICT.items():
        if key in ['量-個','量-特','X的','X的Y','方位']:
            npLIST[0]+= value
            npLIST[1]+= categoryScoreDICT[key]
        if key in ['體貌','結果補語','趨向補語','情態補語','可能補語','數量補語']:
            vpLIST[0]+= value
            vpLIST[1]+= categoryScoreDICT[key]
        if key in ['動前介詞','動後介詞']:
            ppLIST[0]+= value
            ppLIST[1]+= categoryScoreDICT[key]
        if key in ['把字句','被字句','存現句','連謂/兼語','帶連詞複句','緊縮複句','感知/心理狀態']:
            sLIST[0]+= value
            sLIST[1]+= categoryScoreDICT[key]
        scoreLIST = [value,categoryScoreDICT[key]]
        outputData["Scores"].append(scoreLIST)
    calLIST = [npLIST,vpLIST,ppLIST,sLIST,sumLIST]
    for list in calLIST:
        outputData["Scores"].append(list)
    
    generateReportCsv(outputData,outputFilePath)
    
    # "output.json"
    outputJson = json.dumps(outputData,ensure_ascii=False) 
    print(outputJson)