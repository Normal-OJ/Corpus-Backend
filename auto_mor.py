import subprocess
import os
from shutil import copyfile

CMD_PATH = os.getenv("CLANG_CMD_FOLDER")
MOR_LIB = os.getenv("MOR_LIB")

def run_mor(src_path:str, dst_filename = "result00.cha"):
    tmp_fname = "oao.cha"

    copyfile(src_path, tmp_fname) # this is just for demo !!!
    subprocess.check_output([f"{CMD_PATH}/mor", f"-L{MOR_LIB}", tmp_fname]).decode()
    copyfile(tmp_fname, dst_filename)
    os.remove(tmp_fname)
    print("##############################################")

def run_post(src_path = "result00.cha", dst_filename = "result01.cha"):
    tmp_fname = "obo.cha"

    copyfile(src_path, tmp_fname) # this is just for demo !!!
    subprocess.check_output([f"{CMD_PATH}/post", f"+d{MOR_LIB}/post.db", tmp_fname]).decode()
    copyfile(tmp_fname, dst_filename)
    os.remove(tmp_fname)
    print("##############################################")

def run_postmortem(src_path = "result01.cha", dst_filename = "result02.cha"):
    tmp_fname = "oco.cha"

    copyfile(src_path, tmp_fname) # this is just for demo !!!
    subprocess.check_output([f"{CMD_PATH}/postmortem", f"+L{MOR_LIB}", tmp_fname]).decode()
    copyfile(tmp_fname, dst_filename)
    os.remove(tmp_fname)
    print("##############################################")

def run_megrasp(src_path = "result02.cha", dst_filename = "result03.cha"):
    tmp_fname = "odo.cha"

    copyfile(src_path, tmp_fname) # this is just for demo !!!
    subprocess.check_output([f"{CMD_PATH}/megrasp", f"+L{MOR_LIB}", tmp_fname]).decode()
    copyfile(tmp_fname, dst_filename)
    os.remove(tmp_fname)
    print("##############################################")

def run_kideval(src_file = "result00.cha", dst_file = "result.xls", opts = ["+lzho", "+t*CHI"]):
    subprocess.check_output([f"{CMD_PATH}/kideval"]+opts+[src_file]).decode()

def strace_log(cmds:list):
    subprocess.check_output(["ltrace", "-e", "arc4random"] + cmds, stderr=open("strace.log", "w"))

#strace_log(["unix-clan/unix/bin/post","+dunix-clan/unix/bin/zho/post.db","result00.cha"])
run_mor("01_TP_time3.cha")
run_post()
run_postmortem()
run_megrasp()
run_kideval(src_file="result03.cha")