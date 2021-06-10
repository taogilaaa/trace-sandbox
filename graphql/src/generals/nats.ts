import { connect } from 'node-nats-streaming';
import { v4 as uuid } from 'uuid';
import { NATS_CLUSTER, NATS_URL } from './constants';

type SendOptions = {
  channel: string;
  // eslint-disable-next-line @typescript-eslint/ban-types
  message: object;
};

async function sendMessage({ channel, message }: SendOptions): Promise<string> {
  const clientId = uuid();
  const stan = connect(NATS_CLUSTER, clientId, {
    url: NATS_URL,
    name: clientId,
  });
  const jsonMessage = JSON.stringify(message);

  return new Promise((resolve, reject) => {
    stan.on('connect', () => {
      stan.publish(channel, jsonMessage, (err, guid) => {
        stan.close();

        if (err) {
          reject(err);
        }

        resolve(guid);
      });
    });

    stan.on('error', (_reason) => {
      reject();
    });
  });
}

export { sendMessage };
