from typing import *
import requests
import json


class API:
    BASE_URL = "http://localhost:8080"

    def __init__(self, base_url: str = ""):
        if base_url != "":
            self.BASE_URL = base_url

    def request(self, retry: int = 3):
        pass

    def register(self, mobile: str, password: str) -> requests.Response:
        return requests.post(f"{self.BASE_URL}/user/register/mobile", json={
            "mobile": mobile,
            "password": password
        })

    def login(self, mobile: str, password: str):
        return requests.post(f"{self.BASE_URL}/user/login", json={
            "mobile": mobile,
            "password": password
        })
