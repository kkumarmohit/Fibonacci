from flask import Flask, request, jsonify

app = Flask(__name__)

def fibonacci(n):
    if n <= 2:
        return 1
    a, b = 1, 1
    for _ in range(2, n):
        a, b = b, a + b
    return b

def fibonacci_recursive(n):
    memo = {}
    def helper(k):
        if k <= 2:
            return 1
        if k in memo:
            return memo[k]
        memo[k] = helper(k-1) + helper(k-2)
        return memo[k]
    return helper(n)

@app.route('/fibonacci')
def fibonacci_endpoint():
    order = request.args.get('order', type=int)
    if not order or order < 1:
        return 'Invalid or missing order parameter', 400
    value = fibonacci(order)
    return jsonify({'order': order, 'value': value})

@app.route('/recursive-fibonacci')
def recursive_fibonacci_endpoint():
    order = request.args.get('order', type=int)
    if not order or order < 1:
        return 'Invalid or missing order parameter', 400
    value = fibonacci_recursive(order)
    return jsonify({'order': order, 'value': value})

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)
