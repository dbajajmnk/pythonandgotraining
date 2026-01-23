from pathlib import Path

ROOT_DIR = Path("bussines")
JSON_FILE = ROOT_DIR/"client.json"

def createFiles():
    ROOT_DIR.mkdir(exist_ok=True)

createFiles()
