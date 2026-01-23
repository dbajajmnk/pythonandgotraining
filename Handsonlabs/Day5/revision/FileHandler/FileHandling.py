
def writeTextFile(textFileName,data:list[str]):
    with open(textFileName,mode="w",newline="\n") as f:
        for data in data:
            f.write(data)

