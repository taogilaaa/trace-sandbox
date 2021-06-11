import { name, version } from '../../package.json';

// eslint-disable-next-line @typescript-eslint/no-var-requires
require('dotenv').config();

export const APPLICATION_NAME = name;
export const VERSION = version;
export const APPLICATION_NAME_VERSION = `${name}-${version}`;

export const NATS_URL = process.env.NATS_URL ?? 'nats://localhost:4222';
export const NATS_CLUSTER = process.env.NATS_CLUSTER ?? 'test-cluster';
export const SALEORDER_PLACED_CHANNEL = 'sandbox.saleorder.placed';

export const GRPC_URL = process.env.GRPC_URL ?? 'localhost:50041';
