from flask import Flask, jsonify, request

app = Flask(__name__)

# In-memory database
tasks = [
    {"id": 1, "title": "Learn Flask", "done": False},
    {"id": 2, "title": "Build REST API", "done": False}
]

# 1. GET - Retrieve all tasks
@app.route('/tasks', methods=['GET'])
def get_tasks():
    return jsonify(tasks)

# 2. POST - Create a new task
@app.route('/tasks', methods=['POST'])
def create_task():
    new_data = request.get_json()
    new_task = {
        "id": len(tasks) + 1,
        "title": new_data.get("title"),
        "done": False
    }
    tasks.append(new_task)
    return jsonify(new_task), 201

# 3. DELETE - Remove a task
@app.route('/tasks/<int:task_id>', methods=['DELETE'])
def delete_task(task_id):
    global tasks
    tasks = [t for t in tasks if t['id'] != task_id]
    return '', 204

if __name__ == '__main__':
    app.run(debug=True)