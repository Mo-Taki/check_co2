import mh_z19
import time
import timeout_decorator

@timeout_decorator.timeout(10)
def read():
    value = ""
    while value == "":
        out = mh_z19.read()
        value = str(out)[8:-1]
    
    return value


if __name__ == '__main__':
    co2 = read()

    print(co2)
