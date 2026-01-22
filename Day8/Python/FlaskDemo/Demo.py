from flask import Flask, jsonify

app = Flask(__name__)

@app.route('/api/greet', methods=['GET'])
def greet():
    # Returning a dict automatically converts to JSON in Flask 2.0+
    return {"message": "Welcome to the Live Class API!"}, 200

if __name__ == '__main__':
    app.run(debug=True)