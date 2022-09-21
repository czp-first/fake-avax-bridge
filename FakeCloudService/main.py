# -*- coding: UTF-8 -*-
"""
@Summary : docstr
@Author  : Rey
@Time    : 2022-07-18 10:11:48
"""

import json
from pathlib import Path as P

from cryptography.fernet import Fernet
from fastapi import Body, FastAPI, Path, status
from fastapi.responses import JSONResponse
from fastapi.staticfiles import StaticFiles
from pydantic import BaseModel
import uvicorn


app = FastAPI()

app.mount("/static", StaticFiles(directory="static"), name="static")


@app.put("/settings/{file_name}")
async def put_settings(body = Body(...), file_name: str = Path(...)):
    """更新文件"""
    q = P(".") / "static"

    file_path = q / file_name

    if not file_path.exists():
        return JSONResponse(status_code=status.HTTP_404_NOT_FOUND, content={"msg": "not found"})

    with file_path.open("w") as f:
        f.write(json.dumps(body))

    return {}


class EncryptBody(BaseModel):
    key: str
    plaintext: str


@app.post("/encrypt")
async def encrypt(body: EncryptBody = Body(...)):
    """加密"""
    f = Fernet(body.key.encode("utf-8"))
    ciphertext = f.encrypt(body.plaintext.encode("utf-8")).decode("utf-8")
    return {"ciphertext": ciphertext}


class DecryptBody(BaseModel):
    key: str
    ciphertext: str


@app.post("/decrypt")
async def decrypt(body: DecryptBody = Body(...)):
    """解密"""
    f = Fernet(body.key.encode("utf-8"))
    plaintext = f.decrypt(body.ciphertext.encode("utf-8")).decode("utf-8")
    return {"plaintext": plaintext}


if __name__ == '__main__':
    uvicorn.run("main:app", port=5000)
