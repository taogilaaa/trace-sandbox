import { sendMessage } from '../generals/nats';

export type Context = ReturnType<typeof makeContext>;

function makeContext() {
  return {
    date: {
      now: () => new Date(),
    },
    nats: {
      sendMessage,
    },
  };
}

export { makeContext };
