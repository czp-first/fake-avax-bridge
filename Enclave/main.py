# -*- coding: UTF-8 -*-
"""
@Summary : docstr
@Author  : Rey
@Time    : 2022-09-30 11:31:53
"""

from db.init_db import init_db
from server import server


if __name__ == '__main__':
    init_db()
    server()
