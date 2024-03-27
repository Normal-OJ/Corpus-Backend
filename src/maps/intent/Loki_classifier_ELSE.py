#!/usr/bin/env python3
# -*- coding:utf-8 -*-

"""
    Loki module for classifier_ELSE

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
import re
from ArticutAPI import Articut
BASE_PATH = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
try:
    accountInfo = json.load(open(os.path.join(BASE_PATH,"account.info"), encoding="utf-8"))
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
        responseDICT = json.load(open(os.path.join(os.path.dirname(os.path.dirname(__file__)), "reply/reply_classifier_ELSE.json"), encoding="utf-8"))
    except Exception as e:
        print("[ERROR] responseDICT => {}".format(str(e)))

# 將符合句型的參數列表印出。這是 debug 或是開發用的。
def debugInfo(inputSTR, utterance):
    if DEBUG:
        print("[classifier_ELSE] {} ===> {}".format(inputSTR, utterance))

def getResponse(utterance, args):
    resultSTR = ""
    if utterance in responseDICT:
        if len(responseDICT[utterance]):
            resultSTR = sample(responseDICT[utterance], 1)[0].format(*args)

    return resultSTR

def getResult(inputSTR, utterance, args, resultDICT, refDICT, pattern=""):
    debugInfo(inputSTR, utterance)
    if utterance == "一瓶":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            resultDICT["量-特"].append(1)

    if utterance == "多瓶":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            resultDICT["量-特"].append(1)

    if utterance == "很大台":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            pat = re.compile(pattern)  #<DegreeP>([^<]+)</DegreeP>
            try:
                posSTR = articut.parse(inputSTR)["result_pos"][0]
                termSTR = list(pat.finditer(posSTR))[0].group(1)  # termSTR => 很大瓶
                termSTR = f"一{termSTR[-1]}"# termSTR => 一瓶
                posSTR = articut.parse(termSTR)["result_pos"][0]  # posSTR => <ENTITY_classifier>一瓶</ENTITY_classifier>
                if posSTR.startswith("<ENTITY_classifier>"):
                    resultDICT["量-特"].append(1)
                else:
                    pass
            except:
                pass

    if utterance == "很少瓶":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            resultDICT["量-特"].append(1)

    if utterance == "第二關":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            resultDICT["量-特"].append(1)

    if utterance == "一樣東西":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            resultDICT["量-特"].append(1)

    if utterance == "幾塊錢":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            resultDICT["量-特"].append(1)

    if utterance == "第一箱":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            resultDICT["量-特"].append(1)

    if utterance == "那瓶":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            resultDICT["量-特"].append(1)

    return resultDICT