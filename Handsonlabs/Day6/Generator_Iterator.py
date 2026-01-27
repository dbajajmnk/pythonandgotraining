## Write a Content in file
    ## Need to create a functoin with file Path
    ## Need to open a file in Write Mode with Encoding and path
    ## Write the list of items in file with espace character "\n" new line right in next line
## Create Generator
##    Need to open file for read mode of given path
##    read file line by line
##    Check the condition to yield

## Testing 

def writeFileWithData(path):
    list = [
        "Austriala is Good Tourism Place",
        "Austriala is Good Tourism Place",
        "Austriala is Good Tourism Place",
        "Austriala is Good Tourism Place",
        "Austriala is Good Tourism Place",
        "India is Good Tourism Place",
        "India is Good Tourism Place",
        "India is Good Tourism Place",
        "India is Good Tourism Place",
        "India is Good Tourism Place"
    ]
    with open(path,"w",encoding="utf-8") as f:
        for line in list:
            f.write(line+"\n")

def genertorExample(path):
    with open(path,"r",encoding="utf-8") as f:
        for line in f:
            if "India" in line:
                yield line.rstrip("\n")

def main():
    path="myteamisgreat.txt"
    writeFileWithData(path)
    demo = genertorExample(path)
    for d in demo:
        print(d)


if __name__ == "__main__":
    main()
