import os
import unittest

CASE_PATH = "./tests"


def run_tests():
    test_loader = unittest.TestLoader()
    test_suite = test_loader.discover(CASE_PATH, pattern="*_test.py")
    test_runner = unittest.TextTestRunner()
    test_runner.run(test_suite)


if __name__ == "__main__":
    print(os.path.abspath(CASE_PATH))
    run_tests()
