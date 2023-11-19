import sys
import functools
import traceback
import inspect
import unittest


def retry(target=None, max_n=3, func_prefix="test"):
    def decorator(func_or_cls):
        if inspect.isfunction(func_or_cls):
            @functools.wraps(func_or_cls)
            def wrapper(*args, **kwargs):
                n = 0
                while n <= max_n:
                    try:
                        n += 1
                        func_or_cls(*args, **kwargs)
                        return
                    except Exception:
                        if n <= max_n:
                            trace = sys.exc_info()
                            traceback_info = str()
                            for trace_line in traceback.format_exception(trace[0], trace[1], trace[2], 3):
                                traceback_info += trace_line
                            print(traceback_info)
                            args[0].tearDown()
                            args[0].setUp()
                        else:
                            raise

            return wrapper
        elif inspect.isclass(func_or_cls):
            for name, func in list(func_or_cls.__dict__.items()):
                if inspect.isfunction(func) and name.startswith(func_prefix):
                    setattr(func_or_cls, name, decorator(func))
            return func_or_cls
        else:
            raise AttributeError

    if target:
        return decorator(target)
    else:
        return decorator
