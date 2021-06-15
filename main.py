import os

msg = input("Input message: ")
wordList = msg.split(" ")
formattedMsg = ""

for i in range(len(wordList)):
    if i + 1 == len(wordList):
        formattedMsg += wordList[i]
    else:
        formattedMsg += f"{wordList[i]}\\ "

streamOne = os.popen(f"adb shell input text \"{formattedMsg}\"")
outputOne = streamOne.read()
print(outputOne)

streamTwo = os.popen("adb shell input tap 650 750")
outputTwo = streamTwo.read()
print(outputTwo)

print("Message sent successfully!")