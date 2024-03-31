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
        elif len(resultDICT["量-特"])  == 0:
            resultDICT["量-特"].append(1)

    if utterance == "多瓶":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        elif len(resultDICT["量-特"])  == 0:
            resultDICT["量-特"].append(1)

    
    if utterance == "很少瓶":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        elif len(resultDICT["量-特"])  == 0:
            resultDICT["量-特"].append(1)

    if utterance == "第二關":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        elif len(resultDICT["量-特"])  == 0:
            resultDICT["量-特"].append(1)

    if utterance == "一樣東西":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        elif len(resultDICT["量-特"])  == 0:
            resultDICT["量-特"].append(1)

    if utterance == "幾塊錢":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        elif len(resultDICT["量-特"])  == 0:
            resultDICT["量-特"].append(1)

    if utterance == "第一箱":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        elif len(resultDICT["量-特"])  == 0:
            resultDICT["量-特"].append(1)

    if utterance == "那瓶":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        elif len(resultDICT["量-特"])  == 0:
            resultDICT["量-特"].append(1)
    
    if utterance == "一把夾子":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        elif len(resultDICT["量-特"])  == 0:
            resultDICT["量-特"].append(1)

    if utterance == "大一點":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        elif len(resultDICT["量-特"])  == 0:
            resultDICT["量-特"].append(1)

    if utterance == "最上層":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        elif len(resultDICT["量-特"])  == 0:
            resultDICT["量-特"].append(1)

    if utterance == "有點發燒":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        elif len(resultDICT["量-特"])  == 0:
            resultDICT["量-特"].append(1)

    if utterance == "一頭牛":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            resultDICT["量-特"].append(1)

    if utterance == "下一班火車":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            resultDICT["量-特"].append(1)

    if utterance == "下一班車":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            resultDICT["量-特"].append(1)

    if utterance == "這班":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            resultDICT["量-特"].append(1)

    if utterance == "蓋小樓一點":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        elif len(resultDICT["量-特"])  == 0:
            resultDICT["量-特"].append(1)

    return resultDICT