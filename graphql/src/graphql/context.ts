import { sendMessage } from '../generals/nats';
import { v4 as uuid } from 'uuid';

export type Context = ReturnType<typeof makeContext>;

function makeContext() {
  const requestId = uuid();

  return {
    date: {
      now: () => new Date(),
    },
    nats: {
      sendMessage,
    },
    requestId,
  };
}

export { makeContext };
