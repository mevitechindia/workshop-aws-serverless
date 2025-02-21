from flask import Flask, request

app = Flask(__name__)

def get_ip():
    # Check the X-Forwarded-For header for proxies
    xff = request.headers.get("X-Forwarded-For")
    if xff:
        return xff.split(",")[0].strip()
    
    # Check the X-Real-IP header
    xri = request.headers.get("X-Real-IP")
    if xri:
        return xri
    
    # Default to using request.remote_addr
    return request.remote_addr

@app.route('/ipaddress', methods=['GET'])
def ip_address():
    ip = get_ip()
    return f"Your IP address is: {ip}\n"

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)
