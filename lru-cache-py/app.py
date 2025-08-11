from flask import Flask, request, jsonify
import requests
from collections import OrderedDict

app = Flask(__name__)

class LRUCache:
    def __init__(self, capacity=5):
        self.cache = OrderedDict()
        self.capacity = capacity

    def get(self, key):
        if key in self.cache:
            self.cache.move_to_end(key)
            return self.cache[key]
        return None

    def put(self, key, value):
        if key in self.cache:
            self.cache.move_to_end(key)
        self.cache[key] = value
        if len(self.cache) > self.capacity:
            self.cache.popitem(last=False)

cache = LRUCache(5)

BACKEND_URLS = {
    'fibonacci': 'http://fibonacci-api-py:8080/fibonacci',
    'recursive-fibonacci': 'http://fibonacci-api-py:8080/recursive-fibonacci'
}

def proxy_and_cache(order, endpoint):
    cached_value = cache.get(order)
    if cached_value is not None:
        return {'fibonacci': {'order': order, 'value': cached_value}, 'cached': True}
    # Not cached, fetch from backend
    resp = requests.get(BACKEND_URLS[endpoint], params={'order': order})
    if resp.status_code != 200:
        return None
    data = resp.json()
    cache.put(order, data['value'])
    return {'fibonacci': data, 'cached': False}

@app.route('/fibonacci')
def fibonacci():
    order = request.args.get('order', type=int)
    if not order or order < 1:
        return 'Invalid or missing order parameter', 400
    result = proxy_and_cache(order, 'fibonacci')
    if result is None:
        return 'Failed to fetch from backend', 500
    return jsonify(result)

@app.route('/recursive-fibonacci')
def recursive_fibonacci():
    order = request.args.get('order', type=int)
    if not order or order < 1:
        return 'Invalid or missing order parameter', 400
    result = proxy_and_cache(order, 'recursive-fibonacci')
    if result is None:
        return 'Failed to fetch from backend', 500
    return jsonify(result)

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8081)
