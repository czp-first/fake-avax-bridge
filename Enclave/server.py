import json
import socket

from loguru import logger

from routes import ROUTES_MAP


ENCLAVE_PORT = 8000
ip = '127.0.0.1'


def main():
    sk = socket.socket()
    sk.bind((ip, ENCLAVE_PORT))
    logger.debug('start: {}:{}', ip, ENCLAVE_PORT)
    sk.listen()
    while True:
        conn, _addr = sk.accept()
        logger.info('Received new connection')
        payload = conn.recv(4096)

        try:
            client_request = json.loads(payload.decode())
            logger.info('client_request: {}', client_request)
        except Exception as exc:
            msg = f'Exception ({type(exc)}) while loading JSON data: {str(exc)}'
            content = {
                'success': False,
                'error': msg
            }
            conn.send(str.encode(json.dumps(content)))
            conn.close()
            continue
        method = client_request['method']

        handler = ROUTES_MAP.get(method)
        if not handler:
            resp = {}
        else:
            resp = handler(**client_request['body'])

        logger.info({'resp': resp})
        conn.send(str.encode(json.dumps({'content': resp})))
        conn.close()
        logger.info('closed connection')


if __name__ == '__main__':
    main()
