#!/usr/bin/env python3
# -*- coding:utf-8 -*-

"""
    Loki module for directional

    Input:
        inputSTR      str,
        utterance     str,
        args          str[],
        resultDICT    dict,
        refDICT       dict,
        pattern       str

    Output:
        resultDICT    dict
"""

from random import sample
import json
import os
from ArticutAPI import Articut
BASE_PATH = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
try:
    accountInfo = json.load(open(os.path.join(BASE_PATH, "account.info"), encoding="utf-8"))
    USERNAME = accountInfo["username"]
    API_KEY = accountInfo["api_key"]
except Exception as e:
    print("[ERROR] AccountInfo => {}".format(str(e)))
    USERNAME = ""
    API_KEY = ""

articut = Articut(USERNAME, API_KEY)

DEBUG = False
CHATBOT_MODE = False

userDefinedDICT = {}
try:
    userDefinedDICT = json.load(open(os.path.join(os.path.dirname(__file__), "USER_DEFINED.json"), encoding="utf-8"))
except Exception as e:
    print("[ERROR] userDefinedDICT => {}".format(str(e)))

responseDICT = {}
if CHATBOT_MODE:
    try:
        responseDICT = json.load(open(os.path.join(os.path.dirname(os.path.dirname(__file__)), "reply/reply_directional.json"), encoding="utf-8"))
    except Exception as e:
        print("[ERROR] responseDICT => {}".format(str(e)))

# 將符合句型的參數列表印出。這是 debug 或是開發用的。
def debugInfo(inputSTR, utterance):
    if DEBUG:
        print("[directional] {} ===> {}".format(inputSTR, utterance))

def getResponse(utterance, args):
    resultSTR = ""
    if utterance in responseDICT:
        if len(responseDICT[utterance]):
            resultSTR = sample(responseDICT[utterance], 1)[0].format(*args)

    return resultSTR

def getResult(inputSTR, utterance, args, resultDICT, refDICT, pattern=""):
    debugInfo(inputSTR, utterance)
    if utterance == "外面":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            try:
                resultPOS = articut.parse(inputSTR)["result_pos"][0]
                if resultPOS.startswith("<ENTITY") or resultPOS.startswith("<LOCATION"):
                    resultDICT["方位"].append(1)
                else:
                    pass
            except:
                pass
            
    if utterance == "這裡":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            resultDICT["方位"].append(1)
            
    if utterance == "是這裡的嗎":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            resultDICT["方位"].append(1)
            
    if utterance == "嗯這下頭":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            resultDICT["方位"].append(1)        

    return resultDICT