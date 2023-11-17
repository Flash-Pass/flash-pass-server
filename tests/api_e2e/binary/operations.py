from tests.api_e2e.binary.api import API
from tests.api_e2e.binary import generator
import requests

api = API()


def register(mobile: str = "", password: str = "", gen_mobile: bool = True, gen_password: bool = True) -> (str, str, dict):
    if gen_mobile:
        mobile = generator.user_mobile()
    if gen_password:
        password = generator.user_password()

    resp: requests.Response = api.register(mobile, password)
    data = resp.json()

    return mobile, password, data


def login(mobile, password) -> dict:
    resp: requests.Response = api.login(mobile, password)
    data = resp.json()
    return data
