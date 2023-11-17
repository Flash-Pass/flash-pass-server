import string
import random


def user_mobile() -> str:
    return "138" + "".join([
        str(random.randint(0, 9)) for _ in range(8)
    ])


def user_password() -> str:
    return "".join(random.choices(string.ascii_lowercase, k=12))
