import hashlib
import random
import requests
import time


def _generate_header(token: str, secret_key: str, data: dict) -> dict:
    """
    生成请求头
    :param token: token
    :param secret_key: secret_key
    :param data: 参数（GET/POST）
    :return: 请求头字典
    """
    nonce = _nonce()
    return {
        'Nonce': nonce,
        'Token': token,
        'Signature': _sign(token, secret_key, nonce, data)
    }

def _sign(token: str, secret_key: str, nonce: str, data: dict) -> str:
    """
    生成签名
    :param token: token
    :param secret_key: secret_key
    :param nonce: 随机数
    :param data: 参数（GET/POST）
    :return: 签名字符串
    """
    tmp = [token, secret_key, nonce]
    for d, x in data.items():
        tmp.append(str(d) + "=" + str(x))
    return hashlib.sha1(''.join(sorted(tmp)).encode('utf-8')).hexdigest()

def _nonce() -> str:
    """生成随机数"""
    rs = '_'
    data = '124567890abcdefghijklmnopqrstuvwxyz'
    for _ in range(5):
        rs += random.choice(data)
    print(time.time())
    a = str(time.time())[:10] + rs
    print(a)
    return a

def get_demo():
    params = {"show_all":1}
    token = "4bec6394e490aca7acaae197379824d3"
    secret = "k51r7mii94jlebhk4ahq"
    headers = _generate_header(token, secret, params)
    print(headers)
    # 当前委托列表
    a = requests.get('https://openapi.aofex.co/openApi/wallet/list', params=params, headers=headers)
    print(a.json())
    

get_demo()