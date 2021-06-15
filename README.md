# Discord-ADB py

Discord-ADB py is a Discord self bot for Android written in Python. It's been tested for the Samsung Galaxy A01 and passes, no information about how other phones interact with it.

## Usage

Prerequisites:

- Python 3
- Android SDK
- ADB
- Discord for Android

Clone repo branch:

```sh
git clone https://github.com/coffee-disposal/demos.git --branch discord-adb.py discord-adb.py
```

Connecting to Android device:

```sh
sudo adb start-server
sudo adb devices
```

Setting up Discord:
> Navigate to where you want to send the message on your Android device and make sure you have the keyboard open.

Running the script:

If you're on Linux/Unix, the command will be

```py
python3 main.py
```

If you're on Windows, the command will be

```py
python main.py
```
