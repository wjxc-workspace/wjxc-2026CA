from pymodbus.client import ModbusTcpClient
import time
from threading import Thread
import random

# Connect to the Modbus server
client = ModbusTcpClient('10.0.100.40', port=502) # Replace with your server IP and port

count = 0

client.connect()

time.sleep(2)

# Write a single integer value (e.g., 100) to holding register address 0
# Modbus function code 6 (Write Single Register)
def send(n):
    def func(n):
        coils = []
        for i in range(0, 2):
            if i == n:
                coils.append(True)
            else:
                coils.append(False)
        client.write_coils(8, coils)
        client.write_coil(12, counting)
        client.write_coil(14, auto)

    Thread(target=func(n), daemon=True).start()

def enable():
    global counting
    counting = True

def disable():
    global counting
    counting = False

auto = False
counting = False
reset = False
auto_score = 0
teleop_score = 0

AUTO = 0
TELEOP = 1

def read_score(n):
    def func():
        global auto_score, teleop_score
        if n == AUTO:
            auto_score = client.read_holding_registers(1).registers[0]
        elif n == TELEOP:
            teleop_score = client.read_holding_registers(1).registers[0]
    Thread(target=func, daemon=True).start()

OFF = -1
BLUE = 0
BLUE_BLINK = 1
PURPLE = 2
GREEN = 3
COUNTING = 4

def run():
    global auto
    win = random.randint(0, 1) == 1
    enable()
    auto = True
    send(BLUE)
    time.sleep(20)

    time.sleep(3)
    read_score(AUTO)
    auto = False
    auto_score
    if win:
        time.sleep(4)
        send(BLUE_BLINK)
        time.sleep(3)

        disable()
        send(OFF)
        time.sleep(25)

        enable()
        send(BLUE)
        time.sleep(22)
        send(BLUE_BLINK)
        time.sleep(3)

        disable()
        send(OFF)
        time.sleep(25)

        enable()
        send(BLUE)
        time.sleep(52)
        send(BLUE_BLINK)
        time.sleep(3)
        send(OFF)
    else:
        time.sleep(10)

        time.sleep(22)
        send(BLUE_BLINK)
        time.sleep(3)

        disable()
        send(OFF)
        time.sleep(25)

        enable()
        send(BLUE)
        time.sleep(22)
        send(BLUE_BLINK)
        time.sleep(3)

        disable()
        send(OFF)
        time.sleep(25)

        enable()
        send(BLUE)
        time.sleep(27)
        send(BLUE_BLINK)
        time.sleep(3)
        send(OFF)


    time.sleep(3)
    client.write_coil(COUNTING + 8, False)
    read_score(TELEOP)
    time.sleep(3)
    print(f"auto: {auto_score}, teleop: {teleop_score}")

run()
# print(client.read_coils(0, count=14))