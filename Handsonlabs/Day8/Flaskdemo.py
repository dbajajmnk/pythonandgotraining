from __future__ import annotations

from flask import Flask, request, jsonify

app = Flask(__name__)

# @app.get("/home")
# def homeContent():
#     return "Hello World form Flask"

@app.route("/", methods=["GET"])
def homeContent():
    return "Hello World form Flask"

def main()-> None:
    app.run(host="127.122.7.1",port="5000",debug=True)

if __name__ == "__main__":
    main()







