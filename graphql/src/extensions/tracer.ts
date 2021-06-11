import OpentracingExtension, { SpanContext } from 'apollo-opentracing';
import { tracer } from '../generals/tracer';
import { Tags } from 'opentracing';
import { Context } from '../graphql/context';

type FullContext = SpanContext & Context;

function tracerExtension() {
  return new OpentracingExtension<FullContext>({
    server: tracer,
    local: tracer,
    shouldTraceRequest: (info) => {
      if (info.operationName === 'IntrospectionQuery') {
        return false;
      }

      return true;
    },
    onRequestResolve: (span, info) => {
      const requestId = info.context.requestId || '';

      span.addTags({
        [Tags.SPAN_KIND]: Tags.SPAN_KIND_RPC_SERVER,
        [Tags.COMPONENT]: 'apollo',
        [Tags.HTTP_METHOD]: info.request.method,
        [Tags.HTTP_URL]: info.request.url,
        'request.id': requestId,
      });
    },
    onFieldResolveFinish: (error, _result, span) => {
      if (error) {
        span.addTags({
          [Tags.ERROR]: true,
        });
        span.log({
          event: 'error',
          message: 'Request Error',
          'error.kind': error.name,
          'error.message': error.message,
          'error.stack': error.stack,
          'error.object': error,
        });
      }
    },
  });
}

export { tracerExtension };
