import { initTracerFromEnv } from 'jaeger-client';
import { APPLICATION_NAME, VERSION } from './constants';

const baseConfig = {
  serviceName: APPLICATION_NAME,
};

const baseOptions = {
  tags: {
    [`${APPLICATION_NAME}.version`]: VERSION,
  },
};

const tracer = initTracerFromEnv(baseConfig, baseOptions);

export { tracer };
