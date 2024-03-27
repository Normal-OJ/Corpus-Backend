#!/usr/bin/env python3
# -*- coding:utf-8 -*-

"""
    Loki module for compact_sentence

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
        responseDICT = json.load(open(os.path.join(os.path.dirname(os.path.dirname(__file__)), "reply/reply_compact_sentence.json"), encoding="utf-8"))
    except Exception as e:
        print("[ERROR] responseDICT => {}".format(str(e)))

# 將符合句型的參數列表印出。這是 debug 或是開發用的。
def debugInfo(inputSTR, utterance):
    if DEBUG:
        print("[compact_sentence] {} ===> {}".format(inputSTR, utterance))

def getResponse(utterance, args):
    resultSTR = ""
    if utterance in responseDICT:
        if len(responseDICT[utterance]):
            resultSTR = sample(responseDICT[utterance], 1)[0].format(*args)

    return resultSTR

def getResult(inputSTR, utterance, args, resultDICT, refDICT, pattern=""):
    debugInfo(inputSTR, utterance)
    if utterance == " 你有煮什麼我就吃什麼":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            resultDICT["緊縮複句"].append(1)
            pass

    if utterance == "你早點來你就看到他":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            resultDICT["緊縮複句"].append(1)
            pass

    if utterance == "小狗看到爸爸小狗才離開":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            resultDICT["緊縮複句"].append(1)
            pass

    if utterance == "想吃什麼都可以吃":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            resultDICT["緊縮複句"].append(1)
            pass

    if utterance == "想要蓋什麼就蓋什麼":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            resultDICT["緊縮複句"].append(1)
            pass

    if utterance == "你要拿的時候你再裝進去":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            resultDICT["緊縮複句"].append(1)
            pass

    if utterance == "你閉上眼睛我要藏東西":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            resultDICT["緊縮複句"].append(1)
            pass

    if utterance == "固定起來就是門了":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            resultDICT["緊縮複句"].append(1)
            pass

    if utterance == "好了再取消":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            resultDICT["緊縮複句"].append(1)
            pass

    if utterance == "我做一個花園跟你合在一起":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            resultDICT["緊縮複句"].append(1)
            pass

    if utterance == "拉一下會搖":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            resultDICT["緊縮複句"].append(1)
            pass

    if utterance == "放一個圓形就好了":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            resultDICT["緊縮複句"].append(1)
            pass

    if utterance == "火車這樣就可以過去":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            resultDICT["緊縮複句"].append(1)
            pass

    if utterance == "自己會自己蓋自己的":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            resultDICT["緊縮複句"].append(1)
            pass

    if utterance == "要剪剪這樣子寶寶才能吃":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            resultDICT["緊縮複句"].append(1)
            pass

    if utterance == "跳到這裡再跳過來":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            resultDICT["緊縮複句"].append(1)
            pass

    if utterance == "這我的這也是你的":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            resultDICT["緊縮複句"].append(1)
            pass

    if utterance == "那你就先帶著":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            resultDICT["緊縮複句"].append(1)
            pass

    if utterance == "越挖越大":
        if CHATBOT_MODE:
            resultDICT["response"] = getResponse(utterance, args)
        else:
            resultDICT["緊縮複句"].append(1)
            pass

    return resultDICT